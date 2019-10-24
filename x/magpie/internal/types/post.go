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

// PostId represents a unique post id
type PostId uint64

func (id PostId) Valid() bool {
	return id != 0
}

func (id PostId) Next() PostId {
	return id + 1
}

func (id PostId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func ParsePostId(value string) (PostId, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return PostId(0), err
	}

	return PostId(intVal), err
}

// ---------------
// --- Post
// ---------------

// Post is a struct of a Magpie post
type Post struct {
	Id            PostId         `json:"id"`
	ParentId      PostId         `json:"parent_id"`
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
	return strings.TrimSpace(fmt.Sprintf(`Id: %s
Parent Id: %s
Owner: %s
Message: %s
Created: %s
Modified: %s
Likes: %d
Namespace: %s
External Onwer: %s`, p.Id, p.ParentId, p.Owner, p.Message, p.Created, p.Modified, p.Likes, p.Namespace, p.ExternalOwner))
}

func (p Post) Validate() error {
	if !p.Id.Valid() {
		return fmt.Errorf("invalid post id: %s", p.Id)
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
