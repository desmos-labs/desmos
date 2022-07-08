package types

import (
	"encoding/binary"
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacetypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
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
	NextPostIDPrefix       = []byte{0x00}
	PostPrefix             = []byte{0x01}
	PostSectionPrefix      = []byte{0x02}
	NextAttachmentIDPrefix = []byte{0x10}
	AttachmentPrefix       = []byte{0x11}
	UserAnswerPrefix       = []byte{0x20}
	ActivePollQueuePrefix  = []byte{0x21}
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

// NextPostIDStoreKey returns the key uses to store the next post id for the given subspace
func NextPostIDStoreKey(subspaceID uint64) []byte {
	return append(NextPostIDPrefix, subspacetypes.GetSubspaceIDBytes(subspaceID)...)
}

// SubspacePostsPrefix returns the store prefix used to store all the posts related to the given subspace
func SubspacePostsPrefix(subspaceID uint64) []byte {
	return append(PostPrefix, subspacetypes.GetSubspaceIDBytes(subspaceID)...)
}

// PostStoreKey returns the key for a specific post
func PostStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(SubspacePostsPrefix(subspaceID), GetPostIDBytes(postID)...)
}

// SubspaceSectionsPrefix returns the prefix used to store all the section references for the given subspace
func SubspaceSectionsPrefix(subspaceID uint64) []byte {
	return append(PostSectionPrefix, subspacetypes.GetSubspaceIDBytes(subspaceID)...)
}

// SectionPostsPrefix returns the prefix used to store all the section references for the given section
func SectionPostsPrefix(subspaceID uint64, sectionID uint32) []byte {
	return append(SubspaceSectionsPrefix(subspaceID), subspacetypes.GetSectionIDBytes(sectionID)...)
}

// PostSectionStoreKey returns the key used to store the section reference for the given post
func PostSectionStoreKey(subspaceID uint64, sectionID uint32, postID uint64) []byte {
	return append(SectionPostsPrefix(subspaceID, sectionID), GetPostIDBytes(postID)...)
}

var (
	PostSectionPrefixLen = len(PostSectionPrefix)
	SubspaceIDLen        = len(subspacetypes.GetSubspaceIDBytes(1))
	SectionIDLen         = len(subspacetypes.GetSectionIDBytes(1))
	PostIDLen            = len(GetPostIDBytes(1))
)

func SplitPostSectionStoreKey(key []byte) (subspaceID uint64, sectionID uint32, postID uint64) {
	expectedLen := PostSectionPrefixLen + SubspaceIDLen + SectionIDLen + PostIDLen
	if len(key) != expectedLen {
		panic(fmt.Errorf("invalid key length; expected %d but got %d", expectedLen, len(key)))
	}

	key = key[PostSectionPrefixLen:]
	subspaceID = subspacetypes.GetSubspaceIDFromBytes(key[:SubspaceIDLen])
	sectionID = subspacetypes.GetSectionIDFromBytes(key[SubspaceIDLen : SubspaceIDLen+SectionIDLen])
	postID = GetPostIDFromBytes(key[SubspaceIDLen+SectionIDLen:])
	return subspaceID, sectionID, postID
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

// NextAttachmentIDStoreKey returns the store key that is used to store the attachment id to be used next for the given post
func NextAttachmentIDStoreKey(subspaceID uint64, postID uint64) []byte {
	return append(NextAttachmentIDPrefix, GetSubspacePostIDBytes(subspaceID, postID)...)
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

var (
	lenTime       = len(sdk.FormatTimeBytes(time.Now()))
	lenPrefix     = len(ActivePollQueuePrefix)
	lenSubspaceID = 64 / 8
	lenPostID     = 64 / 8
)

// ActivePollByTimeKey gets the active poll queue key by endTime
func ActivePollByTimeKey(endTime time.Time) []byte {
	return append(ActivePollQueuePrefix, sdk.FormatTimeBytes(endTime)...)
}

// ActivePollQueueKey returns the key for a pollID in the activePollQueue
func ActivePollQueueKey(subspaceID uint64, postID uint64, pollID uint32, endTime time.Time) []byte {
	return append(ActivePollByTimeKey(endTime), GetPollIDBytes(subspaceID, postID, pollID)...)
}

// SplitActivePollQueueKey split the active poll key and returns the poll id and endTime
func SplitActivePollQueueKey(key []byte) (subspaceID uint64, postID uint64, pollID uint32, endTime time.Time) {
	if len(key[lenPrefix:]) != 20+lenTime {
		panic(fmt.Errorf("unexpected key length (%d ≠ %d)", len(key[1:]), lenTime+8))
	}

	endTime, err := sdk.ParseTimeBytes(key[lenPrefix : lenPrefix+lenTime])
	if err != nil {
		panic(err)
	}

	subspaceID = subspacetypes.GetSubspaceIDFromBytes(key[lenPrefix+lenTime : lenPrefix+lenTime+lenSubspaceID])
	postID = GetPostIDFromBytes(key[lenPrefix+lenTime+lenSubspaceID : lenPrefix+lenTime+lenSubspaceID+lenPostID])
	pollID = GetAttachmentIDFromBytes(key[lenPrefix+lenTime+lenSubspaceID+lenPostID:])
	return subspaceID, postID, pollID, endTime
}

// GetPollIDBytes returns the byte representation of the provided pollID
func GetPollIDBytes(subspaceID uint64, postID uint64, pollID uint32) []byte {
	return append(GetSubspacePostIDBytes(subspaceID, postID), GetAttachmentIDBytes(pollID)...)
}

// GetPollIDFromBytes returns the pollID in uint32 format from a byte array
func GetPollIDFromBytes(bz []byte) (subspaceID uint64, postID uint64, pollID uint32) {
	if len(bz) != 20 {
		panic(fmt.Errorf("unexpected key length (%d ≠ %d", len(bz), 20))
	}

	subspaceID = subspacetypes.GetSubspaceIDFromBytes(bz[:lenSubspaceID])
	postID = GetPostIDFromBytes(bz[lenSubspaceID : lenSubspaceID+lenPostID])
	pollID = GetAttachmentIDFromBytes(bz[lenSubspaceID+lenPostID:])
	return subspaceID, postID, pollID
}

// PollAnswersPrefix returns the store prefix used to store the polls associated with the given post
func PollAnswersPrefix(subspaceID uint64, postID uint64, pollID uint32) []byte {
	return append(UserAnswerPrefix, GetPollIDBytes(subspaceID, postID, pollID)...)
}

// PollAnswerStoreKey returns the store key used to store the poll answer for the given user
func PollAnswerStoreKey(subspaceID uint64, postID uint64, pollID uint32, user string) []byte {
	return append(PollAnswersPrefix(subspaceID, postID, pollID), []byte(user)...)
}
