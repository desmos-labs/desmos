package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMinFee(messageType string, amount sdk.Coins) MinFee {
	return MinFee{
		MessageType: messageType,
		Amount:      amount,
	}
}

// Validate check if the min fee parameters are valid
func (mf MinFee) Validate() error {
	if mf.MessageType == "" {
		return fmt.Errorf("invalid minimum fee message type")
	}

	if !mf.Amount.IsValid() {
		return fmt.Errorf("invalid minimum fee amount")
	}

	return nil
}

type MinFees []MinFee
