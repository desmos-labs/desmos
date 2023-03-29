package types

import (
	"fmt"
)

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

		if isMinFeeDuplicated(fee, fees) {
			return fmt.Errorf("duplicated min fee for message type %s", fee.MessageType)
		}
	}

	return nil
}

func isMinFeeDuplicated(value MinFee, minFees []MinFee) bool {
	var count = 0
	for _, fee := range minFees {
		if fee.MessageType == value.MessageType {
			count++
		}
	}
	return count > 1
}
