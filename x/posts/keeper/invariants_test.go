package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestValidSubspacesInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "not found next post id breaks invariant",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				}

				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
				}

				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidSubspacesInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidPostsInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "not found subspace breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(false)

			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "not found section breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "not found next post id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "invalid post id compared to next post id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expBroken: true,
		},
		{
			name: "not found next attachment id breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 2)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.DeleteNextAttachmentID(ctx, 1, 1)
			},
			expBroken: true,
		},
		{
			name: "invalid post breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 2)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
				suite.sk.EXPECT().HasSection(gomock.Any(), uint64(1), uint32(0)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextPostID(ctx, 1, 2)

				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidPostsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidAttachmentsInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "not found subspace breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			expBroken: true,
		},
		{
			name: "not found post breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			expBroken: true,
		},
		{
			name: "not found next attachment id returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid attachment id compared to next attachment id returns error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			expBroken: true,
		},
		{
			name: "invalid attachment breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextAttachmentID(ctx, 1, 1, 2)

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("", ""),
				))
			},
			expBroken: true,
		},
		{
			name: "valid data returns no error",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SetNextAttachmentID(ctx, 1, 1, 2)

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidAttachmentsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidUserAnswersInvariant() {
	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "not found subspace breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(false)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expBroken: true,
		},
		{
			name: "not found post breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expBroken: true,
		},
		{
			name: "not found poll breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expBroken: true,
		},
		{
			name: "invalid user answer breaks invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, nil, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expBroken: true,
		},
		{
			name: "valid data does not break invariant",
			setup: func() {
				suite.sk.EXPECT().HasSubspace(gomock.Any(), uint64(1)).Return(true)
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External id",
					"Text",
					"cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)))

				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidUserAnswersInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}

func (suite *KeeperTestSuite) TestValidActivePollsInvariant() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		expBroken bool
	}{
		{
			name: "non nil final results breaks invariant",
			store: func(ctx sdk.Context) {
				poll := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
						types.NewAnswerResult(0, 10),
					}),
				))
				suite.k.SaveAttachment(ctx, poll)
				suite.k.InsertActivePollQueue(ctx, poll)
			},
			expBroken: true,
		},
		{
			name: "valid data returns no error",
			store: func(ctx sdk.Context) {
				poll := types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				))
				suite.k.SaveAttachment(ctx, poll)
				suite.k.InsertActivePollQueue(ctx, poll)
			},
			expBroken: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, broken := keeper.ValidActivePollsInvariant(suite.k)(ctx)
			suite.Require().Equal(tc.expBroken, broken)
		})
	}
}
