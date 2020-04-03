package keeper

import (
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
	store.Set([]byte(moniker), k.Cdc.MustMarshalBinaryBare(&address))
}

// GetMonikerRelatedAddress returns the address associated to the given moniker or nil if it not exists
func (k Keeper) GetMonikerRelatedAddress(ctx sdk.Context, moniker string) (addr sdk.AccAddress) {
	store := ctx.KVStore(k.StoreKey)

	bz := store.Get([]byte(moniker))
	if bz == nil {
		return nil
	}

	k.Cdc.MustUnmarshalBinaryBare(bz, &addr)

	return addr
}

// DeleteMonikerAddressAssociation delete the given moniker association with an address
func (k Keeper) DeleteMonikerAddressAssociation(ctx sdk.Context, moniker string) {
	store := ctx.KVStore(k.StoreKey)
	store.Delete([]byte(moniker))
}

// SaveProfile allows to save the given account inside the current context.
// It assumes that the given account has already been validated.
// It returns an error if an account with the same moniker from a different creator already exists
func (k Keeper) SaveProfile(ctx sdk.Context, acc types.Profile) error {
	store := ctx.KVStore(k.StoreKey)

	key := types.ProfileStoreKey(acc.Creator.String())

	if addr := k.GetMonikerRelatedAddress(ctx, acc.Moniker); addr != nil && !addr.Equals(acc.Creator) {
		return fmt.Errorf("an account with moniker: %s has already been created", acc.Moniker)
	}

	k.AssociateMonikerWithAddress(ctx, acc.Moniker, acc.Creator)
	store.Set(key, k.Cdc.MustMarshalBinaryBare(&acc))

	return nil
}

// DeleteProfile allows to delete an account associated with the given address inside the current context.
// It assumes that the address-related account exists.
func (k Keeper) DeleteProfile(ctx sdk.Context, address string, moniker string) {
	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(address)
	store.Delete(key)
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
func (k Keeper) GetProfile(ctx sdk.Context, address string) (account types.Profile, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := types.ProfileStoreKey(address)

	if bz := store.Get(key); bz != nil {
		k.Cdc.MustUnmarshalBinaryBare(bz, &account)
		return account, true
	}

	return types.Profile{}, false
}
