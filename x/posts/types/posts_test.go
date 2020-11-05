package types_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/x/posts/types"

	"github.com/stretchr/testify/require"
)

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
				Created:        time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				Created:  time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				LastEdited: time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC).
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			).WithAttachments(types.NewAttachments(
				types.NewAttachment("htp:/uri.com", "text/plain", nil)),
			),
			expError: "invalid uri provided",
		},
		{
			name: "Valid post without poll poll",
			post: types.NewPost("",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
			name: "Different optional poll",
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
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
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
				time.Date(2050, 1, 1, 15, 15, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
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

func TestAttachments_Equals(t *testing.T) {
	tests := []struct {
		name      string
		first     types.Attachments
		second    types.Attachments
		expEquals bool
	}{
		{
			name: "Same poll returns true",
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
			name: "different poll returns false",
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
			name: "Same poll returns true",
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
			name:         "Different optional poll returns false",
			optionalData: types.NewOptionalDataEntry("key", "value"),
			otherOpData:  types.NewOptionalDataEntry("key", "val"),
			expBool:      false,
		},
		{
			name:         "Same optional poll returns true",
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
