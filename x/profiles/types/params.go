package types

import (
	"fmt"
	"strings"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

// Default profile paramsModule
var (
	DefaultMinUsernameLength = sdk.NewInt(2)
	DefaultMaxUsernameLength = sdk.NewInt(1000) //longest name on earth count 954 chars
	DefaultRegEx             = `^[A-Za-z0-9_]+$`
	DefaultMinDTagLength     = sdk.NewInt(3)
	DefaultMaxDTagLength     = sdk.NewInt(30)
	DefaultMaxBioLength      = sdk.NewInt(1000)
)

// Parameters store keys
var (
	UsernameLenParamsKey = []byte("UsernameParams")
	DTagLenParamsKey     = []byte("DTagParams")
	MaxBioLenParamsKey   = []byte("MaxBioLen")
)

// ___________________________________________________________________________________________________________________

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new ProfileParams obj
func NewParams(usernameParams UsernameParams, dTagParams DTagParams, maxBioLen sdk.Int) Params {
	return Params{
		UsernameParams: usernameParams,
		DTagParams:     dTagParams,
		MaxBioLength:   maxBioLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		UsernameParams: DefaultUsernameParams(),
		DTagParams:     DefaultDTagParams(),
		MaxBioLength:   DefaultMaxBioLength,
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of profile module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(UsernameLenParamsKey, &params.UsernameParams, ValidateUsernameParams),
		paramstypes.NewParamSetPair(DTagLenParamsKey, &params.DTagParams, ValidateDTagParams),
		paramstypes.NewParamSetPair(MaxBioLenParamsKey, &params.MaxBioLength, ValidateBioParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateUsernameParams(params.UsernameParams); err != nil {
		return err
	}

	if err := ValidateDTagParams(params.DTagParams); err != nil {
		return err
	}

	return ValidateBioParams(params.MaxBioLength)
}

// ___________________________________________________________________________________________________________________

// NewUsernameParams creates a new UsernameParams obj
func NewUsernameParams(minLen, maxLen sdk.Int) UsernameParams {
	return UsernameParams{
		MinUsernameLength: minLen,
		MaxUsernameLength: maxLen,
	}
}

// DefaultUsernameParams return default username params
func DefaultUsernameParams() UsernameParams {
	return NewUsernameParams(
		DefaultMinUsernameLength,
		DefaultMaxUsernameLength,
	)
}

func ValidateUsernameParams(i interface{}) error {
	params, isNameSurnParams := i.(UsernameParams)
	if !isNameSurnParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	minLength := params.MinUsernameLength
	if minLength.IsNil() || minLength.LT(DefaultMinUsernameLength) {
		return fmt.Errorf("invalid minimum username length param: %s", minLength)
	}

	// TODO make sense to cap this? I've done this thinking "what's the sense of having names higher that 1000 chars?"
	maxLength := params.MaxUsernameLength
	if maxLength.IsNil() || maxLength.IsNegative() || maxLength.GT(DefaultMaxUsernameLength) {
		return fmt.Errorf("invalid max username length param: %s", maxLength)
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewDTagParams creates a new DTagParams obj
func NewDTagParams(regEx string, minLen, maxLen sdk.Int) DTagParams {
	return DTagParams{
		RegEx:         regEx,
		MinDTagLength: minLen,
		MaxDTagLength: maxLen,
	}
}

// DefaultDTagParams return default paramsModule
func DefaultDTagParams() DTagParams {
	return NewDTagParams(
		DefaultRegEx,
		DefaultMinDTagLength,
		DefaultMaxDTagLength,
	)
}

func ValidateDTagParams(i interface{}) error {
	params, isUsernameParams := i.(DTagParams)
	if !isUsernameParams {
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
