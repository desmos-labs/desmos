package types

import (
	"fmt"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	// DefaultMaxTextLength represents the default max length for post texts
	DefaultMaxTextLength uint32 = 500
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
func NewParams(maxTextLength uint32) Params {
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
	_, isUint32 := i.(uint32)
	if !isUint32 {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	return nil
}
