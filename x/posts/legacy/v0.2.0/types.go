package v0_2_0

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "posts"
)

// GenesisState contains the data of a v0.2.0 genesis state for the posts module
type GenesisState struct {
	Posts     []Post                `json:"posts"`
	Reactions map[string][]Reaction `json:"reactions"`
}

// PostID represents a unique post id
type PostID uint64

// Post is a struct of a post
type Post struct {
	PostID         PostID            `json:"id"`                      // Unique id
	ParentID       PostID            `json:"parent_id"`               // Post of which this one is a comment
	Message        string            `json:"message"`                 // Message contained inside the post
	Created        sdk.Int           `json:"created"`                 // Block height at which the post has been created
	LastEdited     sdk.Int           `json:"last_edited"`             // Block height at which the post has been edited the last time
	AllowsComments bool              `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string            `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   map[string]string `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Owner          sdk.AccAddress    `json:"owner"`                   // Creator of the Post
}

// Reaction is a struct of a user reaction to a post
type Reaction struct {
	Created sdk.Int        `json:"created"` // Block height at which the reaction was created
	Owner   sdk.AccAddress `json:"owner"`   // User that has created the reaction
	Value   string         `json:"value"`   // Value of the reaction
}
