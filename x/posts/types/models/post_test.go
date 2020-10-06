package models_test

import (
	"github.com/desmos-labs/desmos/x/posts/types/models/common"
	"strconv"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types/models"

	"github.com/stretchr/testify/require"
)

// -------------
// --- PostID
// -------------

func TestPostID_Equals(t *testing.T) {
	tests := []struct {
		name    string
		postID  models.PostID
		otherID models.PostID
		expBool bool
	}{
		{
			name:    "Equal IDs returns true",
			postID:  models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			otherID: models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			expBool: true,
		},
		{
			name:    "Non Equal IDs returns false",
			postID:  models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			otherID: models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			expBool: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			res := test.postID.Equals(test.otherID)
			require.Equal(t, test.expBool, res)
		})
	}

}

func TestPostID_String(t *testing.T) {
	postID := models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	require.Equal(t, "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd", postID.String())
}

// -------------
// --- PostIDs
// -------------

func TestPostIDs_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     models.PostIDs
		second    models.PostIDs
		expEquals bool
	}{
		{
			name: "Different length",
			first: models.PostIDs{
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			},
			second:    models.PostIDs{models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")},
			expEquals: false,
		},
		{
			name: "Different order",
			first: models.PostIDs{
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			},
			second: models.PostIDs{
				models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			},
			expEquals: false,
		},
		{
			name: "Same length and order",
			first: models.PostIDs{
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			},
			second: models.PostIDs{
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			},
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

func TestPostIDs_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name      string
		IDs       models.PostIDs
		newID     models.PostID
		expIDs    models.PostIDs
		expEdited bool
	}{
		{
			name:      "AppendIfMissing dont append anything",
			IDs:       models.PostIDs{models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")},
			newID:     models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
			expIDs:    models.PostIDs{models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")},
			expEdited: false,
		},
		{
			name:  "AppendIfMissing append something",
			IDs:   models.PostIDs{models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")},
			newID: models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			expIDs: models.PostIDs{
				models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"),
			},
			expEdited: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			newIDs, edited := test.IDs.AppendIfMissing(test.newID)
			require.Equal(t, test.expIDs, newIDs)
			require.Equal(t, test.expEdited, edited)
		})
	}
}

// -----------
// --- Post
// -----------

func TestPost_NewPost(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		post       models.Post
		expectedId string
	}{
		{
			post: models.NewPost(
				"1",
				"This is a message",
				true,
				"my_subspace",
				nil,
				time.Date(2020, 1, 1, 12, 0, 0, 0, time.FixedZone("UTC", 0)),
				owner,
			),
			expectedId: "16ee282c835ef2bdfeea5bbd035eb3bea91c2b54f3ce2c28bbccfc9e4173f174",
		},
	}

	for index, test := range tests {
		test := test
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, test.expectedId, test.post.PostID.String())
		})
	}
}

func TestPost_String(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name      string
		post      models.Post
		expString string
	}{
		{
			name: "Post without attachments and poll data",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expString: "[ID] 7238e494e329def3bb2666e8a8fdd4bd3c64654f06c4a8e091be8c7cc441106d [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		},
		{
			name: "Post with attachments and without poll data",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]sdk.AccAddress{owner},
				),
				models.NewAttachment(
					"https://another.com",
					"application/json",
					nil,
				),
			)),
			expString: "[ID] 7b23bc67a9d163134261b0c77eb262da96e783861977ebfc6100bdf408b8a995 [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Post Attachments]:\n [URI] [Mime-Type] [Tags]\n[https://uri.com] [text/plain] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\n] \n[https://another.com] [application/json] []",
		},
		{
			name: "Post without attachments and with poll data",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			).WithPollData(models.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.FixedZone("UTC", 0)),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expString: "[ID] 048813e6e9f93a169234d10999e2f59bac14a364db11d4eb99309ac186cb78d9 [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Poll Data] Question: poll?\nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		},
		{
			name: "Post with attachments and with poll data",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]sdk.AccAddress{owner},
				),
				models.NewAttachment(
					"https://another.com",
					"application/json",
					nil,
				),
			)).WithPollData(models.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.FixedZone("UTC", 0)),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expString: "[ID] 0b67394573f711c7e22d209abeab184b4eaadc7822d797232f959e40dfa3af0f [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Post Attachments]:\n [URI] [Mime-Type] [Tags]\n[https://uri.com] [text/plain] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\n] \n[https://another.com] [application/json] [] [Poll Data] Question: poll?\nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		},
		{
			name: "Post with optional data",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]common.OptionalData{
					{
						"key1",
						"value",
					},
					{
						"key2",
						"value",
					},
					{
						"key3",
						"value",
					},
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expString: "[ID] 4c07f3e94fddbee872f54adeeca2e24b82e7bcc9f04e19f40e03901169d29cc2 [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Optional Data] [[Key] [Value]\n[key1] [value] [Key] [Value]\n[key2] [value] [Key] [Value]\n[key3] [value]]",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expString, test.post.String())
		})
	}
}

func TestPost_Validate(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0))

	tests := []struct {
		name     string
		post     models.Post
		expError string
	}{
		{
			name: "Invalid postID",
			post: models.Post{
				Message:        "Message",
				Created:        date,
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
			},
			expError: "invalid postID: ",
		},
		{
			name: "Invalid post owner",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				nil,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "invalid post owner: ",
		},
		{
			name: "Empty post message, attachment and poll",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Empty post message (blank), attachment and poll",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				" ",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Invalid post creation time",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Time{},
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "invalid post creation time: 0001-01-01 00:00:00 +0000 UTC",
		},
		{
			name: "Invalid post last edit time",
			post: models.Post{
				PostID:     "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				Creator:    owner,
				Message:    "Message",
				Subspace:   "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Created:    date,
				LastEdited: date.AddDate(0, 0, -1),
			},
			expError: "invalid post last edit time: 2019-12-31 12:00:00 +0000 UTC",
		},
		{
			name: "Invalid post subspace",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"",
				nil,
				date,
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post subspace(blank)",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				" ",
				nil,
				date,
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post attachments",
			post: models.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("htp:/uri.com", "text/plain", nil)),
			),
			expError: "invalid uri provided",
		},
		{
			name: "Valid post without poll data",
			post: models.NewPost("",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			).WithAttachments(models.NewAttachments(
				models.NewAttachment("https://uri.com", "text/plain", nil),
			)),
			expError: "",
		},
		{
			name: "Valid post without attachs",
			post: models.NewPost("",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "",
		},
		{
			name: "Valid post without text and attachs, but with poll",
			post: models.NewPost("",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				date,
				owner,
			).WithPollData(models.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				models.NewPollAnswers(
					models.NewPollAnswer(models.AnswerID(1), "Yes"),
					models.NewPollAnswer(models.AnswerID(2), "No"),
				),
				false,
				true,
			)),
			expError: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if len(test.expError) != 0 {
				require.Equal(t, test.expError, test.post.Validate().Error())
			} else {
				require.Nil(t, test.post.Validate())
			}
		})
	}
}

func TestPost_Equals(t *testing.T) {
	id := models.PostID("19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")
	id2 := models.PostID("f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd")
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	otherOwner, err := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 00, 00, 000, timeZone)
	medias := models.Attachments{
		models.Attachment{
			URI:      "https://uri.com",
			MimeType: "text/plain",
		},
	}

	pollData := models.NewPollData(
		"poll?",
		time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
		models.NewPollAnswers(
			models.NewPollAnswer(models.AnswerID(1), "Yes"),
			models.NewPollAnswer(models.AnswerID(2), "No"),
		),
		false,
		true,
	)

	tests := []struct {
		name      string
		first     models.Post
		second    models.Post
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id2,
				ParentID:       id,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: models.Post{
				PostID:         id,
				ParentID:       id,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "Another post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date.AddDate(0, 0, 1),
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 2),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: false,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different subspace",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-1",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-2",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different optional data",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []common.OptionalData{
					{
						Key:   "field",
						Value: "value",
					},
				},
				Creator:     owner,
				Attachments: medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []common.OptionalData{
					{
						Key:   "field",
						Value: "other-value",
					},
				},
				Creator:     owner,
				Attachments: medias,
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        otherOwner,
				Attachments:    medias,
			},
			expEquals: false,
		},
		{
			name: "Different medias",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        otherOwner,
				Attachments:    models.Attachments{},
			},
			expEquals: false,
		},
		{
			name: "Different polls",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
				Attachments:    medias,
				PollData:       nil,
			},
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        otherOwner,
				Attachments:    medias,
				PollData:       &models.PollData{},
			},
			expEquals: false,
		},
		{
			name: "Equals posts",
			first: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
			}.WithAttachments(medias).WithPollData(pollData),
			second: models.Post{
				PostID:         id,
				ParentID:       id2,
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        owner,
			}.WithAttachments(medias).WithPollData(pollData),
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

func TestPost_GetPostHashtags(t *testing.T) {
	owner, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	tests := []struct {
		name        string
		post        models.Post
		expHashtags []string
	}{
		{
			name: "Hashtags in message extracted correctly (spaced hashtags)",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test #desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{"test", "desmos"},
		},
		{
			name: "Hashtags in message extracted correctly (non-spaced hashtags)",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (underscore separated hashtags)",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test_#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (only number hashtag)",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #101112",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "No hashtags in message",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with no hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{},
		},
		{
			name: "No same hashtags inside string array",
			post: models.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with double #hashtag #hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				owner,
			),
			expHashtags: []string{"hashtag"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			hashtags := test.post.GetPostHashtags()
			require.Equal(t, test.expHashtags, hashtags)
		})
	}
}

// -----------
// --- Posts
// -----------
func TestPosts_String(t *testing.T) {
	owner1, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	owner2, err := sdk.AccAddressFromBech32("cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l")
	require.NoError(t, err)

	date := time.Date(2020, 1, 1, 12, 0, 00, 000, time.FixedZone("UTC", 0))

	posts := models.Posts{
		models.NewPost(
			"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			"Post 1",
			false,
			"external-ref-1",
			nil,
			date,
			owner1,
		).WithAttachments(models.NewAttachments(
			models.NewAttachment("https://uri.com", "text/plain", nil),
		)).WithPollData(models.NewPollData(
			"poll?",
			date.Add(time.Hour),
			models.NewPollAnswers(
				models.NewPollAnswer(models.AnswerID(1), "Yes"),
				models.NewPollAnswer(models.AnswerID(2), "No"),
			),
			false,
			true,
		)),
		models.NewPost(
			"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			"Post 2",
			false,
			"external-ref-1",
			nil,
			date,
			owner2,
		).WithAttachments(models.NewAttachments(
			models.NewAttachment("https://uri.com", "text/plain", nil),
		)).WithPollData(models.NewPollData(
			"poll?",
			date.Add(time.Hour),
			models.NewPollAnswers(
				models.NewPollAnswer(models.AnswerID(1), "Yes"),
				models.NewPollAnswer(models.AnswerID(2), "No"),
			),
			false,
			true,
		)),
	}

	expected := `ID - [Creator] Message
66dc5e1537c731ca1c9866940b13027d2aa79277888a61c1bd295e8e944c8054 - [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns] Post 1
4807d764063aa5465901ca66b821a29e4f1b8560299c760b65afbb89d0004d86 - [cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l] Post 2`
	require.Equal(t, expected, posts.String())
}
