package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func TestPost_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		post      types.Post
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			post: types.NewPost(
				0,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			post: types.NewPost(
				1,
				0,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid entities returns error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, ""),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid author address returns error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid post reference returns error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid creation date returns error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Time{},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid last edited date returns error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Tag{
						types.NewTag(1, 1, "tag"),
					},
					[]types.Url{
						types.NewURL(1, 1, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				&time.Time{},
			),
			shouldErr: true,
		},
		{
			name: "valid post returns no error",
			post: types.NewPost(
				1,
				2,
				"External id",
				"Text",
				types.NewEntities(
					[]types.Tag{
						types.NewTag(1, 3, "tag"),
					},
					[]types.Tag{
						types.NewTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				[]types.PostReference{
					types.NewPostReference(types.TYPE_QUOTED, 1),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.post.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPostReference_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		reference types.PostReference
		shouldErr bool
	}{
		{
			name:      "invalid post id returns error",
			reference: types.NewPostReference(types.TYPE_QUOTED, 0),
			shouldErr: true,
		},
		{
			name:      "valid reference returns no error",
			reference: types.NewPostReference(types.TYPE_QUOTED, 1),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.reference.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestEntities_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		entities  *types.Entities
		shouldErr bool
	}{
		{
			name:      "empty entities returns error",
			entities:  types.NewEntities(nil, nil, nil),
			shouldErr: true,
		},
		{
			name: "invalid hashtag returns error",
			entities: types.NewEntities(
				[]types.Tag{
					types.NewTag(0, 0, ""),
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "mention returns error",
			entities: types.NewEntities(
				nil,
				[]types.Tag{
					types.NewTag(0, 0, ""),
				},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "invalid url returns error",
			entities: types.NewEntities(
				nil,
				nil,
				[]types.Url{
					types.NewURL(0, 0, "", ""),
				},
			),
			shouldErr: true,
		},
		{
			name: "overlapping hashtags return error",
			entities: types.NewEntities(
				[]types.Tag{
					types.NewTag(1, 5, "First tag"),
					types.NewTag(4, 10, "Second tag"),
				},
				nil,
				nil,
			),
			shouldErr: true,
		},
		{
			name: "overlapping mentions return error",
			entities: types.NewEntities(
				nil,
				[]types.Tag{
					types.NewTag(1, 5, "First mention"),
					types.NewTag(5, 10, "Second mention"),
				},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "overlapping URLs return error",
			entities: types.NewEntities(
				nil,
				nil,
				[]types.Url{
					types.NewURL(3, 4, "second url", "Second URL"),
					types.NewURL(1, 5, "first url", "First URL"),
				},
			),
			shouldErr: true,
		},
		{
			name: "overlapping hashtag and mention return error",
			entities: types.NewEntities(
				[]types.Tag{
					types.NewTag(1, 10, "First tag"),
				},
				[]types.Tag{
					types.NewTag(9, 15, "First mention"),
				},
				nil,
			),
			shouldErr: true,
		},
		{
			name: "overlapping hashtag and url return error",
			entities: types.NewEntities(
				[]types.Tag{
					types.NewTag(1, 10, "First tag"),
				},
				nil,
				[]types.Url{
					types.NewURL(1, 5, "first url", "First URL"),
				},
			),
			shouldErr: true,
		},
		{
			name: "overlapping mention and url return error",
			entities: types.NewEntities(
				nil,
				[]types.Tag{
					types.NewTag(8, 30, "First mention"),
				},
				[]types.Url{
					types.NewURL(1, 15, "first url", "First URL"),
				},
			),
			shouldErr: true,
		},
		{
			name: "valid entities returns no error",
			entities: types.NewEntities(
				[]types.Tag{
					types.NewTag(1, 2, "first_tag"),
				},
				[]types.Tag{
					types.NewTag(3, 4, "first_mention"),
				},
				[]types.Url{
					types.NewURL(5, 6, "url", "Display URL"),
				},
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.entities.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestTag_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		tag       types.Tag
		shouldErr bool
	}{
		{
			name:      "invalid start and end values return error",
			tag:       types.NewTag(1, 0, "My tag"),
			shouldErr: true,
		},
		{
			name:      "invalid tag value returns error",
			tag:       types.NewTag(1, 10, "   "),
			shouldErr: true,
		},
		{
			name:      "valid tag returns no error",
			tag:       types.NewTag(1, 10, "My tag"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.tag.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestUrl_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		url       types.Url
		shouldErr bool
	}{
		{
			name:      "invalid start and end values returns error",
			url:       types.NewURL(10, 1, "url", ""),
			shouldErr: true,
		},
		{
			name:      "invalid url value returns error",
			url:       types.NewURL(1, 10, "", "Display URL"),
			shouldErr: true,
		},
		{
			name:      "valid url returns no error",
			url:       types.NewURL(1, 10, "ftp://user:password@example.com", "Display URL"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.url.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestAttachment_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		attachment  types.Attachment
		shouldError bool
	}{
		{
			name: "invalid poll attachment returns error",
			attachment: types.NewPollAttachment(0, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			)),
			shouldError: true,
		},
		{
			name: "valid poll attachment returns no error",
			attachment: types.NewPollAttachment(1, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			)),
			shouldError: false,
		},
		{
			name: "invalid media attachment returns error",
			attachment: types.NewMediaAttachment(0, types.NewMedia(
				"ftp://user:password@example.com/image.png",
				"image/png",
			)),
			shouldError: true,
		},
		{
			name: "valid media attachment returns no error",
			attachment: types.NewMediaAttachment(1, types.NewMedia(
				"ftp://user:password@example.com/image.png",
				"image/png",
			)),
			shouldError: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.attachment.Validate()
			if tc.shouldError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestAttachments_Validate(t *testing.T) {
	testCases := []struct {
		name        string
		attachments types.Attachments
		shouldErr   bool
	}{
		{
			name: "duplicated attachment id returns error",
			attachments: types.Attachments{
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
			},
			shouldErr: true,
		},
		{
			name:        "empty attachments return no error",
			attachments: types.Attachments{},
			shouldErr:   false,
		},
		{
			name: "valid attachments return no error",
			attachments: types.Attachments{
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
				types.NewMediaAttachment(2, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
			},
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.attachments.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestMedia_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		media     types.Media
		shouldErr bool
	}{
		{
			name:      "invalid uri returns error",
			media:     types.NewMedia("", "image/png"),
			shouldErr: true,
		},
		{
			name:      "invalid mime type returns error",
			media:     types.NewMedia("ftp://user:password@example.com/image.png", ""),
			shouldErr: true,
		},
		{
			name:      "valid media returns no error",
			media:     types.NewMedia("ftp://user:password@example.com/image.png", "image/png"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.media.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestPoll_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		poll      types.Poll
		shouldErr bool
	}{
		{
			name: "invalid question returns error",
			poll: types.NewPoll(
				"",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "empty provided answers return error",
			poll: types.NewPoll(
				"What animal is best?",
				nil,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "single provided answer returns error",
			poll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "invalid provided answer returns error",
			poll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "duplicated provided answers return error",
			poll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Cat", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "invalid end date returns error",
			poll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Time{},
				false,
				false,
			),
			shouldErr: true,
		},
		{
			name: "valid poll returns no error",
			poll: types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				false,
				false,
			),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.poll.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestPoll_ProvidedAnswer_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		answer    types.Poll_ProvidedAnswer
		shouldErr bool
	}{
		{
			name:      "invalid text returns error",
			answer:    types.NewProvidedAnswer("", nil),
			shouldErr: true,
		},
		{
			name: "invalid attachment returns error",
			answer: types.NewProvidedAnswer("Cat", []types.Attachment{
				types.NewMediaAttachment(0, types.NewMedia("", "")),
			}),
			shouldErr: true,
		},
		{
			name: "duplicated attachment returns error",
			answer: types.NewProvidedAnswer("Cat", []types.Attachment{
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
			}),
			shouldErr: true,
		},
		{
			name: "valid answer returns no error",
			answer: types.NewProvidedAnswer("Cat", []types.Attachment{
				types.NewMediaAttachment(1, types.NewMedia("ftp://user:password@example.com/image.png", "image/png")),
			}),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.answer.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestUserAnswer_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		answer    types.UserAnswer
		shouldErr bool
	}{
		{
			name:      "invalid subspace id returns error",
			answer:    types.NewUserAnswer(0, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			answer:    types.NewUserAnswer(1, 0, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: true,
		},
		{
			name:      "invalid poll id returns error",
			answer:    types.NewUserAnswer(1, 1, 0, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: true,
		},
		{
			name:      "empty answer indexes returns error",
			answer:    types.NewUserAnswer(1, 1, 1, nil, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: true,
		},
		{
			name:      "duplicated answer indexes returns error",
			answer:    types.NewUserAnswer(1, 1, 1, []uint32{1, 1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: true,
		},
		{
			name:      "invalid user address returns error",
			answer:    types.NewUserAnswer(1, 1, 1, []uint32{1}, ""),
			shouldErr: true,
		},
		{
			name:      "valid answer returns no error",
			answer:    types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.answer.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestPollTallyResults_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		results   types.PollTallyResults
		shouldErr bool
	}{
		{
			name: "invalid subspace id returns error",
			results: types.NewPollTallyResults(0, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(2, 10),
			}),
			shouldErr: true,
		},
		{
			name: "invalid post id returns error",
			results: types.NewPollTallyResults(1, 0, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(2, 10),
			}),
			shouldErr: true,
		},
		{
			name: "invalid poll id returns error",
			results: types.NewPollTallyResults(1, 1, 0, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(2, 10),
			}),
			shouldErr: true,
		},
		{
			name:      "empty answer results return error",
			results:   types.NewPollTallyResults(1, 1, 1, nil),
			shouldErr: true,
		},
		{
			name: "duplicated answer results return error",
			results: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(1, 10),
			}),
			shouldErr: true,
		},
		{
			name: "valid tally results return no error",
			results: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(1, 10),
				types.NewAnswerResult(2, 10),
			}),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.results.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
