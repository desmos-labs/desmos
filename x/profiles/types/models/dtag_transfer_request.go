package models

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DTagTransferRequest represent a dtag transfer request between two users
type DTagTransferRequest struct {
	DTagToTrade   string         `json:"dtag_to_trade" yaml:"dtag_to_trade"`
	CurrentOwner  sdk.AccAddress `json:"current_owner" yaml:"current_owner"`
	ReceivingUser sdk.AccAddress `json:"receiving_user" yaml:"receiving_user"`
}

func NewDTagTransferRequest(dtagToTrade string, currentOwner, receivingUser sdk.AccAddress) DTagTransferRequest {
	return DTagTransferRequest{
		DTagToTrade:   dtagToTrade,
		CurrentOwner:  currentOwner,
		ReceivingUser: receivingUser,
	}
}

// Equals returns true if the two requests are equals. False otherwise
func (dtagTR DTagTransferRequest) Equals(other DTagTransferRequest) bool {
	return dtagTR.DTagToTrade == other.DTagToTrade &&
		dtagTR.CurrentOwner.Equals(other.CurrentOwner) &&
		dtagTR.ReceivingUser.Equals(other.ReceivingUser)
}

// String implement fmt.Stringer
func (dtagTR DTagTransferRequest) String() string {
	out := "DTag transfer request:\n"
	out += fmt.Sprintf("[DTagToTrade] %s [Current CurrentOwner] %s [Receiving User] %s",
		dtagTR.DTagToTrade, dtagTR.CurrentOwner, dtagTR.ReceivingUser)
	return out
}

// Validate checks the request validity
func (dtagTR DTagTransferRequest) Validate() error {
	if len(strings.TrimSpace(dtagTR.DTagToTrade)) == 0 {
		return fmt.Errorf("invalid DTag to trade %s", dtagTR.DTagToTrade)
	}

	if dtagTR.CurrentOwner.Empty() {
		return fmt.Errorf("current owner address cannot be empty")
	}

	if dtagTR.ReceivingUser.Empty() {
		return fmt.Errorf("receiving user address cannot be empty")
	}

	if dtagTR.CurrentOwner.Equals(dtagTR.ReceivingUser) {
		return fmt.Errorf("the receiving user and current owner must be different")
	}

	return nil
}
