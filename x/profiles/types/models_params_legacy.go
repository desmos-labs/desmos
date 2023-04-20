package types

import paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

// Parameters store keys
var (
	NicknameParamsKey = []byte("NicknameParams")
	DTagParamsKey     = []byte("DTagParams")
	BioParamsKey      = []byte("MaxBioLen")
	OracleParamsKey   = []byte("OracleParams")
	AppLinksParamsKey = []byte("AppLinksParams")
)

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(NicknameParamsKey, &params.Nickname, ValidateNicknameParams),
		paramstypes.NewParamSetPair(DTagParamsKey, &params.DTag, ValidateDTagParams),
		paramstypes.NewParamSetPair(BioParamsKey, &params.Bio, ValidateBioParams),
		paramstypes.NewParamSetPair(OracleParamsKey, &params.Oracle, ValidateOracleParams),
		paramstypes.NewParamSetPair(AppLinksParamsKey, &params.AppLinks, ValidateAppLinksParams),
	}
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}
