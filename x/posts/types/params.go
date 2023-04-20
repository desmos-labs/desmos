package types

import (
	"fmt"
)

var (
	// DefaultMaxTextLength represents the default max length for post texts
	DefaultMaxTextLength uint32 = 500
)

// -------------------------------------------------------------------------------------------------------------------

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
