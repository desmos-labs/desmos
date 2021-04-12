package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// Default params space for the params keeper
	DefaultParamSpace = ModuleName
)

// Parameters store keys
var (
	MaxPostMessageLengthKey            = []byte("MaxPostMessageLength")
	MaxOptionalDataFieldsNumberKey     = []byte("MaxOptionalDataFieldsNumber")
	MaxOptionalDataFieldValueLengthKey = []byte("MaxOptionalDataFieldValueLength")
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params obj
func NewParams(maxPostMLen, maxOpDataFieldNum, maxOpDataFieldValLen sdk.Int) Params {
	return Params{
		MaxPostMessageLength:            maxPostMLen,
		MaxOptionalDataFieldsNumber:     maxOpDataFieldNum,
		MaxOptionalDataFieldValueLength: maxOpDataFieldValLen,
	}
}

// DefaultParams return default params object
func DefaultParams() Params {
	return Params{
		MaxPostMessageLength:            sdk.NewInt(500),
		MaxOptionalDataFieldsNumber:     sdk.NewInt(10),
		MaxOptionalDataFieldValueLength: sdk.NewInt(200),
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of posts module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MaxPostMessageLengthKey,
			&params.MaxPostMessageLength, ValidateMaxPostMessageLengthParam),
		paramstypes.NewParamSetPair(MaxOptionalDataFieldsNumberKey,
			&params.MaxOptionalDataFieldsNumber, ValidateMaxOptionalDataFieldNumberParam),
		paramstypes.NewParamSetPair(MaxOptionalDataFieldValueLengthKey,
			&params.MaxOptionalDataFieldValueLength, ValidateMaxOptionalDataFieldValueLengthParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	err := ValidateMaxPostMessageLengthParam(params.MaxPostMessageLength)
	if err != nil {
		return err
	}

	err = ValidateMaxOptionalDataFieldNumberParam(params.MaxOptionalDataFieldsNumber)
	if err != nil {
		return err
	}

	err = ValidateMaxOptionalDataFieldValueLengthParam(params.MaxOptionalDataFieldValueLength)
	if err != nil {
		return err
	}

	return nil
}

func ValidateMaxPostMessageLengthParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsZero() || params.IsNegative() {
		return fmt.Errorf("invalid max post message length param: %s", params)
	}

	return nil
}

func ValidateMaxOptionalDataFieldNumberParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsZero() || params.IsNegative() {
		return fmt.Errorf("invalid max optional data fields number param: %s", params)
	}

	return nil
}

func ValidateMaxOptionalDataFieldValueLengthParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsZero() || params.IsNegative() {
		return fmt.Errorf("invalid max optional data fields value length param: %s", params)
	}

	return nil
}
