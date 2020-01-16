package types

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"time"
)

// Represents the interface that each post should implement
type Post interface {
	String() string                     // Returns the string representation of the post
	GetID() PostID                      // Returns the ID of the post
	GetParentID() PostID                // Returns the ParentID of the post
	SetMessage(message string) Post     // Sets the post message and return the post
	GetMessage() string                 // Returns the post message
	CreationTime() time.Time            // Returns the creation time of the post
	GetEditTime() time.Time             // Returns the last modified time of the post
	SetEditTime(time time.Time) Post    // Sets the last modified time of the post and return the post
	CanComment() bool                   // Tells if the post can be commented
	GetSubspace() string                // Returns the subspace of the post
	GetOptionalData() map[string]string //Returns the optional data of the post
	Owner() sdk.AccAddress              // Tells which user owns the post
	Validate() error                    // Tells if the post is valid or not
	Equals(Post) bool                   // Tells if this post has the same contents of the one given
}

type Posts []Post

// String implements fmt.Stringer
func (posts Posts) String() string {
	bytes, err := json.Marshal(&posts)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}
