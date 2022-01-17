package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQuerySubspaceRequest returns a new QuerySubspaceRequest instance
func NewQuerySubspaceRequest(subspaceID uint64) *QuerySubspaceRequest {
	return &QuerySubspaceRequest{SubspaceId: subspaceID}
}

// NewQuerySubspacesRequest returns a new QuerySubspacesRequest instance
func NewQuerySubspacesRequest(pagination *query.PageRequest) *QuerySubspacesRequest {
	return &QuerySubspacesRequest{
		Pagination: pagination,
	}
}
