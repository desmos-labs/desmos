package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	hooks    types.SubspacesHooks

	ak     types.AccountKeeper
	authzk types.AuthzKeeper
}

// NewKeeper creates new instances of the subspaces keeper
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ak types.AccountKeeper, authzKeeper types.AuthzKeeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		ak:       ak,
		authzk:   authzKeeper,
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
