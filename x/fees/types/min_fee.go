package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MinFee contains the minimum amount of coins that should be paid as a fee for each message of the specified type sent
type MinFee struct {
	MessageType string  `json:"message_type" yaml:"message_type"`
	Amount      sdk.Dec `json:"amount" yaml:"amount"`
}

func NewMinFee(messageType string, amount sdk.Dec) MinFee {
	return MinFee{
		MessageType: messageType,
		Amount:      amount,
	}
}

func (mf MinFee) String() string {
	return fmt.Sprintf("Message Type: %s\nAmount: %s\n", mf.MessageType, mf.Amount)
}

// Validate check if the min fee parameters are valid
func (mf MinFee) Validate() error {
	if mf.MessageType == "" {
		return fmt.Errorf("invalid minimum fee message type")
	}

	if mf.Amount.IsNegative() {
		return fmt.Errorf("minimum fee amout cannot be negative")
	}

	return nil
}
