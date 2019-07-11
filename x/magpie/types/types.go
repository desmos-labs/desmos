package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Post is a struct of a Dwiiter post
type Post struct {
	ID      string         `json:"id"`
	Message string         `json:"message"`
	Time    time.Time      `json:"time"`
	Likes   uint           `json:"likes"`
	Owner   sdk.AccAddress `json:"owner"`
}

// NewPost returns an empty Magpie post
func NewPost() Post {
	return Post{}
}

// implement fmt.Stringer
func (p Post) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
Message: %s
Time: %s
Likes: %d`, p.ID, p.Owner, p.Message, p.Time, p.Likes))
}

// Like is a struct of a user like
type Like struct {
	ID     string         `json:"id"`
	PostID string         `json:"post_id"`
	Time   time.Time      `json:"time"`
	Owner  sdk.AccAddress `json:"owner"`
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
Time: %s`, l.ID, l.Owner, l.PostID, l.Time))
}
