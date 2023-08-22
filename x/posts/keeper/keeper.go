package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	hooks    types.PostsHooks

	ak types.ProfilesKeeper
	sk types.SubspacesKeeper
	rk types.RelationshipsKeeper

	// the address capable of executing a MsgUpdateParams message. Typically, this
	// should be the x/gov module account.
	authority string
}

// NewKeeper creates a new instance of the Posts Keeper.
func NewKeeper(
	cdc codec.BinaryCodec, storeKey storetypes.StoreKey,
	ak types.ProfilesKeeper, sk types.SubspacesKeeper, rk types.RelationshipsKeeper, authority string,
) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,

		ak: ak,
		sk: sk,
		rk: rk,

		authority: authority,
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
