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
	// DefaultReasons represents the default set of reasons that can be adopted by subspaces
	DefaultReasons = Reasons{}
)

var (
	// ReasonsKey represents the params key used to store available default reasons
	ReasonsKey = []byte("Reasons")
)

// --------------------------------------------------------------------------------------------------------------------

// ParamKeyTable returns the key declaration for the parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// NewParams returns a new Params instance
func NewParams(reasons Reasons) Params {
	return Params{
		Reasons: reasons,
	}
}

// DefaultParams returns the default params
func DefaultParams() Params {
	return Params{
		Reasons: DefaultReasons,
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of reports module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(ReasonsKey, &params.Reasons, ValidateReasonsParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	return ValidateReasonsParam(params.Reasons)
}

// ValidateReasonsParam validates the reasons params value
func ValidateReasonsParam(i interface{}) error {
	reasons, ok := i.([]Reason)
	if !ok {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	err := NewReasons(reasons...).Validate()
	if err != nil {
		return err
	}

	return nil
}
