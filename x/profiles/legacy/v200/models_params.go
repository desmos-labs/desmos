package v200

import (
	"github.com/cosmos/cosmos-sdk/codec"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(types.NicknameParamsKey, &params.Nickname, types.ValidateNicknameParams),
		paramstypes.NewParamSetPair(types.DTagParamsKey, &params.DTag, types.ValidateDTagParams),
		paramstypes.NewParamSetPair(types.BioParamsKey, &params.Bio, types.ValidateBioParams),
		paramstypes.NewParamSetPair(types.OracleParamsKey, &params.Oracle, types.ValidateOracleParams),
	}
}

func DefaultParams() Params {
	nicknameParams := types.DefaultNicknameParams()
	dTagParams := types.DefaultDTagParams()
	bioParams := types.DefaultBioParams()
	oracleParams := types.DefaultOracleParams()

	return Params{
		Nickname: NicknameParams{
			MinLength: nicknameParams.MinLength,
			MaxLength: nicknameParams.MaxLength,
		},
		DTag: DTagParams{
			RegEx:     dTagParams.RegEx,
			MinLength: dTagParams.MinLength,
			MaxLength: dTagParams.MaxLength,
		},
		Bio: BioParams{
			MaxLength: bioParams.MaxLength,
		},
		Oracle: OracleParams{
			ScriptID:   oracleParams.ScriptID,
			AskCount:   oracleParams.AskCount,
			MinCount:   oracleParams.MinCount,
			PrepareGas: oracleParams.PrepareGas,
			ExecuteGas: oracleParams.ExecuteGas,
			FeeAmount:  oracleParams.FeeAmount,
		},
	}
}

// MustMarshalAppLinksParams serializes the given application links params using the provided BinaryCodec
func MustMarshalAppLinksParams(cdc codec.BinaryCodec, params Params) []byte {
	return cdc.MustMarshal(&params)
}
