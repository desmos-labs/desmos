package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewPost allows to build a new Post instance with the provided data
func NewPost(
	postID string, parentID string, message string, allowsComments bool, subspace string,
	optionalData OptionalData, attachments []Attachment, pollData *PollData,
	lastEdited time.Time, created time.Time, creator string,
) Post {
	return Post{
		PostID:         postID,
		ParentID:       parentID,
		Message:        message,
		Created:        created,
		LastEdited:     lastEdited,
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Attachments:    attachments,
		PollData:       pollData,
		Creator:        creator,
	}
}

// Validate implements validator
func (post Post) Validate() error {
	if !IsValidPostID(post.PostID) {
		return fmt.Errorf("invalid post id: %s", post.PostID)
	}

	if post.PostID == post.ParentID {
		return fmt.Errorf("post id and parent id cannot be the same")
	}

	if len(strings.TrimSpace(post.ParentID)) != 0 && !IsValidPostID(post.ParentID) {
		return fmt.Errorf("invalid parent id: %s", post.ParentID)
	}

	if post.Creator == "" {
		return fmt.Errorf("invalid post owner: %s", post.Creator)
	}

	if len(strings.TrimSpace(post.Message)) == 0 && len(post.Attachments) == 0 && post.PollData == nil {
		return fmt.Errorf("post message, attachments or poll required, they cannot be all empty")
	}

	if !commons.IsValidSubspace(post.Subspace) {
		return fmt.Errorf("post subspace must be a valid sha-256 hash")
	}

	if post.Created.IsZero() {
		return fmt.Errorf("invalid post creation time: %s", post.Created)
	}

	if !post.LastEdited.IsZero() && post.LastEdited.Before(post.Created) {
		return fmt.Errorf("invalid post last edit time: %s", post.LastEdited)
	}

	for _, attachment := range post.Attachments {
		err := attachment.Validate()
		if err != nil {
			return err
		}
	}

	if post.PollData != nil {
		err := post.PollData.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// ContentsEquals returns true if and only if p and other contain the same poll, without considering the ID
func (post Post) ContentsEquals(other Post) bool {
	equalsOptionalData := len(post.OptionalData) == len(other.OptionalData)
	if equalsOptionalData {
		for index, opd := range post.OptionalData {
			equalsOptionalData = equalsOptionalData && opd == other.OptionalData[index]
		}
	}

	return post.ParentID == other.ParentID &&
		post.Message == other.Message &&
		post.Created.Equal(other.Created) &&
		post.LastEdited.Equal(other.LastEdited) &&
		post.AllowsComments == other.AllowsComments &&
		post.Subspace == other.Subspace &&
		equalsOptionalData &&
		post.Creator == other.Creator &&
		post.Attachments.Equal(other.Attachments) &&
		post.PollData.Equal(other.PollData)
}

// GetPostHashtags returns all the post's hashtags without duplicates
func (post Post) GetPostHashtags() []string {
	hashtags := GetTags(post.Message)

	uniqueHashtags := commons.Unique(hashtags)
	withoutHashtag := make([]string, len(uniqueHashtags))

	for index, hashtag := range uniqueHashtags {
		trimmed := strings.TrimLeft(strings.TrimSpace(hashtag), "#")
		if !IsNumeric(trimmed) {
			withoutHashtag[index] = trimmed
		} else {
			withoutHashtag = []string{}
		}
	}

	if len(withoutHashtag) == 0 {
		return []string{}
	}

	return withoutHashtag
}

// ___________________________________________________________________________________________________________________

// AppendIfMissing appends the given id to the ids slice, if not present yet.
// If appended, returns the new slice and true. Otherwise, returns the original slice and false.
func (ids CommentIDs) AppendIfMissing(id string) (CommentIDs, bool) {
	for _, existing := range ids.Ids {
		if existing == id {
			return ids, false
		}
	}
	return CommentIDs{Ids: append(ids.Ids, id)}, true
}

// ___________________________________________________________________________________________________________________

// NewAttachment builds a new Attachment instance with the provided data
func NewAttachment(uri, mimeType string, tags []string) Attachment {
	return Attachment{
		URI:      uri,
		MimeType: mimeType,
		Tags:     tags,
	}
}

// Validate implements validator
func (attachments Attachment) Validate() error {
	if !commons.IsURIValid(attachments.URI) {
		return fmt.Errorf("invalid uri provided")
	}

	if len(strings.TrimSpace(attachments.MimeType)) == 0 {
		return fmt.Errorf("mime type must be specified and cannot be empty")
	}

	for _, address := range attachments.Tags {
		if address == "" {
			return fmt.Errorf("invalid empty tag address: %s", address)
		}
	}

	for _, tag := range attachments.Tags {
		_, err := sdk.AccAddressFromBech32(tag)
		if err != nil {
			return err
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// Attachments represents a slice of Attachment object
type Attachments []Attachment

// NewAttachments builds a new Attachments from the given attachments
func NewAttachments(attachments ...Attachment) Attachments {
	return attachments
}

// Equals returns true iff the atts slice contains the same data in the same order of the other slice
func (attachments Attachments) Equal(other Attachments) bool {
	if len(attachments) != len(other) {
		return false
	}

	for index, attachment := range attachments {
		if !attachment.Equal(other[index]) {
			return false
		}
	}

	return true
}

// AppendIfMissing appends the given otherAttachment to the atts slice if it does not exist inside it yet.
// It returns a new slice of Attachments containing such otherAttachment.
func (attachments Attachments) AppendIfMissing(otherAttachment Attachment) Attachments {
	for _, att := range attachments {
		if att.Equal(otherAttachment) {
			return attachments
		}
	}
	return append(attachments, otherAttachment)
}

// ___________________________________________________________________________________________________________________

type OptionalData []OptionalDataEntry

// NewOptionalDataEntry returns a new OptionalDataEntry object
func NewOptionalDataEntry(key, value string) OptionalDataEntry {
	return OptionalDataEntry{
		Key:   key,
		Value: value,
	}
}
