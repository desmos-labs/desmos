package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	tests := []struct {
		name string
		data struct {
			posts                []types.Post
			userAnswerEntries    []types.UserAnswersEntry
			postReactionsEntries []types.PostReactionsEntry
			registeredReactions  []types.RegisteredReaction
			params               types.Params
		}
		expected *types.GenesisState
	}{
		{
			name: "Default expected state",
			data: struct {
				posts                []types.Post
				userAnswerEntries    []types.UserAnswersEntry
				postReactionsEntries []types.PostReactionsEntry
				registeredReactions  []types.RegisteredReaction
				params               types.Params
			}{
				posts:                nil,
				userAnswerEntries:    nil,
				postReactionsEntries: nil,
				registeredReactions:  nil,
				params:               types.DefaultParams(),
			},
			expected: &types.GenesisState{
				Params: types.DefaultParams(),
			},
		},
		{
			name: "Genesis is exported fully",
			data: struct {
				posts                []types.Post
				userAnswerEntries    []types.UserAnswersEntry
				postReactionsEntries []types.PostReactionsEntry
				registeredReactions  []types.RegisteredReaction
				params               types.Params
			}{
				posts: []types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						true,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswerEntries: []types.UserAnswersEntry{
					types.NewUserAnswersEntry("post_id", []types.UserAnswer{
						types.NewUserAnswer([]string{"1", "2"}, "user"),
					}),
					types.NewUserAnswersEntry("post_id_2", []types.UserAnswer{
						types.NewUserAnswer([]string{"2", "4"}, "user_@"),
					}),
				},
				postReactionsEntries: []types.PostReactionsEntry{
					types.NewPostReactionsEntry("post_id", []types.PostReaction{
						types.NewPostReaction(":emoji:", "post_id", "creator"),
					}),
				},
				registeredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				params: types.DefaultParams(),
			},
			expected: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						true,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types.UserAnswersEntry{
					types.NewUserAnswersEntry("post_id", []types.UserAnswer{
						types.NewUserAnswer([]string{"1", "2"}, "user"),
					}),
					types.NewUserAnswersEntry("post_id_2", []types.UserAnswer{
						types.NewUserAnswer([]string{"2", "4"}, "user_@"),
					}),
				},
				[]types.PostReactionsEntry{
					types.NewPostReactionsEntry("post_id", []types.PostReaction{
						types.NewPostReaction(":emoji:", "post_id", "creator"),
					}),
				},
				[]types.RegisteredReaction{
					types.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				types.DefaultParams(),
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.keeper.SetParams(suite.ctx, test.data.params)

			for _, reaction := range test.data.registeredReactions {
				suite.keeper.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, post := range test.data.posts {
				suite.keeper.SavePost(suite.ctx, post)
			}

			for _, entry := range test.data.userAnswerEntries {
				for _, answer := range entry.UserAnswers {
					suite.keeper.SavePollAnswers(suite.ctx, entry.PostId, answer)
				}
			}

			for _, entry := range test.data.postReactionsEntries {
				for _, reaction := range entry.Reactions {
					err := suite.keeper.SavePostReaction(suite.ctx, entry.PostId, reaction)
					suite.Require().NoError(err)
				}
			}

			exported := suite.keeper.ExportGenesis(suite.ctx)
			suite.Require().Equal(test.expected, exported)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_InitGenesis() {
	tests := []struct {
		name     string
		genesis  *types.GenesisState
		expError bool
		expState struct {
			posts                []types.Post
			userAnswerEntries    []types.UserAnswersEntry
			postReactionsEntries []types.PostReactionsEntry
			registeredReactions  []types.RegisteredReaction
			params               types.Params
		}
	}{
		{
			name:     "Default genesis is initialized properly",
			genesis:  types.DefaultGenesisState(),
			expError: false,
			expState: struct {
				posts                []types.Post
				userAnswerEntries    []types.UserAnswersEntry
				postReactionsEntries []types.PostReactionsEntry
				registeredReactions  []types.RegisteredReaction
				params               types.Params
			}{
				posts:                nil,
				userAnswerEntries:    nil,
				postReactionsEntries: nil,
				registeredReactions:  nil,
				params:               types.DefaultParams(),
			},
		},
		{
			name: "Non default genesis state is imported properly",
			genesis: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						true,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types.UserAnswersEntry{
					types.NewUserAnswersEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types.UserAnswer{
							types.NewUserAnswer(
								[]string{"1", "2"},
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8"),
						},
					),
					types.NewUserAnswersEntry(
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						[]types.UserAnswer{
							types.NewUserAnswer(
								[]string{"2", "4"},
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				[]types.PostReactionsEntry{
					types.NewPostReactionsEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types.PostReaction{
							types.NewPostReaction(
								":emoji:",
								"post_id",
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				[]types.RegisteredReaction{
					types.NewRegisteredReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"value",
						"cosmos1rksrnm9jar2h4ahczmtq5p2m740n4cpf2pjv30",
					),
				},
				types.DefaultParams(),
			),
			expError: false,
			expState: struct {
				posts                []types.Post
				userAnswerEntries    []types.UserAnswersEntry
				postReactionsEntries []types.PostReactionsEntry
				registeredReactions  []types.RegisteredReaction
				params               types.Params
			}{
				posts: []types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						true,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswerEntries: []types.UserAnswersEntry{
					types.NewUserAnswersEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types.UserAnswer{
							types.NewUserAnswer(
								[]string{"1", "2"},
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8"),
						},
					),
					types.NewUserAnswersEntry(
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						[]types.UserAnswer{
							types.NewUserAnswer(
								[]string{"2", "4"},
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				postReactionsEntries: []types.PostReactionsEntry{
					types.NewPostReactionsEntry(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						[]types.PostReaction{
							types.NewPostReaction(
								":emoji:",
								"post_id",
								"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
							),
						},
					),
				},
				registeredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"value",
						"cosmos1rksrnm9jar2h4ahczmtq5p2m740n4cpf2pjv30",
					),
				},
				params: types.DefaultParams(),
			},
		},
		{
			name: "Invalid post",
			genesis: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"",
						true,
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
				types.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Invalid poll answer",
			genesis: types.NewGenesisState(
				nil,
				[]types.UserAnswersEntry{
					types.NewUserAnswersEntry("post_id", []types.UserAnswer{}),
				},
				nil,
				nil,
				types.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Invalid post reactions",
			genesis: types.NewGenesisState(
				nil,
				nil,
				[]types.PostReactionsEntry{
					types.NewPostReactionsEntry("post_id", []types.PostReaction{}),
				},
				nil,
				types.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Double registered registeredReactions",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				[]types.RegisteredReaction{
					types.NewRegisteredReaction("creator", "shortcode", "value", "subspace"),
					types.NewRegisteredReaction("creator", "shortcode", "value", "subspace"),
				},
				types.DefaultParams(),
			),
			expError: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			if test.expError {
				suite.Require().Panics(func() { suite.keeper.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.keeper.InitGenesis(suite.ctx, *test.genesis)

				suite.Require().Equal(test.expState.posts, suite.keeper.GetPosts(suite.ctx))
				suite.Require().Equal(test.expState.registeredReactions, suite.keeper.GetRegisteredReactions(suite.ctx))
				suite.Require().Equal(test.expState.postReactionsEntries, suite.keeper.GetPostReactionsEntries(suite.ctx))
				suite.Require().Equal(test.expState.userAnswerEntries, suite.keeper.GetUserAnswersEntries(suite.ctx))
				suite.Require().Equal(test.expState.params, suite.keeper.GetParams(suite.ctx))
			}
		})
	}
}
