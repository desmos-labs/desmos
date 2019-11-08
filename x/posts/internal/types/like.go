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
	LikeID        LikeID         `json:"id"`
	PostID        PostID         `json:"post_id"`
	Created       int64          `json:"created"` // Block height at which the like was created
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
	bytes, err := json.Marshal(&l)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func (l Like) Validate() error {
	if !l.LikeID.Valid() {
		return fmt.Errorf("invalid like id %s", l.LikeID)
	}

	if l.Owner == nil {
		return fmt.Errorf("invalid like owner: %s", l.Owner)
	}

	if l.Created == 0 {
		return fmt.Errorf("invalid like creation block heigth: %d", l.Created)
	}

	if !l.PostID.Valid() {
		return fmt.Errorf("invalid like post id: %s", l.PostID)
	}

	return nil
}
