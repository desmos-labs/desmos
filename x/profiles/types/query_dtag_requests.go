package types

// NewQueryIncomingDTagTransferRequestsRequest returns a new QueryIncomingDTagTransferRequestsRequest instance
func NewQueryIncomingDTagTransferRequestsRequest(receiver string) *QueryIncomingDTagTransferRequestsRequest {
	return &QueryIncomingDTagTransferRequestsRequest{
		Receiver: receiver,
	}
}
