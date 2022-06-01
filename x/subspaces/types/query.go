package types

// DONTCOVER

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

// NewQuerySectionsRequest returns a new QuerySectionsRequest instance
func NewQuerySectionsRequest(subspaceID uint64, pagination *query.PageRequest) *QuerySectionsRequest {
	return &QuerySectionsRequest{
		SubspaceId: subspaceID,
		Pagination: pagination,
	}
}

// NewQuerySectionRequest returns a new QuerySectionRequest instance
func NewQuerySectionRequest(subspaceID uint64, sectionID uint32) *QuerySectionRequest {
	return &QuerySectionRequest{
		SubspaceId: subspaceID,
		SectionId:  sectionID,
	}
}

// NewQueryUserGroupsRequest returns a new QueryUserGroupsRequest instance
func NewQueryUserGroupsRequest(subspaceID uint64, sectionID uint32, pagination *query.PageRequest) *QueryUserGroupsRequest {
	return &QueryUserGroupsRequest{
		SubspaceId: subspaceID,
		SectionId:  sectionID,
		Pagination: pagination,
	}
}

// NewQueryUserGroupRequest returns a new QueryUserGroupRequest instance
func NewQueryUserGroupRequest(subspaceID uint64, groupID uint32) *QueryUserGroupRequest {
	return &QueryUserGroupRequest{
		SubspaceId: subspaceID,
		GroupId:    groupID,
	}
}

// NewQueryUserGroupMembersRequest returns a new QueryUserGroupMembersRequest instance
func NewQueryUserGroupMembersRequest(
	subspaceID uint64, groupID uint32, pagination *query.PageRequest,
) *QueryUserGroupMembersRequest {
	return &QueryUserGroupMembersRequest{
		SubspaceId: subspaceID,
		GroupId:    groupID,
		Pagination: pagination,
	}
}

// NewQueryUserPermissionsRequest returns a new QueryPermissionsRequest instance
func NewQueryUserPermissionsRequest(subspaceID uint64, user string) *QueryUserPermissionsRequest {
	return &QueryUserPermissionsRequest{
		SubspaceId: subspaceID,
		User:       user,
	}
}

// NewPermissionDetailUser returns a new PermissionDetail for the user with the given address and permission value
func NewPermissionDetailUser(subspaceID uint64, sectionID uint32, user string, permissions Permissions) PermissionDetail {
	return PermissionDetail{
		SubspaceId: subspaceID,
		SectionId:  sectionID,
		Sum: &PermissionDetail_User_{
			User: &PermissionDetail_User{
				User:       user,
				Permission: permissions,
			},
		},
	}
}

// NewPermissionDetailGroup returns a new PermissionDetail for the user with the given id and permission value
func NewPermissionDetailGroup(subspaceID uint64, sectionID uint32, groupID uint32, permissions Permissions) PermissionDetail {
	return PermissionDetail{
		SubspaceId: subspaceID,
		SectionId:  sectionID,
		Sum: &PermissionDetail_Group_{
			Group: &PermissionDetail_Group{
				GroupID:    groupID,
				Permission: permissions,
			},
		},
	}
}
