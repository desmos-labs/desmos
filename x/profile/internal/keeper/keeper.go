package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// AssociateMonikerWithAddress save the relation of moniker and address on chain
func (k Keeper) AssociateMonikerWithAddress(ctx sdk.Context, moniker string, address sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	key := types.MonikerStoreKey(moniker)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&address))
}

// GetMonikerRelatedAddress returns the address associated to the given moniker or nil if it not exists
func (k Keeper) GetMonikerRelatedAddress(ctx sdk.Context, moniker string) (addr sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)
	bz := store.Get(types.MonikerStoreKey(moniker))
	if bz == nil {
		return nil
	}
	k.Cdc.MustUnmarshalBinaryBare(bz, &addr)
	return addr
}

// GetMonikerFromAddress returns the moniker associated with the given address or an empty string if no moniker exists
func (k Keeper) GetMonikerFromAddress(ctx sdk.Context, addr sdk.AccAddress) (moniker string) {
	store := ctx.KVStore(k.StoreKey)
	it := sdk.KVStorePrefixIterator(store, types.MonikerStorePrefix)
	defer it.Close()

	for ; it.Valid(); it.Next() {
		var acc sdk.AccAddress
		k.Cdc.MustUnmarshalBinaryBare(it.Value(), &acc)
		if acc.Equals(addr) {
			return string(bytes.TrimPrefix(it.Key(), types.MonikerStorePrefix))
		}
	}

	return ""
}

// DeleteMonikerAddressAssociation delete the given moniker association with an address
func (k Keeper) DeleteMonikerAddressAssociation(ctx sdk.Context, moniker string) {
	store := ctx.KVStore(k.StoreKey)
	store.Delete(types.MonikerStoreKey(moniker))
}

// replaceMoniker delete the oldMoniker related to the creator address and associate the new one to it
func (k Keeper) replaceMoniker(ctx sdk.Context, oldMoniker, newMoniker string, creator sdk.AccAddress) {
	k.DeleteMonikerAddressAssociation(ctx, oldMoniker)
	k.AssociateMonikerWithAddress(ctx, newMoniker, creator)
}

// SaveProfile allows to save the given account inside the current context.
// It assumes that the given account has already been validated.
// It returns an error if an account with the same moniker from a different creator already exists
func (k Keeper) SaveProfile(ctx sdk.Context, profile types.Profile) error {

	if addr := k.GetMonikerRelatedAddress(ctx, profile.Moniker); addr != nil && !addr.Equals(profile.Creator) {
		return fmt.Errorf("an account with moniker: %s has already been created", profile.Moniker)
	}

	oldMoniker := k.GetMonikerFromAddress(ctx, profile.Creator)
	k.replaceMoniker(ctx, oldMoniker, profile.Moniker, profile.Creator)

	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(profile.Creator)

	store.Set(key, k.Cdc.MustMarshalBinaryBare(&profile))

	return nil
}

// DeleteProfile allows to delete an account associated with the given address inside the current context.
// It assumes that the address-related account exists.
// nolint: interfacer
func (k Keeper) DeleteProfile(ctx sdk.Context, address sdk.AccAddress, moniker string) {
	store := ctx.KVStore(k.StoreKey)
	store.Delete(types.ProfileStoreKey(address))
	k.DeleteMonikerAddressAssociation(ctx, moniker)
}

// GetProfiles returns all the created accounts inside the current context.
func (k Keeper) GetProfiles(ctx sdk.Context) (accounts types.Profiles) {
	accounts = make(types.Profiles, 0)
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ProfileStorePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var acc types.Profile
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &acc)
		accounts = append(accounts, acc)
	}

	return accounts
}

// GetProfile returns the account corresponding to the given address inside the current context.
// nolint: interfacer
func (k Keeper) GetProfile(ctx sdk.Context, address sdk.AccAddress) (account types.Profile, found bool) {
	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(address)
	if bz := store.Get(key); bz != nil {
		k.Cdc.MustUnmarshalBinaryBare(bz, &account)
		return account, true
	}

	return types.Profile{}, false
}
