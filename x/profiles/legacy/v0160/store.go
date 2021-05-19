package v0160

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// MigrateParams migrates the given params subspace to the new types.Params
// This is used because parameters are stored as JSON inside the chain, so when we change the Protobuf field names
// we should also migrate that to the new names
func MigrateParams(subspace paramstypes.Subspace, ctx sdk.Context) types.Params {
	var params Params
	subspace.GetParamSet(ctx, &params)
	return types.NewParams(
		types.NewNicknameParams(
			params.MonikerParams.MinMonikerLength,
			params.MonikerParams.MaxMonikerLength,
		),
		types.NewDTagParams(
			params.DTagParams.RegEx,
			params.DTagParams.MinDTagLength,
			params.DTagParams.MaxDTagLength,
		),
		params.MaxBioLength,
	)
}
