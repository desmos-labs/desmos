package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
)

type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	hooks    types.ReportsHooks

	ak types.ProfilesKeeper
	sk types.SubspacesKeeper
	rk types.RelationshipsKeeper
	pk types.PostsKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new instance of the Posts Keeper.
func NewKeeper(
	cdc codec.BinaryCodec, storeKey storetypes.StoreKey, ak types.ProfilesKeeper,
	sk types.SubspacesKeeper, rk types.RelationshipsKeeper, pk types.PostsKeeper, authority string,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,

		ak: ak,
		sk: sk,
		rk: rk,
		pk: pk,

		authority: authority,
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
