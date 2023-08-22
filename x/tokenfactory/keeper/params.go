package keeper

import (
	"github.com/desmos-labs/desmos/v6/x/tokenfactory/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetParams sets the params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&params)
	store.Set([]byte(types.ParamsPrefixKey), bz)
}

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(types.ParamsPrefixKey))
	if bz == nil {
		return p
	}
	k.cdc.MustUnmarshal(bz, &p)
	return p
}
