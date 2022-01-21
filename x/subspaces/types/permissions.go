package types

import (
	"encoding/binary"
)

// Permission represents a permission that can be set to a user or user group
type Permission = uint32

const (
	// PermissionNothing represents hte permission to do nothing
	PermissionNothing = Permission(0b000000)

	// PermissionWrite identifies users that can create content inside the subspace
	PermissionWrite = Permission(0b000001)

	// PermissionModerateContent allows users to moderate contents of other users (e.g. deleting it)
	PermissionModerateContent = Permission(0b000010)

	// PermissionChangeInfo allows to change the information of the subspace
	PermissionChangeInfo = Permission(0b000100)

	// PermissionManageGroups allows users to manage user groups and members
	PermissionManageGroups = Permission(0b001000)

	// PermissionSetPermissions allows to set other users' permissions (except PermissionSetPermissions).
	// This includes managing user groups and the associated permissions
	PermissionSetPermissions = Permission(0b010000)

	// PermissionEverything allows to do everything.
	// This should usually be reserved only to the owner (which has it by default)
	PermissionEverything = Permission(0b111111)
)

// MarshalPermission marshals the given permission to a byte array
func MarshalPermission(permission Permission) (permissionBytes []byte) {
	permissionBytes = make([]byte, 4)
	binary.BigEndian.PutUint32(permissionBytes, permission)
	return
}

// UnmarshalPermission reads the given byte array as a Permission object
func UnmarshalPermission(bz []byte) (permission Permission) {
	if len(bz) < 4 {
		return PermissionNothing
	}
	return binary.BigEndian.Uint32(bz)
}

// CheckPermission checks whether the given permissions contain the specified permission
func CheckPermission(permissions Permission, permission Permission) bool {
	return (permissions & permission) == permission
}

// CombinePermissions combines all the given permissions into a single Permission object using the OR operator
func CombinePermissions(permissions ...Permission) Permission {
	result := PermissionNothing
	for _, permission := range permissions {
		result = result | permission
	}
	return result
}
