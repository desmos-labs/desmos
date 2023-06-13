package keeper_test

import (
	"time"

	"github.com/golang/mock/gomock"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		setup      func()
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspaces data is exported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					),
					subspacestypes.NewSubspace(
						2,
						"Another text subspace",
						"This is another test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
				suite.k.SetParams(ctx, types.Params{})

				suite.k.SetNextPostID(ctx, 1, 1)

				suite.k.SetNextPostID(ctx, 2, 2)
			},
			expGenesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 1),
				types.NewSubspaceDataEntry(2, 2),
			}, nil, nil, nil, nil, nil, types.Params{}, nil),
		},
		{
			name: "posts are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SetNextAttachmentID(ctx, 1, 1, 1)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))

				suite.k.SetNextAttachmentID(ctx, 1, 2, 3)
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			expGenesis: types.NewGenesisState(
				nil,
				[]types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					),
					types.NewPost(
						1,
						0,
						2,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					),
				},
				[]types.PostDataEntry{
					types.NewPostDataEntry(1, 1, 1),
					types.NewPostDataEntry(1, 2, 3),
				}, nil, nil, nil, types.Params{}, nil),
		},
		{
			name: "attachments are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
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
			expGenesis: types.NewGenesisState(nil, nil, nil, []types.Attachment{
				types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
				types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
			}, nil, nil, types.Params{}, nil),
		},
		{
			name: "active polls are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 2, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)))
				suite.k.InsertActivePollQueue(ctx, types.NewAttachment(1, 1, 2, types.NewPoll(
					"What animal is best?",
					[]types.Poll_ProvidedAnswer{
						types.NewProvidedAnswer("Cat", nil),
						types.NewProvidedAnswer("Dog", nil),
					},
					time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
					false,
					false,
					nil,
				)))
			},
			expGenesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				[]types.Attachment{
					types.NewAttachment(1, 1, 2, types.NewPoll(
						"What animal is best?",
						[]types.Poll_ProvidedAnswer{
							types.NewProvidedAnswer("Cat", nil),
							types.NewProvidedAnswer("Dog", nil),
						},
						time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC),
						false,
						false,
						nil,
					)),
				},
				[]types.ActivePollData{
					types.NewActivePollData(1, 1, 2, time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC)),
				},
				nil,
				types.Params{},
				nil,
			),
		},
		{
			name: "user answers are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1, 2, 3}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, nil, []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				types.NewUserAnswer(1, 1, 2, []uint32{1, 2, 3}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			}, types.Params{}, nil),
		},
		{
			name: "params are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(20))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, nil, nil, types.NewParams(20), nil),
		},
		{
			name: "post owner transfer requests are exported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 2, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, nil, nil, types.Params{}, []types.PostOwnerTransferRequest{
				types.NewPostOwnerTransferRequest(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				types.NewPostOwnerTransferRequest(1, 2, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1eqpa6mv2jgevukaqtjmx5535vhc3mm3cf458zg"),
			}),
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

			genesis := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, genesis)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_ImportGenesis() {
	testCases := []struct {
		name  string
		setup func()
		store func(ctx sdk.Context)
		data  types.GenesisState
		check func(ctx sdk.Context)
	}{
		{
			name: "subspace data is imported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
			data: types.GenesisState{
				SubspacesData: []types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 1),
				},
			},
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextPostID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint64(1), stored)
			},
		},
		{
			name: "genesis post is imported correctly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
			data: types.GenesisState{
				Posts: []types.Post{
					types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						1,
						nil,
						nil,
						nil,
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
						"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
					),
				},
			},
			check: func(ctx sdk.Context) {
				post, found := suite.k.GetPost(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				), post)
			},
		},
		{
			name: "attachment id is imported correctly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
			data: types.GenesisState{
				PostsData: []types.PostDataEntry{
					types.NewPostDataEntry(1, 1, 2),
				},
			},
			check: func(ctx sdk.Context) {
				attachmentID, err := suite.k.GetNextAttachmentID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), attachmentID)
			},
		},
		{
			name: "attachment is imported correctly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
			name: "user answer is imported properly",
			setup: func() {
				subspaces := []subspacestypes.Subspace{
					subspacestypes.NewSubspace(
						1,
						"Test subspace",
						"This is a test subspace",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
						"cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5",
						"cosmos1qzskhrcjnkdz2ln4yeafzsdwht8ch08j4wed69",
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
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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
					types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetUserAnswer(ctx, 1, 1, 1, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st")
				suite.Require().True(found)
				suite.Require().Equal(types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"), stored)
			},
		},
		{
			name: "active polls are imported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			data: types.GenesisState{
				ActivePolls: []types.ActivePollData{
					types.NewActivePollData(1, 1, 2, time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC)),
				},
			},
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.ActivePollQueueKey(1, 1, 2, time.Date(2100, 1, 1, 12, 00, 00, 000, time.UTC))))
			},
		},
		{
			name: "params are imported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			data: types.GenesisState{
				Params: types.NewParams(200),
			},
			check: func(ctx sdk.Context) {
				stored := suite.k.GetParams(ctx)
				suite.Require().Equal(types.NewParams(200), stored)
			},
		},
		{
			name: "post transfer owner requests are imported properly",
			setup: func() {
				suite.sk.EXPECT().IterateSubspaces(gomock.Any(), gomock.Any())
			},
			data: types.GenesisState{
				PostOwnerTransferRequests: []types.PostOwnerTransferRequest{
					types.NewPostOwnerTransferRequest(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetPostOwnerTransferRequest(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewPostOwnerTransferRequest(1, 1, "cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd", "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
					stored)
			},
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

			suite.k.InitGenesis(ctx, tc.data)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
