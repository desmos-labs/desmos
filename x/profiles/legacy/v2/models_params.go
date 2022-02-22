package v2

import (
	"fmt"
	"strings"

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
		paramstypes.NewParamSetPair(types.NicknameParamsKey, &params.Nickname, ValidateNicknameParams),
		paramstypes.NewParamSetPair(types.DTagParamsKey, &params.DTag, ValidateDTagParams),
		paramstypes.NewParamSetPair(types.BioParamsKey, &params.Bio, ValidateBioParams),
		paramstypes.NewParamSetPair(types.OracleParamsKey, &params.Oracle, ValidateOracleParams),
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

func ValidateNicknameParams(i interface{}) error {
	params, areNicknameParams := i.(NicknameParams)
	if !areNicknameParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	minLength := params.MinLength
	if minLength.IsNil() || minLength.LT(types.DefaultMinNicknameLength) {
		return fmt.Errorf("invalid minimum nickname length param: %s", minLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	maxLength := params.MaxLength
	if maxLength.IsNil() || maxLength.IsNegative() || maxLength.GT(types.DefaultMaxNicknameLength) {
		return fmt.Errorf("invalid max nickname length param: %s", maxLength)
	}

	return nil
}

func ValidateBioParams(i interface{}) error {
	bioParams, isBioParams := i.(BioParams)
	if !isBioParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if bioParams.MaxLength.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", bioParams.MaxLength)
	}

	return nil
}

func ValidateDTagParams(i interface{}) error {
	params, isDtagParams := i.(DTagParams)
	if !isDtagParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty dTag regEx param")
	}

	if params.MinLength.IsNegative() || params.MinLength.LT(types.DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinLength)
	}

	if params.MaxLength.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxLength)
	}

	return nil
}

func ValidateOracleParams(i interface{}) error {
	params, isOracleParams := i.(OracleParams)
	if !isOracleParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.AskCount < params.MinCount {
		return fmt.Errorf("invalid ask count: %d, min count: %d", params.AskCount, params.MinCount)
	}

	if params.MinCount <= 0 {
		return fmt.Errorf("invalid min count: %d", params.MinCount)
	}

	if params.PrepareGas <= 0 {
		return fmt.Errorf("invalid prepare gas: %d", params.PrepareGas)
	}

	if params.ExecuteGas <= 0 {
		return fmt.Errorf("invalid execute gas: %d", params.ExecuteGas)
	}

	err := params.FeeAmount.Validate()
	if err != nil {
		return err
	}

	return nil
}
