package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResPost is a result struct
type QueryResPost struct {
	ID      string         `json:"ID"`
	Message string         `json:"message"`
	Owner   sdk.AccAddress `json:"owner"`
	Time    time.Time      `json:"time"`
	Likes   uint           `json:"likes"`
}

func (r QueryResPost) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
Message: %s
Time: %s
Likes: %d`, r.ID, r.Owner, r.Message, r.Time, r.Likes))
}

// QueryResLike is a result struct
type QueryResLike struct {
	ID     string         `json:"ID"`
	PostID string         `json:"post_id"`
	Owner  sdk.AccAddress `json:"owner"`
	Time   time.Time      `json:"time"`
}

func (r QueryResLike) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
PostID: %s`, r.ID, r.Owner, r.PostID))
}

// // Query Result Payload for a resolve query
// type QueryResResolve struct {
// 	Value string `json:"value"`
// }

// // implement fmt.Stringer
// func (r QueryResResolve) String() string {
// 	return r.Value
// }

// // Query Result Payload for a names query
// type QueryResNames []string

// // implement fmt.Stringer
// func (n QueryResNames) String() string {
// 	return strings.Join(n[:], "\n")
// }
