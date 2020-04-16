package types_test

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	"github.com/stretchr/testify/require"

	"testing"
)

// ---------------
// --- PollAnswer
// ---------------

func TestPollAnswer_String(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	require.Equal(t, `Answer - ID: 1 ; Text: Yes`, answer.String())
}

func TestPollAnswer_Validate(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(0), Text: ""}
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
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
			answer:      types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: types.AnswerID(2), Text: "Yes"},
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"},
			otherAnswer: types.PollAnswer{ID: types.AnswerID(1), Text: "No"},
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      types.PollAnswer{ID: types.AnswerID(1), Text: "yes"},
			otherAnswer: types.PollAnswer{ID: types.AnswerID(1), Text: "yes"},
			expEquals:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.answer.Equals(test.otherAnswer))
		})
	}
}

// ---------------
// --- UsersPollAnswers
// ---------------

func TestPollAnswers_String(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}
	answers := types.PollAnswers{answer, answer2}

	require.Equal(t, "Provided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]", answers.String())
}

func TestPollAnswers_Validate(t *testing.T) {
	tests := []struct {
		answers types.PollAnswers
		expErr  string
	}{
		{
			answers: types.PollAnswers{},
			expErr:  "poll answers must be at least two",
		},
		{
			answers: types.PollAnswers{types.NewPollAnswer(types.AnswerID(0), ""), types.NewPollAnswer(types.AnswerID(1), "")},
			expErr:  "answer text must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expErr, test.answers.Validate().Error())
	}
}

func TestPollAnswers_Equals(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}

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
			require.Equal(t, test.expEquals, test.answers.Equals(test.others))
		})
	}
}

func TestPollAnswers_AppendIfMissing(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}

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
			require.Equal(t, test.expLen, len(test.answers.AppendIfMissing(test.answer)))
		})
	}
}

func TestPollAnswers_ExtractAnswersIDs(t *testing.T) {
	answer := types.PollAnswer{ID: types.AnswerID(1), Text: "Yes"}
	answer2 := types.PollAnswer{ID: types.AnswerID(2), Text: "No"}

	expectedIDs := []types.AnswerID{1, 2}
	pollAnswers := types.PollAnswers{answer, answer2}

	actual := pollAnswers.ExtractAnswersIDs()

	require.Equal(t, expectedIDs, actual)
}
