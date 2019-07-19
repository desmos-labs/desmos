package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Should a new namespace be registered before accepting posting?

// Post is a struct of a Magpie post
type Post struct {
	ID            string         `json:"id"`
	ParentID      string         `json:"parent_id"`
	Message       string         `json:"message"`
	Created       time.Time      `json:"created"`
	Modified      time.Time      `json:"modified"`
	Likes         uint           `json:"likes"`
	Owner         sdk.AccAddress `json:"owner"`
	Namespace     string         `json:"namespace"`      // External service namespace, e.g. cosmos
	ExternalOwner sdk.AccAddress `json:"external_owner"` // External owner address of the post
}

// NewPost returns an empty Magpie post
func NewPost() Post {
	return Post{}
}

// implement fmt.Stringer
func (p Post) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Parent ID: %s
Owner: %s
Message: %s
Created: %s
Modified: %s
Likes: %d
Namespace: %s
External Onwer: %s`, p.ID, p.ParentID, p.Owner, p.Message, p.Created, p.Modified, p.Likes, p.Namespace, p.ExternalOwner))
}

// Like is a struct of a user like
type Like struct {
	ID      string         `json:"id"`
	PostID  string         `json:"post_id"`
	Created time.Time      `json:"Created"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewLike returns an empty Like
func NewLike() Like {
	return Like{}
}

// implement fmt.Stringer
func (l Like) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
PostID: %s
Created: %s`, l.ID, l.Owner, l.PostID, l.Created))
}

// Session is a struct of a user session
type Session struct {
	ID            string         `json:"ID"`
	Owner         sdk.AccAddress `json:"onwer"`
	Created       time.Time      `json:"created"`
	Expiry        time.Time      `json:"expiry"`
	Namesapce     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
	Signature     string         `json:"signature"`
}

// NewSession return an empty Session
func NewSession() Session {
	return Session{}
}

// implement fmt.Stringer
func (s Session) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Created: %s
Expiry: %s
Namespace: %s
External Owner: %s
Signature: %s`, s.Owner, s.Created, s.Expiry, s.Namesapce, s.ExternalOwner, s.Signature))
}
