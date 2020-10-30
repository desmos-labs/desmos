package keeper

import "github.com/desmos-labs/desmos/x/profiles/types"

// NewDTagOwner returns a DTagOwner instance wrapping the given address
func NewDTagOwner(address string) DTagOwner {
	return DTagOwner{Address: address}
}

// NewDTagRequests returns a DTagRequests instance wrapping the given requests
func NewDTagRequests(requests []types.DTagTransferRequest) DTagRequests {
	return DTagRequests{Requests: requests}
}
