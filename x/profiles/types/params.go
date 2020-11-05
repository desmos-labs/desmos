package types

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// default paramspace for paramsModule keeper
	DefaultParamspace = ModuleName
)

// Default profile paramsModule
var (
	DefaultMinMonikerLength = sdk.NewInt(2)
	DefaultMaxMonikerLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultRegEx            = `^[A-Za-z0-9_]+$`
	DefaultMinDTagLength    = sdk.NewInt(3)
	DefaultMaxDTagLength    = sdk.NewInt(30)
	DefaultMaxBioLength     = sdk.NewInt(1000)
)

// Parameters store keys
var (
	MonikerLenParamsKey = []byte("MonikerParams")
	DtagLenParamsKey    = []byte("DtagParams")
	MaxBioLenParamsKey  = []byte("MaxBioLen")
)

// ___________________________________________________________________________________________________________________

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new ProfileParams obj
func NewParams(monikerLen MonikerParams, dtagLen DTagParams, maxBioLen sdk.Int) Params {
	return Params{
		MonikerParams: monikerLen,
		DtagParams:    dtagLen,
		MaxBioLength:  maxBioLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		MonikerParams: DefaultMonikerParams(),
		DtagParams:    DefaultDtagParams(),
		MaxBioLength:  DefaultMaxBioLength,
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MonikerLenParamsKey, &params.MonikerParams, ValidateMonikerParams),
		paramstypes.NewParamSetPair(DtagLenParamsKey, &params.DtagParams, ValidateDtagParams),
		paramstypes.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLength, ValidateBioParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMonikerParams(params.MonikerParams); err != nil {
		return err
	}

	if err := ValidateDtagParams(params.DtagParams); err != nil {
		return err
	}

	return ValidateBioParams(params.MaxBioLength)
}

// ___________________________________________________________________________________________________________________

// NewMonikerParams creates a new MonikerParams obj
func NewMonikerParams(minLen, maxLen sdk.Int) MonikerParams {
	return MonikerParams{
		MinMonikerLength: minLen,
		MaxMonikerLength: maxLen,
	}
}

// DefaultMonikerParams return default moniker params
func DefaultMonikerParams() MonikerParams {
	return NewMonikerParams(DefaultMinMonikerLength,
		DefaultMaxMonikerLength)
}

func ValidateMonikerParams(i interface{}) error {
	params, isNameSurnParams := i.(MonikerParams)

	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.MinMonikerLength.IsNegative() || params.MinMonikerLength.LT(DefaultMinMonikerLength) {
		return fmt.Errorf("invalid minimum moniker length param: %s", params.MinMonikerLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	if params.MaxMonikerLength.IsNegative() || params.MaxMonikerLength.GT(DefaultMaxMonikerLength) {
		return fmt.Errorf("invalid max moniker length param: %s", params.MaxMonikerLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewDtagParams creates a new DtagParams obj
func NewDtagParams(regEx string, minLen, maxLen sdk.Int) DTagParams {
	return DTagParams{
		RegEx:         regEx,
		MinDtagLength: minLen,
		MaxDtagLength: maxLen,
	}
}

// DefaultDtagParams return default paramsModule
func DefaultDtagParams() DTagParams {
	return NewDtagParams(
		DefaultRegEx,
		DefaultMinDTagLength,
		DefaultMaxDTagLength,
	)
}

func ValidateDtagParams(i interface{}) error {
	params, isMonikerParams := i.(DTagParams)
	if !isMonikerParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty dTag regEx param")
	}

	if params.MinDtagLength.IsNegative() || params.MinDtagLength.LT(DefaultMinDTagLength) {
		return fmt.Errorf("invalid minimum dTag length param: %s", params.MinDtagLength)
	}

	if params.MaxDtagLength.IsNegative() {
		return fmt.Errorf("invalid max dTag length param: %s", params.MaxDtagLength)
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
