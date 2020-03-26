package keeper

import (
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
// It assumes that the given account wasn't already been inserted.
func (k Keeper) SaveAccount(ctx sdk.Context, acc types.Account) {
	store := ctx.KVStore(k.StoreKey)

	store.Set(types.AccountStoreKey(acc.Moniker), k.Cdc.MustMarshalBinaryBare(&acc))

}
