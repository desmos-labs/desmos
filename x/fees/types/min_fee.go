package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewMinFee allows to build a MinFee instance based on the given message type and fee amount
func NewMinFee(messageType string, amount sdk.Coins) MinFee {
	return MinFee{
		MessageType: messageType,
		Amount:      amount,
	}
}

// Validate checks if minimum fee represents a valid instance
func (mf MinFee) Validate() error {
	if !strings.HasPrefix(mf.MessageType, "/") {
		return fmt.Errorf("invalid message type")
	}

	if !mf.Amount.IsValid() {
		return fmt.Errorf("invalid minimum fee amount")
	}

	return nil
}

// ContainsMinFee returns true iff the given min fee slice contains a min fee for the given message type
func ContainsMinFee(minFees []MinFee, msgType string) bool {
	for _, minFee := range minFees {
		if minFee.MessageType == msgType {
			return true
		}
	}
	return false
}
