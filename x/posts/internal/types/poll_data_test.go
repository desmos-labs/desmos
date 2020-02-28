package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	"testing"
	"time"
)

// ---------------
// --- PollData
// ---------------

var pollEndDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)
var answer = types.NewPollAnswer(types.AnswerID(1), "Yes")
var answer2 = types.NewPollAnswer(types.AnswerID(2), "No")

func pollDataPointer(data types.PollData) *types.PollData {
	return &data
}

func TestPollData_String(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}
	pollData := types.NewPollData("poll?", pollEndDate, types.PollAnswers{answer, answer2}, true, false, true)

	require.Equal(t, "Question: poll? \nOpen: true \nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		pollData.String())
}

func TestPollData_Validate(t *testing.T) {
	tests := []struct {
		pollData types.PollData
		expError string
	}{
		{
			pollData: types.NewPollData("", pollEndDate, types.PollAnswers{}, true, true, true),
			expError: "missing poll title",
		},
		{
			pollData: types.NewPollData("title", time.Time{}, types.PollAnswers{}, true, true, true),
			expError: "invalid poll's end date",
		},
		{
			pollData: types.NewPollData("title", pollEndDate, types.PollAnswers{}, true, true, true),
			expError: "poll answers must be at least two",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expError, test.pollData.Validate().Error())
	}
}

func TestArePollDataEquals(t *testing.T) {
	tests := []struct {
		name      string
		first     *types.PollData
		second    *types.PollData
		expEquals bool
	}{
		{
			name:      "Different titles",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different open",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), false, false, true)),
			expEquals: false,
		},
		{
			name:      "Different end date",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", time.Now().UTC(), types.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different provided answers",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different edits answer option",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, false)),
			expEquals: false,
		},
		{
			name:      "Different multiple answers option",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, true, true)),
			expEquals: false,
		},
		{
			name:      "Equals poll data",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: true,
		},
		{
			name:      "First nil",
			first:     nil,
			second:    pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Second nil",
			first:     pollDataPointer(types.NewPollData("poll?", pollEndDate, types.NewPollAnswers(answer), true, false, true)),
			second:    nil,
			expEquals: false,
		},
		{
			name:      "Both nil",
			first:     nil,
			second:    nil,
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, types.ArePollDataEquals(test.first, test.second))
		})
	}
}

// ---------------
// --- UserAnswer
// ---------------

func TestUserAnswer_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	userPollAnswers := types.NewUserAnswer(answers, user)

	require.Equal(t, "User: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", userPollAnswers.String())
}

func TestUserAnswer_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}

	tests := []struct {
		name            string
		userPollAnswers types.UserAnswer
		expErr          string
	}{
		{
			name:            "Empty user returns error",
			userPollAnswers: types.NewUserAnswer(answers, nil),
			expErr:          "user cannot be empty",
		},
		{
			name:            "Empty answers returns error",
			userPollAnswers: types.NewUserAnswer(nil, user),
			expErr:          "answers cannot be empty",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if err := test.userPollAnswers.Validate(); err != nil {
				require.Equal(t, test.expErr, err.Error())
			}
		})
	}
}

func TestUserAnswer_Equals(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}
	answers2 := []types.AnswerID{types.AnswerID(1)}

	tests := []struct {
		name      string
		first     types.UserAnswer
		second    types.UserAnswer
		expEquals bool
	}{
		{
			name:      "Different users returns false",
			first:     types.NewUserAnswer(answers, user),
			second:    types.NewUserAnswer(answers, user2),
			expEquals: false,
		},
		{
			name:      "Different answers lengths returns false",
			first:     types.NewUserAnswer(answers, user),
			second:    types.NewUserAnswer(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Different answers return false",
			first:     types.NewUserAnswer(answers, user),
			second:    types.NewUserAnswer(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Equals userPollAnswers returns true",
			first:     types.NewUserAnswer(answers, user),
			second:    types.NewUserAnswer(answers, user),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

// ------------------
// --- UserAnswers
// ------------------

func TestUserAnswers_AppendIfMissingOrIfUsersEquals(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	user2, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	require.NoError(t, err)

	answers := []types.AnswerID{types.AnswerID(1), types.AnswerID(2)}
	answers2 := []types.AnswerID{types.AnswerID(3)}

	tests := []struct {
		name        string
		usersAD     types.UserAnswers
		ansDet      types.UserAnswer
		expUsersAD  types.UserAnswers
		expAppended bool
	}{
		{
			name:    "Missing user answers details appended correctly",
			usersAD: types.UserAnswers{types.NewUserAnswer(answers, user)},
			ansDet:  types.NewUserAnswer(answers, user2),
			expUsersAD: types.UserAnswers{
				types.NewUserAnswer(answers, user),
				types.NewUserAnswer(answers, user2),
			},
			expAppended: true,
		},
		{
			name:        "Same user with different answers replace previous ones",
			usersAD:     types.UserAnswers{types.NewUserAnswer(answers, user)},
			ansDet:      types.NewUserAnswer(answers2, user),
			expUsersAD:  types.UserAnswers{types.NewUserAnswer(answers2, user)},
			expAppended: true,
		},
		{
			name:        "Equals user answers details returns the same users answers details",
			usersAD:     types.UserAnswers{types.NewUserAnswer(answers, user)},
			ansDet:      types.NewUserAnswer(answers, user),
			expUsersAD:  types.UserAnswers{types.NewUserAnswer(answers, user)},
			expAppended: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.usersAD.AppendIfMissingOrIfUsersEquals(test.ansDet)
			require.Equal(t, test.expUsersAD, actual)
			require.Equal(t, test.expAppended, appended)
		})
	}
}
