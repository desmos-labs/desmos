package v010

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "posts"
)

// GenesisState contains the data of a v0.1.0 genesis state for the posts module
type GenesisState struct {
	Posts []Post            `json:"posts"`
	Likes map[string][]Like `json:"likes"`
}

// PostID represents a unique post id
type PostID uint64

// Post is a struct of a post
type Post struct {
	PostID            PostID         `json:"id"`                 // Unique id
	ParentID          PostID         `json:"parent_id"`          // Post of which this one is a comment
	Message           string         `json:"message"`            // Message contained inside the post
	Created           sdk.Int        `json:"created"`            // Block height at which the post has been created
	LastEdited        sdk.Int        `json:"last_edited"`        // Block height at which the post has been edited the last time
	AllowsComments    bool           `json:"allows_comments"`    // Tells if users can reference this PostID as the parent
	ExternalReference string         `json:"external_reference"` // Used to know when to display this post
	Owner             sdk.AccAddress `json:"owner"`              // Creator of the post
}

// PostReaction is a struct of a user like
type Like struct {
	Created sdk.Int        `json:"created"` // Block height at which the like was created
	Owner   sdk.AccAddress `json:"owner"`   // User that has inserted the like
}
