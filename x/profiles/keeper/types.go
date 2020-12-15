package keeper

import "github.com/desmos-labs/desmos/x/profiles/types"

// NewWrappedDTagOwner returns a DTagOwner instance wrapping the given address
func NewWrappedDTagOwner(address string) WrappedDTagOwner {
	return WrappedDTagOwner{Address: address}
}

// NewWrappedDTagTransferRequests returns a DTagRequests instance wrapping the given requests
func NewWrappedDTagTransferRequests(requests []types.DTagTransferRequest) WrappedDTagTransferRequests {
	return WrappedDTagTransferRequests{Requests: requests}
}
