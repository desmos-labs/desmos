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

// NewQueryUserGroupsRequest returns a new QueryUserGroupsRequest instance
func NewQueryUserGroupsRequest(subspaceID uint64, pagination *query.PageRequest) *QueryUserGroupsRequest {
	return &QueryUserGroupsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQueryUserGroupMembersRequest returns a new QueryUserGroupMembersRequest instance
func NewQueryUserGroupMembersRequest(
	subspaceID uint64, groupName string, pagination *query.PageRequest,
) *QueryUserGroupMembersRequest {
	return &QueryUserGroupMembersRequest{
		SubspaceId: subspaceID,
		GroupName:  groupName,
		Pagination: pagination,
	}
}

// NewQueryPermissionsRequest returns a new QueryPermissionsRequest instance
func NewQueryPermissionsRequest(subspaceID uint64, target string) *QueryPermissionsRequest {
	return &QueryPermissionsRequest{
		SubspaceId: subspaceID,
		Target:     target,
	}
}
