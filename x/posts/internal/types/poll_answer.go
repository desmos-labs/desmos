package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// ---------------
// --- AnswerID
// ---------------

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

// ---------------
// --- PollAnswer
// ---------------

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   AnswerID `json:"id" yaml:"id"`     // Unique id inside the post, serialized as a string for Javascript compatibility
	Text string   `json:"text" yaml:"text"` // Text of the answer
}

// NewPollAnswer returns a new PollAnswer object
func NewPollAnswer(id AnswerID, text string) PollAnswer {
	return PollAnswer{
		ID:   id,
		Text: text,
	}
}

// String implements fmt.Stringer
func (pa PollAnswer) String() string {
	formattedID := strconv.FormatUint(uint64(pa.ID), 10)
	return fmt.Sprintf("Answer - ID: %s ; Text: %s", formattedID, pa.Text)
}

// Validate implements validator
func (pa PollAnswer) Validate() error {
	if strings.TrimSpace(pa.Text) == "" {
		return fmt.Errorf("answer text must be specified and cannot be empty")
	}

	return nil
}

// Equals allows to check whether the contents of p are the same of other
func (pa PollAnswer) Equals(other PollAnswer) bool {
	return pa.ID == other.ID && pa.Text == other.Text
}

// ---------------
// --- PollAnswers
// ---------------

// PollAnswers represents a slice of poll answers
type PollAnswers []PollAnswer

// NewPollAnswers builds a new PollAnswers object starting from the given answers
func NewPollAnswers(answers ...PollAnswer) PollAnswers {
	return answers
}

// Strings implements fmt.Stringer
func (answers PollAnswers) String() string {
	out := "Provided Answers:\n[ID] [Text]\n"
	for _, answer := range answers {
		out += fmt.Sprintf("[%s] [%s]\n",
			strconv.FormatUint(uint64(answer.ID), 10), answer.Text)
	}
	return strings.TrimSpace(out)
}

// Validate implements validator
func (answers PollAnswers) Validate() error {
	if len(answers) < 2 {
		return fmt.Errorf("poll answers must be at least two")
	}

	for _, answer := range answers {
		if err := answer.Validate(); err != nil {
			return err
		}
	}

	return nil
}

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

// AppendIfMissing appends the given answer to the answers slice if it does not exist inside it yet.
// It returns a new slice of PollAnswers containing such PollAnswer.
func (answers PollAnswers) AppendIfMissing(newAnswer PollAnswer) PollAnswers {
	for _, answer := range answers {
		if answer.Equals(newAnswer) {
			return answers
		}
	}
	return append(answers, newAnswer)
}

// ExtractAnswersIDs appends every answer ID to a slice of IDs.
//It returns a slice of answers IDs.
func (answers PollAnswers) ExtractAnswersIDs() (answersIDs []AnswerID) {
	for _, answer := range answers {
		answersIDs = append(answersIDs, answer.ID)
	}
	return answersIDs
}
