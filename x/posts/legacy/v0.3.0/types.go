package v030

import (
	"regexp"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "posts"
)

var (
	SubspaceRegEx = regexp.MustCompile("^[a-fA-F0-9]{64}$")
)

// GenesisState contains the data of a v0.3.0 genesis state for the posts module
type GenesisState struct {
	Posts     []Post                `json:"posts"`
	Reactions map[string][]Reaction `json:"reactions"`
}

// PostID represents a unique post id
type PostID uint64

type OptionalData map[string]string

// Post is a struct of a post
type Post struct {
	PostID         PostID         `json:"id"`                      // Unique id
	ParentID       PostID         `json:"parent_id"`               // Post of which this one is a comment
	Message        string         `json:"message"`                 // Message contained inside the post
	Created        time.Time      `json:"created"`                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited"`             // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments"`         // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace"`                // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator"`                 // Creator of the Post
}

// PostReaction is a struct of a user reaction to a post
type Reaction struct {
	Owner sdk.AccAddress `json:"owner"` // User that has created the reaction
	Value string         `json:"value"` // PostReaction of the reaction
}
