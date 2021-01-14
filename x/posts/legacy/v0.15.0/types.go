package v0150

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GenesisState contains the data of a v0.15.0 genesis state for the posts module
type GenesisState struct {
	Posts               []Post               `json:"posts"`
	UsersPollAnswers    []UserAnswersEntry   `json:"users_poll_answers"`
	PostReactions       []PostReactionsEntry `json:"post_reactions"`
	RegisteredReactions []RegisteredReaction `json:"registered_reactions"`
	Params              Params               `json:"params"`
}

func (state GenesisState) FindUserAnswerEntryForPostID(postID string) (bool, *UserAnswersEntry) {
	for _, entry := range state.UsersPollAnswers {
		if entry.PostID == postID {
			return true, &entry
		}
	}
	return false, nil
}

func (state GenesisState) FindPostReactionEntryForPostID(postID string) (bool, *PostReactionsEntry) {
	for _, entry := range state.PostReactions {
		if entry.PostID == postID {
			return true, &entry
		}
	}
	return false, nil
}

// ----------------------------------------------------------------------------------------------------------------

type Post struct {
	PostID         string              `json:"id"`
	ParentID       string              `json:"parent_id"`
	Message        string              `json:"message"`
	Created        time.Time           `json:"created"`
	LastEdited     time.Time           `json:"last_edited"`
	AllowsComments bool                `json:"allows_comments"`
	Subspace       string              `json:"subspace"`
	OptionalData   []OptionalDataEntry `json:"optional_data,omitempty"`
	Creator        string              `json:"creator"`
	Attachments    []Attachment        `json:"attachments,omitempty"`
	PollData       *PollData           `json:"poll_data,omitempty"`
}

type OptionalDataEntry struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Attachment struct {
	URI      string   `json:"uri"`
	MimeType string   `json:"mime_type"`
	Tags     []string `json:"tags,omitempty"`
}

type PollData struct {
	Question              string       `json:"question,omitempty"`
	ProvidedAnswers       []PollAnswer `json:"provided_answers"`
	EndDate               time.Time    `json:"end_date"`
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"`
	AllowsAnswerEdits     bool         `json:"allows_answer_edits"`
}

type PollAnswer struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// ----------------------------------------------------------------------------------------------------------------

type UserAnswersEntry struct {
	PostID      string       `json:"post_id,omitempty"`
	UserAnswers []UserAnswer `json:"user_answers"`
}

type UserAnswer struct {
	User    string   `json:"user,omitempty"`
	Answers []string `json:"answers,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type PostReactionsEntry struct {
	PostID    string         `json:"post_id,omitempty"`
	Reactions []PostReaction `json:"reactions"`
}

type PostReaction struct {
	ShortCode string `json:"short_code"`
	Value     string `json:"value,omitempty"`
	Owner     string `json:"owner,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type RegisteredReaction struct {
	ShortCode string `json:"short_code"`
	Value     string `json:"value,omitempty"`
	Subspace  string `json:"subspace,omitempty"`
	Creator   string `json:"creator,omitempty"`
}

// ----------------------------------------------------------------------------------------------------------------

type Params struct {
	MaxPostMessageLength            sdk.Int `json:"max_post_message_length"`
	MaxOptionalDataFieldsNumber     sdk.Int `json:"max_optional_data_fields_number"`
	MaxOptionalDataFieldValueLength sdk.Int `json:"max_optional_data_field_value_length"`
}
