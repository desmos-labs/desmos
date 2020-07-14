package polls_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/types/models/polls"
	"github.com/stretchr/testify/require"

	"testing"
	"time"
)

// ---------------
// --- PollData
// ---------------

func pollDataPointer(data polls.PollData) *polls.PollData {
	return &data
}

func TestPollData_String(t *testing.T) {
	var timeZone, _ = time.LoadLocation("UTC")
	var pollEndDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)

	pollData := polls.NewPollData(
		"poll?",
		pollEndDate,
		polls.NewPollAnswers(
			polls.NewPollAnswer(polls.AnswerID(1), "Yes"),
			polls.NewPollAnswer(polls.AnswerID(2), "No"),
		),
		true,
		false,
		true,
	)

	require.Equal(t, "Question: poll? \nOpen: true \nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		pollData.String())
}

func TestPollData_Validate(t *testing.T) {
	var timeZone, _ = time.LoadLocation("UTC")
	var pollEndDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)

	tests := []struct {
		pollData polls.PollData
		expError string
	}{
		{
			pollData: polls.NewPollData("", pollEndDate, polls.PollAnswers{}, true, true, true),
			expError: "missing poll title",
		},
		{
			pollData: polls.NewPollData("title", time.Time{}, polls.PollAnswers{}, true, true, true),
			expError: "invalid poll's end date",
		},
		{
			pollData: polls.NewPollData("title", pollEndDate, polls.PollAnswers{}, true, true, true),
			expError: "poll answers must be at least two",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expError, test.pollData.Validate().Error())
	}
}

func TestArePollDataEquals(t *testing.T) {
	var timeZone, _ = time.LoadLocation("UTC")
	var pollEndDate = time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone)
	var answer = polls.NewPollAnswer(polls.AnswerID(1), "Yes")
	var answer2 = polls.NewPollAnswer(polls.AnswerID(2), "No")

	tests := []struct {
		name      string
		first     *polls.PollData
		second    *polls.PollData
		expEquals bool
	}{
		{
			name:      "Different titles",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different open",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), false, false, true)),
			expEquals: false,
		},
		{
			name:      "Different end date",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", time.Now().UTC(), polls.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different provided answers",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Different edits answer option",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, false)),
			expEquals: false,
		},
		{
			name:      "Different multiple answers option",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, true, true)),
			expEquals: false,
		},
		{
			name:      "Equals poll data",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: true,
		},
		{
			name:      "First nil",
			first:     nil,
			second:    pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer, answer2), true, false, true)),
			expEquals: false,
		},
		{
			name:      "Second nil",
			first:     pollDataPointer(polls.NewPollData("poll?", pollEndDate, polls.NewPollAnswers(answer), true, false, true)),
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
			require.Equal(t, test.expEquals, polls.ArePollDataEquals(test.first, test.second))
		})
	}
}

// ---------------
// --- UserAnswer
// ---------------

func TestUserAnswer_String(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := []polls.AnswerID{polls.AnswerID(1), polls.AnswerID(2)}

	userPollAnswers := polls.NewUserAnswer(answers, user)

	require.Equal(t, "User: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", userPollAnswers.String())
}

func TestUserAnswer_Validate(t *testing.T) {
	user, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	answers := []polls.AnswerID{polls.AnswerID(1), polls.AnswerID(2)}

	tests := []struct {
		name            string
		userPollAnswers polls.UserAnswer
		expErr          string
	}{
		{
			name:            "Empty user returns error",
			userPollAnswers: polls.NewUserAnswer(answers, nil),
			expErr:          "user cannot be empty",
		},
		{
			name:            "Empty answers returns error",
			userPollAnswers: polls.NewUserAnswer(nil, user),
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

	answers := []polls.AnswerID{polls.AnswerID(1), polls.AnswerID(2)}
	answers2 := []polls.AnswerID{polls.AnswerID(1)}

	tests := []struct {
		name      string
		first     polls.UserAnswer
		second    polls.UserAnswer
		expEquals bool
	}{
		{
			name:      "Different users returns false",
			first:     polls.NewUserAnswer(answers, user),
			second:    polls.NewUserAnswer(answers, user2),
			expEquals: false,
		},
		{
			name:      "Different answers lengths returns false",
			first:     polls.NewUserAnswer(answers, user),
			second:    polls.NewUserAnswer(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Different answers return false",
			first:     polls.NewUserAnswer(answers, user),
			second:    polls.NewUserAnswer(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Equals userPollAnswers returns true",
			first:     polls.NewUserAnswer(answers, user),
			second:    polls.NewUserAnswer(answers, user),
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

	answers := []polls.AnswerID{polls.AnswerID(1), polls.AnswerID(2)}
	answers2 := []polls.AnswerID{polls.AnswerID(3)}

	tests := []struct {
		name        string
		usersAD     polls.UserAnswers
		ansDet      polls.UserAnswer
		expUsersAD  polls.UserAnswers
		expAppended bool
	}{
		{
			name:    "Missing user answers details appended correctly",
			usersAD: polls.UserAnswers{polls.NewUserAnswer(answers, user)},
			ansDet:  polls.NewUserAnswer(answers, user2),
			expUsersAD: polls.UserAnswers{
				polls.NewUserAnswer(answers, user),
				polls.NewUserAnswer(answers, user2),
			},
			expAppended: true,
		},
		{
			name:        "Same user with different answers replace previous ones",
			usersAD:     polls.UserAnswers{polls.NewUserAnswer(answers, user)},
			ansDet:      polls.NewUserAnswer(answers2, user),
			expUsersAD:  polls.UserAnswers{polls.NewUserAnswer(answers2, user)},
			expAppended: true,
		},
		{
			name:        "Equals user answers details returns the same users answers details",
			usersAD:     polls.UserAnswers{polls.NewUserAnswer(answers, user)},
			ansDet:      polls.NewUserAnswer(answers, user),
			expUsersAD:  polls.UserAnswers{polls.NewUserAnswer(answers, user)},
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
