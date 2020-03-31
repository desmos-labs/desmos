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

// SaveAccount allows to save the given account inside the current context.
// It assumes that the given account has already been validated.
// It returns an error if an account with the same moniker from a different creator already exists
func (k Keeper) SaveAccount(ctx sdk.Context, acc types.Profile) error {
	store := ctx.KVStore(k.StoreKey)

	key := types.ProfileStoreKey(acc.Moniker)

	if store.Has(key) {
		bz := store.Get(key)
		var savedAcc types.Profile
		k.Cdc.MustUnmarshalBinaryBare(bz, &savedAcc)
		if !savedAcc.Creator.Equals(acc.Creator) {
			return fmt.Errorf("an account with moniker: %s has already been created", acc.Moniker)
		}
	}

	store.Set(key, k.Cdc.MustMarshalBinaryBare(&acc))
	return nil
}

// DeleteAccount allows to delete an account with the given moniker inside the current context.
// It assumes that the moniker-related account exists.
func (k Keeper) DeleteAccount(ctx sdk.Context, moniker string) {
	store := ctx.KVStore(k.StoreKey)
	key := types.ProfileStoreKey(moniker)
	store.Delete(key)
}

// GetAccounts returns all the created accounts inside the current context.
func (k Keeper) GetAccounts(ctx sdk.Context) (accounts types.Profiles) {
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

// GetAccount returns the account corresponding to the given moniker inside the current context.
func (k Keeper) GetAccount(ctx sdk.Context, moniker string) (account types.Profile, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := types.ProfileStoreKey(moniker)

	if bz := store.Get(key); bz != nil {
		k.Cdc.MustUnmarshalBinaryBare(bz, &account)
		return account, true
	}

	return types.Profile{}, false
}
