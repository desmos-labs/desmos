package v2

// Permission represents a permission that can be set to a user or user group
type Permission = uint32

const (
	// PermissionNothing represents the permission to do nothing
	PermissionNothing = Permission(0b000000)

	// PermissionWrite identifies users that can create content inside the subspace
	PermissionWrite = Permission(0b000001)

	// PermissionModerateContent allows users to moderate contents of other users (e.g. deleting it)
	PermissionModerateContent = Permission(0b000010)

	// PermissionChangeInfo allows to change the information of the subspace
	PermissionChangeInfo = Permission(0b000100)

	// PermissionManageGroups allows users to manage user groups and members
	PermissionManageGroups = Permission(0b001000)

	// PermissionSetPermissions allows users to set other users' permissions (except PermissionSetPermissions).
	// This includes managing user groups and the associated permissions
	PermissionSetPermissions = Permission(0b010000)

	// PermissionDeleteSubspace allows users to delete the subspace.
	PermissionDeleteSubspace = Permission(0b100000)

	// PermissionEverything allows to do everything.
	// This should usually be reserved only to the owner (which has it by default)
	PermissionEverything = Permission(0b111111)
)

var (
	permissionsMap = map[Permission]string{
		PermissionNothing:         "Nothing",
		PermissionWrite:           "Write",
		PermissionModerateContent: "ModerateContent",
		PermissionChangeInfo:      "ChangeInfo",
		PermissionManageGroups:    "ManageGroups",
		PermissionSetPermissions:  "SetUserPermissions",
		PermissionEverything:      "Everything",
	}
)

// CombinePermissions combines all the given permissions into a single Permission object using the OR operator
func CombinePermissions(permissions ...Permission) Permission {
	result := PermissionNothing
	for _, permission := range permissions {
		result |= permission
	}
	return result
}

// SanitizePermission sanitizes the given permission to remove any unwanted bits set to 1
func SanitizePermission(permission Permission) Permission {
	mask := PermissionNothing
	for perm := range permissionsMap {
		mask = CombinePermissions(mask, perm)
	}

	return permission & mask
}
