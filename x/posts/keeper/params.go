package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types2.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types2.Params) {
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}
