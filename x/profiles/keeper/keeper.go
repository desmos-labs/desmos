package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/subspace"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	// The reference to the ParamsStore to get and set profile specific params
	paramSubspace params.Subspace

	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, paramSpace params.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		paramSubspace: paramSpace,
		StoreKey:      storeKey,
		Cdc:           cdc,
	}
}

// AssociateDtagWithAddress save the relation of dtag and address on chain
func (k Keeper) AssociateDtagWithAddress(ctx sdk.Context, dtag string, address sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	key := types.DtagStoreKey(dtag)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&address))
}

// GetDtagRelatedAddress returns the address associated to the given dtag or nil if it not exists
func (k Keeper) GetDtagRelatedAddress(ctx sdk.Context, dtag string) (addr sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	bz := store.Get(types.DtagStoreKey(dtag))
	if bz == nil {
		return nil
	}
	k.Cdc.MustUnmarshalBinaryBare(bz, &addr)
	return addr
}

// GetDtagFromAddress returns the dtag associated with the given address or an empty string if no dtag exists
func (k Keeper) GetDtagFromAddress(ctx sdk.Context, addr sdk.AccAddress) (dtag string) {
	store := ctx.KVStore(k.StoreKey)
	it := sdk.KVStorePrefixIterator(store, types.DtagStorePrefix)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var acc sdk.AccAddress
		k.Cdc.MustUnmarshalBinaryBare(it.Value(), &acc)
		if acc.Equals(addr) {
			return string(bytes.TrimPrefix(it.Key(), types.DtagStorePrefix))
		}
	}

	return ""
}

// DeleteDtagAddressAssociation delete the given dtag association with an address
func (k Keeper) DeleteDtagAddressAssociation(ctx sdk.Context, dtag string) {
	store := ctx.KVStore(k.StoreKey)
	store.Delete(types.DtagStoreKey(dtag))
}

// replaceDtag delete the oldDtag related to the creator address and associate the new one to it
func (k Keeper) replaceDtag(ctx sdk.Context, oldDtag, newDtag string, creator sdk.AccAddress) {
	k.DeleteDtagAddressAssociation(ctx, oldDtag)
	k.AssociateDtagWithAddress(ctx, newDtag, creator)
}

// SaveProfile allows to save the given profile inside the current context.
// It assumes that the given profile has already been validated.
// It returns an error if a profile with the same dtag from a different creator already exists
func (k Keeper) SaveProfile(ctx sdk.Context, profile types.Profile) error {

	if addr := k.GetDtagRelatedAddress(ctx, profile.DTag); addr != nil && !addr.Equals(profile.Creator) {
		return fmt.Errorf("a profile with dtag: %s has already been created", profile.DTag)
	}

	oldDtag := k.GetDtagFromAddress(ctx, profile.Creator)
	k.replaceDtag(ctx, oldDtag, profile.DTag, profile.Creator)

	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(profile.Creator)

	store.Set(key, k.Cdc.MustMarshalBinaryBare(&profile))

	return nil
}

// DeleteProfile allows to delete a profile associated with the given address inside the current context.
// It assumes that the address-related profile exists.
// nolint: interfacer
func (k Keeper) DeleteProfile(ctx sdk.Context, address sdk.AccAddress, dtag string) {
	store := ctx.KVStore(k.StoreKey)
	store.Delete(types.ProfileStoreKey(address))
	k.DeleteDtagAddressAssociation(ctx, dtag)
}

// GetProfiles returns all the created profiles inside the current context.
func (k Keeper) GetProfiles(ctx sdk.Context) (profiles types.Profiles) {
	profiles = make(types.Profiles, 0)
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProfileStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var acc types.Profile
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &acc)
		profiles = append(profiles, acc)
	}

	return profiles
}

// GetProfile returns the profile corresponding to the given address inside the current context.
// nolint: interfacer
func (k Keeper) GetProfile(ctx sdk.Context, address sdk.AccAddress) (profile types.Profile, found bool) {
	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(address)
	if bz := store.Get(key); bz != nil {
		k.Cdc.MustUnmarshalBinaryBare(bz, &profile)
		return profile, true
	}

	return types.Profile{}, false
}
