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
	PostId         string         `json:"id" `
	ParentId       string         `json:"parent_id"`
	Message        string         `json:"message"`
	Created        time.Time      `json:"created"`
	LastEdited     time.Time      `json:"last_edited"`
	AllowsComments bool           `json:"allows_comments"`
	Subspace       string         `json:"subspace"`
	OptionalData   OptionalData   `json:"optional_data,omitempty"`
	Creator        sdk.AccAddress `json:"creator"`
	Attachments    []Attachment   `json:"attachments,omitempty"`
	PollData       *PollData      `json:"poll_data,omitempty"`
}

type OptionalData map[string]string

type PostReaction struct {
	Owner     sdk.AccAddress `json:"owner"`
	Shortcode string         `json:"shortcode"`
	Value     string         `json:"value"`
}

type Attachment struct {
	URI      string           `json:"uri"`
	MimeType string           `json:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty"`
}

type PollData struct {
	Question              string       `json:"question"`
	ProvidedAnswers       []PollAnswer `json:"provided_answers"`
	EndDate               time.Time    `json:"end_date"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"`
	AllowsAnswerEdits     bool         `json:"allows_answer_edits"`
}

type PollAnswer struct {
	ID   uint64 `json:"id"`
	Text string `json:"text"`
}

// ----------------------------------------------------------------------------------------------------------------

type RegisteredReaction struct {
	ShortCode string         `json:"shortcode"`
	Value     string         `json:"value"`
	Subspace  string         `json:"subspace"`
	Creator   sdk.AccAddress `json:"creator"`
}

// ----------------------------------------------------------------------------------------------------------------

type UserAnswer struct {
	Answers []uint64       `json:"answers"`
	User    sdk.AccAddress `json:"user"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length"`
}
