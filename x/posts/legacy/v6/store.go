package v6

import (
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

// MigrateStore migrates the x/posts module state from the consensus version 5 to version 6.
// Specifically, it takes the parameters that are currently stored
// and managed by the x/params module and stores them directly into the x/posts
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
