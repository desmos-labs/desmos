package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ---------------
// --- PollData
// ---------------

// PollData contains the information of a poll that is associated to a post
type PollData struct {
	Title                 string      `json:"title"`                   // Describes what poll is about
	ProvidedAnswers       PollAnswers `json:"provided_answers"`        // Lists of answers provided by the creator
	EndDate               time.Time   `json:"end_date"`                // RFC3339 date at which the poll will no longer accept new answers
	Open                  bool        `json:"open"`                    // Tells if the poll is still accepting answers
	AllowsMultipleAnswers bool        `json:"allows_multiple_answers"` // Tells if the poll is a single or multiple answers one
	AllowsAnswerEdits     bool        `json:"allows_answer_edits"`     // Tells if the poll allows answer edits
}

func NewPollData(title string, endDate time.Time, providedAnswers PollAnswers, open, allowMultipleAnswers, allowsAnswerEdits bool) *PollData {
	return &PollData{
		Title:                 title,
		EndDate:               endDate,
		ProvidedAnswers:       providedAnswers,
		Open:                  open,
		AllowsMultipleAnswers: allowMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
	}
}

// String implements fmt.Stringer
func (pd PollData) String() string {
	bytes, err := json.Marshal(&pd)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (pd PollData) MarshalJSON() ([]byte, error) {
	type temp PollData
	return json.Marshal(temp(pd))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (pd *PollData) UnmarshalJSON(data []byte) error {
	type pollData PollData
	var temp pollData
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*pd = PollData(temp)
	return nil
}

func (pd PollData) Validate() error {
	if strings.TrimSpace(pd.Title) == "" {
		return fmt.Errorf("missing poll title")
	}

	if pd.EndDate.Before(time.Now().UTC()) {
		return fmt.Errorf("end date cannot be in the past")
	}

	if err := pd.ProvidedAnswers.Validate(); err != nil {
		return err
	}

	return nil
}

func (pd *PollData) Equals(other *PollData) bool {
	if pd != nil && other == nil || pd == nil && other != nil {
		return false
	}

	if pd == nil && other == nil {
		return true
	}

	return pd.Title == other.Title &&
		pd.Open == other.Open &&
		pd.EndDate == other.EndDate &&
		pd.ProvidedAnswers.Equals(other.ProvidedAnswers) &&
		pd.AllowsMultipleAnswers == other.AllowsMultipleAnswers &&
		pd.AllowsAnswerEdits == other.AllowsAnswerEdits
}

// ---------------
// --- AnswersDetails
// ---------------
type AnswersDetails struct {
	Answers []uint64       `json:"answers"`
	User    sdk.AccAddress `json:"user"`
}

func NewAnswersDetails(answers []uint64, user sdk.AccAddress) AnswersDetails {
	return AnswersDetails{
		Answers: answers,
		User:    user,
	}
}

// Strings implements fmt.Stringer
func (userPollAnswers AnswersDetails) String() string {
	out := fmt.Sprintf("User: %s \nAnswers IDs:\n", userPollAnswers.User.String())
	for _, answer := range userPollAnswers.Answers {
		out += strconv.FormatUint(answer, 10) + " "
	}

	return strings.TrimSpace(out)
}

// Validate implements validator
func (userPollAnswers AnswersDetails) Validate() error {
	if userPollAnswers.User.Empty() {
		return fmt.Errorf("user cannot be empty")
	}

	if len(userPollAnswers.Answers) == 0 {
		return fmt.Errorf("answers cannot be empty")
	}

	return nil
}

// Equals returns true iff the userPollAnswers contains the same
// data of the other userPollAnswers
func (userPollAnswers AnswersDetails) Equals(other AnswersDetails) bool {
	if !userPollAnswers.User.Equals(other.User) {
		return false
	}

	if len(userPollAnswers.Answers) != len(other.Answers) {
		return false
	}

	for index, answer := range userPollAnswers.Answers {
		if answer != other.Answers[index] {
			return false
		}
	}

	return true
}

// ---------------
// --- PollAnswers
// ---------------
type PollAnswers []PollAnswer

// Strings implements fmt.Stringer
func (answers PollAnswers) String() string {
	out := "Answers\n[ID] [Text]\n"
	for _, answer := range answers {
		out += fmt.Sprintf("[%s] [%s]\n",
			strconv.FormatUint(answer.ID, 10), answer.Text)
	}
	return strings.TrimSpace(out)
}

// Validate implements validator
func (answers PollAnswers) Validate() error {
	if len(answers) == 0 {
		return fmt.Errorf("answers cannot be empty")
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

// ---------------
// --- PollAnswer
// ---------------

// PollAnswer contains the data of a single poll answer inserted by the creator
type PollAnswer struct {
	ID   uint64 `json:"id"`   // Unique id inside the post
	Text string `json:"text"` // Text of the answer
}

// String implements fmt.Stringer
func (pa PollAnswer) String() string {
	formattedID := strconv.FormatUint(pa.ID, 10)
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
