package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramsModule "github.com/cosmos/cosmos-sdk/x/params/subspace"
)

const (
	// default paramspace for paramsModule keeper
	DefaultParamspace = ModuleName
)

var (
	DefaultFeeDenom = sdk.DefaultBondDenom
	DefaultMinFees  []MinFee
)

// Parameters store keys
var (
	FeeDenomStoreKey = []byte("FeeDenom")
	MinFeesStoreKey  = []byte("MinFees")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	FeeDenom string   `json:"fee_denom" yaml:"fee_denom"`
	MinFees  []MinFee `json:"min_fees" yaml:"min_fees"`
}

// NewParams create a new params object with the given data
func NewParams(feeDenom string, minFees []MinFee) Params {
	return Params{
		FeeDenom: feeDenom,
		MinFees:  minFees,
	}
}

// DefaultParams return default params object
func DefaultParams() Params {
	return Params{
		FeeDenom: DefaultFeeDenom,
		MinFees:  DefaultMinFees,
	}
}

// String implements Stringer
func (params Params) String() string {
	out := "Fee parameters:\n"
	out += fmt.Sprintf("FeeDenom: %s\nMinFees: %s\n", params.FeeDenom, params.MinFees)
	return strings.TrimSpace(out)
}

func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(FeeDenomStoreKey, &params.FeeDenom, ValidateFeeDenomParam),
		paramsModule.NewParamSetPair(MinFeesStoreKey, &params.MinFees, ValidateMinFeesParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateFeeDenomParam(params.FeeDenom); err != nil {
		return err
	}

	if err := ValidateMinFeesParam(params.MinFees); err != nil {
		return err
	}

	return nil
}

func ValidateFeeDenomParam(i interface{}) error {
	param, isCorrectParam := i.(string)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameter type: %s", i)
	}

	if len(strings.TrimSpace(param)) == 0 {
		return fmt.Errorf("invalid fee denom param, it shouldn't be empty")
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
