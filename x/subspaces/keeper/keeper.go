package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
}

// NewKeeper creates new instances of the subspaces keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryCodec) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
