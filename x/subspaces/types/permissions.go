package types

import (
	"fmt"
	"strings"
)

// Permission represents a permissions that can be set to a user or user group
type Permission = string

// newPermission returns a new Permission containing the given value
func newPermission(permissionName string) Permission {
	return Permission(strings.ToUpper(strings.ReplaceAll(permissionName, " ", "_")))
}

var (
	// PermissionWrite identifies users that can create content inside the subspace
	// TODO: Remove this
	PermissionWrite = RegisterPermission("write")

	// PermissionEditSubspace allows to change the information of the subspace
	PermissionEditSubspace = RegisterPermission("edit subspace")

	// PermissionDeleteSubspace allows users to delete the subspace.
	PermissionDeleteSubspace = RegisterPermission("delete subspace")

	// PermissionManageGroups allows users to manage user groups and members
	PermissionManageGroups = RegisterPermission("manage groups")

	// PermissionSetPermissions allows users to set other users' permissions (except PermissionSetPermissions).
	// This includes managing user groups and the associated permissions
	PermissionSetPermissions = RegisterPermission("set permissions")

	// PermissionInteractWithContent allows users to interact with content inside the subspace (eg. polls)
	PermissionInteractWithContent = Permission(0b1000000)

	// PermissionEditOwnContent allows users to edit their own content inside the subspace
	PermissionEditOwnContent = Permission(0b10000000)

	// PermissionEverything allows to do everything.
	// This should usually be reserved only to the owner (which has it by default)
	PermissionEverything = RegisterPermission("everything")
)

var (
	// registeredPermissions represents the list of permissions that are registered and should be considered valid
	registeredPermissions []Permission
)

// containsPermission tells whether the given permissions slice contains the provided permissions
func containsPermission(slice []Permission, permission Permission) bool {
	for _, element := range slice {
		if element == permission {
			return true
		}
	}
	return false
}

// isPermissionRegistered checks if the given permissions is registered or not
func isPermissionRegistered(permission Permission) bool {
	return containsPermission(registeredPermissions, permission)
}

// RegisterPermission registers the permissions with the given name and returns its value
func RegisterPermission(permissionName string) Permission {
	permission := newPermission(permissionName)
	if isPermissionRegistered(permission) {
		panic(fmt.Errorf("permissions %s has already been registered", permission))
	}

	registeredPermissions = append(registeredPermissions, permission)
	return permission
}

// ParsePermission parses the given permissions string as a single Permissions instance
func ParsePermission(permission string) Permission {
	return newPermission(permission)
}

// SerializePermission serializes the given permissions to a string value
func SerializePermission(permission Permission) string {
	return string(permission)
}

type Permissions []Permission

// NewPermissions allows to build a new Permissions instance
func NewPermissions(permissions ...Permission) Permissions {
	return permissions
}

// Equals returns true iff the given permissions slice is equals to this one
func (p Permissions) Equals(other Permissions) bool {
	if len(p) != len(other) {
		return false
	}

	for i, element := range p {
		if element != other[i] {
			return false
		}
	}

	return true
}

// CheckPermission checks whether the given permissions contain the specified permissions
func CheckPermission(permissions Permissions, permission Permission) bool {
	// If PermissionEverything is set, every permissions will be valid
	if containsPermission(permissions, PermissionEverything) {
		return true
	}

	// Otherwise, check for the specific permissions
	return containsPermission(permissions, permission)
}

func CheckPermissions(permissions Permissions, toCheck Permissions) bool {
	for _, permission := range toCheck {
		if !CheckPermission(permissions, permission) {
			return false
		}
	}
	return true
}

// CombinePermissions combines all the given permissions into a single Permission object using the OR operator
func CombinePermissions(permissions ...Permission) Permissions {
	// If the given slice contains PermissionEverything, then just return that one
	if containsPermission(permissions, PermissionEverything) {
		return Permissions{PermissionEverything}
	}
	return SanitizePermission(permissions)
}

// SanitizePermission sanitizes the given permissions to remove any duplicate
func SanitizePermission(permissions Permissions) Permissions {
	var sanitized Permissions
	existing := map[Permission]bool{}
	for _, permission := range permissions {
		if !isPermissionRegistered(permission) {
			// Skip invalid permissions
			continue
		}

		if _, exists := existing[permission]; exists {
			// If a permission already exists, skip it
			continue
		}

		sanitized = append(sanitized, permission)
		existing[permission] = true
	}

	return sanitized
}

// ArePermissionsValid checks whether the given value represents a valid permissions or not
func ArePermissionsValid(permissions Permissions) bool {
	return SanitizePermission(permissions).Equals(permissions)
}
