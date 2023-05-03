package keeper

import (
	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey storetypes.StoreKey
	cdc      codec.BinaryCodec
	sk       types.SubspacesKeeper
}

// NewKeeper creates new instances of the relationships Keeper.
func NewKeeper(cdc codec.BinaryCodec, storeKey storetypes.StoreKey, sk types.SubspacesKeeper) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
		sk:       sk,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

// DoesSubspaceExist tells if the subspace with the given id exists
func (k Keeper) DoesSubspaceExist(ctx sdk.Context, subspaceID uint64) bool {
	// The subspace with id 0 always exists (as it represents all the subspaces)
	return subspaceID == 0 || k.sk.HasSubspace(ctx, subspaceID)
}
