package types_test

import (
	"fmt"
	types2 "github.com/desmos-labs/desmos/x/posts/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"

	"testing"
	"time"
)

func TestPollAnswer_Validate(t *testing.T) {
	answer := types2.NewPollAnswer("0", "")
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

// ___________________________________________________________________________________________________________________

func TestPollAnswers_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name     string
		answers  types2.PollAnswers
		toAppend types2.PollAnswer
		expSlice types2.PollAnswers
	}{
		{
			name:     "Answer is appended to empty slice",
			answers:  nil,
			toAppend: types2.NewPollAnswer("id", "answer"),
			expSlice: types2.NewPollAnswers(
				types2.NewPollAnswer("id", "answer"),
			),
		},
		{
			name: "Answer is appended to non empty slice",
			answers: types2.NewPollAnswers(
				types2.NewPollAnswer("id_1", "answer_1"),
			),
			toAppend: types2.NewPollAnswer("id_2", "answer_2"),
			expSlice: types2.NewPollAnswers(
				types2.NewPollAnswer("id_1", "answer_1"),
				types2.NewPollAnswer("id_2", "answer_2"),
			),
		},
		{
			name: "Answer is not appended if existing",
			answers: types2.NewPollAnswers(
				types2.NewPollAnswer("id", "answer"),
			),
			toAppend: types2.NewPollAnswer("id", "answer"),
			expSlice: types2.NewPollAnswers(
				types2.NewPollAnswer("id", "answer"),
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
		poll     *types2.PollData
		expError error
	}{
		{
			name: "missing poll question",
			poll: types2.NewPollData(
				"",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types2.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("missing poll question"),
		},
		{
			name: "invalid poll end date",
			poll: types2.NewPollData(
				"title",
				time.Time{},
				types2.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("invalid poll end date"),
		},
		{
			name: "not enough poll answer",
			poll: types2.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types2.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("poll answer must be at least two"),
		},
		{
			name: "invalid answer",
			poll: types2.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types2.NewPollAnswers(
					types2.NewPollAnswer("id_1", "answer_1"),
					types2.NewPollAnswer("id_2", ""),
				),
				true,
				true,
			),
			expError: fmt.Errorf("answer text must be specified and cannot be empty"),
		},
		{
			name: "valid poll",
			poll: types2.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types2.NewPollAnswers(
					types2.NewPollAnswer("id_1", "answer_1"),
					types2.NewPollAnswer("id_2", "answer_2"),
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
		answer    types2.UserAnswer
		shouldErr bool
	}{
		{
			name:      "Empty post id returns error",
			answer:    types2.NewUserAnswer("", "", []string{}),
			shouldErr: true,
		},
		{
			name:      "Empty user returns error",
			answer:    types2.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "", []string{}),
			shouldErr: true,
		},
		{
			name:      "Empty answers returns error",
			answer:    types2.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{}),
			shouldErr: true,
		},
		{
			name:      "Invalid answer returns error",
			answer:    types2.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{""}),
			shouldErr: true,
		},
		{
			name:      "Valid answer returns no error",
			answer:    types2.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", []string{"1", "2"}),
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
	answer := types2.NewUserAnswer("9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "user", []string{"1", "2"})

	marshaled := types2.MustMarshalUserAnswer(cdc, answer)
	unmarshaled := types2.MustUnmarshalUserAnswer(cdc, marshaled)
	require.Equal(t, answer, unmarshaled)
}
