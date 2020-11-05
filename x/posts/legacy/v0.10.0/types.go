package v0100

// DONTCOVER

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	v040posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.4.0"
	v060 "github.com/desmos-labs/desmos/x/posts/legacy/v0.6.0"
	v080posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.8.0"
)

// GenesisState contains the data of a v0.10.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post                            `json:"posts"`
	UsersPollAnswers    map[string][]v040posts.UserAnswer `json:"users_poll_answers"`
	PostReactions       map[string][]v060.PostReaction    `json:"post_reactions"`
	RegisteredReactions []v040posts.Reaction              `json:"registered_reactions"`
	Params              v080posts.Params                  `json:"params"`
}

// Post is a struct of a post
type Post struct {
	PostID         v040posts.PostID       `json:"id" yaml:"id" `                                          // Unique id
	ParentID       v040posts.PostID       `json:"parent_id" yaml:"parent_id"`                             // Post of which this one is a comment
	Message        string                 `json:"message" yaml:"message"`                                 // Message contained inside the post
	Created        time.Time              `json:"created" yaml:"created"`                                 // RFC3339 date at which the post has been created
	LastEdited     time.Time              `json:"last_edited" yaml:"last_edited"`                         // RFC3339 date at which the post has been edited the last time
	AllowsComments bool                   `json:"allows_comments" yaml:"allows_comments"`                 // Tells if users can reference this PostID as the parent
	Subspace       string                 `json:"subspace" yaml:"subspace"`                               // Identifies the application that has posted the message
	OptionalData   v040posts.OptionalData `json:"optional_data,omitempty" yaml:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress         `json:"creator" yaml:"creator"`                                 // Creator of the Post
	Attachments    []Attachment           `json:"attachments,omitempty" yaml:"attachments,omitempty"`     // Contains all the attachments that are shared with the post
	PollData       *v040posts.PollData    `json:"poll_data,omitempty" yaml:"poll_data"`                   // Contains the poll details, if existing
}

type Attachment struct {
	URI      string           `json:"uri" yaml:"uri"`
	MimeType string           `json:"mime_type" yaml:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty" yaml:"tags,omitempty"`
}
