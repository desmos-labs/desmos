package v2

import (
	"encoding/binary"
	"fmt"
	"sort"
	"strings"
)

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
		PermissionDeleteSubspace:  "DeleteSubspace",
		PermissionEverything:      "Everything",
	}
)

// ParsePermission parses the given permission string as a single Permissions instance
func ParsePermission(permission string) (Permission, error) {
	// Check inside the map if we have anything here
	for permValue, permString := range permissionsMap {
		if strings.EqualFold(permission, permString) {
			return permValue, nil
		}
	}

	return 0, fmt.Errorf("invalid permission value: %s", permission)
}

// SerializePermission serializes the given permission to a string value
func SerializePermission(permission Permission) string {
	return permissionsMap[permission]
}

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
		result |= permission
	}
	return result
}

// getValidPermissions returns the valid permissions slice
func getValidPermissions() []Permission {
	validPermissions := make([]Permission, len(permissionsMap))

	i := 0
	for perm := range permissionsMap {
		validPermissions[i] = perm
		i++
	}

	// Sort the permissions
	sort.Slice(validPermissions, func(i, j int) bool {
		return validPermissions[i] < validPermissions[j]
	})

	return validPermissions
}

// SplitPermissions splits the given combined permission value into its individual values
func SplitPermissions(permission Permission) []Permission {
	if permission == PermissionNothing {
		return nil
	}

	if permission == PermissionEverything {
		return []Permission{PermissionEverything}
	}

	var permissions []Permission
	for _, perm := range getValidPermissions() {
		if perm == PermissionNothing || perm == PermissionEverything {
			continue
		}

		if (permission & perm) == perm {
			permissions = append(permissions, perm)
		}
	}

	return permissions
}

// SanitizePermission sanitizes the given permission to remove any unwanted bits set to 1
func SanitizePermission(permission Permission) Permission {
	mask := PermissionNothing
	for perm := range permissionsMap {
		mask = CombinePermissions(mask, perm)
	}

	return permission & mask
}

// IsPermissionValid checks whether the given value represents a valid permission or not
func IsPermissionValid(permission Permission) bool {
	return SanitizePermission(permission) == permission
}
