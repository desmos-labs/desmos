package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Like id
// ---------------

// LikeId represents a unique like id
type LikeId uint64

func (id LikeId) Valid() bool {
	return id != 0
}

func (id LikeId) Next() LikeId {
	return id + 1
}

func (id LikeId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func ParseLikeId(value string) (LikeId, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return LikeId(0), err
	}

	return LikeId(intVal), err
}

// ---------------
// --- Like
// ---------------

// Like is a struct of a user like
type Like struct {
	LikeID        LikeId         `json:"id"`
	PostID        PostId         `json:"post_id"`
	Created       time.Time      `json:"created"`
	Owner         sdk.AccAddress `json:"owner"`
	Namespace     string         `json:"namespace"`
	ExternalOwner string         `json:"external_owner"`
}

// NewLike returns an empty Like
func NewLike() Like {
	return Like{}
}

// implement fmt.Stringer
func (l Like) String() string {
	return strings.TrimSpace(fmt.Sprintf(`PostID: %s
Owner: %s
PostID: %s
Created: %s
Namespace: %s
External Owner: %s`, l.LikeID, l.Owner, l.PostID, l.Created, l.Namespace, l.ExternalOwner))
}

func (l Like) Validate() error {
	if !l.LikeID.Valid() {
		return fmt.Errorf("invalid like id %s", l.LikeID)
	}

	if l.Owner == nil {
		return fmt.Errorf("invalid like owner: %s", l.Owner)
	}

	if l.Created.String() == "" {
		return fmt.Errorf("invalid like creation time: %s", l.Created)
	}

	if !l.PostID.Valid() {
		return fmt.Errorf("invalid like post id: %s", l.PostID)
	}

	return nil
}
