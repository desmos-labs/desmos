package types

type SubspacesMsg struct {
	CreateSubspace *MsgCreateSubspace `json:"create_subspace"`
}

type SubspacesQueryRequest struct {
	Subspaces        *QuerySubspacesRequest        `json:"subspaces"`
	Subspace         *QuerySubspaceRequest         `json:"subspace"`
	UserGroups       *QueryUserGroupsRequest       `json:"user_groups"`
	UserGroup        *QueryUserGroupRequest        `json:"user_group"`
	UserGroupMembers *QueryUserGroupMembersRequest `json:"user_group_members"`
	UserPermissions  *QueryUserPermissionsRequest  `json:"user_permissions"`
}
