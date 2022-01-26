package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
	pk       types.ProfileKeeper
	sk       types.SubspacesKeeper
}

// NewKeeper creates new instances of the relationships Keeper.
func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, pk types.ProfileKeeper, sk types.SubspacesKeeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		pk:       pk,
		sk:       sk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// HasProfile tells whether the given user has a profile or not
func (k Keeper) HasProfile(ctx sdk.Context, user string) bool {
	return k.pk.HasProfile(ctx, user)
}

// DoesSubspaceExist tells if the subspace with the given id exists
func (k Keeper) DoesSubspaceExist(ctx sdk.Context, subspaceID uint64) bool {
	return k.sk.HasSubspace(ctx, subspaceID)
}
