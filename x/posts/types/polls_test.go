package types_test

import (
	"fmt"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/app"

	"github.com/desmos-labs/desmos/x/posts/types"

	"testing"
	"time"
)

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.NewPollAnswer("0", "")
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

// ___________________________________________________________________________________________________________________

func TestPollAnswers_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name     string
		answers  types.PollAnswers
		toAppend types.PollAnswer
		expSlice types.PollAnswers
	}{
		{
			name:     "Answer is appended to empty slice",
			answers:  nil,
			toAppend: types.NewPollAnswer("id", "answer"),
			expSlice: types.NewPollAnswers(
				types.NewPollAnswer("id", "answer"),
			),
		},
		{
			name: "Answer is appended to non empty slice",
			answers: types.NewPollAnswers(
				types.NewPollAnswer("id_1", "answer_1"),
			),
			toAppend: types.NewPollAnswer("id_2", "answer_2"),
			expSlice: types.NewPollAnswers(
				types.NewPollAnswer("id_1", "answer_1"),
				types.NewPollAnswer("id_2", "answer_2"),
			),
		},
		{
			name: "Answer is not appended if existing",
			answers: types.NewPollAnswers(
				types.NewPollAnswer("id", "answer"),
			),
			toAppend: types.NewPollAnswer("id", "answer"),
			expSlice: types.NewPollAnswers(
				types.NewPollAnswer("id", "answer"),
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
		poll     *types.PollData
		expError error
	}{
		{
			name: "missing poll question",
			poll: types.NewPollData(
				"",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("missing poll question"),
		},
		{
			name: "invalid poll end date",
			poll: types.NewPollData(
				"title",
				time.Time{},
				types.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("invalid poll end date"),
		},
		{
			name: "not enough poll answer",
			poll: types.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.PollAnswers{},
				true,
				true,
			),
			expError: fmt.Errorf("poll answer must be at least two"),
		},
		{
			name: "invalid answer",
			poll: types.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.NewPollAnswers(
					types.NewPollAnswer("id_1", "answer_1"),
					types.NewPollAnswer("id_2", ""),
				),
				true,
				true,
			),
			expError: fmt.Errorf("answer text must be specified and cannot be empty"),
		},
		{
			name: "valid poll",
			poll: types.NewPollData(
				"title",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
				types.NewPollAnswers(
					types.NewPollAnswer("id_1", "answer_1"),
					types.NewPollAnswer("id_2", "answer_2"),
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
		name   string
		answer types.UserAnswer
		expErr error
	}{
		{
			name:   "Empty user returns error",
			answer: types.NewUserAnswer([]string{"1", "2"}, ""),
			expErr: fmt.Errorf("user cannot be empty"),
		},
		{
			name:   "Empty answer returns error",
			answer: types.NewUserAnswer(nil, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expErr: fmt.Errorf("answer cannot be empty"),
		},
		{
			name:   "Valid answer returns no error",
			answer: types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expErr, test.answer.Validate())
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestAppendIfMissingOrIfUsersEquals(t *testing.T) {
	tests := []struct {
		name             string
		answers          []types.UserAnswer
		answer           types.UserAnswer
		expectedSlice    []types.UserAnswer
		expectedAppended bool
	}{
		{
			name: "Missing user answer appended correctly",
			answers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"1", "2"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			expectedSlice: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			expectedAppended: true,
		},
		{
			name: "Same user with different answer replace previous ones",
			answers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"3"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expectedSlice: []types.UserAnswer{
				types.NewUserAnswer([]string{"3"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			expectedAppended: true,
		},
		{
			name: "Equals user answers returns the same users answer details",
			answers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			answer: types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			expectedSlice: []types.UserAnswer{
				types.NewUserAnswer([]string{"1", "2"}, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
			},
			expectedAppended: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := types.AppendIfMissingOrIfUsersEquals(test.answers, test.answer)
			require.Equal(t, test.expectedSlice, actual)
			require.Equal(t, test.expectedAppended, appended)
		})
	}
}

func TestUserAnswersMarshaling(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	answers := []types.UserAnswer{
		types.NewUserAnswer([]string{"1", "2"}, "user"),
		types.NewUserAnswer([]string{"3", "4"}, "user_2"),
	}
	marshaled := types.MustMarshalUserAnswers(cdc, answers)
	unmarshaled := types.MustUnmarshalUserAnswers(cdc, marshaled)
	require.Equal(t, answers, unmarshaled)
}
