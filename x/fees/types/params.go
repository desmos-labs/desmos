package types

import (
	"fmt"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// default paramspace for paramsModule keeper
	DefaultParamspace = ModuleName
)

var (
	DefaultMinFees []MinFee
)

// Parameters store keys
var (
	MinFeesStoreKey = []byte("MinFees")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams create a new params object with the given data
func NewParams(minFees []MinFee) Params {
	return Params{
		MinFees: minFees,
	}
}

// DefaultParams return default params object
func DefaultParams() Params {
	return Params{
		MinFees: DefaultMinFees,
	}
}

func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MinFeesStoreKey, &params.MinFees, ValidateMinFeesParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMinFeesParam(params.MinFees); err != nil {
		return err
	}

	return nil
}

func ValidateMinFeesParam(i interface{}) error {
	fees, isCorrectParam := i.([]MinFee)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameter type: %s", i)
	}

	for _, fee := range fees {
		if err := fee.Validate(); err != nil {
			return err
		}
	}

	return nil
}
