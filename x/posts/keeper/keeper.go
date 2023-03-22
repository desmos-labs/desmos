package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey       storetypes.StoreKey
	cdc            codec.BinaryCodec
	paramsSubspace paramstypes.Subspace
	hooks          types.PostsHooks

	ak types.ProfilesKeeper
	sk types.SubspacesKeeper
	rk types.RelationshipsKeeper
}

// NewKeeper creates a new instance of the Posts Keeper.
func NewKeeper(
	cdc codec.BinaryCodec, storeKey storetypes.StoreKey, paramsSubspace paramstypes.Subspace,
	ak types.ProfilesKeeper, sk types.SubspacesKeeper, rk types.RelationshipsKeeper,
) Keeper {
	if !paramsSubspace.HasKeyTable() {
		paramsSubspace = paramsSubspace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:       storeKey,
		cdc:            cdc,
		paramsSubspace: paramsSubspace,

		ak: ak,
		sk: sk,
		rk: rk,
	}
}

// Logger returns a module-specific logger
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetHooks allows to set the posts hooks
func (k *Keeper) SetHooks(sh types.PostsHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set posts hooks twice")
	}

	k.hooks = sh
	return k
}
