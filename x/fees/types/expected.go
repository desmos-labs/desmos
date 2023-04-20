package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// DONTCOVER

// ParamsSubspace defines an interface that implements the legacy x/params ParamsSubspace
// type.
//
// NOTE: This is used solely for migration of x/params managed parameters.
type ParamsSubspace interface {
	GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet)
}
