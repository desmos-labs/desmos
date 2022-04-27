package types

import (
	"encoding/binary"

	subspacetypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName
)

var (
	PostIDPrefix       = []byte{0x00}
	PostPrefix         = []byte{0x01}
	AttachmentIDPrefix = []byte{0x02}
)

// GetPostIDBytes returns the byte representation of the postID
func GetPostIDBytes(postID uint64) (postIDBz []byte) {
	postIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(postIDBz, postID)
	return
}

// GetPostIDFromBytes returns postID in uint64 format from a byte array
func GetPostIDFromBytes(bz []byte) (postID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// GetSubspacePostIDBytes returns the byte representation of the subspaceID merged with the postID
func GetSubspacePostIDBytes(subspaceID uint64, postID uint64) []byte {
	return append(subspacetypes.GetSubspaceIDBytes(subspaceID), GetPostIDBytes(postID)...)
}

// PostIDStoreKey returns the key uses to store the next post id for the given subspace
func PostIDStoreKey(subspaceID uint64) []byte {
	return append(PostIDPrefix, subspacetypes.GetSubspaceIDBytes(subspaceID)...)
}

// PostStoreKey returns the key for a specific post
func PostStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(PostPrefix, GetSubspacePostIDBytes(subspaceID, postID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetAttachmentIDBytes returns the byte representation of the attachmentID
func GetAttachmentIDBytes(attachmentID uint32) (attachmentIDBz []byte) {
	attachmentIDBz = make([]byte, 4)
	binary.BigEndian.PutUint32(attachmentIDBz, attachmentID)
	return
}

// GetAttachmentIDFromBytes returns the attachmentID in uint32 format from a byte array
func GetAttachmentIDFromBytes(bz []byte) (attachmentID uint32) {
	return binary.BigEndian.Uint32(bz)
}

// AttachmentIDStoreKey returns the store key that is used to store the attachment id to be used next for the given post
func AttachmentIDStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(AttachmentIDPrefix, GetSubspacePostIDBytes(subspaceID, postID)...)
}
