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
			expError: "invalid post id: ",
		},
		{
			name: "Invalid post owner",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				types.NewAttachments(
					types.NewAttachment("https://uri.com", "text/plain", nil),
				),
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"",
			),
			expError: "invalid post owner: ",
		},
		{
			name: "Empty post message, attachment and poll",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				nil,
				nil,
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Empty post message (blank), attachment and poll",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				" ",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				nil,
				nil,
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post message, attachments or poll required, they cannot be all empty",
		},
		{
			name: "Invalid post creation time",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				types.NewAttachments(
					types.NewAttachment("https://uri.com", "text/plain", nil),
				),
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Time{},
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
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
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"",
				nil,
				types.NewAttachments(
					types.NewAttachment("https://uri.com", "text/plain", nil),
				),
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post subspace(blank)",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				" ",
				nil,
				types.NewAttachments(
					types.NewAttachment("https://uri.com", "text/plain", nil),
				),
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "post subspace must be a valid sha-256 hash",
		},
		{
			name: "Invalid post attachments",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
				"e1ba4807a15d8579f79cfd90a07fc015e6125565c9271eb94aded0b2ebf86163",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				types.NewAttachments(
					types.NewAttachment("htp:/uri.com", "text/plain", nil),
				),
				nil,
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "invalid uri provided",
		},
		{
			name: "Valid post without poll data",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				types.NewAttachments(
					types.NewAttachment("https://uri.com", "text/plain", nil),
				),
				nil,
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "",
		},
		{
			name: "Valid post without attachs",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "",
				"Message",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				nil,
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			expError: "",
		},
		{
			name: "Valid post without text and attachs, but with poll",
			post: types.NewPost(
				"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08", "",
				"",
				true,
				"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				nil,
				nil,
				types.NewPollData(
					"poll?",
					time.Now().UTC().Add(time.Hour),
					types.NewPollAnswers(
						types.NewPollAnswer("1", "Yes"),
						types.NewPollAnswer("2", "No"),
					),
					false,
					true,
				),
				time.Time{},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
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

func TestPost_GetPostHashtags(t *testing.T) {
	tests := []struct {
		name        string
		post        types.Post
		expHashtags []string
	}{
		{
			name: "Hashtags in message extracted correctly (spaced hashtags)",
			post: types.Post{
				Message: "Post with #test #desmos",
			},
			expHashtags: []string{"test", "desmos"},
		},
		{
			name: "Hashtags in message extracted correctly (non-spaced hashtags)",
			post: types.Post{
				Message: "Post with #test#desmos",
			},
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (underscore separated hashtags)",
			post: types.Post{
				Message: "Post with #test_#desmos",
			},
			expHashtags: []string{},
		},
		{
			name: "Hashtags in message extracted correctly (only number hashtag)",
			post: types.Post{
				Message: "Post with #101112",
			},
			expHashtags: []string{},
		},
		{
			name: "No hashtags in message",
			post: types.Post{
				Message: "Post with no hashtag",
			},
			expHashtags: []string{},
		},
		{
			name: "No same hashtags inside string array",
			post: types.Post{
				Message: "Post with double #hashtag #hashtag",
			},
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

func TestAttachments_Equal(t *testing.T) {
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

func TestAttachment_Validate(t *testing.T) {
	tests := []struct {
		name       string
		attachment types.Attachment
		expErr     string
	}{
		{
			name: "Empty URI",
			attachment: types.NewAttachment(
				"",
				"text/plain",
				nil,
			),
			expErr: "invalid uri provided",
		},
		{
			name: "Invalid URI",
			attachment: types.NewAttachment(
				"htt://example.com",
				"text/plain",
				nil,
			),
			expErr: "invalid uri provided",
		},
		{
			name: "Empty mime type",
			attachment: types.NewAttachment(
				"https://example.com",
				"",
				nil,
			),
			expErr: "mime type must be specified and cannot be empty",
		},
		{
			name: "Invalid Tags",
			attachment: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{""},
			),
			expErr: "invalid empty tag address: ",
		},
		{
			name: "No errors attachment (with tags)",
			attachment: types.NewAttachment(
				"https://example.com",
				"text/plain",
				[]string{"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"},
			),
			expErr: "",
		},
		{
			name: "No errors attachment (without tags)",
			attachment: types.NewAttachment(
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
				require.Equal(t, test.expErr, test.attachment.Validate().Error())
			} else {
				require.Nil(t, test.attachment.Validate())
			}
		})
	}
}
