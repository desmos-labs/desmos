package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacetypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// DONTCOVER

const (
	ModuleName = "posts"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionCreatePost = "create_post"
	ActionEditPost   = "edit_post"
)

var (
	PostIDPrefix          = []byte{0x00}
	PostPrefix            = []byte{0x01}
	AttachmentIDPrefix    = []byte{0x02}
	AttachmentPrefix      = []byte{0x03}
	PollAnswerPrefix      = []byte{0x04}
	PollTallyResultPrefix = []byte{0x05}
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

// PostAttachmentsPrefix returns the store prefix used to store all the given post attachments
func PostAttachmentsPrefix(subspaceID uint64, postID uint64) []byte {
	return append(AttachmentPrefix, GetSubspacePostIDBytes(subspaceID, postID)...)
}

// AttachmentStoreKey returns the store key that is used to store the attachment having the given id
func AttachmentStoreKey(subspaceID uint64, postID uint64, attachmentID uint32) []byte {
	return append(PostAttachmentsPrefix(subspaceID, postID), GetAttachmentIDBytes(attachmentID)...)
}

// --------------------------------------------------------------------------------------------------------------------

// GetPollIDBytes returns the byte representation of the provided pollID
func GetPollIDBytes(subspaceID uint64, postID uint64, pollID uint32) []byte {
	return append(GetSubspacePostIDBytes(subspaceID, postID), GetAttachmentIDBytes(pollID)...)
}

// PollAnswersPrefix returns the store prefix used to store the polls associated with the given post
func PollAnswersPrefix(subspaceID uint64, postID uint64, pollID uint32) []byte {
	return append(PollAnswerPrefix, GetPollIDBytes(subspaceID, postID, pollID)...)
}

// PollAnswerStoreKey returns the store key used to store the poll answer for the given user
func PollAnswerStoreKey(subspaceID uint64, postID uint64, pollID uint32, user sdk.AccAddress) []byte {
	return append(PollAnswersPrefix(subspaceID, postID, pollID), user...)
}

// PollTallyResultsStoreKey returns the store key used to store the tally results for the given poll
func PollTallyResultsStoreKey(subspaceID uint64, postID uint64, pollID uint32) []byte {
	return append(PollTallyResultPrefix, GetPollIDBytes(subspaceID, postID, pollID)...)
}
