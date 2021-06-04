package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewDTagTransferRequest(dTagToTrade string, sender, receiver string) DTagTransferRequest {
	return DTagTransferRequest{
		DTagToTrade: dTagToTrade,
		Receiver:    receiver,
		Sender:      sender,
	}
}

// Validate checks the request validity
func (request DTagTransferRequest) Validate() error {
	_, err := sdk.AccAddressFromBech32(request.Sender)
	if err != nil {
		return fmt.Errorf("invalid DTag transfer request sender address: %s", request.Sender)
	}

	_, err = sdk.AccAddressFromBech32(request.Receiver)
	if err != nil {
		return fmt.Errorf("invalid receiver address: %s", request.Receiver)
	}

	if request.Receiver == request.Sender {
		return fmt.Errorf("the sender and receiver must be different")
	}

	if strings.TrimSpace(request.DTagToTrade) == "" {
		return fmt.Errorf("invalid DTag to trade: %s", request.DTagToTrade)
	}

	return nil
}

// NewDTagTransferRequests returns a DTagTransferRequests instance wrapping the given requests
func NewDTagTransferRequests(requests []DTagTransferRequest) DTagTransferRequests {
	return DTagTransferRequests{Requests: requests}
}
