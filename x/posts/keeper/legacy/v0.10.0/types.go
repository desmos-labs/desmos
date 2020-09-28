package v0100

import (
	"sort"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// PostID represents a unique post id
type PostID string

// OptionalData represents a Posts' optional data and allows for custom
// Amino and JSON serialization and deserialization.
type OptionalData map[string]string

// KeyValue is a simple key/value representation of one field of a OptionalData.
type KeyValue struct {
	Key   string
	Value string
}

// MarshalAmino transforms the OptionalData to an array of key/value.
func (m OptionalData) MarshalAmino() ([]KeyValue, error) {
	fieldKeys := make([]string, len(m))
	i := 0
	for key := range m {
		fieldKeys[i] = key
		i++
	}

	sort.Stable(sort.StringSlice(fieldKeys))

	p := make([]KeyValue, len(m))
	for i, key := range fieldKeys {
		p[i] = KeyValue{
			Key:   key,
			Value: m[key],
		}
	}

	return p, nil
}

// UnmarshalAmino transforms the key/value array to a OptionalData.
func (m *OptionalData) UnmarshalAmino(keyValues []KeyValue) error {
	tempMap := make(map[string]string, len(keyValues))
	for _, p := range keyValues {
		tempMap[p.Key] = p.Value
	}

	*m = tempMap

	return nil
}

// Attachment contains the information representing any type of file provided with a post.
// This file can be an image or a multimedia file (vocals, video, documents, etc.).
type Attachment struct {
	URI      string           `json:"uri" yaml:"uri"`
	MimeType string           `json:"mime_type" yaml:"mime_type"`
	Tags     []sdk.AccAddress `json:"tags,omitempty" yaml:"tags,omitempty"`
}

// PollData contains the information of a poll that is associated to a post
type PollData struct {
	Question              string       `json:"question" yaml:"question"`                               // Describes what poll is about
	ProvidedAnswers       []PollAnswer `json:"provided_answers" yaml:"provided_answers"`               // Lists of answers provided by the creator
	EndDate               time.Time    `json:"end_date" yaml:"end_date"`                               // RFC3339 date at which the poll will no longer accept new answers
	Open                  bool         `json:"is_open" yaml:"is_open"`                                 // Tells if the poll is still accepting answers
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers" yaml:"allows_multiple_answers"` // Tells if the poll is a single or multiple answers one
	AllowsAnswerEdits     bool         `json:"allows_answer_edits" yaml:"allows_answer_edits"`         // Tells if the poll allows answer edits
}

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   AnswerID `json:"id" yaml:"id"`     // Unique id inside the post, serialized as a string for Javascript compatibility
	Text string   `json:"text" yaml:"text"` // Text of the answer
}

// AnswerID represents a unique answer id
type AnswerID uint64

// Post is a struct of a post
type Post struct {
	PostID         PostID         `json:"id" yaml:"id" `                                          // Unique id
	ParentID       PostID         `json:"parent_id" yaml:"parent_id"`                             // Post of which this one is a comment
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
