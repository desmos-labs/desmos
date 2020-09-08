package models

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DTagTransferRequest represent a dtag transfer request between two users
type DTagTransferRequest struct {
	CurrentOwner  sdk.AccAddress `json:"current_owner" yaml:"current_owner"`
	ReceivingUser sdk.AccAddress `json:"receiving_user" yaml:"receiving_user"`
}

func NewDTagTransferRequest(currentOwner, receivingUser sdk.AccAddress) DTagTransferRequest {
	return DTagTransferRequest{
		CurrentOwner:  currentOwner,
		ReceivingUser: receivingUser,
	}
}

// Equals returns true if the two requests are equals. False otherwise
func (dtagTR DTagTransferRequest) Equals(other DTagTransferRequest) bool {
	return dtagTR.CurrentOwner.Equals(other.CurrentOwner) &&
		dtagTR.ReceivingUser.Equals(other.ReceivingUser)
}

// String implement fmt.Stringer
func (dtagTR DTagTransferRequest) String() string {
	out := "DTag transfer request:\n"
	out += fmt.Sprintf("[Current CurrentOwner] %s [Receiving User] %s", dtagTR.CurrentOwner, dtagTR.ReceivingUser)
	return out
}

// Validate checks the request validity
func (dtagTR DTagTransferRequest) Validate() error {
	if dtagTR.CurrentOwner.Empty() {
		return fmt.Errorf("current owner address cannot be empty")
	}

	if dtagTR.ReceivingUser.Empty() {
		return fmt.Errorf("receiving user address cannot be empty")
	}

	return nil
}
