package keeper_test

import (
	"time"

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
			}, nil, nil, nil, nil, types.Params{}),
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
			}, nil, nil, nil, types.Params{}),
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
			}, nil, nil, types.Params{}),
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
			}, nil, types.Params{}),
		},
		{
			name: "tally results are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.Params{})
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 100),
				}))
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 2, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(0, 10),
					types.NewAnswerResult(1, 50),
					types.NewAnswerResult(2, 2),
				}))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, []types.PollTallyResults{
				types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(1, 100),
				}),
				types.NewPollTallyResults(1, 1, 2, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(0, 10),
					types.NewAnswerResult(1, 50),
					types.NewAnswerResult(2, 2),
				}),
			}, types.Params{}),
		},
		{
			name: "params are exported properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(20))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, nil, types.NewParams(20)),
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
