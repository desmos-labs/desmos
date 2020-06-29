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

// Default profile paramsModule
var (
	DefaultMaxPostMessageLength            = sdk.NewInt(500)
	DefaultMaxOptionalDataFieldsNumber     = sdk.NewInt(10)
	DefaultMaxOptionalDataFieldValueLength = sdk.NewInt(200)
)

// Parameters store keys
var (
	MaxPostMessageLengthKey            = []byte("MaxPostMessageLength")
	MaxOptionalDataFieldsNumberKey     = []byte("MaxOptionalDataFieldsNumber")
	MaxOptionalDataFieldValueLengthKey = []byte("MaxOptionalDataFieldValueLength")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramsModule.KeyTable {
	return paramsModule.NewKeyTable().RegisterParamSet(&Params{})
}

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length" yaml:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number" yaml:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length" yaml:"max_optional_data_field_value_length"`
}

// NewParams creates a new Params obj
func NewParams(maxPostMLen, maxOpDataFieldNum, maxOpDataFieldValLen sdk.Int) Params {
	return Params{
		MaxPostMessageLength:            maxPostMLen,
		MaxOptionalDataFieldsNumber:     maxOpDataFieldNum,
		MaxOptionalDataFieldValueLength: maxOpDataFieldValLen,
	}
}

// DefaultParams return default paramsModule
func DefaultParams() Params {
	return Params{
		MaxPostMessageLength:            DefaultMaxPostMessageLength,
		MaxOptionalDataFieldsNumber:     DefaultMaxOptionalDataFieldsNumber,
		MaxOptionalDataFieldValueLength: DefaultMaxOptionalDataFieldValueLength,
	}
}

// String implements Stringer
func (params Params) String() string {
	out := "Posts parameters:\n"
	out += fmt.Sprintf("MaxPostMessageLength: %s\nMaxOptionalDataFieldsNumber: %s\nMaxOptionalDataFieldValueLength: %s\n",
		params.MaxPostMessageLength,
		params.MaxOptionalDataFieldsNumber,
		params.MaxOptionalDataFieldValueLength,
	)

	return strings.TrimSpace(out)
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of posts module's parameters.
func (params *Params) ParamSetPairs() paramsModule.ParamSetPairs {
	return paramsModule.ParamSetPairs{
		paramsModule.NewParamSetPair(MaxPostMessageLengthKey, &params.MaxPostMessageLength, ValidateMaxPostMessageLengthParam),
		paramsModule.NewParamSetPair(MaxOptionalDataFieldsNumberKey, &params.MaxOptionalDataFieldsNumber, ValidateMaxOptionalDataFieldNumberParam),
		paramsModule.NewParamSetPair(MaxOptionalDataFieldValueLengthKey, &params.MaxOptionalDataFieldValueLength, ValidateMaxOptionalDataFieldValueLengthParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	if err := ValidateMaxPostMessageLengthParam(params.MaxPostMessageLength); err != nil {
		return err
	}

	if err := ValidateMaxOptionalDataFieldNumberParam(params.MaxOptionalDataFieldsNumber); err != nil {
		return err
	}

	if err := ValidateMaxOptionalDataFieldValueLengthParam(params.MaxOptionalDataFieldValueLength); err != nil {
		return err
	}

	return nil
}

func ValidateMaxPostMessageLengthParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsNegative() || params.LT(DefaultMaxPostMessageLength) {
		return fmt.Errorf("invalid max post message length param: %s", params)
	}

	return nil
}

func ValidateMaxOptionalDataFieldNumberParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsNegative() || params.LT(DefaultMaxOptionalDataFieldsNumber) {
		return fmt.Errorf("invalid max optional data fields number param: %s", params)
	}

	return nil
}

func ValidateMaxOptionalDataFieldValueLengthParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsNegative() || params.LT(DefaultMaxOptionalDataFieldValueLength) {
		return fmt.Errorf("invalid max optional data fields value length param: %s", params)
	}

	return nil
}
