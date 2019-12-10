package types

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- Like
// ---------------

// Like is a struct of a user like
type Like struct {
	Created sdk.Int        `json:"created"` // Block height at which the like was created
	Owner   sdk.AccAddress `json:"owner"`
}

// NewLike returns an empty Like
func NewLike(created int64, owner sdk.AccAddress) Like {
	return Like{Created: sdk.NewInt(created), Owner: owner}
}

// String implements fmt.Stringer
func (l Like) String() string {
	bytes, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

// Validate implements validator
func (l Like) Validate() error {
	if l.Owner.Empty() {
		return fmt.Errorf("invalid like owner: %s", l.Owner)
	}

	if l.Created.Equal(sdk.ZeroInt()) {
		return fmt.Errorf("invalid like creation block height: %s", l.Created)
	}

	return nil
}

// Equals returns true if l and other contain the same data
func (l Like) Equals(other Like) bool {
	return l.Created == other.Created &&
		l.Owner.Equals(other.Owner)
}

// ------------
// --- Likes
// ------------

// Likes represents a slice of Like objects
type Likes []Like

// AppendIfMissing returns a new slice of Like objects containing
// the given like if it wasn't already present.
// It also returns the result of the append.
func (likes Likes) AppendIfMissing(other Like) (Likes, bool) {
	for _, like := range likes {
		if like.Equals(other) {
			return likes, false
		}
	}
	return append(likes, other), true
}

// ContainsOwnerLike returns true if the likes slice contain
// a like from the given owner, false otherwise
func (likes Likes) ContainsOwnerLike(owner sdk.AccAddress) bool {
	for _, like := range likes {
		if like.Owner.Equals(owner) {
			return true
		}
	}
	return false
}

// IndexOfByOwner returns the index of the like having the
// given owner inside the likes slice.
func (likes Likes) IndexOfByOwner(owner sdk.AccAddress) int {
	for index, like := range likes {
		if like.Owner.Equals(owner) {
			return index
		}
	}
	return -1
}

// RemoveLikeOfOwner returns a new Likes slice not containing the
// like of the given owner.
// If the like was removed properly, true is also returned. Otherwise,
// if no like was found, false is returned instead.
func (likes Likes) RemoveLikeOfOwner(owner sdk.AccAddress) (Likes, bool) {
	index := likes.IndexOfByOwner(owner)
	if index == -1 {
		return likes, false
	}

	return append(likes[:index], likes[index+1:]...), true
}
