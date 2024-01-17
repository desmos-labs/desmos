package v2

import (
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

type SubspacesKeeper interface {
	IterateSubspaces(ctx sdk.Context, fn func(subspaces subspacestypes.Subspace) (stop bool))
}

// MigrateStore performs in-place store migrations from v1 to v2
// The things done here are the following:
// 1. setting up the next reason id and report id keys for existing subspaces
// 2. setting up the module params
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, paramsSubspace types.ParamsSubspace, sk SubspacesKeeper) error {
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
