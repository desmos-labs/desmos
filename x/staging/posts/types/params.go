package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

const (
	// Default params space for the params keeper
	DefaultParamSpace = ModuleName
)

// Parameters store keys
var (
	MaxPostMessageLengthKey                    = []byte("MaxPostMessageLength")
	MaxAdditionalAttributesFieldsNumberKey     = []byte("MaxAdditionalAttributesFieldsNumber")
	MaxAdditionalAttributesFieldValueLengthKey = []byte("MaxAdditionalAttributesFieldValueLength")
	MaxAdditionalAttributesFieldKeyLengthKey   = []byte("MaxAdditionalAttributesFieldKeyLength")
	ReportTypesKey                             = []byte("ReportTypesKey")

	DefaultReportReasons = []string{
		"nudity",
		"violence",
		"intimidation",
		"harassment",
		"suicide",
		"self_harm",
		"fake_information",
		"spam",
		"unauthorized_sales",
		"terrorism",
		"scam",
	}
)

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params obj
func NewParams(maxPostMLen, maxOpDataFieldNum, maxOpDataFieldValLen, maxOpDataFieldKeyLen sdk.Int, reportReasons []string) Params {
	return Params{
		MaxPostMessageLength:                    maxPostMLen,
		MaxAdditionalAttributesFieldsNumber:     maxOpDataFieldNum,
		MaxAdditionalAttributesFieldValueLength: maxOpDataFieldValLen,
		MaxAdditionalAttributesFieldKeyLength:   maxOpDataFieldKeyLen,
		ReportReasons:                           reportReasons,
	}
}

// DefaultParams return default params object
func DefaultParams() Params {
	return Params{
		MaxPostMessageLength:                    sdk.NewInt(500),
		MaxAdditionalAttributesFieldsNumber:     sdk.NewInt(10),
		MaxAdditionalAttributesFieldValueLength: sdk.NewInt(200),
		MaxAdditionalAttributesFieldKeyLength:   sdk.NewInt(10),
		ReportReasons:                           DefaultReportReasons,
	}
}

// ParamSetPairs implements the ParamSet interface and returns the key/value pairs
// of posts module's parameters.
func (params *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.NewParamSetPair(MaxPostMessageLengthKey,
			&params.MaxPostMessageLength, ValidateMaxPostMessageLengthParam),
		paramstypes.NewParamSetPair(MaxAdditionalAttributesFieldsNumberKey,
			&params.MaxAdditionalAttributesFieldsNumber, ValidateMaxAdditionalAttributesFieldNumberParam),
		paramstypes.NewParamSetPair(MaxAdditionalAttributesFieldValueLengthKey,
			&params.MaxAdditionalAttributesFieldValueLength, ValidateMaxAdditionalAttributesFieldValueLengthParam),
		paramstypes.NewParamSetPair(MaxAdditionalAttributesFieldKeyLengthKey,
			&params.MaxAdditionalAttributesFieldKeyLength, ValidateMaxAdditionalAttributesFieldKeyLengthParam),
		paramstypes.NewParamSetPair(ReportTypesKey,
			&params.ReportReasons, ValidateReportReasonsParam),
	}
}

// Validate perform basic checks on all parameters to ensure they are correct
func (params Params) Validate() error {
	err := ValidateMaxPostMessageLengthParam(params.MaxPostMessageLength)
	if err != nil {
		return err
	}

	err = ValidateMaxAdditionalAttributesFieldNumberParam(params.MaxAdditionalAttributesFieldsNumber)
	if err != nil {
		return err
	}

	err = ValidateMaxAdditionalAttributesFieldValueLengthParam(params.MaxAdditionalAttributesFieldValueLength)
	if err != nil {
		return err
	}

	err = ValidateMaxAdditionalAttributesFieldKeyLengthParam(params.MaxAdditionalAttributesFieldKeyLength)
	if err != nil {
		return err
	}

	err = ValidateReportReasonsParam(params.ReportReasons)
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

func ValidateMaxAdditionalAttributesFieldNumberParam(i interface{}) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsZero() || params.IsNegative() {
		return fmt.Errorf("invalid max additional attributes fields number param: %s", params)
	}

	return nil
}

func validateAdditionalAttributesFieldLengthParam(i interface{}, paramName string) error {
	params, isCorrectParam := i.(sdk.Int)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if params.IsZero() || params.IsNegative() {
		return fmt.Errorf("invalid max additional attributes fields %s length param: %s", paramName, params)
	}

	return nil
}

func ValidateMaxAdditionalAttributesFieldValueLengthParam(i interface{}) error {
	return validateAdditionalAttributesFieldLengthParam(i, "value")
}

func ValidateMaxAdditionalAttributesFieldKeyLengthParam(i interface{}) error {
	return validateAdditionalAttributesFieldLengthParam(i, "key")
}

func ValidateReportReasonsParam(i interface{}) error {
	reportReasonsParam, isCorrectParam := i.([]string)

	if !isCorrectParam {
		return fmt.Errorf("invalid parameters type: %s", i)
	}

	if len(reportReasonsParam) == 0 {
		return fmt.Errorf("invalid report reasons param length")
	}

	for _, reportReason := range reportReasonsParam {
		if strings.TrimSpace(reportReason) == "" {
			return fmt.Errorf("invalid empty report reason inside report reasons param")
		}
	}

	return nil
}
