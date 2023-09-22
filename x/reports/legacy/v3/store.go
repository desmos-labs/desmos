package v3

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/reports/types"
)

// MigrateStore migrates the x/reports module state from the consensus version 2 to version 3.
// Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/reports
// module state.
func MigrateStore(ctx sdk.Context, storeKey storetypes.StoreKey, legacySubspace types.ParamsSubspace, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)

	var params types.Params
	legacySubspace.GetParamSet(ctx, &params)

	if err := params.Validate(); err != nil {
		return err
	}

	bz := cdc.MustMarshal(&params)
	store.Set(types.ParamsKey, bz)

	return nil
}
