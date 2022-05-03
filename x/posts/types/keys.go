package types

import (
	"encoding/binary"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacetypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

// DONTCOVER

const (
	ModuleName   = "posts"
	RouterKey    = ModuleName
	StoreKey     = ModuleName
	QuerierRoute = ModuleName

	ActionCreatePost           = "create_post"
	ActionEditPost             = "edit_post"
	ActionAddPostAttachment    = "add_post_attachment"
	ActionRemovePostAttachment = "remove_post_attachment"
	ActionDeletePost           = "delete_post"
	ActionAnswerPoll           = "answer_poll"

	DoNotModify = "[do-not-modify]"
)

var (
	PostIDPrefix       = []byte{0x00}
	PostPrefix         = []byte{0x01}
	AttachmentIDPrefix = []byte{0x02}
	AttachmentPrefix   = []byte{0x03}
	PollAnswerPrefix   = []byte{0x04}

	ActivePollQueuePrefix   = []byte{0x10}
	InactivePollQueuePrefix = []byte{0x11}
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

// SubspacePostsPrefix returns the store prefix used to store all the posts related to the given subspace
func SubspacePostsPrefix(subspaceID uint64) []byte {
	return append(PostPrefix, subspacetypes.GetSubspaceIDBytes(subspaceID)...)
}

// PostStoreKey returns the key for a specific post
func PostStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(SubspacePostsPrefix(subspaceID), GetPostIDBytes(postID)...)
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

var lenTime = len(sdk.FormatTimeBytes(time.Now()))

// ActivePollByTimeKey gets the active poll queue key by endTime
func ActivePollByTimeKey(endTime time.Time) []byte {
	return append(ActivePollQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

// ActivePollQueueKey returns the key for a pollID in the activePollQueue
func ActivePollQueueKey(subspaceID uint64, postID uint64, pollID uint32, endTime time.Time) []byte {
	return append(ActivePollByTimeKey(endTime), GetPollIDBytes(subspaceID, postID, pollID)...)
}

// SplitActivePollQueueKey split the active proposal key and returns the poll id and endTime
func SplitActivePollQueueKey(key []byte) (subspaceID uint64, postID uint64, pollID uint32, endTime time.Time) {
	if len(key[1:]) != 20+lenTime {
		panic(fmt.Sprintf("unexpected key length (%d ≠ %d)", len(key[1:]), lenTime+8))
	}

	endTime, err := sdk.ParseTimeBytes(key[1 : 1+lenTime])
	if err != nil {
		panic(err)
	}

	subspaceID = subspacetypes.GetSubspaceIDFromBytes(key[1+lenTime : 1+lenTime+8])
	postID = GetPostIDFromBytes(key[1+8+lenTime : 1+16+lenTime])
	pollID = GetAttachmentIDFromBytes(key[1+16+lenTime:])
	return subspaceID, postID, pollID, endTime
}

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
