package v0160

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// DONTCOVER

// MigrateParams migrates the given params subspace to the new types.Params
// This is used because parameters are stored as JSON inside the chain, so when we change the Protobuf field names
// we should also migrate that to the new names
func MigrateParams(ctx sdk.Context, amino *codec.LegacyAmino, subspace paramstypes.Subspace) (types.Params, error) {
	var monikerParams MonikerParams
	bz := subspace.GetRaw(ctx, MonikerLenParamsKey)
	err := amino.UnmarshalJSON(bz, &monikerParams)
	if err != nil {
		return types.Params{}, err
	}

	var dTagParams DTagParams
	bz = subspace.GetRaw(ctx, DTagParamsKey)
	err = amino.UnmarshalJSON(bz, &dTagParams)
	if err != nil {
		return types.Params{}, err
	}

	var maxBioLen sdk.Int
	bz = subspace.GetRaw(ctx, MaxBioLenParamsKey)
	err = amino.UnmarshalJSON(bz, &maxBioLen)
	if err != nil {
		return types.Params{}, err
	}

	return types.NewParams(
		types.NewNicknameParams(
			monikerParams.MinMonikerLength,
			monikerParams.MaxMonikerLength,
		),
		types.NewDTagParams(
			dTagParams.RegEx,
			dTagParams.MinDTagLength,
			dTagParams.MaxDTagLength,
		),
		maxBioLen,
	), nil
}
