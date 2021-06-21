package keeper_test

import (
	types2 "github.com/desmos-labs/desmos/x/posts/types"
	"time"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	tests := []struct {
		name string
		data struct {
			posts                []types2.Post
			userAnswers          []types2.UserAnswer
			postReactionsEntries []types2.PostReactionsEntry
			registeredReactions  []types2.RegisteredReaction
			reports              []types2.Report
			params               types2.Params
		}
		expected *types2.GenesisState
	}{
		{
			name: "Default expected state",
			data: struct {
				posts                []types2.Post
				userAnswers          []types2.UserAnswer
				postReactionsEntries []types2.PostReactionsEntry
				registeredReactions  []types2.RegisteredReaction
				reports              []types2.Report
				params               types2.Params
			}{
				posts:                nil,
				userAnswers:          nil,
				postReactionsEntries: nil,
				registeredReactions:  nil,
				reports:              nil,
				params:               types2.DefaultParams(),
			},
			expected: &types2.GenesisState{
				Params: types2.DefaultParams(),
			},
		},
		{
			name: "Genesis is exported fully",
			data: struct {
				posts                []types2.Post
				userAnswers          []types2.UserAnswer
				postReactionsEntries []types2.PostReactionsEntry
				registeredReactions  []types2.RegisteredReaction
				reports              []types2.Report
				params               types2.Params
			}{
				posts: []types2.Post{
					types2.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types2.CommentsStateBlocked,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswers: []types2.UserAnswer{
					types2.NewUserAnswer("post_id_1", "user", []string{"1", "2"}),
					types2.NewUserAnswer("post_id_2", "user_@", []string{"2", "4"}),
				},
				postReactionsEntries: []types2.PostReactionsEntry{
					types2.NewPostReactionsEntry("post_id", []types2.PostReaction{
						types2.NewPostReaction(":emoji:", "post_id", "creator"),
					}),
				},
				registeredReactions: []types2.RegisteredReaction{
					types2.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				reports: []types2.Report{
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				params: types2.DefaultParams(),
			},
			expected: types2.NewGenesisState(
				[]types2.Post{
					types2.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types2.CommentsStateBlocked,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types2.UserAnswer{
					types2.NewUserAnswer("post_id_1", "user", []string{"1", "2"}),
					types2.NewUserAnswer("post_id_2", "user_@", []string{"2", "4"}),
				},
				[]types2.PostReactionsEntry{
					types2.NewPostReactionsEntry("post_id", []types2.PostReaction{
						types2.NewPostReaction(":emoji:", "post_id", "creator"),
					}),
				},
				[]types2.RegisteredReaction{
					types2.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				[]types2.Report{
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types2.DefaultParams(),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.k.SetParams(suite.ctx, test.data.params)

			for _, reaction := range test.data.registeredReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.data.posts {
				suite.k.SavePost(suite.ctx, post)
			}

			for _, answer := range test.data.userAnswers {
				suite.k.SaveUserAnswer(suite.ctx, answer)
			}

			for _, entry := range test.data.postReactionsEntries {
				for _, reaction := range entry.Reactions {
					err := suite.k.SavePostReaction(suite.ctx, entry.PostID, reaction)
					suite.Require().NoError(err)
				}
			}

			for _, report := range test.data.reports {
				err := suite.k.SaveReport(suite.ctx, report)
				suite.Require().NoError(err)
			}

			exported := suite.k.ExportGenesis(suite.ctx)
			suite.Require().Equal(test.expected, exported)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	tests := []struct {
		name     string
		genesis  *types2.GenesisState
		expError bool
		expState struct {
			posts                []types2.Post
			userAnswers          []types2.UserAnswer
			postReactionsEntries []types2.PostReactionsEntry
			registeredReactions  []types2.RegisteredReaction
			reports              []types2.Report
			params               types2.Params
		}
	}{
		{
			name:     "Default genesis is initialized properly",
			genesis:  types2.DefaultGenesisState(),
			expError: false,
			expState: struct {
				posts                []types2.Post
				userAnswers          []types2.UserAnswer
				postReactionsEntries []types2.PostReactionsEntry
				registeredReactions  []types2.RegisteredReaction
				reports              []types2.Report
				params               types2.Params
			}{
				posts:                nil,
				userAnswers:          nil,
				postReactionsEntries: nil,
				registeredReactions:  nil,
				reports:              nil,
				params:               types2.DefaultParams(),
			},
		},
		{
			name: "Non default genesis state is imported properly",
			genesis: types2.NewGenesisState(
				[]types2.Post{
					types2.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types2.CommentsStateBlocked,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types2.UserAnswer{
					types2.NewUserAnswer(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
						[]string{"1", "2"},
					),
					types2.NewUserAnswer(
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
						[]string{"2", "4"},
					),
				},
				[]types2.PostReactionsEntry{
					types2.NewPostReactionsEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types2.PostReaction{
							types2.NewPostReaction(
								":emoji:",
								"post_id",
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				[]types2.RegisteredReaction{
					types2.NewRegisteredReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"value",
						"cosmos1rksrnm9jar2h4ahczmtq5p2m740n4cpf2pjv30",
					),
				},
				[]types2.Report{
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types2.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types2.DefaultParams(),
			),
			expError: false,
			expState: struct {
				posts                []types2.Post
				userAnswers          []types2.UserAnswer
				postReactionsEntries []types2.PostReactionsEntry
				registeredReactions  []types2.RegisteredReaction
				reports              []types2.Report
				params               types2.Params
			}{
				posts: []types2.Post{
					types2.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types2.CommentsStateBlocked,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswers: []types2.UserAnswer{
					types2.NewUserAnswer("a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd", "cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8", []string{"1", "2"}),
					types2.NewUserAnswer("b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9", "cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8", []string{"2", "4"}),
				},
				postReactionsEntries: []types2.PostReactionsEntry{
					types2.NewPostReactionsEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types2.PostReaction{
							types2.NewPostReaction(
								":emoji:",
								"post_id",
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				registeredReactions: []types2.RegisteredReaction{
					types2.NewRegisteredReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"value",
						"cosmos1rksrnm9jar2h4ahczmtq5p2m740n4cpf2pjv30",
					),
				},
				params: types2.DefaultParams(),
			},
		},
		{
			name: "Invalid post",
			genesis: types2.NewGenesisState(
				[]types2.Post{
					types2.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"",
						types2.CommentsStateBlocked,
						"",
						nil,
						nil,
						nil,
						time.Time{},
						time.Now(),
						"",
					),
				},
				nil,
				nil,
				nil,
				nil,
				types2.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Invalid user answer",
			genesis: types2.NewGenesisState(
				nil,
				[]types2.UserAnswer{
					types2.NewUserAnswer("post_id", "", []string{}),
				},
				nil,
				nil,
				nil,
				types2.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Invalid post reactions",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				[]types2.PostReactionsEntry{
					types2.NewPostReactionsEntry("post_id", []types2.PostReaction{}),
				},
				nil,
				nil,
				types2.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Double registered registeredReactions",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				[]types2.RegisteredReaction{
					types2.NewRegisteredReaction("creator", "shortcode", "value", "subspace"),
					types2.NewRegisteredReaction("creator", "shortcode", "value", "subspace"),
				},
				nil,
				types2.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Double reports",
			genesis: types2.NewGenesisState(
				nil,
				nil,
				nil,
				nil,
				[]types2.Report{
					types2.NewReport("post_id", "type", "message", "user"),
					types2.NewReport("post_id", "type", "message", "user"),
				},
				types2.DefaultParams(),
			),
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, *test.genesis)

				suite.Require().Equal(test.expState.posts, suite.k.GetPosts(suite.ctx))
				suite.Require().Equal(test.expState.registeredReactions, suite.k.GetRegisteredReactions(suite.ctx))
				suite.Require().Equal(test.expState.postReactionsEntries, suite.k.GetPostReactionsEntries(suite.ctx))
				suite.Require().Equal(test.expState.userAnswers, suite.k.GetAllUserAnswers(suite.ctx))
				suite.Require().Equal(test.expState.params, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
