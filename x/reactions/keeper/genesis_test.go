package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"

	poststypes "github.com/desmos-labs/desmos/v5/x/posts/types"
	"github.com/desmos-labs/desmos/v5/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	testCases := []struct {
		name       string
		setup      func()
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "subspaces data entries are exported properly",
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

				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(subspace subspacestypes.Subspace) (stop bool)) {
						for _, subspace := range subspaces {
							fn(subspace)
						}
					})

				suite.pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextRegisteredReactionID(ctx, 1, 2)
			},
			expGenesis: types.NewGenesisState([]types.SubspaceDataEntry{
				types.NewSubspaceDataEntry(1, 2),
			}, nil, nil, nil, nil),
		},
		{
			name: "registered reactions are exported properly",
			setup: func() {
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any())

				suite.pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				))
			},
			expGenesis: types.NewGenesisState(nil, []types.RegisteredReaction{
				types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				),
			}, nil, nil, nil),
		},
		{
			name: "post data entries are exported properly",
			setup: func() {
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any())

				posts := []poststypes.Post{
					poststypes.NewPost(
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
						poststypes.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					),
				}
				suite.pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any()).
					Do(func(ctx sdk.Context, fn func(post poststypes.Post) (stop bool)) {
						for _, post := range posts {
							fn(post)
						}
					})
			},
			store: func(ctx sdk.Context) {
				suite.k.SetNextReactionID(ctx, 1, 1, 2)
			},
			expGenesis: types.NewGenesisState(nil, nil, []types.PostDataEntry{
				types.NewPostDataEntry(1, 1, 2),
			}, nil, nil),
		},
		{
			name: "reactions are exported properly",
			setup: func() {
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any())

				suite.pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, []types.Reaction{
				types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				),
			}, nil),
		},
		{
			name: "reactions params are exported properly",
			setup: func() {
				suite.sk.EXPECT().
					IterateSubspaces(gomock.Any(), gomock.Any())

				suite.pk.EXPECT().
					IteratePosts(gomock.Any(), gomock.Any())
			},
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				))
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, nil, []types.SubspaceReactionsParams{
				types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				),
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
		name    string
		store   func(ctx sdk.Context)
		genesis types.GenesisState
		check   func(ctx sdk.Context)
	}{
		{
			name: "subspaces data entry is imported properly",
			genesis: types.GenesisState{
				SubspacesData: []types.SubspaceDataEntry{
					types.NewSubspaceDataEntry(1, 2),
				},
			},
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextRegisteredReactionID(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), stored)
			},
		},
		{
			name: "registered reaction is imported properly",
			genesis: types.GenesisState{
				RegisteredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						1,
						1,
						":hello:",
						"https://example.com?image=hello.png",
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetRegisteredReaction(ctx, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewRegisteredReaction(
					1,
					1,
					":hello:",
					"https://example.com?image=hello.png",
				), stored)
			},
		},
		{
			name: "post data entry is imported properly",
			genesis: types.GenesisState{
				PostsData: []types.PostDataEntry{
					types.NewPostDataEntry(1, 1, 2),
				},
			},
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetNextReactionID(ctx, 1, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(uint32(2), stored)
			},
		},
		{
			name: "reaction is imported properly",
			genesis: types.GenesisState{
				Reactions: []types.Reaction{
					types.NewReaction(
						1,
						1,
						1,
						types.NewRegisteredReactionValue(1),
						"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, found := suite.k.GetReaction(ctx, 1, 1, 1)
				suite.Require().True(found)
				suite.Require().Equal(types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1qewk97fp49vzssrfnc997jpztc5nzr7xsd8zdc",
				), stored)
			},
		},
		{
			name: "reactions params is imported properly",
			genesis: types.GenesisState{
				SubspacesParams: []types.SubspaceReactionsParams{
					types.NewSubspaceReactionsParams(
						1,
						types.NewRegisteredReactionValueParams(true),
						types.NewFreeTextValueParams(true, 10, ""),
					),
				},
			},
			check: func(ctx sdk.Context) {
				stored, err := suite.k.GetSubspaceReactionsParams(ctx, 1)
				suite.Require().NoError(err)
				suite.Require().Equal(types.NewSubspaceReactionsParams(
					1,
					types.NewRegisteredReactionValueParams(true),
					types.NewFreeTextValueParams(true, 10, ""),
				), stored)
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

			suite.k.InitGenesis(ctx, tc.genesis)
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
