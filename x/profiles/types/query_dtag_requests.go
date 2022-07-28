package types

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQueryIncomingDTagTransferRequestsRequest returns a new QueryIncomingDTagTransferRequestsRequest instance
func NewQueryIncomingDTagTransferRequestsRequest(
	receiver string, pagination *query.PageRequest,
) *QueryIncomingDTagTransferRequestsRequest {
	return &QueryIncomingDTagTransferRequestsRequest{
		Receiver:   receiver,
		Pagination: pagination,
	}
}
