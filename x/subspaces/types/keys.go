package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DONTCOVER

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace          = "create_subspace"
	ActionEditSubspace            = "edit_subspace"
	ActionDeleteSubspace          = "delete_subspace"
	ActionCreateSection           = "create_section"
	ActionEditSection             = "edit_section"
	ActionMoveSection             = "move_section"
	ActionDeleteSection           = "delete_section"
	ActionCreateUserGroup         = "create_user_group"
	ActionEditUserGroup           = "edit_user_group"
	ActionMoveUserGroup           = "move_user_group"
	ActionSetUserGroupPermissions = "set_user_group_permissions"
	ActionDeleteUserGroup         = "delete_user_group"
	ActionAddUserToUserGroup      = "add_user_to_user_group"
	ActionRemoveUserFromUserGroup = "remove_user_from_user_group"
	ActionSetUserPermissions      = "set_user_permissions"

	QuerierRoute = ModuleName

	DoNotModify = "[do-not-modify]"
)

var (
	SubspaceIDKey              = []byte{0x00}
	SubspacePrefix             = []byte{0x01}
	GroupIDPrefix              = []byte{0x02}
	GroupsPrefix               = []byte{0x03}
	GroupMembersStorePrefix    = []byte{0x04}
	UserPermissionsStorePrefix = []byte{0x05}
	SectionIDPrefix            = []byte{0x06}
	SectionsPrefix             = []byte{0x07}
)

// GetSubspaceIDBytes returns the byte representation of the subspaceID
func GetSubspaceIDBytes(subspaceID uint64) (subspaceIDBz []byte) {
	subspaceIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(subspaceIDBz, subspaceID)
	return
}

// GetSubspaceIDFromBytes returns subspaceID in uint64 format from a byte array
func GetSubspaceIDFromBytes(bz []byte) (subspaceID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// SubspaceKey returns the key for a specific subspace
func SubspaceKey(subspaceID uint64) []byte {
	return append(SubspacePrefix, GetSubspaceIDBytes(subspaceID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetSectionIDBytes returns the byte representation of the sectionID
func GetSectionIDBytes(sectionID uint32) (sectionIDBz []byte) {
	sectionIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(sectionIDBz, sectionID)
	return
}

// GetSectionIDFromBytes returns sectionID in uint32 format from a byte array
func GetSectionIDFromBytes(bz []byte) (sectionID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// NextSectionIDStoreKey returns the key used to store the next section id for the given subspace
func NextSectionIDStoreKey(subspaceID uint64) []byte {
	return append(SectionIDPrefix, GetSubspaceIDBytes(subspaceID)...)
}

// SubspaceSectionsPrefix returns the prefix used to store all the sections for the given subspace
func SubspaceSectionsPrefix(subspaceId uint64) []byte {
	return append(SectionsPrefix, GetSubspaceIDBytes(subspaceId)...)
}

// SectionStoreKey returns the key used to store the given section
func SectionStoreKey(subspaceID uint64, sectionID uint32) []byte {
	return append(SubspaceSectionsPrefix(subspaceID), GetSectionIDBytes(sectionID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// NextGroupIDStoreKey returns the store key that is used to store the group id to be used next for the given subspace
func NextGroupIDStoreKey(subspaceID uint64) []byte {
	return append(GroupIDPrefix, GetSubspaceIDBytes(subspaceID)...)
}

// GetGroupIDBytes returns the byte representation of the groupID
func GetGroupIDBytes(groupID uint32) (groupIDBz []byte) {
	groupIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(groupIDBz, groupID)
	return
}

// GetGroupIDFromBytes returns groupID in uint32 format from a byte array
func GetGroupIDFromBytes(bz []byte) (subspaceID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// SubspaceGroupsPrefix returns the store prefix used to store all the groups of a given subspace
func SubspaceGroupsPrefix(subspaceID uint64) []byte {
	return append(GroupsPrefix, GetSubspaceIDBytes(subspaceID)...)
}

// SectionGroupsPrefix returns the prefix used to store all the groups for the given section
func SectionGroupsPrefix(subspaceID uint64, sectionID uint32) []byte {
	return append(SubspaceGroupsPrefix(subspaceID), GetSectionIDBytes(sectionID)...)
}

// GroupStoreKey returns the key used to store the group having the given id inside the specified section
func GroupStoreKey(subspaceID uint64, sectionID uint32, groupID uint32) []byte {
	return append(SectionGroupsPrefix(subspaceID, sectionID), GetGroupIDBytes(groupID)...)
}

// GroupMembersStoreKey returns the key used to store all the members of the given group inside the given subspace
func GroupMembersStoreKey(subspaceID uint64, groupID uint32) []byte {
	return append(append(GroupMembersStorePrefix, GetSubspaceIDBytes(subspaceID)...), GetGroupIDBytes(groupID)...)
}

// GroupMemberStoreKey returns the key used to store the membership of the given user to the
// specified group inside the provided subspace
func GroupMemberStoreKey(subspaceID uint64, groupID uint32, user sdk.AccAddress) []byte {
	return append(GroupMembersStoreKey(subspaceID, groupID), GetAddressBytes(user)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetAddressBytes returns the given user address as a byte array
func GetAddressBytes(user sdk.AccAddress) []byte {
	return user
}

// GetAddressFromBytes returns the sdk.AccAddress representation of the given user address
func GetAddressFromBytes(bz []byte) sdk.AccAddress {
	return bz
}

// SubspacePermissionsPrefix returns the prefix used to store user permissions for the given subspace
func SubspacePermissionsPrefix(subspaceID uint64) []byte {
	return append(UserPermissionsStorePrefix, GetSubspaceIDBytes(subspaceID)...)
}

// SectionPermissionsPrefix returns the prefix used to store the permissions for the given section
func SectionPermissionsPrefix(subspaceID uint64, sectionID uint32) []byte {
	return append(SubspacePermissionsPrefix(subspaceID), GetSectionIDBytes(sectionID)...)
}

// UserPermissionStoreKey returns the key used to store the permission for the given user inside the given subspace
func UserPermissionStoreKey(subspaceID uint64, sectionID uint32, user sdk.AccAddress) []byte {
	return append(SectionPermissionsPrefix(subspaceID, sectionID), GetAddressBytes(user)...)
}

// SplitUserAddressPermissionKey splits a UserPermissionStoreKey into the subspace id, section id and user address
func SplitUserAddressPermissionKey(key []byte) (subspaceID uint64, sectionID uint32, user sdk.AccAddress) {
	key = key[1:] // Remove the prefix

	subspaceID = GetSubspaceIDFromBytes(key[8:])
	key = key[8:] // Remove the subspace id

	sectionID = GetSectionIDFromBytes(key[4:])
	key = key[4:] // Remove the section id

	return subspaceID, sectionID, GetAddressFromBytes(key)
}
