package v2

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DONTCOVER

var (
	SubspacePrefix             = []byte{0x01}
	GroupIDPrefix              = []byte{0x02}
	GroupsPrefix               = []byte{0x03}
	GroupMembersStorePrefix    = []byte{0x04}
	UserPermissionsStorePrefix = []byte{0x05}
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

// PermissionsStoreKey returns the key used to store the entire ACL for a given subspace
func PermissionsStoreKey(subspaceID uint64) []byte {
	return append(UserPermissionsStorePrefix, GetSubspaceIDBytes(subspaceID)...)
}

func GetAddressBytes(user sdk.AccAddress) []byte {
	return user
}

func GetAddressFromBytes(bz []byte) sdk.AccAddress {
	return bz
}

// --------------------------------------------------------------------------------------------------------------------

// GroupIDStoreKey returns the store key that is used to store the group id to be used next for the given subspace
func GroupIDStoreKey(subspaceID uint64) []byte {
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

// GroupsStoreKey returns the key used to store all the groups of a given subspace
func GroupsStoreKey(subspaceID uint64) []byte {
	return append(GroupsPrefix, GetSubspaceIDBytes(subspaceID)...)
}

// GroupStoreKey returns the key used to store a group for a subspace
func GroupStoreKey(subspaceID uint64, groupID uint32) []byte {
	return append(GroupsStoreKey(subspaceID), GetGroupIDBytes(groupID)...)
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

// UserPermissionStoreKey returns the key used to store the permission for the given user inside the given subspace
func UserPermissionStoreKey(subspaceID uint64, user sdk.AccAddress) []byte {
	return append(PermissionsStoreKey(subspaceID), GetAddressBytes(user)...)
}

// SplitUserPermissionKey splits the provided UserPermissionStoreKey into the related subspace id and user address
func SplitUserPermissionKey(key []byte) (subspaceID uint64, user sdk.AccAddress) {
	key = key[1:]
	subspaceID = GetSubspaceIDFromBytes(key[:8])
	user = GetAddressFromBytes(key[8:])
	return subspaceID, user
}
