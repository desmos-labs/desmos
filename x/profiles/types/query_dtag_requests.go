package types

// NewQueryDTagTransfersRequest returns a new QueryDTagTransfersRequest containing the given data
func NewQueryDTagTransfersRequest(user string) *QueryDTagTransfersRequest {
	return &QueryDTagTransfersRequest{
		User: user,
	}
}
