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

// ---------------
// --- Post
// ---------------

// Post is a struct of a Magpie post
type Post struct {
	PostID     PostID         `json:"id"`
	ParentID   PostID         `json:"parent_id"`
	Message    string         `json:"message"`
	Created    int64          `json:"created"`     // Block height at which the post has been created
	LastEdited int64          `json:"last_edited"` // Block height at which the post has been edited the last time
	Owner      sdk.AccAddress `json:"owner"`
}

func NewPost(ID, parentID PostID, message string, created int64, owner sdk.AccAddress) Post {
	return Post{
		PostID:     ID,
		ParentID:   parentID,
		Message:    message,
		Created:    created,
		LastEdited: 0,
		Owner:      owner,
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

	if p.Created == 0 {
		return fmt.Errorf("invalid post creation block heigth: %d", p.Created)
	}

	if p.LastEdited == 0 || p.LastEdited < p.Created {
		return fmt.Errorf("invalid Post edit time %d", p.LastEdited)
	}

	return nil
}

func (p Post) Equals(other Post) bool {
	return p.PostID == other.PostID &&
		p.ParentID == other.ParentID &&
		p.Message == other.Message &&
		p.Created == other.Created &&
		p.LastEdited == other.LastEdited &&
		p.Owner.Equals(other.Owner)
}

// -------------
// --- Posts
// -------------

type Posts []Post

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
