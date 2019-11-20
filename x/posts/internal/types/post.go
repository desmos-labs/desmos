package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Post id
// ---------------

// PostID represents a unique post id
type PostID uint64

func (id PostID) Valid() bool {
	return id != 0
}

func (id PostID) Next() PostID {
	return id + 1
}

func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

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

// Post is a struct of a Magpie post
type Post struct {
	PostID         PostID         `json:"id,string"`
	ParentID       PostID         `json:"parent_id,string"`
	Message        string         `json:"message"`
	Created        sdk.Int        `json:"created"`     // Block height at which the post has been created
	LastEdited     sdk.Int        `json:"last_edited"` // Block height at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments"`
	Owner          sdk.AccAddress `json:"owner"`
}

func NewPost(id, parentID PostID, message string, allowsComments bool, created int64, owner sdk.AccAddress) Post {
	return Post{
		PostID:         id,
		ParentID:       parentID,
		Message:        message,
		Created:        sdk.NewInt(created),
		LastEdited:     sdk.ZeroInt(),
		AllowsComments: allowsComments,
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

	if p.Message == "" {
		return fmt.Errorf("invalid post message: %s", p.Message)
	}

	if p.Created.Equal(sdk.ZeroInt()) {
		return fmt.Errorf("invalid post creation block heigth: %s", p.Created)
	}

	if p.LastEdited.Equal(sdk.ZeroInt()) || p.LastEdited.LT(p.Created) {
		return fmt.Errorf("invalid Post edit time %s", p.LastEdited)
	}

	return nil
}

func (p Post) Equals(other Post) bool {
	return p.PostID == other.PostID &&
		p.ParentID == other.ParentID &&
		p.Message == other.Message &&
		p.Created == other.Created &&
		p.LastEdited == other.LastEdited &&
		p.AllowsComments == other.AllowsComments &&
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
