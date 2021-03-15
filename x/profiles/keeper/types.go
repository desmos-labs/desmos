package keeper

import "github.com/desmos-labs/desmos/x/profiles/types"

// NewWrappedDTagTransferRequests returns a DTagRequests instance wrapping the given requests
func NewWrappedDTagTransferRequests(requests []types.DTagTransferRequest) WrappedDTagTransferRequests {
	return WrappedDTagTransferRequests{Requests: requests}
}
