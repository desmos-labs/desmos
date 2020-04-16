package v040

import (
	"crypto/sha256"
	"encoding/hex"
	"regexp"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "posts"
)

var (
	Sha256RegEx = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
)

// GenesisState contains the data of a v0.4.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post                    `json:"posts"`
	UsersPollAnswers    map[string][]UserAnswer   `json:"users_poll_answers"`
	PostReactions       map[string][]PostReaction `json:"post_reactions"`
	RegisteredReactions []Reaction                `json:"registered_reactions"`
}

// PostID represents a unique post id
type PostID string

// ComputeID returns a sha256 hash of the given data concatenated together
// nolint: interfacer
func ComputeID(creationDate time.Time, creator sdk.AccAddress, subspace string) PostID {
	hash := sha256.Sum256([]byte(creationDate.String() + creator.String() + subspace))
	return PostID(hex.EncodeToString(hash[:]))
}

// OptionalData represents a Posts' optional data and allows for custom
// Amino and JSON serialization and deserialization.
type OptionalData map[string]string

// PostMedia is a struct of a post media
type PostMedia struct {
	URI      string `json:"uri"`
	MimeType string `json:"mime_type"`
}

// AnswerID represents a unique answer id
type AnswerID uint64

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   AnswerID `json:"id"`   // Unique id inside the post, serialized as a string for Javascript compatibility
	Text string   `json:"text"` // Text of the answer
}

// PollData contains the information of a poll that is associated to a post
type PollData struct {
	Question              string       `json:"question"`                // Describes what poll is about
	ProvidedAnswers       []PollAnswer `json:"provided_answers"`        // Lists of answers provided by the creator
	EndDate               time.Time    `json:"end_date"`                // RFC3339 date at which the poll will no longer accept new answers
	Open                  bool         `json:"is_open"`                 // Tells if the poll is still accepting answers
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"` // Tells if the poll is a single or multiple answers one
	AllowsAnswerEdits     bool         `json:"allows_answer_edits"`     // Tells if the poll allows answer edits
}

// UserAnswer contains the data of a user's answer submission to a post's poll
type UserAnswer struct {
	Answers []AnswerID     `json:"answers"`
	User    sdk.AccAddress `json:"user"`
}

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
	Medias         []PostMedia    `json:"medias,omitempty"`        // Contains all the medias that are shared with the post
	PollData       *PollData      `json:"poll_data,omitempty"`     // Contains the poll details, if existing
}

// PostReaction is a struct of a user reaction to a post
type PostReaction struct {
	Owner sdk.AccAddress `json:"owner"` // Creator that has created the reaction
	Value string         `json:"value"` // PostReaction of the reaction
}

// Reaction represents a registered reaction that can be referenced
// by its shortCode inside post reactions
type Reaction struct {
	ShortCode string
	Value     string
	Subspace  string
	Creator   sdk.AccAddress
}
