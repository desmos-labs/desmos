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
	if strings.TrimSpace(mf.MessageType) == "" {
		return fmt.Errorf("invalid minimum fee message type")
	}

	if !mf.Amount.IsValid() {
		return fmt.Errorf("invalid minimum fee amount")
	}

	return nil
}
