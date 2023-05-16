package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

type mockHooks struct {
	CalledMap map[string]bool
}

func newMockHooks() *mockHooks {
	return &mockHooks{CalledMap: make(map[string]bool)}
}

var _ types.PostsHooks = &mockHooks{}

func (h *mockHooks) AfterPostSaved(ctx sdk.Context, subspaceID uint64, postID uint64) {
	h.CalledMap["AfterPostSaved"] = true
}

func (h *mockHooks) AfterPostDeleted(ctx sdk.Context, subspaceID uint64, postID uint64) {
	h.CalledMap["AfterPostDeleted"] = true
}

func (h *mockHooks) AfterAttachmentSaved(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	h.CalledMap["AfterAttachmentSaved"] = true
}

func (h *mockHooks) AfterAttachmentDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, attachmentID uint32) {
	h.CalledMap["AfterAttachmentDeleted"] = true
}

func (h *mockHooks) AfterPollAnswerSaved(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	h.CalledMap["AfterPollAnswerSaved"] = true
}

func (h *mockHooks) AfterPollAnswerDeleted(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32, user string) {
	h.CalledMap["AfterPollAnswerDeleted"] = true
}

func (h *mockHooks) AfterPollVotingPeriodEnded(ctx sdk.Context, subspaceID uint64, postID uint64, pollID uint32) {
	h.CalledMap["AfterPollVotingPeriodEnded"] = true
}

func (suite *KeeperTestSuite) TestHooks_AfterPostSaved() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		post  types.Post
	}{
		{
			name: "AfterPostSaved is called properly",
			post: types.NewPost(
				1,
				0,
				2,
				"External id",
				"Text",
				"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
				1,
				types.NewEntities(
					[]types.TextTag{
						types.NewTextTag(1, 3, "tag"),
					},
					[]types.TextTag{
						types.NewTextTag(4, 6, "tag"),
					},
					[]types.Url{
						types.NewURL(7, 9, "URL", "Display URL"),
					},
				),
				[]string{"generic"},
				[]types.PostReference{
					types.NewPostReference(types.POST_REFERENCE_TYPE_QUOTE, 1, 0),
				},
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
			),
		}}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.SavePost(ctx, tc.post)

			suite.Require().True(hooks.CalledMap["AfterPostSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterPostDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
	}{
		{
			name: "AfterPostDeleted is called properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			subspaceID: 1,
			postID:     2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.DeletePost(ctx, tc.subspaceID, tc.postID)

			suite.Require().True(hooks.CalledMap["AfterPostDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterAttachmentSaved() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		attachment types.Attachment
		check      func(ctx sdk.Context)
	}{
		{
			name: "AfterAttachmentSaved is called properly",
			attachment: types.NewAttachment(1, 1, 1, types.NewMedia(
				"ftp://user:password@example.com/image.png",
				"image/png",
			)),
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.SaveAttachment(ctx, tc.attachment)

			suite.Require().True(hooks.CalledMap["AfterAttachmentSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterAttachmentDeleted() {
	testCases := []struct {
		name         string
		store        func(ctx sdk.Context)
		subspaceID   uint64
		postID       uint64
		attachmentID uint32
	}{
		{
			name: "AfterAttachmentDeleted is called properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			subspaceID:   1,
			postID:       1,
			attachmentID: 1,
		},
	}
	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.DeleteAttachment(ctx, tc.subspaceID, tc.postID, tc.attachmentID)

			suite.Require().True(hooks.CalledMap["AfterAttachmentDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterPollAnswerSaved() {
	testCases := []struct {
		name   string
		answer types.UserAnswer
	}{
		{
			name:   "AfterPollAnswerSaved is called properly",
			answer: types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.SaveUserAnswer(ctx, tc.answer)

			suite.Require().True(hooks.CalledMap["AfterPollAnswerSaved"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterPollAnswerDeleted() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		subspaceID uint64
		postID     uint64
		pollID     uint32
		user       string
		check      func(ctx sdk.Context)
	}{
		{
			name: "AfterPollAnswerDeleted is called properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw"))
			},
			subspaceID: 1,
			postID:     1,
			pollID:     1,
			user:       "cosmos1jseuux3pktht0kkhlcsv4kqff3mql65udqs4jw",
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.DeleteUserAnswer(ctx, tc.subspaceID, tc.postID, tc.pollID, tc.user)

			suite.Require().True(hooks.CalledMap["AfterPollAnswerDeleted"])
		})
	}
}

func (suite *KeeperTestSuite) TestHooks_AfterPollVotingPeriodEnded() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		poll  types.Attachment
		check func(ctx sdk.Context)
	}{
		{
			name: "AfterPollVotingPeriodEnded is called properly",
			store: func(ctx sdk.Context) {
				attachment := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
						types.NewProvidedAnswer("No one of the above", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					true,
					false,
					nil,
				))
				suite.k.SaveAttachment(ctx, attachment)
				suite.k.InsertActivePollQueue(ctx, attachment)

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
			},
			poll: types.NewAttachment(1, 1, 1, types.NewPoll(
				"What animal is best?",
				[]types.Poll_ProvidedAnswer{
					types.NewProvidedAnswer("Cat", nil),
					types.NewProvidedAnswer("Dog", nil),
					types.NewProvidedAnswer("No one of the above", nil),
				},
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				true,
				false,
				nil,
			)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			hooks := newMockHooks()
			suite.k.SetHooks(types.NewMultiPostsHooks(hooks))

			suite.k.EndPoll(ctx, tc.poll)

			suite.Require().True(hooks.CalledMap["AfterPollVotingPeriodEnded"])
		})
	}
}
