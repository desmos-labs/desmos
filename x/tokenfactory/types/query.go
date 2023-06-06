package types

// NewQuerySubspaceDenomsRequest returns a new QuerySubspaceDenomsRequest instance
func NewQuerySubspaceDenomsRequest(subspaceID uint64) *QuerySubspaceDenomsRequest {
	return &QuerySubspaceDenomsRequest{
		SubspaceId: subspaceID,
	}
}

// NewQueryParamsRequest returns a new QueryParamsRequest instance
func NewQueryParamsRequest() *QueryParamsRequest {
	return &QueryParamsRequest{}
}
