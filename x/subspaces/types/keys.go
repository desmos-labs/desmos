package types

import "encoding/binary"

// DONTCOVER

const (
	ModuleName = "subspaces"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreateSubspace          = "create_subspace"
	ActionEditSubspace            = "edit_subspace"
	ActionCreateUserGroup         = "create_user_group"
	ActionDeleteUserGroup         = "delete_user_group"
	ActionAddUserToUserGroup      = "add_user_to_user_group"
	ActionRemoveUserFromUserGroup = "remove_user_from_user_group"
	ActionSetPermissions          = "set_permissions"

	QuerierRoute = ModuleName

	DoNotModify = "[do-not-modify]"
)

var (
	SubspacePrefix          = []byte{0x00}
	SubspaceIDKey           = []byte{0x01}
	ACLStorePrefix          = []byte{0x02}
	GroupsPrefix            = []byte{0x03}
	GroupMembersStorePrefix = []byte{0x4}
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

func ACLStoreKey(subspaceID uint64) []byte {
	return append(ACLStorePrefix, GetSubspaceIDBytes(subspaceID)...)
}

func PermissionStoreKey(subspaceID uint64, target string) []byte {
	return append(ACLStoreKey(subspaceID), []byte(target)...)
}

func GroupsStoreKey(subspaceID uint64) []byte {
	return append(GroupsPrefix, GetSubspaceIDBytes(subspaceID)...)
}

func GetGroupNameBytes(groupName string) []byte {
	return []byte(groupName)
}

func GetGroupNameFromBytes(bz []byte) string {
	return string(bz)
}

// GroupStoreKey returns the key used to store a group for a subspace
func GroupStoreKey(subspaceID uint64, groupName string) []byte {
	return append(GroupsStoreKey(subspaceID), GetGroupNameBytes(groupName)...)
}

func GroupMembersStoreKey(subspaceID uint64, groupName string) []byte {
	return append(append(GroupMembersStorePrefix, GetSubspaceIDBytes(subspaceID)...), GetGroupNameBytes(groupName)...)
}

func GroupMemberStoreKey(subspaceID uint64, groupName string, user string) []byte {
	return append(GroupStoreKey(subspaceID, groupName), []byte(user)...)
}
