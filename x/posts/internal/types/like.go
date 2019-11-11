package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Like id
// ---------------

// LikeID represents a unique like id
type LikeID uint64

func (id LikeID) Valid() bool {
	return id != 0
}

func (id LikeID) Next() LikeID {
	return id + 1
}

func (id LikeID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

func ParseLikeID(value string) (LikeID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return LikeID(0), err
	}

	return LikeID(intVal), err
}

// ---------------
// --- Like
// ---------------

// Like is a struct of a user like
type Like struct {
	Created int64          `json:"created"` // Block height at which the like was created
	Owner   sdk.AccAddress `json:"owner"`
}

// NewLike returns an empty Like
func NewLike() Like {
	return Like{}
}

// implement fmt.Stringer
func (l Like) String() string {
	bytes, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (l Like) Validate() error {
	if l.Owner.Empty() {
		return fmt.Errorf("invalid like owner: %s", l.Owner)
	}

	if l.Created == 0 {
		return fmt.Errorf("invalid like creation block heigth: %d", l.Created)
	}

	return nil
}

func (l Like) Equals(other Like) bool {
	return l.Created == other.Created &&
		l.Owner.Equals(other.Owner)
}

// ------------
// --- Likes
// ------------

type Likes []Like

func (likes Likes) AppendIfMissing(like Like) (Likes, bool) {
	for _, like := range likes {
		if like.Equals(like) {
			return likes, false
		}
	}
	return append(likes, like), true
}

func (likes Likes) ContainsOwnerLike(owner sdk.AccAddress) bool {
	for _, like := range likes {
		if like.Owner.Equals(owner) {
			return true
		}
	}
	return false
}
