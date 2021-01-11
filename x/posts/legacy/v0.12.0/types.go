package v0120

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "posts"
)

// GenesisState contains the data of a v0.12.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post                    `json:"posts"`
	UsersPollAnswers    map[string][]UserAnswer   `json:"users_poll_answers"`
	PostReactions       map[string][]PostReaction `json:"post_reactions"`
	RegisteredReactions []RegisteredReaction      `json:"registered_reactions"`
	Params              Params                    `json:"params"`
}

// ----------------------------------------------------------------------------------------------------------------

// ComputeID returns a sha256 hash of the given data concatenated together
//nolint: interfacer
func ComputeID(creationDate time.Time, creator sdk.AccAddress, subspace string) string {
	hash := sha256.Sum256([]byte(creationDate.String() + creator.String() + subspace))
	return hex.EncodeToString(hash[:])
}

type Post struct {
	PostID         string         `json:"id" yaml:"id" `                                          // Unique id
	ParentID       string         `json:"parent_id" yaml:"parent_id"`                             // Post of which this one is a comment
	Message        string         `json:"message" yaml:"message"`                                 // Message contained inside the post
	Created        time.Time      `json:"created" yaml:"created"`                                 // RFC3339 date at which the post has been created
	LastEdited     time.Time      `json:"last_edited" yaml:"last_edited"`                         // RFC3339 date at which the post has been edited the last time
	AllowsComments bool           `json:"allows_comments" yaml:"allows_comments"`                 // Tells if users can reference this PostID as the parent
	Subspace       string         `json:"subspace" yaml:"subspace"`                               // Identifies the application that has posted the message
	OptionalData   OptionalData   `json:"optional_data,omitempty" yaml:"optional_data,omitempty"` // Arbitrary data that can be used from the developers
	Creator        sdk.AccAddress `json:"creator" yaml:"creator"`                                 // Creator of the Post
	Attachments    []Attachment   `json:"attachments,omitempty" yaml:"attachments,omitempty"`     // Contains all the attachments that are shared with the post
	PollData       *PollData      `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`         // Contains the poll details, if existing
}

type OptionalData map[string]string

type PostReaction struct {
	Owner     sdk.AccAddress `json:"owner" yaml:"owner"`
	Shortcode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
}

type Attachment struct {
	URI      string           `json:"uri" yaml:"uri"`
	MimeType string           `json:"mime_type" yaml:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty" yaml:"tags,omitempty"`
}

type PollData struct {
	Question              string       `json:"question" yaml:"question"`
	ProvidedAnswers       []PollAnswer `json:"provided_answers" yaml:"provided_answers"`
	EndDate               time.Time    `json:"end_date" yaml:"end_date"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers" yaml:"allows_multiple_answers"`
	AllowsAnswerEdits     bool         `json:"allows_answer_edits" yaml:"allows_answer_edits"`
}

type PollAnswer struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
}

// ----------------------------------------------------------------------------------------------------------------

type RegisteredReaction struct {
	ShortCode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

// ----------------------------------------------------------------------------------------------------------------

type UserAnswer struct {
	Answers []uint64       `json:"answers"`
	User    sdk.AccAddress `json:"user"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length" yaml:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number" yaml:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length" yaml:"max_optional_data_field_value_length"`
}
