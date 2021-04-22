package keeper

import (
	"fmt"
	"regexp"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	relationshipskeeper "github.com/desmos-labs/desmos/x/staging/relationships/keeper"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           codec.BinaryMarshaler
	paramSubspace paramstypes.Subspace

	ak authkeeper.AccountKeeper
	rk relationshipskeeper.Keeper
}

// NewKeeper creates new instances of the profiles Keeper.
// This k stores the profile data using two different associations:
// 1. Address -> Profile
//    This is used to easily retrieve the profile of a user based on an address
// 2. DTag -> Address
//    This is used to get the address of a user based on a DTag
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey, paramSpace paramstypes.Subspace,
	rk relationshipskeeper.Keeper, ak authkeeper.AccountKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramSubspace: paramSpace,
		rk:            rk,
		ak:            ak,
	}
}

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.rk.HasUserBlocked(ctx, blocker, blocked, "")
}

// StoreProfile stores the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same DTag from a different creator already exists
func (k Keeper) StoreProfile(ctx sdk.Context, profile *types.Profile) error {
	addr := k.GetAddressFromDTag(ctx, profile.DTag)
	if addr != "" && addr != profile.GetAddress().String() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with dtag %s has already been created", profile.DTag)
	}

	store := ctx.KVStore(k.storeKey)

	// Remove the previous DTag association (if the DTag has changed)
	oldProfile, found, err := k.GetProfile(ctx, profile.GetAddress().String())
	if err != nil {
		return err
	}
	if found && oldProfile.DTag != profile.DTag {
		store.Delete(types.DTagStoreKey(oldProfile.DTag))
	}

	// Store the DTag -> Address association
	store.Set(types.DTagStoreKey(profile.DTag), profile.GetAddress())

	// Store the account inside the auth keeper
	k.ak.SetAccount(ctx, profile)
	return nil
}

// GetProfile returns the profile corresponding to the given address inside the current context.
func (k Keeper) GetProfile(ctx sdk.Context, address string) (profile *types.Profile, found bool, err error) {
	sdkAcc, err := sdk.AccAddressFromBech32(address)
	if err != nil {
		return nil, false, err
	}

	stored, ok := k.ak.GetAccount(ctx, sdkAcc).(*types.Profile)
	if !ok {
		return nil, false, nil
	}

	return stored, true, nil
}

// GetAddressFromDTag returns the address associated to the given DTag or an empty string if it does not exists
func (k Keeper) GetAddressFromDTag(ctx sdk.Context, dTag string) (addr string) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.DTagStoreKey(dTag))
	if bz == nil {
		return ""
	}

	return sdk.AccAddress(bz).String()
}

// RemoveProfile allows to delete a profile associated with the given address inside the current context.
// It assumes that the address-related profile exists.
func (k Keeper) RemoveProfile(ctx sdk.Context, address string) error {
	profile, found, err := k.GetProfile(ctx, address)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"no profile associated with the following address found: %s", address)
	}

	// Delete the DTag -> Address association
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DTagStoreKey(profile.DTag))

	// Delete the profile data by replacing the stored account
	k.ak.SetAccount(ctx, profile.GetAccount())
	return nil
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func (k Keeper) ValidateProfile(ctx sdk.Context, profile *types.Profile) error {
	params := k.GetParams(ctx)

	minMonikerLen := params.MonikerParams.MinMonikerLength.Int64()
	maxMonikerLen := params.MonikerParams.MaxMonikerLength.Int64()

	if profile.Moniker != "" {
		nameLen := int64(len(profile.Moniker))
		if nameLen < minMonikerLen {
			return fmt.Errorf("profile moniker cannot be less than %d characters", minMonikerLen)
		}
		if nameLen > maxMonikerLen {
			return fmt.Errorf("profile moniker cannot exceed %d characters", maxMonikerLen)
		}
	}

	dTagRegEx := regexp.MustCompile(params.DTagParams.RegEx)
	minDTagLen := params.DTagParams.MinDTagLength.Int64()
	maxDTagLen := params.DTagParams.MaxDTagLength.Int64()
	dTagLen := int64(len(profile.DTag))

	if !dTagRegEx.MatchString(profile.DTag) {
		return fmt.Errorf("invalid profile dtag, it should match the following regEx %s", dTagRegEx)
	}

	if dTagLen < minDTagLen {
		return fmt.Errorf("profile dtag cannot be less than %d characters", minDTagLen)
	}

	if dTagLen > maxDTagLen {
		return fmt.Errorf("profile dtag cannot exceed %d characters", maxDTagLen)
	}

	maxBioLen := params.MaxBioLength.Int64()
	if profile.Bio != "" && int64(len(profile.Bio)) > maxBioLen {
		return fmt.Errorf("profile biography cannot exceed %d characters", maxBioLen)
	}

	return profile.Validate()
}

// ___________________________________________________________________________________________________________________

// SaveDTagTransferRequest save the given request associating it to the request recipient.
// It returns an error if the same request already exists.
func (k Keeper) SaveDTagTransferRequest(ctx sdk.Context, request types.DTagTransferRequest) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(request.Receiver)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	for _, req := range requests.Requests {
		if req.Sender == request.Sender && req.Receiver == request.Receiver {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				request.Sender, request.Receiver)
		}
	}

	requests = types.NewDTagTransferRequests(append(requests.Requests, request))
	store.Set(key, k.cdc.MustMarshalBinaryBare(&requests))
	return nil
}

// GetUserIncomingDTagTransferRequests returns all the request made to the given user inside the current context.
func (k Keeper) GetUserIncomingDTagTransferRequests(ctx sdk.Context, user string) []types.DTagTransferRequest {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(user)

	var requests types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	return requests.Requests
}

// GetDTagTransferRequests returns all the requests inside the given context
func (k Keeper) GetDTagTransferRequests(ctx sdk.Context) (requests []types.DTagTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestsPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var userRequests types.DTagTransferRequests
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &userRequests)
		requests = append(requests, userRequests.Requests...)
	}

	return requests
}

// DeleteAllDTagTransferRequests delete all the requests made to the given user
func (k Keeper) DeleteAllDTagTransferRequests(ctx sdk.Context, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DTagTransferRequestStoreKey(user))
}

// DeleteDTagTransferRequest deletes the transfer requests made from the sender towards the recipient
func (k Keeper) DeleteDTagTransferRequest(ctx sdk.Context, sender, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DTagTransferRequestStoreKey(recipient)

	var wrapped types.DTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &wrapped)

	for index, request := range wrapped.Requests {
		if request.Sender == sender {
			requests := append(wrapped.Requests[:index], wrapped.Requests[index+1:]...)
			if len(requests) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.cdc.MustMarshalBinaryBare(&types.DTagTransferRequests{Requests: requests}))
			}
			return nil
		}
	}

	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", sender, recipient)
}
