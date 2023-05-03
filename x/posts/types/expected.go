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
	SetParamSet(ctx sdk.Context, ps paramstypes.ParamSet)
	GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet)
}

// PostsHooksWrapper is a wrapper for modules to inject StakingHooks using depinject.
type PostsHooksWrapper struct{ Hooks PostsHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (PostsHooksWrapper) IsOnePerModuleType() {}
