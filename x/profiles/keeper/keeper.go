package keeper

import (
	"fmt"
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	relationshipskeeper "github.com/desmos-labs/desmos/x/relationships/keeper"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryMarshaler

	relKeeper     relationshipskeeper.Keeper
	paramSubspace paramstypes.Subspace
}

// NewKeeper creates new instances of the profiles Keeper.
// This k stores the profile data using two different associations:
// 1. Address -> Profile
//    This is used to easily retrieve the profile of a user based on an address
// 2. DTag -> Address
//    This is used to get the address of a user based on a DTag
func NewKeeper(
	cdc codec.BinaryMarshaler, storeKey sdk.StoreKey,
	paramSpace paramstypes.Subspace, relKeeper relationshipskeeper.Keeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		paramSubspace: paramSpace,
		relKeeper:     relKeeper,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// IsUserBlocked tells if the given blocker has blocked the given blocked user
func (k Keeper) IsUserBlocked(ctx sdk.Context, blocker, blocked string) bool {
	return k.relKeeper.HasUserBlocked(ctx, blocker, blocked, "")
}

// associateDtagWithAddress save the relation of dtag and address on chain
func (k Keeper) associateDtagWithAddress(ctx sdk.Context, dtag string, address string) {
	store := ctx.KVStore(k.storeKey)
	wrapped := WrappedDTagOwner{Address: address}
	store.Set(types.DtagStoreKey(dtag), k.cdc.MustMarshalBinaryBare(&wrapped))
}

// GetDTagRelatedAddress returns the address associated to the given dtag or an empty string if it does not exists
func (k Keeper) GetDTagRelatedAddress(ctx sdk.Context, dtag string) (addr string) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.DtagStoreKey(dtag))
	if bz == nil {
		return ""
	}

	var owner WrappedDTagOwner
	k.cdc.MustUnmarshalBinaryBare(bz, &owner)
	return owner.Address
}

// GetDtagFromAddress returns the dtag associated with the given address or an empty string if no dtag exists
func (k Keeper) GetDtagFromAddress(ctx sdk.Context, addr string) (dtag string) {
	profile, found := k.GetProfile(ctx, addr)
	if !found {
		return ""
	}

	return profile.Dtag
}

// deleteDtagAddressAssociation delete the given dtag association with an address
func (k Keeper) deleteDtagAddressAssociation(ctx sdk.Context, dtag string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.DtagStoreKey(dtag))
}

// replaceDtag delete the oldDtag related to the creator address and associate the new one to it
func (k Keeper) replaceDtag(ctx sdk.Context, oldDtag, newDtag string, creator string) {
	k.deleteDtagAddressAssociation(ctx, oldDtag)
	k.associateDtagWithAddress(ctx, newDtag, creator)
}

// StoreProfile stores the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same dtag from a different creator already exists
func (k Keeper) StoreProfile(ctx sdk.Context, profile types.Profile) error {

	addr := k.GetDTagRelatedAddress(ctx, profile.Dtag)
	if addr != "" && addr != profile.Creator {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"a profile with dtag %s has already been created", profile.Dtag)
	}

	oldDtag := k.GetDtagFromAddress(ctx, profile.Creator)
	k.replaceDtag(ctx, oldDtag, profile.Dtag, profile.Creator)

	store := ctx.KVStore(k.storeKey)
	key := types.ProfileStoreKey(profile.Creator)
	store.Set(key, k.cdc.MustMarshalBinaryBare(&profile))

	return nil
}

// GetProfile returns the profile corresponding to the given address inside the current context.
func (k Keeper) GetProfile(ctx sdk.Context, address string) (profile types.Profile, found bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.ProfileStoreKey(address))
	if bz != nil {
		k.cdc.MustUnmarshalBinaryBare(bz, &profile)
		return profile, true
	}

	return types.Profile{}, false
}

// RemoveProfile allows to delete a profile associated with the given address inside the current context.
// It assumes that the address-related profile exists.
func (k Keeper) RemoveProfile(ctx sdk.Context, address string) error {
	profile, found := k.GetProfile(ctx, address)
	if !found {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"no profile associated with the following address found: %s", address)
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.ProfileStoreKey(address))
	k.deleteDtagAddressAssociation(ctx, profile.Dtag)
	return nil
}

// GetProfiles returns all the created profiles inside the current context.
func (k Keeper) GetProfiles(ctx sdk.Context) []types.Profile {
	var profiles []types.Profile

	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProfileStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var profile types.Profile
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &profile)
		profiles = append(profiles, profile)
	}

	return profiles
}

// ValidateProfile checks if the given profile is valid according to the current profile's module params
func (k Keeper) ValidateProfile(ctx sdk.Context, profile types.Profile) error {
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
