package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/desmos-labs/desmos/x/fees/types"
)

type Keeper struct {
	// The reference to the ParamsStore to get and set params
	paramSubspace params.Subspace
	Cdc           *codec.Codec
}

func NewKeeper(cdc *codec.Codec, paramSpace params.Subspace) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		Cdc:           cdc,
		paramSubspace: paramSpace,
	}
}

// SetParams sets params on the store
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSubspace.SetParamSet(ctx, &params)
}

// GetParams returns the params from the store
func (k Keeper) GetParams(ctx sdk.Context) (p types.Params) {
	k.paramSubspace.GetParamSet(ctx, &p)
	return p
}
