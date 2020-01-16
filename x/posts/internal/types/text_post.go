package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- TextPost id
// ---------------

// PostID represents a unique post id
type PostID uint64

// Valid tells if the id can be used safely
func (id PostID) Valid() bool {
	return id != 0
}

// Next returns the subsequent id to this one
func (id PostID) Next() PostID {
	return id + 1
}

// String implements fmt.Stringer
func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// checkPostsEqual compares two PostID instances
func (id PostID) Equals(other PostID) bool {
	return id == other
}

// MarshalJSON implements Marshaler
func (id PostID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON implements Unmarshaler
func (id *PostID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	postID, err := ParsePostID(s)
	if err != nil {
		return err
	}

	*id = postID
	return nil
}

// ParsePostID returns the PostID represented inside the provided
// value, or an error if no id could be parsed properly
func ParsePostID(value string) (PostID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return PostID(0), err
	}

	return PostID(intVal), err
}

// ----------------
// --- TextPost IDs
// ----------------

// PostIDs represents a slice of PostID objects
type PostIDs []PostID

// checkPostsEqual returns true iff the ids slice and the other
// one contain the same data in the same order
func (ids PostIDs) Equals(other PostIDs) bool {
	if len(ids) != len(other) {
		return false
	}

	for index, id := range ids {
		if id != other[index] {
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
		if ele.Equals(id) {
			return ids, false
		}
	}
	return append(ids, id), true
}

// ---------------
// --- TextPost
// ---------------

// TextPost is a struct of a post
type TextPost struct {
	PostID         PostID         `json:"id"`                      // Unique id
	ParentID       PostID         `json:"parent_id"`               // Post of which this one is a comment
	Message        string         `json:"message"`                 // Message contained inside the post
	Created        time.Time      `json:"created"`                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited"`             // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator"`                 // Creator of the Post
}

func NewTextPost(id, parentID PostID, message string, allowsComments bool, subspace string, optionalData map[string]string,
	created time.Time, creator sdk.AccAddress) TextPost {
	return TextPost{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        created,
		LastEdited:     time.Time{},
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        creator,
	}
}

func NewTextPostComplete(id, parentID PostID, message string, created, lastEdited time.Time, allowsComments bool,
	subspace string, optionalData map[string]string, creator sdk.AccAddress) TextPost {
	return TextPost{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        created,
		LastEdited:     lastEdited,
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Creator:        creator,
	}
}

// String implements fmt.Stringer
func (p TextPost) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// GetID implements Post GetID
func (p TextPost) GetID() PostID {
	return p.PostID
}

// GetParentID implements Post GetParentID
func (p TextPost) GetParentID() PostID {
	return p.ParentID
}

func (p TextPost) SetMessage(message string) Post {
	p.Message = message
	return p
}

func (p TextPost) GetMessage() string {
	return p.Message
}

func (p TextPost) CreationTime() time.Time {
	return p.Created
}

func (p TextPost) SetEditTime(time time.Time) Post {
	p.LastEdited = time
	return p
}

func (p TextPost) GetEditTime() time.Time {
	return p.LastEdited
}

func (p TextPost) CanComment() bool {
	return p.AllowsComments
}

func (p TextPost) GetSubspace() string {
	return p.Subspace
}

func (p TextPost) GetOptionalData() map[string]string {
	return p.OptionalData
}

func (p TextPost) Owner() sdk.AccAddress {
	return p.Creator
}

// Validate implements Post Validate
func (p TextPost) Validate() error {
	if !p.PostID.Valid() {
		return fmt.Errorf("invalid post id: %s", p.PostID)
	}

	if p.Creator == nil {
		return fmt.Errorf("invalid post owner: %s", p.Creator)
	}

	if len(strings.TrimSpace(p.Message)) == 0 {
		return fmt.Errorf("post message must be non empty and non blank")
	}

	if len(p.Message) > MaxPostMessageLength {
		return fmt.Errorf("post message cannot be longer than %d characters", MaxPostMessageLength)
	}

	if len(strings.TrimSpace(p.Subspace)) == 0 {
		return fmt.Errorf("post subspace must be non empty and non blank")
	}

	if p.Created.IsZero() {
		return fmt.Errorf("invalid post creation time: %s", p.Created)
	}

	if p.Created.After(time.Now().UTC()) {
		return fmt.Errorf("post creation date cannot be in the future")
	}

	if !p.LastEdited.IsZero() && p.LastEdited.Before(p.Created) {
		return fmt.Errorf("invalid post last edit time: %s", p.LastEdited)
	}

	if !p.LastEdited.IsZero() && p.LastEdited.After(time.Now().UTC()) {
		return fmt.Errorf("post last edit date cannot be in the future")
	}

	if len(p.OptionalData) > MaxOptionalDataFieldsNumber {
		return fmt.Errorf("post optional data cannot contain more than %d key-value pairs", MaxOptionalDataFieldsNumber)
	}

	for key, value := range p.OptionalData {
		if len(value) > MaxOptionalDataFieldValueLength {
			return fmt.Errorf(
				"post optional data values cannot exceed %d characters. %s of post with id %s is longer than this",
				MaxOptionalDataFieldValueLength, key, p.PostID,
			)
		}
	}

	return nil
}

// Equals implements Post Equals
func (p TextPost) Equals(other Post) bool {
	// Cast and delegate
	if otherPost, ok := other.(TextPost); ok {
		return checkPostsEqual(p, otherPost)
	}

	return false
}

func checkPostsEqual(first TextPost, second TextPost) bool {
	equalsOptionalData := len(first.OptionalData) == len(second.OptionalData)
	if equalsOptionalData {
		for key := range first.OptionalData {
			equalsOptionalData = equalsOptionalData && first.OptionalData[key] == second.OptionalData[key]
		}
	}

	return first.PostID.Equals(second.PostID) &&
		first.ParentID.Equals(second.ParentID) &&
		first.Message == second.Message &&
		first.Created.Equal(second.Created) &&
		first.LastEdited.Equal(second.LastEdited) &&
		first.AllowsComments == second.AllowsComments &&
		first.Subspace == second.Subspace &&
		equalsOptionalData &&
		first.Creator.Equals(second.Creator)
}

// -------------
// --- TextPosts
// -------------

// TextPosts represents a slice of TextPost objects
type TextPosts []TextPost

// checkPostsEqual returns true iff the p slice contains the same
// data in the same order of the other slice
func (p TextPosts) Equals(other TextPosts) bool {
	if len(p) != len(other) {
		return false
	}

	for index, post := range p {
		if !post.Equals(other[index]) {
			return false
		}
	}

	return true
}

// String implements stringer interface
func (p TextPosts) String() string {
	out := "ID - [Creator] Message\n"
	for _, post := range p {
		out += fmt.Sprintf("%d - [%s] %s\n",
			post.PostID, post.Creator, post.Message)
	}
	return strings.TrimSpace(out)
}
