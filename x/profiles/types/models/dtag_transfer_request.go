package models

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DTagTransferRequest represent a dtag transfer request between two users
type DTagTransferRequest struct {
	DTagToTrade string         `json:"dtag_to_trade" yaml:"dtag_to_trade"`
	Receiver    sdk.AccAddress `json:"receiver" yaml:"receiver"`
	Sender      sdk.AccAddress `json:"sender" yaml:"sender"`
}

func NewDTagTransferRequest(dtagToTrade string, receiver, sender sdk.AccAddress) DTagTransferRequest {
	return DTagTransferRequest{
		DTagToTrade: dtagToTrade,
		Receiver:    receiver,
		Sender:      sender,
	}
}

// Equals returns true if the two requests are equals. False otherwise
func (request DTagTransferRequest) Equals(other DTagTransferRequest) bool {
	return request.DTagToTrade == other.DTagToTrade &&
		request.Receiver.Equals(other.Receiver) &&
		request.Sender.Equals(other.Sender)
}

// String implement fmt.Stringer
func (request DTagTransferRequest) String() string {
	out := "DTag transfer request:\n"
	out += fmt.Sprintf("[DTag to trade] %s [Request Receiver] %s [Request Sender] %s",
		request.DTagToTrade, request.Receiver, request.Sender)
	return out
}

// Validate checks the request validity
func (request DTagTransferRequest) Validate() error {
	if len(strings.TrimSpace(request.DTagToTrade)) == 0 {
		return fmt.Errorf("invalid DTag to trade %s", request.DTagToTrade)
	}

	if request.Receiver.Empty() {
		return fmt.Errorf("receiver address cannot be empty")
	}

	if request.Sender.Empty() {
		return fmt.Errorf("sender address cannot be empty")
	}

	if request.Receiver.Equals(request.Sender) {
		return fmt.Errorf("the sender and receiver must be different")
	}

	return nil
}
