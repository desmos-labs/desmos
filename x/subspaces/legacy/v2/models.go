package v2

// NewUserGroup returns a new UserGroup instance
func NewUserGroup(subspaceID uint64, id uint32, name, description string, permissions Permission) UserGroup {
	return UserGroup{
		SubspaceID:  subspaceID,
		ID:          id,
		Name:        name,
		Description: description,
		Permissions: permissions,
	}
}

// --------------------------------------------------------------------------------------------------------------------

// NewPermissionDetailUser returns a new PermissionDetail for the user with the given address and permission value
func NewPermissionDetailUser(user string, permission Permission) PermissionDetail {
	return PermissionDetail{
		Sum: &PermissionDetail_User_{
			User: &PermissionDetail_User{
				User:       user,
				Permission: permission,
			},
		},
	}
}

// NewPermissionDetailGroup returns a new PermissionDetail for the user with the given id and permission value
func NewPermissionDetailGroup(groupID uint32, permission Permission) PermissionDetail {
	return PermissionDetail{
		Sum: &PermissionDetail_Group_{
			Group: &PermissionDetail_Group{
				GroupID:    groupID,
				Permission: permission,
			},
		},
	}
}
