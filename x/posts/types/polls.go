package types

import (
	"fmt"
	"strings"
	"time"
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

// NewPollAnswers builds a new PollAnswers object from the given answers
func NewPollAnswers(answers ...PollAnswer) PollAnswers {
	return answers
}

// AppendIfMissing appends the given answer to the answers slice if it does not exist inside it yet.
// It returns a new slice of PollAnswers containing such PollAnswer.
func (answers PollAnswers) AppendIfMissing(newAnswer PollAnswer) PollAnswers {
	for _, answer := range answers {
		if answer.Equal(newAnswer) {
			return answers
		}
	}
	return append(answers, newAnswer)
}

// ExtractAnswersIDs appends every answer ID to a slice of IDs.
//It returns a slice of answers IDs.
func (answers PollAnswers) ExtractAnswersIDs() (answersIDs []string) {
	for _, answer := range answers {
		answersIDs = append(answersIDs, answer.ID)
	}
	return answersIDs
}

// ___________________________________________________________________________________________________________________

// NewPollData returns a new PollData object pointer containing the given data
func NewPollData(
	question string, endDate time.Time, providedAnswers []PollAnswer, allowMultipleAnswers, allowsAnswerEdits bool,
) PollData {
	return PollData{
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
		return fmt.Errorf("missing poll title")
	}

	if data.EndDate.IsZero() {
		return fmt.Errorf("invalid poll's end date")
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

// NewUserAnswer returns a new AnswerDetails object containing the given data
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
		return fmt.Errorf("answers cannot be empty")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// UserAnswers represents a slice of UserAnswer
type UserAnswers []UserAnswer

// NewUserAnswers allows to create a new UserAnswers object from the given answers
func NewUserAnswers(answers ...UserAnswer) UserAnswers {
	return answers
}

// AppendIfMissingOrIfUserEquals appends the given answer to the user's answers slice if it does not exist inside it yet
// or if the user of the answer details is the same.
// It returns a new slice of containing such answer and a boolean indicating if the slice has been modified or not.
func (answers UserAnswers) AppendIfMissingOrIfUsersEquals(answer UserAnswer) (UserAnswers, bool) {
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
