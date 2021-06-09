package types

import (
	"github.com/cosmos/cosmos-sdk/types/query"
)

// NewQuerySubspaceRequest returns a new QuerySubspaceRequest instance
func NewQuerySubspaceRequest(subspaceID string) *QuerySubspaceRequest {
	return &QuerySubspaceRequest{SubspaceId: subspaceID}
}

// NewQuerySubspacesRequest allows to build a new QuerySubspacesRequest instance
func NewQuerySubspacesRequest(pagination *query.PageRequest) *QuerySubspacesRequest {
	return &QuerySubspacesRequest{
		Pagination: pagination,
	}
}

// NewQuerySubspaceAdminsRequest returns a new QuerySubspaceAdminsRequest instance
func NewQuerySubspaceAdminsRequest(subspaceID string, pagination *query.PageRequest) *QuerySubspaceAdminsRequest {
	return &QuerySubspaceAdminsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQuerySubspaceRegisteredUsersRequest returns a new QuerySubspaceRegisteredUsersRequest instance
func NewQuerySubspaceRegisteredUsersRequest(subspaceID string, pagination *query.PageRequest) *QuerySubspaceRegisteredUsersRequest {
	return &QuerySubspaceRegisteredUsersRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQuerySubspaceBannedUsersRequest returns a new QuerySubspaceBannedUsersRequest instance
func NewQuerySubspaceBannedUsersRequest(subspaceID string, pagination *query.PageRequest) *QuerySubspaceBannedUsersRequest {
	return &QuerySubspaceBannedUsersRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}
