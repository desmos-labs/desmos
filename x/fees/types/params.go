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
	DefaultFeeDenom    = "udaric"
	DefaultRequiredFee = sdk.NewDecWithPrec(1, 2)
)

// Parameters store keys
var (
	FeeDenomStoreKey    = []byte("FeeDenom")
	RequiredFeeStoreKey = []byte("RequiredFee")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	FeeDenom    string  `json:"fee_denom" yaml:"fee_denom"`
	RequiredFee sdk.Dec `json:"required_fee" yaml:"required_fee"`
}

// NewParams create a new params object with the given data
func NewParams(feeDenom string, requiredFee sdk.Dec) Params {
	return Params{
		FeeDenom:    feeDenom,
		RequiredFee: requiredFee,
	}
}

// DefaultParams return default params object
func DefaultParams() Params {
	return Params{
		FeeDenom:    DefaultFeeDenom,
		RequiredFee: DefaultRequiredFee,
	}
}

// String implements Stringer
func (params Params) String() string {
	out := "Fee parameters:\n"
	out += fmt.Sprintf("FeeDenom: %s\nRequiredFee: %d\n", params.FeeDenom, params.RequiredFee)
	return strings.TrimSpace(out)
}

func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(FeeDenomStoreKey, &params.FeeDenom, ValidateFeeDenomParam),
		paramsModule.NewParamSetPair(RequiredFeeStoreKey, &params.RequiredFee, ValidateRequiredFeeParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateFeeDenomParam(params.FeeDenom); err != nil {
		return err
	}

	if err := ValidateRequiredFeeParam(params.RequiredFee); err != nil {
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

func ValidateRequiredFeeParam(i interface{}) error {
	param, isCorrectParam := i.(sdk.Dec)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameter type: %s", i)
	}

	if param.IsNegative() {
		return fmt.Errorf("invalid minimum fee required, it shouldn't be a negative number")
	}

	return nil
}
