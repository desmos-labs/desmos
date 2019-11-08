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
	PostID        PostID         `json:"id"`
	ParentID      PostID         `json:"parent_id"`
	Message       string         `json:"message"`
	Created       int64          `json:"created"`     // Block height at which the post has been created
	LastEdited    int64          `json:"last_edited"` // Block height at which the post has been edited the last time
	Owner         sdk.AccAddress `json:"owner"`
	Namespace     string         `json:"namespace"`      // External service namespace, e.g. cosmos
	ExternalOwner string         `json:"external_owner"` // External owner address of the post
}

// NewPost returns an empty Magpie post
func NewPost() Post {
	return Post{}
}

// implement fmt.Stringer
func (p Post) String() string {
	bytes, err := json.Marshal(&p)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

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

	if p.Namespace == "" {
		return fmt.Errorf("invalid post namespace: %s", p.Namespace)
	}

	if p.ExternalOwner == "" {
		return fmt.Errorf("invalid post external owner: %s", p.ExternalOwner)
	}

	return nil
}
