package keeper

import (
	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	hooks    types.ReactionsHooks

	ak types.ProfilesKeeper
	sk types.SubspacesKeeper
	rk types.RelationshipsKeeper
	pk types.PostsKeeper
}

// NewKeeper creates a new instance of the reactions Keeper.
func NewKeeper(
	cdc codec.BinaryCodec, storeKey storetypes.StoreKey,
	ak types.ProfilesKeeper, sk types.SubspacesKeeper, rk types.RelationshipsKeeper, pk types.PostsKeeper,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,

		ak: ak,
		sk: sk,
		rk: rk,
		pk: pk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetHooks allows to set the reactions hooks
func (k *Keeper) SetHooks(rs types.ReactionsHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set reactions hooks twice")
	}

	k.hooks = rs
	return k
}
