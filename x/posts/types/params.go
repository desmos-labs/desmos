package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	// DefaultMaxTextLength represents the default max length for post texts
	DefaultMaxTextLength = sdk.NewInt(500)
)

var (
	// MaxTextLengthKey represents the key used to store the max length for posts texts
	MaxTextLengthKey = []byte("MaxTextLength")
)

// -------------------------------------------------------------------------------------------------------------------

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// NewParams returns a new Params instance
func NewParams(maxTextLength sdk.Int) Params {
	return Params{
		MaxTextLength: maxTextLength,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		MaxTextLength: DefaultMaxTextLength,
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of posts module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MaxTextLengthKey, &params.MaxTextLength, ValidateMaxTextLength),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	return ValidateMaxTextLength(params.MaxTextLength)
}

// -------------------------------------------------------------------------------------------------------------------

func ValidateMaxTextLength(i interface{}) error {
	maxLength, isDtagParams := i.(sdk.Int)
	if !isDtagParams {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if maxLength.IsNegative() {
		return fmt.Errorf("invalid max text length param: %s", maxLength)
	}

	return nil
}
