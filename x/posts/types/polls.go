package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

// NewPollAnswer returns a new PollAnswer object
func NewPollAnswer(id string, text string) PollAnswer {
	return PollAnswer{
		ID:   id,
		Text: text,
	}
}

// Validate implements validator
func (answer PollAnswer) Validate() error {
	if strings.TrimSpace(answer.Text) == "" {
		return fmt.Errorf("answer text must be specified and cannot be empty")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// PollAnswers represent a slice of PollAnswer objects
type PollAnswers []PollAnswer

// NewPollAnswers builds a new PollAnswers object from the given answer
func NewPollAnswers(answers ...PollAnswer) PollAnswers {
	return answers
}

// AppendIfMissing appends the given answer to the answer slice if it does not exist inside it yet.
// It returns a new slice of PollAnswers containing such PollAnswer.
func (answers PollAnswers) AppendIfMissing(answer PollAnswer) PollAnswers {
	for _, existing := range answers {
		if existing.Equal(answer) {
			return answers
		}
	}
	return append(answers, answer)
}

// ___________________________________________________________________________________________________________________

// NewPollData returns a new PollData object pointer containing the given poll
func NewPollData(
	question string, endDate time.Time, providedAnswers []PollAnswer, allowMultipleAnswers, allowsAnswerEdits bool,
) *PollData {
	return &PollData{
		Question:              question,
		EndDate:               endDate,
		ProvidedAnswers:       providedAnswers,
		AllowsMultipleAnswers: allowMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
	}
}

// Validate implements the validator interface
func (data PollData) Validate() error {
	if strings.TrimSpace(data.Question) == "" {
		return fmt.Errorf("missing poll question")
	}

	if data.EndDate.IsZero() {
		return fmt.Errorf("invalid poll end date")
	}

	if len(data.ProvidedAnswers) < 2 {
		return fmt.Errorf("poll answer must be at least two")
	}

	for _, answer := range data.ProvidedAnswers {
		err := answer.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// NewUserAnswer returns a new UserAnswer object containing the given poll
func NewUserAnswer(answers []string, user string) UserAnswer {
	return UserAnswer{
		Answers: answers,
		User:    user,
	}
}

// Validate implements validator
func (answers UserAnswer) Validate() error {
	if answers.User == "" {
		return fmt.Errorf("user cannot be empty")
	}

	if len(answers.Answers) == 0 {
		return fmt.Errorf("answer cannot be empty")
	}

	for _, answer := range answers.Answers {
		if strings.TrimSpace(answer) == "" {
			return fmt.Errorf("invalid answer")
		}
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// AppendIfMissingOrIfUserEquals appends the given answer to the user's answer slice if it does not exist inside it yet
// or if the user of the answer details is the same.
// It returns a new slice of containing such answer and a boolean indicating if the slice has been modified or not.
func AppendIfMissingOrIfUsersEquals(answers []UserAnswer, answer UserAnswer) ([]UserAnswer, bool) {
	for index, ad := range answers {
		if ad.Equal(answer) {
			return answers, false
		}

		if ad.User == answer.User {
			answers[index] = answer
			return answers, true
		}

	}

	return append(answers, answer), true
}

// ___________________________________________________________________________________________________________________

func MustMarshalUserAnswers(cdc codec.BinaryMarshaler, answer []UserAnswer) []byte {
	return cdc.MustMarshalBinaryBare(&UserAnswers{Answers: answer})
}

func MustUnmarshalUserAnswers(cdc codec.BinaryMarshaler, bz []byte) []UserAnswer {
	var answers UserAnswers
	cdc.MustUnmarshalBinaryBare(bz, &answers)
	return answers.Answers
}
