package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/v4/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(subspaces subspacestypes.Subspace) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, paramsSubspace paramstypes.Subspace, sk SubspacesKeeper) error {
	store := ctx.KVStore(storeKey)

	// Set the next reason id and report id for all the subspaces
	sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		store.Set(types.NextReasonIDStoreKey(subspace.ID), types.GetReasonIDBytes(1))
		store.Set(types.NextReportIDStoreKey(subspace.ID), types.GetReportIDBytes(1))
		return false
	})

	// Set the module params
	params := types.DefaultParams()
	paramsSubspace.SetParamSet(ctx, &params)

	return nil
}
