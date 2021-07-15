package types_test

import (
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/staging/posts/types"

	"testing"
	"time"
)

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.NewProvidedAnswer("0", "")
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

// ___________________________________________________________________________________________________________________

func TestPollAnswers_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name     string
		answers  types.ProvidedAnswers
		toAppend types.ProvidedAnswer
		expSlice types.ProvidedAnswers
	}{
		{
			name:     "Answer is appended to empty slice",
			answers:  nil,
			toAppend: types.NewProvidedAnswer("id", "answer"),
			expSlice: types.NewPollAnswers(
				types.NewProvidedAnswer("id", "answer"),
			),
		},
		{
			name: "Answer is appended to non empty slice",
			answers: types.NewPollAnswers(
				types.NewProvidedAnswer("id_1", "answer_1"),
			),
			toAppend: types.NewProvidedAnswer("id_2", "answer_2"),
			expSlice: types.NewPollAnswers(
				types.NewProvidedAnswer("id_1", "answer_1"),
				types.NewProvidedAnswer("id_2", "answer_2"),
			),
		},
		{
			name: "Answer is not appended if existing",
			answers: types.NewPollAnswers(
				types.NewProvidedAnswer("id", "answer"),
			),
			toAppend: types.NewProvidedAnswer("id", "answer"),
			expSlice: types.NewPollAnswers(
				types.NewProvidedAnswer("id", "answer"),
			),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			slice := test.answers.AppendIfMissing(test.toAppend)
			require.Equal(t, test.expSlice, slice)
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestPollData_Validate(t *testing.T) {
	tests := []struct {
		name     string
		poll     *types.Poll
		expError error
	}{
		{
			name: "missing poll question",
			poll: types.NewPoll(
				"",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.ProvidedAnswers{},
				true,
				true,
			),
			expError: types.ErrPollEmptyQuestion,
		},
		{
			name: "invalid poll end date",
			poll: types.NewPoll(
				"title",
				time.Time{},
				types.ProvidedAnswers{},
				true,
				true,
			),
			expError: types.ErrPollEndDate,
		},
		{
			name: "not enough poll answer",
			poll: types.NewPoll(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.ProvidedAnswers{},
				true,
				true,
			),
			expError: types.ErrPollInvalidAnswersMinNumber,
		},
		{
			name: "empty answer",
			poll: types.NewPoll(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.NewPollAnswers(
					types.NewProvidedAnswer("id_1", "answer_1"),
					types.NewProvidedAnswer("id_2", ""),
				),
				true,
				true,
			),
			expError: types.ErrPollEmptyAnswer,
		},
		{
			name: "valid poll",
			poll: types.NewPoll(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.NewPollAnswers(
					types.NewProvidedAnswer("id_1", "answer_1"),
					types.NewProvidedAnswer("id_2", "answer_2"),
				),
				true,
				true,
			),
			expError: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expError, test.poll.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestUserAnswer_Validate(t *testing.T) {
	tests := []struct {
		name      string
		answer    types.UserAnswer
		shouldErr bool
	}{
		{
			name:      "Empty post id returns error",
			answer:    types.NewUserAnswer("", "", []string{}),
			shouldErr: true,
		},
		{
			name:      "Empty user returns error",
			answer:    types.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "", []string{}),
			shouldErr: true,
		},
		{
			name:      "Empty answers returns error",
			answer:    types.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{}),
			shouldErr: true,
		},
		{
			name:      "Invalid answer returns error",
			answer:    types.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{""}),
			shouldErr: true,
		},
		{
			name:      "Valid answer returns no error",
			answer:    types.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
			shouldErr: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldErr {
				require.Error(t, test.answer.Validate())
			} else {
				require.NoError(t, test.answer.Validate())
			}

		})
	}
}

func TestUserAnswersMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	answer := types.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "user", []string{"1", "2"})

	marshaled := types.MustMarshalUserAnswer(cdc, answer)
	unmarshaled := types.MustUnmarshalUserAnswer(cdc, marshaled)
	require.Equal(t, answer, unmarshaled)
}
