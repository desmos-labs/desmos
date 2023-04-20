package types

import (
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	// ReasonsKey represents the params key used to store available default reasons
	ReasonsKey = []byte("StandardReasons")
)

// ParamKeyTable returns the key declaration for the parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of reports module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(ReasonsKey, &params.StandardReasons, ValidateStandardReasonsParam),
	}
}
