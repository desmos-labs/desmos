package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

// Default profile paramsModule
var (
	DefaultRegEx         = `^[A-Za-z0-9_]+$`
	DefaultMinNameLength = sdk.NewInt(3)
	DefaultMaxNameLength = sdk.NewInt(10)
)

// Parameters store keys
var (
	NameLenParamsKey = []byte("NameParams")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// NewParams is the constructor for Params
func NewParams(nameParams NameParams) Params {
	return Params{
		NameParams: nameParams,
	}
}

// DefaultParams returns default paramsModule
func DefaultParams() Params {
	return Params{
		NameParams: DefaultNameParams(),
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of subspaces module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(NameLenParamsKey, &params.NameParams, ValidateNameParams),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	return ValidateNameParams(params.NameParams)
}

// NewNameParams is a constructor for NameParams
func NewNameParams(regEx string, minLen, maxLen sdk.Int) NameParams {
	return NameParams{
		RegEx:     regEx,
		MinLength: minLen,
		MaxLength: maxLen,
	}
}

// DefaultNameParams return default NameParams
func DefaultNameParams() NameParams {
	return NewNameParams(DefaultRegEx, DefaultMinNameLength, DefaultMaxNameLength)
}

func ValidateNameParams(i interface{}) error {
	params, isDtagParams := i.(NameParams)
	if !isDtagParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(strings.TrimSpace(params.RegEx)) == 0 {
		return fmt.Errorf("empty name regEx param")
	}

	if params.MinLength.IsNegative() || params.MinLength.LT(DefaultMinNameLength) {
		return fmt.Errorf("invalid minimum name length param: %s", params.MinLength)
	}

	if params.MaxLength.IsNegative() {
		return fmt.Errorf("invalid max name length param: %s", params.MaxLength)
	}

	return nil
}
