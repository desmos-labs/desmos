package polls_test

import (
	"github.com/desmos-labs/desmos/x/posts/internal/types/models/polls"
	"github.com/stretchr/testify/require"

	"testing"
)

// ---------------
// --- PollAnswer
// ---------------

func TestPollAnswer_String(t *testing.T) {
	answer := polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"}
	require.Equal(t, `Answer - ID: 1 ; Text: Yes`, answer.String())
}

func TestPollAnswer_Validate(t *testing.T) {
	answer := polls.PollAnswer{ID: polls.AnswerID(0), Text: ""}
	require.Equal(t, "answer text must be specified and cannot be empty", answer.Validate().Error())
}

func TestPollAnswer_Equals(t *testing.T) {
	tests := []struct {
		name        string
		answer      polls.PollAnswer
		otherAnswer polls.PollAnswer
		expEquals   bool
	}{
		{
			name:        "Different answers ID",
			answer:      polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"},
			otherAnswer: polls.PollAnswer{ID: polls.AnswerID(2), Text: "Yes"},
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"},
			otherAnswer: polls.PollAnswer{ID: polls.AnswerID(1), Text: "No"},
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      polls.PollAnswer{ID: polls.AnswerID(1), Text: "yes"},
			otherAnswer: polls.PollAnswer{ID: polls.AnswerID(1), Text: "yes"},
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
// --- PollAnswers
// ---------------

func TestPollAnswers_String(t *testing.T) {
	answer := polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"}
	answer2 := polls.PollAnswer{ID: polls.AnswerID(2), Text: "No"}
	answers := polls.PollAnswers{answer, answer2}

	require.Equal(t, "Provided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]", answers.String())
}

func TestPollAnswers_Validate(t *testing.T) {
	tests := []struct {
		answers polls.PollAnswers
		expErr  string
	}{
		{
			answers: polls.PollAnswers{},
			expErr:  "poll answers must be at least two",
		},
		{
			answers: polls.PollAnswers{polls.NewPollAnswer(polls.AnswerID(0), ""), polls.NewPollAnswer(polls.AnswerID(1), "")},
			expErr:  "answer text must be specified and cannot be empty",
		},
	}

	for _, test := range tests {
		require.Equal(t, test.expErr, test.answers.Validate().Error())
	}
}

func TestPollAnswers_Equals(t *testing.T) {
	answer := polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"}
	answer2 := polls.PollAnswer{ID: polls.AnswerID(2), Text: "No"}

	tests := []struct {
		name      string
		answers   polls.PollAnswers
		others    polls.PollAnswers
		expEquals bool
	}{
		{
			name:      "Different lengths",
			answers:   polls.PollAnswers{answer},
			others:    polls.PollAnswers{},
			expEquals: false,
		},
		{
			name:      "Different answers",
			answers:   polls.PollAnswers{answer},
			others:    polls.PollAnswers{answer2},
			expEquals: false,
		},
		{
			name:      "Equals answers",
			answers:   polls.PollAnswers{answer},
			others:    polls.PollAnswers{answer},
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
	answer := polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"}
	answer2 := polls.PollAnswer{ID: polls.AnswerID(2), Text: "No"}

	tests := []struct {
		name    string
		answers polls.PollAnswers
		answer  polls.PollAnswer
		expLen  int
	}{
		{
			name:    "Appended new answer",
			answers: polls.PollAnswers{answer},
			answer:  answer2,
			expLen:  2,
		},
		{
			name:    "Not appended an existing answer",
			answers: polls.PollAnswers{answer},
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
	answer := polls.PollAnswer{ID: polls.AnswerID(1), Text: "Yes"}
	answer2 := polls.PollAnswer{ID: polls.AnswerID(2), Text: "No"}

	expectedIDs := []polls.AnswerID{1, 2}
	pollAnswers := polls.PollAnswers{answer, answer2}

	actual := pollAnswers.ExtractAnswersIDs()

	require.Equal(t, expectedIDs, actual)
}
