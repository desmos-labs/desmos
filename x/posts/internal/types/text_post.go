package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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
	PostID         PostID            `json:"id"`                      // Unique id
	ParentID       PostID            `json:"parent_id"`               // TextPost of which this one is a comment
	Message        string            `json:"message"`                 // Message contained inside the post
	Created        sdk.Int           `json:"created"`                 // Block height at which the post has been created
	LastEdited     sdk.Int           `json:"last_edited"`             // Block height at which the post has been edited the last time
	AllowsComments bool              `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string            `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   map[string]string `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress    `json:"creator"`                 // Creator of the TextPost
}

func NewTextPost(id, parentID PostID, message string, allowsComments bool, subspace string, optionalData map[string]string,
	created int64, creator sdk.AccAddress) TextPost {
	return TextPost{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        sdk.NewInt(created),
		LastEdited:     sdk.ZeroInt(),
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

func (p TextPost) CreationTime() sdk.Int {
	return p.Created
}

func (p TextPost) SetEditTime(time sdk.Int) Post {
	p.LastEdited = time
	return p
}

func (p TextPost) GetEditTime() sdk.Int {
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

	if len(strings.TrimSpace(p.Subspace)) == 0 {
		return fmt.Errorf("post subspace must be non empty and non blank")
	}

	if sdk.ZeroInt().Equal(p.Created) {
		return fmt.Errorf("invalid post creation block height: %s", p.Created)
	}

	if p.LastEdited.GT(sdk.ZeroInt()) && p.Created.GT(p.LastEdited) {
		return fmt.Errorf("invalid post last edit block height: %s", p.LastEdited)
	}

	if len(p.OptionalData) > MaxOptionalDataFieldsNumber {
		return fmt.Errorf("post optional data cannot contain more than 10 key-value pairs")
	}

	for key, value := range p.OptionalData {
		if len(value) > MaxOptionalDataFieldValueLength {
			return fmt.Errorf("post optional data values cannot exceed 200 characters. %s of post with id %s is longer than this", key, p.PostID)
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

// MarshalJSON implements Marshaler
func (p TextPost) MarshalJSON() ([]byte, error) {
	type textPostJSON TextPost
	return json.Marshal(textPostJSON(p))
}

// UnmarshalJSON implements Unmarshaler
func (p *TextPost) UnmarshalJSON(data []byte) error {
	type textPostJSON TextPost
	var temp textPostJSON
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	*p = TextPost(temp)
	return nil
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
