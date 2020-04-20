package v030_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v030posts "github.com/desmos-labs/desmos/x/posts/legacy/v0.3.0"
	"github.com/stretchr/testify/require"
)

// POST ID

func TestParsePostID(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		expID    v030posts.PostID
		expError string
	}{
		{
			name:     "Invalid id returns error",
			value:    "id",
			expID:    v030posts.PostID(0),
			expError: "strconv.ParseUint: parsing \"id\": invalid syntax",
		},
		{
			name:     "Valid id returns proper value",
			value:    "10",
			expID:    v030posts.PostID(10),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			id, err := v030posts.ParsePostID(test.value)

			if err == nil {
				require.Equal(t, test.expID, id)
			} else {
				require.Equal(t, test.expError, err.Error())
			}
		})
	}
}

func TestPostID_String(t *testing.T) {
	postID := v030posts.PostID(10)
	actual := postID.String()
	require.Equal(t, "10", actual)
}

// POST

func TestPost_ConflictWith(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	post030 := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		OptionalData:   map[string]string{},
		Created:        time.Now().UTC(),
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner,
	}

	post2 := v030posts.Post{
		PostID:         v030posts.PostID(2),
		ParentID:       v030posts.PostID(1),
		Message:        "Message",
		AllowsComments: true,
		Subspace:       "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
		OptionalData:   map[string]string{},
		Created:        time.Now().UTC(),
		LastEdited:     time.Now().UTC().Add(time.Hour),
		Creator:        owner,
	}

	tests := []struct {
		name      string
		post      v030posts.Post
		otherPost v030posts.Post
		expBool   bool
	}{
		{
			name:      "non conflict posts",
			post:      post030,
			otherPost: post2,
			expBool:   false,
		},
		{
			name:      "conflict posts",
			post:      post030,
			otherPost: post030,
			expBool:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.post.ConflictsWith(test.otherPost))
		})
	}
}

// POST MEDIA

func TestPostMedia_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     v030posts.PostMedia
		second    v030posts.PostMedia
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: v030posts.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: v030posts.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			expEquals: true,
		},
		{
			name: "Different URI returns false",
			first: v030posts.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: v030posts.PostMedia{
				URI:      "https://another.com",
				MimeType: "text/plain",
			},
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: v030posts.PostMedia{
				URI:      "https://example.com",
				MimeType: "text/plain",
			},
			second: v030posts.PostMedia{
				URI:      "https://example.com",
				MimeType: "application/json",
			},
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

func TestPostMedias_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     v030posts.PostMedias
		second    v030posts.PostMedias
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: true,
		},
		{
			name: "different data returns false",
			first: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			expEquals: false,
		},
		{
			name: "different length returns false",
			first: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
			second: v030posts.PostMedias{
				v030posts.PostMedia{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equals(test.second))
		})
	}
}

// POLL ANSWER

func TestPollAnswer_Equals(t *testing.T) {
	tests := []struct {
		name        string
		answer      v030posts.PollAnswer
		otherAnswer v030posts.PollAnswer
		expEquals   bool
	}{
		{
			name:        "Different answers ID",
			answer:      v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "Yes"},
			otherAnswer: v030posts.PollAnswer{ID: v030posts.AnswerID(2), Text: "Yes"},
			expEquals:   false,
		},
		{
			name:        "Different answers Text",
			answer:      v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "Yes"},
			otherAnswer: v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "No"},
			expEquals:   false,
		},
		{
			name:        "Equals answers",
			answer:      v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "yes"},
			otherAnswer: v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "yes"},
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

func TestPollAnswers_Equals(t *testing.T) {
	answer := v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "Yes"}
	answer2 := v030posts.PollAnswer{ID: v030posts.AnswerID(2), Text: "No"}

	tests := []struct {
		name      string
		answers   v030posts.PollAnswers
		others    v030posts.PollAnswers
		expEquals bool
	}{
		{
			name:      "Different lengths",
			answers:   v030posts.PollAnswers{answer},
			others:    v030posts.PollAnswers{},
			expEquals: false,
		},
		{
			name:      "Different answers",
			answers:   v030posts.PollAnswers{answer},
			others:    v030posts.PollAnswers{answer2},
			expEquals: false,
		},
		{
			name:      "Equals answers",
			answers:   v030posts.PollAnswers{answer},
			others:    v030posts.PollAnswers{answer},
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

// POLL DATA

func TestArePollDataEquals(t *testing.T) {
	answer := v030posts.PollAnswer{ID: v030posts.AnswerID(1), Text: "Yes"}
	answer2 := v030posts.PollAnswer{ID: v030posts.AnswerID(2), Text: "No"}

	tests := []struct {
		name      string
		first     v030posts.PollData
		second    v030posts.PollData
		expEquals bool
	}{
		{
			name: "Different titles",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll??",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			expEquals: false,
		},
		{
			name: "Different open",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              false,
				AllowsAnswerEdits: true,
			},
			expEquals: false,
		},
		{
			name: "Different end date",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC().Add(time.Hour),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			expEquals: false,
		},
		{
			name: "Different provided answers",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			expEquals: false,
		},
		{
			name: "Different edits answer option",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: false,
			},
			expEquals: false,
		},
		{
			name: "Different multiple answers option",
			first: v030posts.PollData{
				Question:              "poll?",
				ProvidedAnswers:       v030posts.PollAnswers{answer, answer2},
				EndDate:               time.Now().UTC(),
				Open:                  true,
				AllowsAnswerEdits:     true,
				AllowsMultipleAnswers: true,
			},
			second: v030posts.PollData{
				Question:              "poll?",
				ProvidedAnswers:       v030posts.PollAnswers{answer, answer2},
				EndDate:               time.Now().UTC(),
				Open:                  true,
				AllowsAnswerEdits:     true,
				AllowsMultipleAnswers: false,
			},
			expEquals: false,
		},
		{
			name: "Equals poll data",
			first: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			second: v030posts.PollData{
				Question:          "poll?",
				ProvidedAnswers:   v030posts.PollAnswers{answer, answer2},
				EndDate:           time.Now().UTC(),
				Open:              true,
				AllowsAnswerEdits: true,
			},
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, v030posts.ArePollDataEquals(&test.first, &test.second))
		})
	}
}
