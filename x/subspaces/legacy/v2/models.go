package v2

// DONTCOVER

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
