package types_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

// ---------------
// --- PollData
// ---------------

func TestPollData_String(t *testing.T) {
	answer := types.PollAnswer{ID: uint(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint(2), Text: "No"}
	pollData := types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true)

	assert.Equal(t, "Question: poll? \nOpen: true \nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		pollData.String())
}

func TestPollData_Validate(t *testing.T) {
	timeZone, _ := time.LoadLocation("UTC")

	tests := []struct {
		pollData *types.PollData
		expError string
	}{
		{
			pollData: types.NewPollData("", testPostEndPollDate, types.PollAnswers{}, true, true, true),
			expError: "missing poll title",
		},
		{
			pollData: types.NewPollData("title", time.Date(2019, 1, 1, 0, 0, 00, 000, timeZone), types.PollAnswers{}, true, true, true),
			expError: "end date cannot be in the past",
		},
		{
			pollData: types.NewPollData("title", testPostEndPollDate, types.PollAnswers{}, true, true, true),
			expError: "answers cannot be empty",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expError, test.pollData.Validate().Error())
	}
}

func TestPollData_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     *types.PollData
		second    *types.PollData
		expEquals bool
	}{
		{
			name:      "Different titles",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			expEquals: false,
		},
		{
			name:      "Different open",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, false, false, true),
			expEquals: false,
		},
		{
			name:      "Different end date",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll?", time.Now().UTC(), types.PollAnswers{answer, answer2}, true, false, true),
			expEquals: false,
		},
		{
			name:      "Different provided answers",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer}, true, false, true),
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			expEquals: false,
		},
		{
			name:      "Different edits answer option",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, false),
			expEquals: false,
		},
		{
			name:      "Different multiple answers option",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, true, true),
			expEquals: false,
		},
		{
			name:      "Equals poll data",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			expEquals: true,
		},
		{
			name:      "First nil",
			first:     nil,
			second:    types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true),
			expEquals: false,
		},
		{
			name:      "Second nil",
			first:     types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer}, true, false, true),
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
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

// ---------------
// --- PollAnswers
// ---------------

func TestPollAnswers_String(t *testing.T) {
	answer := types.PollAnswer{ID: uint(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint(2), Text: "No"}
	answers := types.PollAnswers{answer, answer2}

	assert.Equal(t, "Provided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]", answers.String())
}

func TestPollAnswers_Validate(t *testing.T) {
	tests := []struct {
		answers types.PollAnswers
		expErr  string
	}{
		{
			answers: types.PollAnswers{},
			expErr:  "answers cannot be empty",
		},
		{
			answers: types.PollAnswers{types.PollAnswer{ID: uint(0), Text: ""}},
			expErr:  "answer text must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expErr, test.answers.Validate().Error())
	}
}

func TestPollAnswers_Equals(t *testing.T) {
	answer := types.PollAnswer{ID: uint(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint(2), Text: "No"}

	tests := []struct {
		name      string
		answers   types.PollAnswers
		others    types.PollAnswers
		expEquals bool
	}{
		{
			name:      "Different lengths",
			answers:   types.PollAnswers{answer},
			others:    types.PollAnswers{},
			expEquals: false,
		},
		{
			name:      "Different answers",
			answers:   types.PollAnswers{answer},
			others:    types.PollAnswers{answer2},
			expEquals: false,
		},
		{
			name:      "Equals answers",
			answers:   types.PollAnswers{answer},
			others:    types.PollAnswers{answer},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.answers.Equals(test.others))
		})
	}
}

func TestPollAnswers_AppendIfMissing(t *testing.T) {
	answer := types.PollAnswer{ID: uint(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint(2), Text: "No"}

	tests := []struct {
		name    string
		answers types.PollAnswers
		answer  types.PollAnswer
		expLen  int
	}{
		{
			name:    "Appended new answer",
			answers: types.PollAnswers{answer},
			answer:  answer2,
			expLen:  2,
		},
		{
			name:    "Not appended an existing answer",
			answers: types.PollAnswers{answer},
			answer:  answer,
			expLen:  1,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expLen, len(test.answers.AppendIfMissing(test.answer)))
		})
	}
}

// ---------------
// --- PollAnswer
// ---------------

func TestPollAnswer_String(t *testing.T) {
	answer := types.PollAnswer{ID: uint(1), Text: "Yes"}
	assert.Equal(t, `Answer - ID: 1 ; Text: Yes`, answer.String())
}

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.PollAnswer{ID: uint(0), Text: ""}
	assert.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

func TestPollAnswer_Equals(t *testing.T) {
	tests := []struct {
		name        string
		answer      types.PollAnswer
		otherAnswer types.PollAnswer
		expEquals   bool
	}{
		{
			name:        "Different answers ID",
			answer:      types.PollAnswer{ID: uint(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: uint(2), Text: "Yes"},
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      types.PollAnswer{ID: uint(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: uint(1), Text: "No"},
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      types.PollAnswer{ID: uint(1), Text: "yes"},
			otherAnswer: types.PollAnswer{ID: uint(1), Text: "yes"},
			expEquals:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.answer.Equals(test.otherAnswer))
		})
	}
}

// ---------------
// --- AnswersDetails
// ---------------
func TestUserPollAnswers_String(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	answers := []uint{uint(1), uint(2)}

	userPollAnswers := types.NewAnswersDetails(answers, user)

	assert.Equal(t, "User: cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns \nAnswers IDs: 1 2", userPollAnswers.String())
}

func TestUserPollAnswers_Validate(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	answers := []uint{uint(1), uint(2)}

	tests := []struct {
		name            string
		userPollAnswers types.AnswersDetails
		expErr          string
	}{
		{
			name:            "Empty user returns error",
			userPollAnswers: types.NewAnswersDetails(answers, nil),
			expErr:          "user cannot be empty",
		},
		{
			name:            "Empty answers returns error",
			userPollAnswers: types.NewAnswersDetails(nil, user),
			expErr:          "answers cannot be empty",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if err := test.userPollAnswers.Validate(); err != nil {
				assert.Equal(t, test.expErr, err.Error())
			}
		})
	}
}

func TestUserPollAnswers_Equals(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	user2, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	answers := []uint{uint(1), uint(2)}
	answers2 := []uint{uint(1)}

	tests := []struct {
		name      string
		first     types.AnswersDetails
		second    types.AnswersDetails
		expEquals bool
	}{
		{
			name:      "Different users returns false",
			first:     types.NewAnswersDetails(answers, user),
			second:    types.NewAnswersDetails(answers, user2),
			expEquals: false,
		},
		{
			name:      "Different answers lengths returns false",
			first:     types.NewAnswersDetails(answers, user),
			second:    types.NewAnswersDetails(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Different answers return false",
			first:     types.NewAnswersDetails(answers, user),
			second:    types.NewAnswersDetails(answers2, user2),
			expEquals: false,
		},
		{
			name:      "Equals userPollAnswers returns true",
			first:     types.NewAnswersDetails(answers, user),
			second:    types.NewAnswersDetails(answers, user),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

// ---------------
// --- UsersAnswersDetails
// ---------------

func TestUsersAnswersDetails(t *testing.T) {
	user, _ := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	user2, _ := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	answers := []uint{uint(1), uint(2)}
	answers2 := []uint{uint(3)}

	tests := []struct {
		name        string
		usersAD     types.UsersAnswersDetails
		ansDet      types.AnswersDetails
		expUsersAD  types.UsersAnswersDetails
		expAppended bool
	}{
		{
			name:    "Missing user answers details appended correctly",
			usersAD: types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
			ansDet:  types.NewAnswersDetails(answers, user2),
			expUsersAD: types.UsersAnswersDetails{
				types.NewAnswersDetails(answers, user),
				types.NewAnswersDetails(answers, user2),
			},
			expAppended: true,
		},
		{
			name:        "Same user with different answers replace previous ones",
			usersAD:     types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
			ansDet:      types.NewAnswersDetails(answers2, user),
			expUsersAD:  types.UsersAnswersDetails{types.NewAnswersDetails(answers2, user)},
			expAppended: true,
		},
		{
			name:        "Equals user answers details returns the same users answers details",
			usersAD:     types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
			ansDet:      types.NewAnswersDetails(answers, user),
			expUsersAD:  types.UsersAnswersDetails{types.NewAnswersDetails(answers, user)},
			expAppended: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual, appended := test.usersAD.AppendIfMissingOrIfUsersEquals(test.ansDet)
			assert.Equal(t, test.expUsersAD, actual)
			assert.Equal(t, test.expAppended, appended)
		})
	}
}
