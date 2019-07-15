package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// QueryResPost is a result struct
type QueryResPost struct {
	ID            string         `json:"ID"`
	Message       string         `json:"message"`
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Modified      time.Time      `jsond:"modified"`
	Likes         uint           `json:"likes"`
	Namespace     string         `json:"namespace"`
	ExternalOwner sdk.AccAddress `json:"external_owner"`
}

func (r QueryResPost) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
Message: %s
Created: %s
Modified: %s
Likes: %d
Namespace: %s
ExternalOwner: %s`, r.ID, r.Owner, r.Message, r.Created, r.Modified, r.Likes, r.Namespace, r.ExternalOwner))
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

// QueryResSession is a result struct
type QueryResSession struct {
	ID            string         `json:"ID"`
	Owner         sdk.AccAddress `json:"owner"`
	Created       time.Time      `json:"created"`
	Expiry        time.Time      `json:"Expiry"`
	Namespace     string         `json:"namespace"`
	ExternalOwner sdk.AccAddress `json:"external_owner"`
}

func (r QueryResSession) String() string {
	return strings.TrimSpace(fmt.Sprintf(`ID: %s
Owner: %s
Created: %s
Expiry: %s
Namesapce: %s
External Owner: %s`, r.ID, r.Owner, r.Created, r.Expiry, r.Namespace, r.ExternalOwner))
}
