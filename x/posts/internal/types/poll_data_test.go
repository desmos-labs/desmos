package types_test

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/assert"

	"testing"
	"time"
)

// ---------------
// --- PollData
// ---------------

func TestPollData_String(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}
	pollData := types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true)

	assert.Equal(t, `{"title":"poll?","provided_answers":[{"id":1,"text":"Yes"},{"id":2,"text":"No"}],"end_date":"2050-01-01T15:15:00Z","open":true,"allows_multiple_answers":false,"allows_answer_edits":true}`,
		pollData.String())
}

func TestPollData_MarshalJSON(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}
	pollData := types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true)

	json := types.ModuleCdc.MustMarshalJSON(pollData)
	assert.Equal(t, `{"title":"poll?","provided_answers":[{"id":1,"text":"Yes"},{"id":2,"text":"No"}],"end_date":"2050-01-01T15:15:00Z","open":true,"allows_multiple_answers":false,"allows_answer_edits":true}`, string(json))
}

func TestPollData_UnmarshalJSON(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}
	pollData := types.NewPollData("poll?", testPostEndPollDate, types.PollAnswers{answer, answer2}, true, false, true)

	tests := []struct {
		name        string
		strPollData string
		expPollData types.PollData
		expErr      string
	}{
		{
			name:        "Invalid poll data returns error",
			strPollData: "invalid poll data",
			expPollData: types.PollData{},
			expErr:      "invalid character 'i' looking for beginning of value",
		}, {
			name:        "Valid poll data is read properly",
			strPollData: `{"title":"poll?","provided_answers":[{"id":1,"text":"Yes"},{"id":2,"text":"No"}],"end_date":"2050-01-01T15:15:00Z","open":true,"allows_multiple_answers":false,"allows_answer_edits":true}`,
			expPollData: *pollData,
			expErr:      "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			var pollData types.PollData
			err := types.ModuleCdc.UnmarshalJSON([]byte(test.strPollData), &pollData)
			if err != nil {
				assert.Equal(t, test.expErr, err.Error())
			} else {
				assert.Equal(t, pollData, test.expPollData)
			}
		})
	}
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
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expEquals, test.first.Equals(*test.second))
		})
	}
}

// ---------------
// --- PollAnswers
// ---------------

func TestPollAnswers_String(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}
	answers := types.PollAnswers{answer, answer2}

	assert.Equal(t, "Answers\n[ID] [Text]\n[1] [Yes]\n[2] [No]", answers.String())
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
			answers: types.PollAnswers{types.PollAnswer{ID: uint64(0), Text: ""}},
			expErr:  "answer text must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expErr, test.answers.Validate().Error())
	}
}

func TestPollAnswers_Equals(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}

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
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: uint64(2), Text: "No"}

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
	answer := types.PollAnswer{ID: uint64(1), Text: "Yes"}
	assert.Equal(t, `Answer - ID: 1 ; Text: Yes`, answer.String())
}

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.PollAnswer{ID: uint64(1), Text: ""}
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
			answer:      types.PollAnswer{ID: uint64(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: uint64(2), Text: "Yes"},
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      types.PollAnswer{ID: uint64(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: uint64(1), Text: "No"},
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      types.PollAnswer{ID: uint64(1), Text: "yes"},
			otherAnswer: types.PollAnswer{ID: uint64(1), Text: "yes"},
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
