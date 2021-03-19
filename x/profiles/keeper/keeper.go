package keeper

import (
	"fmt"
	"regexp"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

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
// It returns an error if a profile with the same dtag from a different creator already exists
func (k Keeper) StoreProfile(ctx sdk.Context, profile *types.Profile) error {
	addr := k.GetAddressFromDtag(ctx, profile.Dtag)
	if addr != "" && addr != profile.GetAddress().String() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with dtag %s has already been created", profile.Dtag)
	}

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

	dTagRegEx := regexp.MustCompile(params.DtagParams.RegEx)
	minDtagLen := params.DtagParams.MinDtagLength.Int64()
	maxDtagLen := params.DtagParams.MaxDtagLength.Int64()
	dtagLen := int64(len(profile.Dtag))

	if !dTagRegEx.MatchString(profile.Dtag) {
		return fmt.Errorf("invalid profile dtag, it should match the following regEx %s", dTagRegEx)
	}

	if dtagLen < minDtagLen {
		return fmt.Errorf("profile dtag cannot be less than %d characters", minDtagLen)
	}

	if dtagLen > maxDtagLen {
		return fmt.Errorf("profile dtag cannot exceed %d characters", maxDtagLen)
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
	key := types.DtagTransferRequestStoreKey(request.Receiver)

	var requests WrappedDTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	for _, req := range requests.Requests {
		if req.Sender == request.Sender && req.Receiver == request.Receiver {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"the transfer request from %s to %s has already been made",
				request.Sender, request.Receiver)
		}
	}

	requests = NewWrappedDTagTransferRequests(append(requests.Requests, request))
	store.Set(key, k.cdc.MustMarshalBinaryBare(&requests))
	return nil
}

// GetUserIncomingDTagTransferRequests returns all the request made to the given user inside the current context.
func (k Keeper) GetUserIncomingDTagTransferRequests(ctx sdk.Context, user string) []types.DTagTransferRequest {
	store := ctx.KVStore(k.storeKey)
	key := types.DtagTransferRequestStoreKey(user)

	var requests WrappedDTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &requests)
	return requests.Requests
}

// GetDTagTransferRequests returns all the requests inside the given context
func (k Keeper) GetDTagTransferRequests(ctx sdk.Context) (requests []types.DTagTransferRequest) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DTagTransferRequestsPrefix)

	for ; iterator.Valid(); iterator.Next() {
		var userRequests WrappedDTagTransferRequests
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &userRequests)
		requests = append(requests, userRequests.Requests...)
	}

	return requests
}

// DeleteAllDTagTransferRequests delete all the requests made to the given user
func (k Keeper) DeleteAllDTagTransferRequests(ctx sdk.Context, user string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DtagTransferRequestStoreKey(user))
}

// DeleteDTagTransferRequest deletes the transfer requests made from the sender towards the recipient
func (k Keeper) DeleteDTagTransferRequest(ctx sdk.Context, sender, recipient string) error {
	store := ctx.KVStore(k.storeKey)
	key := types.DtagTransferRequestStoreKey(recipient)

	var wrapped WrappedDTagTransferRequests
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &wrapped)

	for index, request := range wrapped.Requests {
		if request.Sender == sender {
			requests := append(wrapped.Requests[:index], wrapped.Requests[index+1:]...)
			if len(requests) == 0 {
				store.Delete(key)
			} else {
				store.Set(key, k.cdc.MustMarshalBinaryBare(&WrappedDTagTransferRequests{Requests: requests}))
			}
			return nil
		}
	}

	return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", sender, recipient)
}
