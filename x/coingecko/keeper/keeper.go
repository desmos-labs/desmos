package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/coingecko/types"
	"github.com/tendermint/tendermint/libs/log"
)

type Keeper struct {
	cdc codec.BinaryCodec
}

// NewKeeper creates new instances of the supply keeper
func NewKeeper(cdc codec.BinaryCodec) Keeper {
	return Keeper{
		cdc: cdc,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}
