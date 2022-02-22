package types

type SubspacesQueryRoutes struct {
	Subspaces SubspacesQueryRequest `json:"subspaces"`
}

type SubspacesQueryRequest struct {
	Subspaces        *QuerySubspacesRequest        `json:"subspaces"`
	Subspace         *QuerySubspaceRequest         `json:"subspace"`
	UserGroups       *QueryUserGroupsRequest       `json:"user_groups"`
	UserGroup        *QueryUserGroupRequest        `json:"user_group"`
	UserGroupMembers *QueryUserGroupMembersRequest `json:"user_group_members"`
	UserPermissions  *QueryUserPermissionsRequest  `json:"user_permissions"`
}
