package v030

// DONTCOVER

import (
	"encoding/json"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	ModuleName = "posts"
)

// GenesisState contains the data of a v0.3.0 genesis state for the posts module
type GenesisState struct {
	Posts       []Post                  `json:"posts"`
	PollAnswers map[string][]UserAnswer `json:"poll_answers_details"`
	Reactions   map[string][]Reaction   `json:"reactions"`
}

// PostID represents a unique post id
type PostID uint64

// String implements fmt.Stringer
func (id PostID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON implements Marshaler
func (id PostID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON implements Unmarshaler
func (id *PostID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	postID, err := ParsePostID(s)
	if err != nil {
		return err
	}

	*id = postID
	return nil
}

// ParsePostID returns the PostID represented inside the provided
// value, or an error if no id could be parsed properly
func ParsePostID(value string) (PostID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return PostID(0), err
	}

	return PostID(intVal), err
}

type OptionalData map[string]string

// PostMedia is a struct of a post media
type PostMedia struct {
	URI      string `json:"uri"`
	MimeType string `json:"mime_type"`
}

// AnswerID represents a unique answer id
type AnswerID uint64

// String implements fmt.Stringer
func (id AnswerID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

// MarshalJSON implements Marshaler
func (id AnswerID) MarshalJSON() ([]byte, error) {
	return json.Marshal(id.String())
}

// UnmarshalJSON implements Unmarshaler
func (id *AnswerID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	postID, err := ParseAnswerID(s)
	if err != nil {
		return err
	}

	*id = postID
	return nil
}

// ParseAnswerID returns the AnswerID represented inside the provided
// value, or an error if no id could be parsed properly
func ParseAnswerID(value string) (AnswerID, error) {
	intVal, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		return AnswerID(0), err
	}

	return AnswerID(intVal), err
}

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   AnswerID `json:"id"`   // Unique id inside the post, serialized as a string for Javascript compatibility
	Text string   `json:"text"` // Text of the answer
}

// PollData contains the information of a poll that is associated to a post
type PollData struct {
	Question              string      `json:"question"`                // Describes what poll is about
	ProvidedAnswers       PollAnswers `json:"provided_answers"`        // Lists of answers provided by the creator
	EndDate               time.Time   `json:"end_date"`                // RFC3339 date at which the poll will no longer accept new answers
	Open                  bool        `json:"is_open"`                 // Tells if the poll is still accepting answers
	AllowsMultipleAnswers bool        `json:"allows_multiple_answers"` // Tells if the poll is a single or multiple answers one
	AllowsAnswerEdits     bool        `json:"allows_answer_edits"`     // Tells if the poll allows answer edits
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
	Medias         PostMedias     `json:"medias,omitempty"`        // Contains all the medias that are shared with the post
	PollData       *PollData      `json:"poll_data,omitempty"`     // Contains the poll details, if existing
}

// Reaction is a struct of a user reaction to a post
type Reaction struct {
	Owner sdk.AccAddress `json:"owner"` // User that has created the reaction
	Value string         `json:"value"` // Value of the reaction
}

func (p Post) ConflictsWith(other Post) bool {
	return p.Created.Equal(other.Created) &&
		p.Subspace == other.Subspace &&
		p.Creator.Equals(other.Creator)
}

// ContentsEquals returns true if and only if p and other contain the same data, without considering the ID
func (p Post) ContentsEquals(other Post) bool {
	equalsOptionalData := len(p.OptionalData) == len(other.OptionalData)
	if equalsOptionalData {
		for key := range p.OptionalData {
			equalsOptionalData = equalsOptionalData && p.OptionalData[key] == other.OptionalData[key]
		}
	}

	return p.ParentID == other.ParentID &&
		p.Message == other.Message &&
		p.Created.Equal(other.Created) &&
		p.LastEdited.Equal(other.LastEdited) &&
		p.AllowsComments == other.AllowsComments &&
		p.Subspace == other.Subspace &&
		equalsOptionalData &&
		p.Creator.Equals(other.Creator) &&
		p.Medias.Equals(other.Medias) &&
		ArePollDataEquals(p.PollData, other.PollData)
}

// Equals allows to check whether the contents of pm are the same of other
func (pm PostMedia) Equals(other PostMedia) bool {
	return pm.URI == other.URI && pm.MimeType == other.MimeType
}

type PostMedias []PostMedia

// Equals returns true iff the pms slice contains the same
// data in the same order of the other slice
func (pms PostMedias) Equals(other PostMedias) bool {
	if len(pms) != len(other) {
		return false
	}

	for index, postMedia := range pms {
		if !postMedia.Equals(other[index]) {
			return false
		}
	}

	return true
}

// Equals allows to check whether the contents of p are the same of other
func (pa PollAnswer) Equals(other PollAnswer) bool {
	return pa.ID == other.ID && pa.Text == other.Text
}

type PollAnswers []PollAnswer

// Equals returns true iff the answers slice contains the same
// data in the same order of the other slice
func (answers PollAnswers) Equals(other PollAnswers) bool {
	if len(answers) != len(other) {
		return false
	}

	for index, answer := range answers {
		if !answer.Equals(other[index]) {
			return false
		}
	}

	return true
}

// Equals returns true if this poll data object has the same contents of the other given.
// It assumes neither pd or other are null.
// To check the equality between possibly null values use ArePollDataEquals instead.
func (pd PollData) Equals(other PollData) bool {
	return pd.Question == other.Question &&
		pd.Open == other.Open &&
		pd.EndDate == other.EndDate &&
		pd.ProvidedAnswers.Equals(other.ProvidedAnswers) &&
		pd.AllowsMultipleAnswers == other.AllowsMultipleAnswers &&
		pd.AllowsAnswerEdits == other.AllowsAnswerEdits
}

// ArePollDataEquals check whether the first and second pointers
// to a PollData object represents the same poll or not.
func ArePollDataEquals(first, second *PollData) bool {
	if first != nil && second != nil {
		return first.Equals(*second)
	}

	return first == second
}
