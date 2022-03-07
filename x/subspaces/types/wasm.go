package types

import "encoding/json"

type SubspacesMsgRoute struct {
	Msg SubspacesMsg `json:"subspaces"`
}

type SubspacesMsg struct {
	CreateSubspace          *MsgCreateSubspace          `json:"create_subspace"`
	EditSubspace            *MsgEditSubspace            `json:"edit_subspace"`
	DeleteSubspace          *MsgDeleteSubspace          `json:"delete_subspace"`
	CreateUserGroup         *MsgCreateUserGroup         `json:"create_user_group"`
	EditUserGroup           *MsgEditUserGroup           `json:"edit_user_group"`
	SetUserGroupPermissions *MsgSetUserGroupPermissions `json:"set_user_group_permissions"`
	DeleteUserGroup         *MsgDeleteUserGroup         `json:"delete_user_group"`
	AddUserToUserGroup      *MsgAddUserToUserGroup      `json:"add_user_to_user_group"`
	RemoveUserFromUserGroup *MsgRemoveUserFromUserGroup `json:"remove_user_from_user_group"`
	SetUserPermissions      *MsgSetUserPermissions      `json:"set_user_permissions"`
}

type SubspacesQueryRoute struct {
	Query SubspacesQueryRequest `json:"subspaces"`
}

type SubspacesQueryRequest struct {
	Subspaces        json.RawMessage               `json:"subspaces"`
	Subspace         json.RawMessage               `json:"subspace"`
	UserGroups       *QueryUserGroupsRequest       `json:"user_groups"`
	UserGroup        *QueryUserGroupRequest        `json:"user_group"`
	UserGroupMembers *QueryUserGroupMembersRequest `json:"user_group_members"`
	UserPermissions  *QueryUserPermissionsRequest  `json:"user_permissions"`
}
