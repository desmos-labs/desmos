package types

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// -----------------
// --- PollData
// -----------------

// PollData contains the information of a poll that is associated to a post
type PollData struct {
	Question              string      `json:"question" yaml:"question"`                               // Describes what poll is about
	ProvidedAnswers       PollAnswers `json:"provided_answers" yaml:"provided_answers"`               // Lists of answers provided by the creator
	EndDate               time.Time   `json:"end_date" yaml:"end_date"`                               // RFC3339 date at which the poll will no longer accept new answers
	Open                  bool        `json:"is_open" yaml:"is_open"`                                 // Tells if the poll is still accepting answers
	AllowsMultipleAnswers bool        `json:"allows_multiple_answers" yaml:"allows_multiple_answers"` // Tells if the poll is a single or multiple answers one
	AllowsAnswerEdits     bool        `json:"allows_answer_edits" yaml:"allows_answer_edits"`         // Tells if the poll allows answer edits
}

// NewPollData returns a new PollData object pointer containing the given data
func NewPollData(question string, endDate time.Time, providedAnswers PollAnswers,
	open, allowMultipleAnswers, allowsAnswerEdits bool) PollData {
	return PollData{
		Question:              question,
		EndDate:               endDate,
		ProvidedAnswers:       providedAnswers,
		Open:                  open,
		AllowsMultipleAnswers: allowMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
	}
}

// String implements fmt.Stringer
func (pd PollData) String() string {
	out := fmt.Sprintf("Question: %s \nOpen: %s \nEndDate: %s\nAllow multiple answers: %s \nAllow answer edits: %s \n",
		pd.Question,
		strconv.FormatBool(pd.Open),
		pd.EndDate,
		strconv.FormatBool(pd.AllowsMultipleAnswers),
		strconv.FormatBool(pd.AllowsAnswerEdits),
	)

	out += pd.ProvidedAnswers.String()

	return out
}

// Validate implements the validator interface
func (pd PollData) Validate() error {
	if strings.TrimSpace(pd.Question) == "" {
		return fmt.Errorf("missing poll title")
	}

	if pd.EndDate.IsZero() {
		return fmt.Errorf("invalid poll's end date")
	}

	if err := pd.ProvidedAnswers.Validate(); err != nil {
		return err
	}

	return nil
}

// ArePollDataEquals check whether the first and second pointers
// to a PollData object represents the same poll or not.
func ArePollDataEquals(first, second *PollData) bool {
	if first != nil && second != nil {
		return first.Equals(*second)
	}

	return first == second
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

// -----------------
// --- UserAnswer
// -----------------

// UserAnswer contains the data of a user's answer submission
type UserAnswer struct {
	Answers []AnswerID     `json:"answers" yaml:"answers"`
	User    sdk.AccAddress `json:"user" yaml:"user"`
}

// NewUserAnswer returns a new AnswerDetails object containing the given data
func NewUserAnswer(answers []AnswerID, user sdk.AccAddress) UserAnswer {
	return UserAnswer{
		Answers: answers,
		User:    user,
	}
}

// Strings implements fmt.Stringer
func (userAnswers UserAnswer) String() string {
	out := fmt.Sprintf("User: %s \nAnswers IDs: ", userAnswers.User.String())
	for _, answer := range userAnswers.Answers {
		out += strconv.FormatUint(uint64(answer), 10) + " "
	}

	return strings.TrimSpace(out)
}

// Validate implements validator
func (userAnswers UserAnswer) Validate() error {
	if userAnswers.User.Empty() {
		return fmt.Errorf("user cannot be empty")
	}

	if len(userAnswers.Answers) == 0 {
		return fmt.Errorf("answers cannot be empty")
	}

	return nil
}

// Equals returns true iff the userPollAnswers contains the same
// data of the other userPollAnswers
func (userAnswers UserAnswer) Equals(other UserAnswer) bool {
	if !userAnswers.User.Equals(other.User) {
		return false
	}

	if len(userAnswers.Answers) != len(other.Answers) {
		return false
	}

	for index, answer := range userAnswers.Answers {
		if answer != other.Answers[index] {
			return false
		}
	}

	return true
}

// ---------------
// --- UserAnswers
// ---------------

type UserAnswers []UserAnswer

// NewUserAnswers allows to create a new UserAnswers object from the given answers
func NewUserAnswers(answers ...UserAnswer) UserAnswers {
	return answers
}

// AppendIfMissingOrIfUserEquals appends the given answer to the user's answers slice if it does not exist inside it yet
// or if the user of the answer details is the same.
// It returns a new slice of containing such answer and a boolean indicating if the slice has been modified or not.
func (ua UserAnswers) AppendIfMissingOrIfUsersEquals(answer UserAnswer) (UserAnswers, bool) {
	for index, ad := range ua {

		if ad.Equals(answer) {
			return ua, false
		}

		if ad.User.Equals(answer.User) {
			ua[index] = answer
			return ua, true
		}

	}

	return append(ua, answer), true
}
