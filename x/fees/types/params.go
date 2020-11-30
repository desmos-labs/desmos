package types

import (
	"fmt"
	"strings"

	paramsModule "github.com/cosmos/cosmos-sdk/x/params/subspace"
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
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	MinFees []MinFee `json:"min_fees" yaml:"min_fees"`
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

// String implements Stringer
func (params Params) String() string {
	out := "Fee parameters:\n"
	out += fmt.Sprintf("MinFees: %s\n", params.MinFees)
	return strings.TrimSpace(out)
}

func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(MinFeesStoreKey, &params.MinFees, ValidateMinFeesParam),
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
