package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Post id
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

// Equals compares two PostID instances
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
// --- Post IDs
// ----------------

// PostIDs represents a slice of PostID objects
type PostIDs []PostID

// Equals returns true iff the ids slice and the other
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

// ---------------
// --- Post
// ---------------

// Post is a struct of a post
type Post struct {
	PostID         PostID            `json:"id"`                      // Unique id
	ParentID       PostID            `json:"parent_id"`               // Post of which this one is a comment
	Message        string            `json:"message"`                 // Message contained inside the post
	Created        sdk.Int           `json:"created"`                 // Block height at which the post has been created
	LastEdited     sdk.Int           `json:"last_edited"`             // Block height at which the post has been edited the last time
	AllowsComments bool              `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string            `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   map[string]string `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Owner          sdk.AccAddress    `json:"owner"`                   // Creator of the Post
}

func NewPost(id, parentID PostID, message string, allowsComments bool, subspace string, optionalData map[string]string,
	created int64, owner sdk.AccAddress) Post {
	return Post{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        sdk.NewInt(created),
		LastEdited:     sdk.ZeroInt(),
		AllowsComments: allowsComments,
		Subspace:       subspace,
		OptionalData:   optionalData,
		Owner:          owner,
	}
}

// String implements fmt.Stringer
func (p Post) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// Validate implements validator
func (p Post) Validate() error {
	if !p.PostID.Valid() {
		return fmt.Errorf("invalid post id: %s", p.PostID)
	}

	if p.Owner == nil {
		return fmt.Errorf("invalid post owner: %s", p.Owner)
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

func (p Post) Equals(other Post) bool {
	equalsOptionalData := len(p.OptionalData) == len(other.OptionalData)
	if equalsOptionalData {
		for key := range p.OptionalData {
			equalsOptionalData = equalsOptionalData && p.OptionalData[key] == other.OptionalData[key]
		}
	}

	return p.PostID.Equals(other.PostID) &&
		p.ParentID.Equals(other.ParentID) &&
		p.Message == other.Message &&
		p.Created.Equal(other.Created) &&
		p.LastEdited.Equal(other.LastEdited) &&
		p.Subspace == other.Subspace &&
		equalsOptionalData &&
		p.AllowsComments == other.AllowsComments &&
		p.ExternalReference == other.ExternalReference &&
		p.Owner.Equals(other.Owner)
}

// -------------
// --- Posts
// -------------

// Posts represents a slice of Post objects
type Posts []Post

// Equals returns true iff the p slice contains the same
// data in the same order of the other slice
func (p Posts) Equals(other Posts) bool {
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
func (p Posts) String() string {
	out := "ID - [Creator] Message\n"
	for _, post := range p {
		out += fmt.Sprintf("%d - [%s] %s\n",
			post.PostID, post.Owner, post.Message)
	}
	return strings.TrimSpace(out)
}
