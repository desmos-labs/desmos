package keeper_test

import (
	"time"

	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func (suite *KeeperTestsuite) TestKeeper_ExportGenesis() {
	user, err := sdk.AccAddressFromBech32("cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st")
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspaces data is exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SetPostID(ctx, 1, 1)
				suite.k.SetPostID(ctx, 2, 2)
			},
			expGenesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 1),
				types.NewSubspaceDataEntry(2, 2),
			}, nil, nil, nil, types.Params{}),
		},
		{
			name: "posts are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SetAttachmentID(ctx, 1, 1, 1)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				suite.k.SetAttachmentID(ctx, 1, 2, 3)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			expGenesis: types.NewGenesisState(nil, []types.GenesisPost{
				types.NewGenesisPost(1, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				)),
				types.NewGenesisPost(3, types.NewPost(
					1,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				)),
			}, nil, nil, types.Params{}),
		},
		{
			name: "attachments are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
			},
			expGenesis: types.NewGenesisState(nil, nil, []types.Attachment{
				types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
				types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
			}, nil, types.Params{}),
		},
		{
			name: "user answers are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, user))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1, 2, 3}, user))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, user),
				types.NewUserAnswer(1, 1, 2, []uint32{1, 2, 3}, user),
			}, types.Params{}),
		},
		{
			name: "params are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(20))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, types.NewParams(20)),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestsuite) TestKeeper_ImportGenesis() {
	user, err := sdk.AccAddressFromBech32("cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st")
	suite.Require().NoError(err)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		data      types.GenesisState
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existing subspace returns error while importing data",
			data: types.GenesisState{
				SubspacesData: []types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 1),
				},
			},
			shouldErr: true,
		},
		{
			name: "subspace data is imported properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			data: types.GenesisState{
				SubspacesData: []types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 1),
				},
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "non existing subspace returns error while importing post",
			data: types.GenesisState{
				GenesisPosts: []types.GenesisPost{
					types.NewGenesisPost(1, types.NewPost(
						1,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					)),
				},
			},
			shouldErr: true,
		},
		{
			name: "genesis post is imported correctly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			data: types.GenesisState{
				GenesisPosts: []types.GenesisPost{
					types.NewGenesisPost(2, types.NewPost(
						1,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					)),
				},
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				post, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				), post)

				attachmentID, err := suite.k.GetAttachmentID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), attachmentID)
			},
		},
		{
			name: "non existing post returns error while importing attachment",
			data: types.GenesisState{
				Attachments: []types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
			},
			shouldErr: true,
		},
		{
			name: "attachment is imported correctly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			data: types.GenesisState{
				Attachments: []types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewMedia(
						"ftp://user:password@example.com/image.png",
						"image/png",
					)),
				},
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetAttachment(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)), stored)
			},
		},
		{
			name: "poll attachment with non null results and future end date returns error",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			data: types.GenesisState{
				Attachments: []types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(3000, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						types.NewPollTallyResults([]types.PollTallyResults_AnswerResult{
							types.NewAnswerResult(0, 1),
							types.NewAnswerResult(1, 2),
						}),
					)),
				},
			},
			shouldErr: true,
		},
		{
			name: "poll attachment is added to active poll queue properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			data: types.GenesisState{
				Attachments: []types.Attachment{
					types.NewAttachment(1, 1, 1, types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(3000, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					)),
				},
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetAttachment(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewAttachment(1, 1, 1, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(3000, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)), stored)

				store := ctx.KVStore(suite.storeKey)
				endDate := time.Date(3000, 1, 1, 12, 00, 00, 000, time.UTC)
				suite.Require().True(store.Has(types.ActivePollQueueKey(1, 1, 1, endDate)))
			},
		},
		{
			name: "non existing poll returns error when importing user answer",
			data: types.GenesisState{
				UserAnswers: []types.UserAnswer{
					types.NewUserAnswer(1, 1, 1, []uint32{1}, user),
				},
			},
			shouldErr: true,
		},
		{
			name: "user answer is imported properly",
			store: func(ctx sdk.Context) {
				suite.sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
					"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
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
			},
			data: types.GenesisState{
				UserAnswers: []types.UserAnswer{
					types.NewUserAnswer(1, 1, 1, []uint32{1}, user),
				},
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, user)
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1}, user), stored)
			},
		},
		{
			name: "params are imported properly",
			data: types.GenesisState{
				Params: types.NewParams(200),
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(types.NewParams(200), stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			if tc.shouldErr {
				suite.Require().Panics(func() { suite.k.InitGenesis(ctx, tc.data) })
			} else {
				suite.Require().NotPanics(func() { suite.k.InitGenesis(ctx, tc.data) })
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
