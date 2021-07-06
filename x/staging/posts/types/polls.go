package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewAnswer returns a new Answer object
func NewAnswer(id string, text string) Answer {
	return Answer{
		ID:   id,
		Text: text,
	}
}

// Validate implements validator
func (answer Answer) Validate() error {
	if strings.TrimSpace(answer.Text) == "" {
		return fmt.Errorf("answer text must be specified and cannot be empty")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// Answers represent a slice of Answer objects
type Answers []Answer

// NewPollAnswers builds a new Answers object from the given answer
func NewPollAnswers(answers ...Answer) Answers {
	return answers
}

// AppendIfMissing appends the given answer to the answer slice if it does not exist inside it yet.
// It returns a new slice of Answers containing such PollAnswer.
func (answers Answers) AppendIfMissing(answer Answer) Answers {
	for _, existing := range answers {
		if existing.Equal(answer) {
			return answers
		}
	}
	return append(answers, answer)
}

// ___________________________________________________________________________________________________________________

// NewPoll returns a new Poll object pointer containing the given poll
func NewPoll(
	question string, endDate time.Time, providedAnswers []Answer, allowMultipleAnswers, allowsAnswerEdits bool,
) *Poll {
	return &Poll{
		Question:              question,
		EndDate:               endDate,
		ProvidedAnswers:       providedAnswers,
		AllowsMultipleAnswers: allowMultipleAnswers,
		AllowsAnswerEdits:     allowsAnswerEdits,
	}
}

// Validate implements the validator interface
func (data Poll) Validate() error {
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
func NewUserAnswer(postID, user string, answers []string) UserAnswer {
	return UserAnswer{
		PostID:  postID,
		User:    user,
		Answers: answers,
	}
}

// Validate implements validator
func (ua UserAnswer) Validate() error {
	if !IsValidPostID(ua.PostID) {
		return fmt.Errorf("invalid post id: %s", ua.PostID)
	}

	if _, err := sdk.AccAddressFromBech32(ua.User); err != nil {
		return fmt.Errorf("invalid user address: %s", ua.User)
	}

	if len(ua.Answers) == 0 {
		return fmt.Errorf("answers cannot be empty")
	}

	for _, answer := range ua.Answers {
		if strings.TrimSpace(answer) == "" {
			return fmt.Errorf("invalid answer")
		}

	}

	return nil
}

// MustMarshalUserAnswer serializes the given user answer using the provided BinaryMarshaler
func MustMarshalUserAnswer(cdc codec.BinaryMarshaler, answer UserAnswer) []byte {
	return cdc.MustMarshalBinaryBare(&answer)
}

// MustUnmarshalUserAnswer deserializes the given byte array as a user answer using
// the provided BinaryMarshaler
func MustUnmarshalUserAnswer(cdc codec.BinaryMarshaler, bz []byte) UserAnswer {
	var answer UserAnswer
	cdc.MustUnmarshalBinaryBare(bz, &answer)
	return answer
}
