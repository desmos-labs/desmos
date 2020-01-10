package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// Represents the interface that each post should implement
type Post interface {
	String() string                 // Returns the string representation of the post
	GetID() PostID                  // Returns the ID of the post
	GetParentID() PostID            // Returns the ParentID of the post
	SetMessage(message string) Post // Sets the post message and return the post
	GetMessage() string             // Returns the post message
	CreationTime() sdk.Int          // Returns the creation time of the post
	GetEditTime() sdk.Int           // Returns the last modified time of the post
	SetEditTime(time sdk.Int) Post  // Sets the last modified time of the post and return the post
	CanComment() bool               // Tells if the post can be commented
	Owner() sdk.AccAddress          // Tells which user owns the post
	Validate() error                // Tells if the post is valid or not
	Equals(Post) bool               // Tells if this post has the same contents of the one given
}
