package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResPost is a result struct
type QueryResPost struct {
	ID            string         `json:"id"`
	ParentID      string         `json:"parent_id"`
	Message       string         `json:"message"`
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Modified      time.Time      `json:"modified"`
	Likes         uint           `json:"likes"`
	Namespace     string         `json:"namespace"`
	ExternalOwner sdk.AccAddress `json:"external_owner"`
}

func (r QueryResPost) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
ParentID: %s
Owner: %s
Message: %s
Created: %s
Modified: %s
Likes: %d
Namespace: %s
External Owner: %s`, r.ID, r.ParentID, r.Owner, r.Message, r.Created, r.Modified, r.Likes, r.Namespace, r.ExternalOwner))
}

// QueryResLike is a result struct
type QueryResLike struct {
	ID     string         `json:"id"`
	PostID string         `json:"post_id"`
	Owner  sdk.AccAddress `json:"owner"`
	Time   time.Time      `json:"time"`
}

func (r QueryResLike) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
PostID: %s`, r.ID, r.Owner, r.PostID))
}

// QueryResSession is a result struct
type QueryResSession struct {
	ID            string         `json:"id"`
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Expiry        time.Time      `json:"expiry"`
	Namespace     string         `json:"namespace"`
	ExternalOwner sdk.AccAddress `json:"external_owner"`
	Signature     string         `json:"signature"`
}

func (r QueryResSession) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Owner: %s
Created: %s
Expiry: %s
Namesapce: %s
External Owner: %s
Signature: %s`, r.Owner, r.Created, r.Expiry, r.Namespace, r.ExternalOwner, r.Signature))
}
