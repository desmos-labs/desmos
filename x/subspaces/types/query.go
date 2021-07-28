package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQuerySubspaceRequest returns a new QuerySubspaceRequest instance
func NewQuerySubspaceRequest(subspaceID string) *QuerySubspaceRequest {
	return &QuerySubspaceRequest{SubspaceId: subspaceID}
}

// NewQuerySubspacesRequest returns a new QuerySubspacesRequest instance
func NewQuerySubspacesRequest(pagination *query.PageRequest) *QuerySubspacesRequest {
	return &QuerySubspacesRequest{
		Pagination: pagination,
	}
}

// NewQueryAdminsRequest returns a new QuerySubspaceAdminsRequest instance
func NewQueryAdminsRequest(subspaceID string, pagination *query.PageRequest) *QueryAdminsRequest {
	return &QueryAdminsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryRegisteredUsersRequest returns a new QuerySubspaceRegisteredUsersRequest instance
func NewQueryRegisteredUsersRequest(subspaceID string, pagination *query.PageRequest) *QueryRegisteredUsersRequest {
	return &QueryRegisteredUsersRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryBannedUsersRequest returns a new QuerySubspaceBannedUsersRequest instance
func NewQueryBannedUsersRequest(subspaceID string, pagination *query.PageRequest) *QueryBannedUsersRequest {
	return &QueryBannedUsersRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryTokenomicsRequest returns a new QueryTokenomicsRequest instance
func NewQueryTokenomicsRequest(subspaceID string) *QueryTokenomicsRequest {
	return &QueryTokenomicsRequest{
		SubspaceId: subspaceID,
	}
}

// NewQueryAllTokenomicsRequest returns a new QueryAllTokenomicsRequest instance
func NewQueryAllTokenomicsRequest(pagination *query.PageRequest) *QueryAllTokenomicsRequest {
	return &QueryAllTokenomicsRequest{
		Pagination: pagination,
	}
}
