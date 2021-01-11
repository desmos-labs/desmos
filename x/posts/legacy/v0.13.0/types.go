package v0130

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState contains the data of a v0.13.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post                    `json:"posts"`
	UsersPollAnswers    map[string][]UserAnswer   `json:"users_poll_answers"`
	PostReactions       map[string][]PostReaction `json:"post_reactions"`
	RegisteredReactions []RegisteredReaction      `json:"registered_reactions"`
	Params              Params                    `json:"params"`
}

// ----------------------------------------------------------------------------------------------------------------

type Post struct {
	PostID         string              `json:"id" yaml:"id" `
	ParentID       string              `json:"parent_id" yaml:"parent_id"`
	Message        string              `json:"message" yaml:"message"`
	Created        time.Time           `json:"created" yaml:"created"`
	LastEdited     time.Time           `json:"last_edited" yaml:"last_edited"`
	AllowsComments bool                `json:"allows_comments" yaml:"allows_comments"`
	Subspace       string              `json:"subspace" yaml:"subspace"`
	OptionalData   []OptionalDataEntry `json:"optional_data,omitempty" yaml:"optional_data,omitempty"`
	Creator        sdk.AccAddress      `json:"creator" yaml:"creator"`
	Attachments    []Attachment        `json:"attachments,omitempty" yaml:"attachments,omitempty"`
	PollData       *PollData           `json:"poll_data,omitempty" yaml:"poll_data,omitempty"`
}

type OptionalDataEntry struct {
	Key   string `json:"key" yaml:"key"`
	Value string `json:"value" yaml:"value"`
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

type PostReaction struct {
	Owner     sdk.AccAddress `json:"owner" yaml:"owner"`
	Shortcode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
}

// ----------------------------------------------------------------------------------------------------------------

type UserAnswer struct {
	Answers []uint64       `json:"answers"`
	User    sdk.AccAddress `json:"user"`
}

// ----------------------------------------------------------------------------------------------------------------

type RegisteredReaction struct {
	ShortCode string         `json:"shortcode" yaml:"shortcode"`
	Value     string         `json:"value" yaml:"value"`
	Subspace  string         `json:"subspace" yaml:"subspace"`
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length" yaml:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number" yaml:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length" yaml:"max_optional_data_field_value_length"`
}
