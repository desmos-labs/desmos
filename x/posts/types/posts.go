package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/commons"
	"strings"
	"time"
)

// NewPostID returns a new PostID instance
func NewPostID(value string) PostID {
	return PostID{Id: value}
}

// Valid tells if the id can be used safely
func (id PostID) Valid() bool {
	return strings.TrimSpace(id.String()) != "" && IsValidPostID(id.String())
}

// ___________________________________________________________________________________________________________________

// PostIDs represents a slice of PostID objects
type PostIDs []PostID

// Equals returns true iff the ids slice and the other
// one contain the same data in the same order
func (ids PostIDs) Equals(other PostIDs) bool {
	if len(ids) != len(other) {
		return false
	}

	for index, id := range ids {
		if !id.Equal(other[index]) {
			return false
		}
	}

	return true
}

// AppendIfMissing appends the given postID to the ids slice if it does not exist inside it yet.
// It returns a new slice of PostIDs containing such ID and a boolean indicating whether or not the original
// slice has been modified.
func (ids PostIDs) AppendIfMissing(id PostID) (PostIDs, bool) {
	for _, ele := range ids {
		if ele.Equal(id) {
			return ids, false
		}
	}
	return append(ids, id), true
}

// String implements fmt.Stringer
func (ids PostIDs) String() string {
	var stringIDs = make([]string, len(ids))
	for index, id := range ids {
		stringIDs[index] = id.String()
	}

	out := strings.Join(stringIDs, ", ")
	return fmt.Sprintf("[%s]", strings.TrimSpace(out))
}

// ___________________________________________________________________________________________________________________

func NewPost(
	parentID string, message string, allowsComments bool, subspace string,
	optionalData OptionalData, created time.Time, creator string,
) Post {
	post := Post{
		PostID:         "",
		ParentID:       parentID,
		Message:        message,
		Created:        created,
		LastEdited:     time.Time{},
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        creator,
	}

	// postID calculation
	post.PostID = ComputeID(post)

	return post
}

// WithAttachments allows to easily set the given attachments as the multimedia files associated with the p Post
func (post Post) WithAttachments(attachments []Attachment) Post {
	post.Attachments = attachments
	post.PostID = ComputeID(post)
	return post
}

// WithPollData allows to easily set the given data as the poll data files associated with the p Post
func (post Post) WithPollData(data PollData) Post {
	post.PollData = &data
	post.PostID = ComputeID(post)
	return post
}

// Validate implements validator
func (post Post) Validate() error {
	if !IsValidPostID(post.PostID) {
		return fmt.Errorf("invalid postID: %s", post.PostID)
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
		if err := post.PollData.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ContentsEquals returns true if and only if p and other contain the same data, without considering the ID
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

// Posts represents a slice of Post objects
type Posts []Post

// String implements stringer interface
func (posts Posts) String() string {
	out := "ID - [Creator] Message\n"
	for _, post := range posts {
		out += fmt.Sprintf("%s - [%s] %s\n",
			post.PostID, post.Creator, post.Message)
	}
	return strings.TrimSpace(out)
}

// Len implements sort.Interface
func (posts Posts) Len() int {
	return len(posts)
}

// Swap implements sort.Interface
func (posts Posts) Swap(i, j int) {
	posts[i], posts[j] = posts[j], posts[i]
}

// Less implements sort.Interface
func (posts Posts) Less(i, j int) bool {
	return posts[i].Created.Before(posts[j].Created)
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

// Equals returns true iff the atts slice contains the same
// data in the same order of the other slice
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
func (atts Attachments) AppendIfMissing(otherAttachment Attachment) Attachments {
	for _, att := range atts {
		if att.Equal(otherAttachment) {
			return atts
		}
	}
	return append(atts, otherAttachment)
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
