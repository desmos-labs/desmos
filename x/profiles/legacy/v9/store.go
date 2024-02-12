package v9

import (
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/profiles/types"
)

// MigrateStore migrates the x/profiles module state from the consensus version 8 to version 9.
// Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/profiles
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
