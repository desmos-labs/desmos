package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	ak       AccountKeeper
	bk       BankKeeper
	cdc      codec.BinaryCodec
	hooks    types.SubspacesHooks
}

// NewKeeper creates new instances of the subspaces keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak AccountKeeper, bk BankKeeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		ak:       ak,
		bk:       bk,
		cdc:      cdc,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetHooks allows to set the subspaces hooks
func (k *Keeper) SetHooks(sh types.SubspacesHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set subspaces hooks twice")
	}

	k.hooks = sh
	return k
}
