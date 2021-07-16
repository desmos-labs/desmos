package types

import (
	"fmt"
	"strings"
	"time"

	subspacestypes "github.com/desmos-labs/desmos/x/staging/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/commons"
)

// NewPost allows to build a new Post instance with the provided data
func NewPost(
	postID string, parentID string, message string, commentsState CommentsState, subspace string,
	additionalAttributes []Attribute, attachments []Attachment, poll *Poll,
	lastEdited time.Time, created time.Time, creator string,
) Post {
	return Post{
		PostID:               postID,
		ParentID:             parentID,
		Message:              message,
		Created:              created,
		LastEdited:           lastEdited,
		CommentsState:        commentsState,
		Subspace:             subspace,
		AdditionalAttributes: additionalAttributes,
		Attachments:          attachments,
		Poll:                 poll,
		Creator:              creator,
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

	if len(strings.TrimSpace(post.Message)) == 0 && len(post.Attachments) == 0 && post.Poll == nil {
		return fmt.Errorf("post message, attachments or poll required, they cannot be all empty")
	}

	if !IsValidCommentsState(post.CommentsState) {
		return fmt.Errorf("invalid comments state: %s", post.CommentsState)
	}

	if !subspacestypes.IsValidSubspace(post.Subspace) {
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

	if post.Poll != nil {
		err := post.Poll.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// CommentsStateFromString convert a string in the corresponding CommentsState
func CommentsStateFromString(comState string) (CommentsState, error) {
	commentState, ok := CommentsState_value[comState]
	if !ok {
		return CommentsStateUnspecified, fmt.Errorf("'%s' is not a valid comments state", comState)
	}
	return CommentsState(commentState), nil
}

// NormalizeCommentsState - normalize user specified comments state
func NormalizeCommentsState(comState string) string {
	switch strings.ToLower(comState) {
	case "allowed":
		return CommentsStateAllowed.String()
	case "blocked":
		return CommentsStateBlocked.String()
	default:
		return comState
	}
}

// IsValidCommentsState checks if the commentsState given correspond to one of the valid ones
func IsValidCommentsState(commentsState CommentsState) bool {
	return commentsState == CommentsStateAllowed || commentsState == CommentsStateBlocked
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

// NewAttribute returns a new Attribute object
func NewAttribute(key, value string) Attribute {
	return Attribute{
		Key:   key,
		Value: value,
	}
}
