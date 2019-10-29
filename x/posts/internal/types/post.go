package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

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
	Created       time.Time      `json:"created"`
	Modified      time.Time      `json:"modified"`
	Likes         uint           `json:"likes"`
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
	return strings.TrimSpace(fmt.Sprintf(`PostID: %s
Parent PostID: %s
Owner: %s
Message: %s
Created: %s
Modified: %s
Likes: %d
Namespace: %s
External Onwer: %s`, p.PostID, p.ParentID, p.Owner, p.Message, p.Created, p.Modified, p.Likes, p.Namespace, p.ExternalOwner))
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

	if p.Created.String() == "" {
		return fmt.Errorf("invalid post creation time: %s", p.Created)
	}

	if p.Modified.String() == "" {
		return fmt.Errorf("invalid Post edit time %s", p.Modified)
	}

	if p.Namespace == "" {
		return fmt.Errorf("invalid post namespace: %s", p.Namespace)
	}

	if p.ExternalOwner == "" {
		return fmt.Errorf("invalid post external owner: %s", p.ExternalOwner)
	}

	return nil
}
