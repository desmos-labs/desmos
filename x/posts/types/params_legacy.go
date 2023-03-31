package types

import paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

const (
	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	// MaxTextLengthKey represents the key used to store the max length for posts texts
	MaxTextLengthKey = []byte("MaxTextLength")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().
		RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of posts module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MaxTextLengthKey, &params.MaxTextLength, ValidateMaxTextLength),
	}
}
