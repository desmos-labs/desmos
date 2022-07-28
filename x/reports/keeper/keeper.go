package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/v4/x/reports/types"
)

type Keeper struct {
	storeKey       sdk.StoreKey
	cdc            codec.BinaryCodec
	paramsSubspace paramstypes.Subspace
	hooks          types.ReportsHooks

	ak types.ProfilesKeeper
	sk types.SubspacesKeeper
	rk types.RelationshipsKeeper
	pk types.PostsKeeper
}

// NewKeeper creates a new instance of the Posts Keeper.
func NewKeeper(
	cdc codec.BinaryCodec, storeKey sdk.StoreKey, paramsSubspace paramstypes.Subspace,
	ak types.ProfilesKeeper, sk types.SubspacesKeeper, rk types.RelationshipsKeeper, pk types.PostsKeeper,
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
		pk: pk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// SetHooks allows to set the reports hooks
func (k *Keeper) SetHooks(rs types.ReportsHooks) *Keeper {
	if k.hooks != nil {
		panic("cannot set reports hooks twice")
	}

	k.hooks = rs
	return k
}
