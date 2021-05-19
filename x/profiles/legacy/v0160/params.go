package v0160

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Default profile paramsModule
var (
	DefaultMinMonikerLength = sdk.NewInt(2)
	DefaultMaxMonikerLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultMinDTagLength    = sdk.NewInt(3)
)

// Parameters store keys
var (
	MonikerLenParamsKey = []byte("MonikerParams")
	DTagLenParamsKey    = []byte("DTagParams")
	MaxBioLenParamsKey  = []byte("MaxBioLen")
)

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MonikerLenParamsKey, &params.MonikerParams, ValidateMonikerParams),
		paramstypes.NewParamSetPair(DTagLenParamsKey, &params.DTagParams, ValidateDTagParams),
		paramstypes.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLength, ValidateBioParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMonikerParams(params.MonikerParams); err != nil {
		return err
	}

	if err := ValidateDTagParams(params.DTagParams); err != nil {
		return err
	}

	return ValidateBioParams(params.MaxBioLength)
}

// ___________________________________________________________________________________________________________________

func ValidateMonikerParams(i interface{}) error {
	params, isNameSurnParams := i.(MonikerParams)
	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	minLength := params.MinMonikerLength
	if minLength.IsNil() || minLength.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", minLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	maxLength := params.MaxMonikerLength
	if maxLength.IsNil() || maxLength.IsNegative() || maxLength.GT(DefaultMaxMonikerLength) {
		return fmt.Errorf("invalid max moniker length param: %s", maxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

func ValidateDTagParams(i interface{}) error {
	params, isMonikerParams := i.(DTagParams)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty dTag regEx param")
	}

	if params.MinDTagLength.IsNegative() || params.MinDTagLength.LT(DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinDTagLength)
	}

	if params.MaxDTagLength.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxDTagLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

func ValidateBioParams(i interface{}) error {
	bioLen, isBioLenParams := i.(sdk.Int)
	if !isBioLenParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if bioLen.IsNegative() {
		return fmt.Errorf("invalid max bio length param: %s", bioLen)
	}

	return nil
}
