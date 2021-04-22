package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewPictures is a constructor function for Pictures
func NewPictures(profile, cover string) Pictures {
	return Pictures{
		Profile: profile,
		Cover:   cover,
	}
}

// Validate check the validity of the Pictures
func (pic Pictures) Validate() error {
	if pic.Profile != "" {
		valid := commons.IsURIValid(pic.Profile)
		if !valid {
			return fmt.Errorf("invalid profile picture uri provided")
		}
	}

	if pic.Cover != "" {
		valid := commons.IsURIValid(pic.Cover)
		if !valid {
			return fmt.Errorf("invalid profile cover uri provided")
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

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
