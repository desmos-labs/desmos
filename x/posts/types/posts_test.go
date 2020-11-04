package types_test

import (
	"github.com/desmos-labs/desmos/x/posts/types"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPost_NewPost(t *testing.T) {
	tests := []struct {
		post       types.Post
		expectedId string
	}{
		{
			post: types.NewPost(
				"1",
				"This is a message",
				true,
				"my_subspace",
				nil,
				time.Date(2020, 1, 1, 12, 0, 0, 0, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expectedId: "16ee282c835ef2bdfeea5bbd035eb3bea91c2b54f3ce2c28bbccfc9e4173f174",
		},
	}

	for index, test := range tests {
		test := test
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			require.Equal(t, test.expectedId, test.post.PostID)
		})
	}
}

func TestPost_String(t *testing.T) {
	tests := []struct {
		name      string
		post      types.Post
		expString string
	}{
		{
			name: "Post without attachments and poll data",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expString: "[ID] 7238e494e329def3bb2666e8a8fdd4bd3c64654f06c4a8e091be8c7cc441106d [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		},
		{
			name: "Post with attachments and without poll data",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
				types.NewAttachment(
					"https://another.com",
					"application/json",
					nil,
				),
			)),
			expString: "[ID] 7b23bc67a9d163134261b0c77eb262da96e783861977ebfc6100bdf408b8a995 [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Post Attachments]:\n [URI] [Mime-Type] [Tags]\n[https://uri.com] [text/plain] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\n] \n[https://another.com] [application/json] []",
		},
		{
			name: "Post without attachments and with poll data",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithPollData(types.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.FixedZone("UTC", 0)),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expString: "[ID] 048813e6e9f93a169234d10999e2f59bac14a364db11d4eb99309ac186cb78d9 [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Poll Data] Question: poll?\nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		},
		{
			name: "Post with attachments and with poll data",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
				types.NewAttachment(
					"https://another.com",
					"application/json",
					nil,
				),
			)).WithPollData(types.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.FixedZone("UTC", 0)),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expString: "[ID] 0b67394573f711c7e22d209abeab184b4eaadc7822d797232f959e40dfa3af0f [Parent ID] e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163 [Message] My post message [Creation Time] 2020-01-01 12:00:00 +0000 UTC [Edited Time] 0001-01-01 00:00:00 +0000 UTC [Allows Comments] true [Subspace] 4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e [Creator] cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns [Post Attachments]:\n [URI] [Mime-Type] [Tags]\n[https://uri.com] [text/plain] [cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns,\n] \n[https://another.com] [application/json] [] [Poll Data] Question: poll?\nEndDate: 2050-01-01 15:15:00 +0000 UTC\nAllow multiple answers: false \nAllow answer edits: true \nProvided Answers:\n[ID] [Text]\n[1] [Yes]\n[2] [No]",
		},
		{
			name: "Post with optional data",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"My post message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				[]types.OptionalDataEntry{
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
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
	tests := []struct {
		name     string
		post     types.Post
		expError string
	}{
		{
			name: "Invalid postID",
			post: types.Post{
				Message:        "Message",
				Created:        time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			expError: "invalid postID: ",
		},
		{
			name: "Invalid post owner",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expError: "invalid post owner: ",
		},
		{
			name: "Empty post message, attachment and poll",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Empty post message (blank), attachment and poll",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				" ",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Invalid post creation time",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Time{},
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expError: "invalid post creation time: 0001-01-01 00:00:00 +0000 UTC",
		},
		{
			name: "Invalid post last edit time",
			post: types.Post{
				PostID:   "dd065b70feb810a8c6f535cf670fe6e3534085221fa964ed2660ebca93f910d1",
				Creator:  "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Message:  "Message",
				Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Created:  time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				LastEdited: time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)).
					AddDate(0, 0, -1),
			},
			expError: "invalid post last edit time: 2019-12-31 12:00:00 +0000 UTC",
		},
		{
			name: "Invalid post subspace",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post subspace(blank)",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				" ",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("https://uri.com", "text/plain", nil),
			)).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post attachments",
			post: types.NewPost(
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("htp:/uri.com", "text/plain", nil)),
			),
			expError: "invalid uri provided",
		},
		{
			name: "Valid post without poll data",
			post: types.NewPost("",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("https://uri.com", "text/plain", nil),
			)),
			expError: "",
		},
		{
			name: "Valid post without attachs",
			post: types.NewPost("",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expError: "",
		},
		{
			name: "Valid post without text and attachs, but with poll",
			post: types.NewPost("",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithPollData(types.NewPollData(
				"poll?",
				time.Now().UTC().Add(time.Hour),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
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
	date := time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTF", 0))

	tests := []struct {
		name      string
		first     types.Post
		second    types.Post
		expEquals bool
	}{
		{
			name: "Different post ID",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different parent ID",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different message",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "Another post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different creation time",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date.AddDate(0, 0, 1),
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different last edited",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 2),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different allows comments",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: false,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different subspace",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-1",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "desmos-2",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different optional data",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []types.OptionalDataEntry{
					{
						Key:   "field",
						Value: "value",
					},
				},
				Creator: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData: []types.OptionalDataEntry{
					{
						Key:   "field",
						Value: "other-value",
					},
				},
				Creator: "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different owner",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			expEquals: false,
		},
		{
			name: "Different medias",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Attachments:    types.Attachments{},
			},
			expEquals: false,
		},
		{
			name: "Different polls",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
				PollData: nil,
			},
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Attachments: types.Attachments{
					types.NewAttachment("https://uri.com", "text/plain", nil),
				},
				PollData: &types.PollData{},
			},
			expEquals: false,
		},
		{
			name: "Equals posts",
			first: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			}.WithAttachments(types.Attachments{
				types.NewAttachment("https://uri.com", "text/plain", nil),
			}).WithPollData(types.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			second: types.Post{
				PostID:         "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				ParentID:       "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Message:        "My post message",
				Created:        date,
				LastEdited:     date.AddDate(0, 0, 1),
				AllowsComments: true,
				Subspace:       "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				OptionalData:   nil,
				Creator:        "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			}.WithAttachments(types.Attachments{
				types.NewAttachment("https://uri.com", "text/plain", nil),
			}).WithPollData(types.NewPollData(
				"poll?",
				time.Date(2050, 1, 1, 15, 15, 00, 000, timeZone),
				types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
				),
				false,
				true,
			)),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equal(test.second))
		})
	}
}

func TestPost_GetPostHashtags(t *testing.T) {
	tests := []struct {
		name        string
		post        types.Post
		expHashtags []string
	}{
		{
			name: "Hashtags in message extracted correctly (spaced hashtags)",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test #desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expHashtags: []string{"test", "desmos"},
		},
		{
			name: "Hashtags in message extracted correctly (non-spaced hashtags)",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (underscore separated hashtags)",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #test_#desmos",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (only number hashtag)",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with #101112",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expHashtags: []string{},
		},
		{
			name: "No hashtags in message",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with no hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expHashtags: []string{},
		},
		{
			name: "No same hashtags inside string array",
			post: types.NewPost(
				"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				"Post with double #hashtag #hashtag",
				false,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.FixedZone("UTC", 0)),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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

// ___________________________________________________________________________________________________________________

func TestPosts_String(t *testing.T) {
	date := time.Date(2020, 1, 1, 12, 0, 00, 000, time.FixedZone("UTC", 0))

	posts := types.Posts{
		types.NewPost(
			"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			"Post 1",
			false,
			"external-ref-1",
			nil,
			date,
			"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
		).WithAttachments(types.NewAttachments(
			types.NewAttachment("https://uri.com", "text/plain", nil),
		)).WithPollData(types.NewPollData(
			"poll?",
			date.Add(time.Hour),
			types.NewPollAnswers(
				types.NewPollAnswer("1", "Yes"),
				types.NewPollAnswer("2", "No"),
			),
			false,
			true,
		)),
		types.NewPost(
			"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
			"Post 2",
			false,
			"external-ref-1",
			nil,
			date,
			"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
		).WithAttachments(types.NewAttachments(
			types.NewAttachment("https://uri.com", "text/plain", nil),
		)).WithPollData(types.NewPollData(
			"poll?",
			date.Add(time.Hour),
			types.NewPollAnswers(
				types.NewPollAnswer("1", "Yes"),
				types.NewPollAnswer("2", "No"),
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

// ___________________________________________________________________________________________________________________

func TestAttachments_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Attachments
		second    types.Attachments
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
			),
			second: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
			),
			expEquals: true,
		},
		{
			name: "different data returns false",
			first: types.NewAttachments(
				types.NewAttachment("uri", "text/plain", nil),
			),
			second: types.NewAttachments(
				types.NewAttachment("uri", "application/json", nil),
			),
			expEquals: false,
		},
		{
			name: "different length returns false",
			first: types.NewAttachments(
				types.NewAttachment("uri", "text/plain", nil),
				types.NewAttachment("uri", "application/json", nil),
			),
			second: types.NewAttachments(
				types.NewAttachment("uri", "text/plain", nil),
			),
			expEquals: false,
		},
		{
			name: "different tags length returns false",
			first: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
			),
			second: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
			),
			expEquals: false,
		},
		{
			name: "different tags returns false",
			first: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
			),
			second: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
				),
				types.NewAttachment(
					"uri",
					"application/json",
					[]string{"cosmos16vphdl9nhm26murvfrrp8gdsknvfrxctl6y29h"},
				),
			),
			expEquals: false,
		},
		{
			name: "nil tags returns true",
			first: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					nil,
				),
				types.NewAttachment(
					"uri",
					"application/json",
					nil,
				),
			),
			second: types.NewAttachments(
				types.NewAttachment(
					"uri",
					"text/plain",
					nil,
				),
				types.NewAttachment(
					"uri",
					"application/json",
					nil,
				),
			),
			expEquals: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equal(test.second))
		})
	}
}

func TestAttachments_AppendIfMissing(t *testing.T) {
	tests := []struct {
		name        string
		medias      types.Attachments
		newMedia    types.Attachment
		expMedias   types.Attachments
		expAppended bool
	}{
		{
			name: "append a new attachment and returns true",
			medias: types.Attachments{
				types.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.Attachment{
				URI:      "uri",
				MimeType: "application/json",
			},
			expMedias: types.Attachments{
				types.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
				types.Attachment{
					URI:      "uri",
					MimeType: "application/json",
				},
			},
		},
		{
			name: "not append an existing attachment and returns false",
			medias: types.Attachments{
				types.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
			newMedia: types.Attachment{
				URI:      "uri",
				MimeType: "text/plain",
			},
			expMedias: types.Attachments{
				types.Attachment{
					URI:      "uri",
					MimeType: "text/plain",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			medias := test.medias.AppendIfMissing(test.newMedia)
			require.Equal(t, test.expMedias, medias)
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestPostMedia_Validate(t *testing.T) {
	tests := []struct {
		name      string
		postMedia types.Attachment
		expErr    string
	}{
		{
			name: "Empty URI",
			postMedia: types.NewAttachment(
				"",
				"text/plain",
				nil,
			),
			expErr: "invalid uri provided",
		},
		{
			name: "Invalid URI",
			postMedia: types.NewAttachment(
				"htt://example.com",
				"text/plain",
				nil,
			),
			expErr: "invalid uri provided",
		},
		{
			name: "Empty mime type",
			postMedia: types.NewAttachment(
				"https://example.com",
				"",
				nil,
			),
			expErr: "mime type must be specified and cannot be empty",
		},
		{
			name: "Invalid Tags",
			postMedia: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{""},
			),
			expErr: "invalid empty tag address: ",
		},
		{
			name: "No errors attachment (with tags)",
			postMedia: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			),
			expErr: "",
		},
		{
			name: "No errors attachment (without tags)",
			postMedia: types.NewAttachment(
				"https://example.com",
				"text/plain",
				nil,
			),
			expErr: "",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if len(test.expErr) != 0 {
				require.Equal(t, test.expErr, test.postMedia.Validate().Error())
			} else {
				require.Nil(t, test.postMedia.Validate())
			}
		})
	}
}

func TestPostMedia_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Attachment
		second    types.Attachment
		expEquals bool
	}{
		{
			name: "Same data returns true",
			first: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			),
			second: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			),
			expEquals: true,
		},
		{
			name: "Different URI returns false",
			first: types.NewAttachment(
				"https://example.com",
				"text/plain",
				nil,
			),
			second: types.NewAttachment(
				"https://another.com",
				"text/plain",
				nil,
			),
			expEquals: false,
		},
		{
			name: "Different mime type returns false",
			first: types.NewAttachment(
				"https://example.com",
				"text/plain",
				nil,
			),
			second: types.NewAttachment(
				"https://example.com",
				"application/json",
				nil,
			),
			expEquals: false,
		},
		{
			name: "Different tags returns false",
			first: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			),
			second: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{},
			),
			expEquals: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expEquals, test.first.Equal(test.second))
		})
	}
}

// ___________________________________________________________________________________________________________________

func TestOptionalData_Equals(t *testing.T) {
	tests := []struct {
		name         string
		optionalData types.OptionalDataEntry
		otherOpData  types.OptionalDataEntry
		expBool      bool
	}{
		{
			name:         "Different optional data returns false",
			optionalData: types.NewOptionalDataEntry("key", "value"),
			otherOpData:  types.NewOptionalDataEntry("key", "val"),
			expBool:      false,
		},
		{
			name:         "Same optional data returns true",
			optionalData: types.NewOptionalDataEntry("key", "value"),
			otherOpData:  types.NewOptionalDataEntry("key", "value"),
			expBool:      true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expBool, test.optionalData.Equal(test.otherOpData))
		})
	}
}

func TestOptionalData_String(t *testing.T) {
	opt := types.NewOptionalDataEntry("optional", "data")
	require.Equal(t, "[Key] [Value]\n[optional] [data]", opt.String())
}
