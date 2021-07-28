package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewProvidedAnswer returns a new ProvidedAnswer object
func NewProvidedAnswer(id string, text string) ProvidedAnswer {
	return ProvidedAnswer{
		ID:   id,
		Text: text,
	}
}

// Validate implements validator
func (answer ProvidedAnswer) Validate() error {
	if strings.TrimSpace(answer.Text) == "" {
		return sdkerrors.Wrap(ErrInvalidPostPoll, "answer text must be specified and cannot be empty")
	}

	return nil
}

// ___________________________________________________________________________________________________________________

// ProvidedAnswers represent a slice of ProvidedAnswers objects
type ProvidedAnswers []ProvidedAnswer

// NewPollAnswers builds a new ProvidedAnswers object from the given answer
func NewPollAnswers(providedAnswers ...ProvidedAnswer) ProvidedAnswers {
	return providedAnswers
}

// AppendIfMissing appends the given answer to the answers slice if it does not exist inside it yet.
// It returns a new slice of ProvidedAnswers containing such PollAnswer.
func (answers ProvidedAnswers) AppendIfMissing(answer ProvidedAnswer) ProvidedAnswers {
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
	question string, endDate time.Time, providedAnswers []ProvidedAnswer, allowMultipleAnswers, allowsAnswerEdits bool,
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
		return sdkerrors.Wrap(ErrInvalidPostPoll, "missing poll question")
	}

	if data.EndDate.IsZero() {
		return sdkerrors.Wrap(ErrInvalidPostPoll, "invalid poll end date")
	}

	if len(data.ProvidedAnswers) < 2 {
		return sdkerrors.Wrap(ErrInvalidPollAnswers, "poll answers must be at least two")
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
