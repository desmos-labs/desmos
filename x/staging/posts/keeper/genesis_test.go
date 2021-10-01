package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v2/x/staging/posts/types"
)

func (suite *KeeperTestSuite) TestKeeper_ExportGenesis() {
	tests := []struct {
		name string
		data struct {
			posts               []types.Post
			userAnswers         []types.UserAnswer
			postReactions       []types.PostReaction
			registeredReactions []types.RegisteredReaction
			reports             []types.Report
			params              types.Params
		}
		expected *types.GenesisState
	}{
		{
			name: "Default expected state",
			data: struct {
				posts               []types.Post
				userAnswers         []types.UserAnswer
				postReactions       []types.PostReaction
				registeredReactions []types.RegisteredReaction
				reports             []types.Report
				params              types.Params
			}{
				posts:               nil,
				userAnswers:         nil,
				postReactions:       nil,
				registeredReactions: nil,
				reports:             nil,
				params:              types.DefaultParams(),
			},
			expected: &types.GenesisState{
				Params: types.DefaultParams(),
			},
		},
		{
			name: "Genesis is exported fully",
			data: struct {
				posts               []types.Post
				userAnswers         []types.UserAnswer
				postReactions       []types.PostReaction
				registeredReactions []types.RegisteredReaction
				reports             []types.Report
				params              types.Params
			}{
				posts: []types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types.CommentsStateBlocked,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswers: []types.UserAnswer{
					types.NewUserAnswer("post_id_1", "user", []string{"1", "2"}),
					types.NewUserAnswer("post_id_2", "user_@", []string{"2", "4"}),
				},
				postReactions: []types.PostReaction{
					types.NewPostReaction("post_id", ":emoji:", "post_id", "creator"),
				},
				registeredReactions: []types.RegisteredReaction{
					types.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				reports: []types.Report{
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				params: types.DefaultParams(),
			},
			expected: types.NewGenesisState(
				[]types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types.CommentsStateBlocked,
						"subspace",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types.UserAnswer{
					types.NewUserAnswer("post_id_1", "user", []string{"1", "2"}),
					types.NewUserAnswer("post_id_2", "user_@", []string{"2", "4"}),
				},
				[]types.PostReaction{
					types.NewPostReaction("post_id", ":emoji:", "post_id", "creator"),
				},
				[]types.RegisteredReaction{
					types.NewRegisteredReaction("creator", ":emoji:", "value", "subspace"),
				},
				[]types.Report{
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types.DefaultParams(),
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

			for _, reaction := range test.data.postReactions {
				err := suite.k.SavePostReaction(suite.ctx, reaction)
				suite.Require().NoError(err)
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
		genesis  *types.GenesisState
		expError bool
		expState struct {
			posts               []types.Post
			userAnswers         []types.UserAnswer
			postReactions       []types.PostReaction
			registeredReactions []types.RegisteredReaction
			reports             []types.Report
			params              types.Params
		}
	}{
		{
			name:     "Default genesis is initialized properly",
			genesis:  types.DefaultGenesisState(),
			expError: false,
			expState: struct {
				posts               []types.Post
				userAnswers         []types.UserAnswer
				postReactions       []types.PostReaction
				registeredReactions []types.RegisteredReaction
				reports             []types.Report
				params              types.Params
			}{
				posts:               nil,
				userAnswers:         nil,
				postReactions:       nil,
				registeredReactions: nil,
				reports:             nil,
				params:              types.DefaultParams(),
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
						types.CommentsStateBlocked,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				[]types.UserAnswer{
					types.NewUserAnswer(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
						[]string{"1", "2"},
					),
					types.NewUserAnswer(
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
						[]string{"2", "4"},
					),
				},
				[]types.PostReaction{
					types.NewPostReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"post_id",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
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
				[]types.Report{
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"scam",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
					types.NewReport(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"",
						"message",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types.DefaultParams(),
			),
			expError: false,
			expState: struct {
				posts               []types.Post
				userAnswers         []types.UserAnswer
				postReactions       []types.PostReaction
				registeredReactions []types.RegisteredReaction
				reports             []types.Report
				params              types.Params
			}{
				posts: []types.Post{
					types.NewPost(
						"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
						"",
						"message",
						types.CommentsStateBlocked,
						"b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9",
						nil,
						nil,
						nil,
						time.Time{},
						time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
						"creator",
					),
				},
				userAnswers: []types.UserAnswer{
					types.NewUserAnswer("a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd", "cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8", []string{"1", "2"}),
					types.NewUserAnswer("b459afddb3a09621ee29b78b3968e566d7fb0001d96395d54030eb703b0337a9", "cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8", []string{"2", "4"}),
				},
				postReactions: []types.PostReaction{
					types.NewPostReaction(
						"a56145270ce6b3bebd1dd012b73948677dd618d496488bc608a3cb43ce3547dd",
						":emoji:",
						"post_id",
						"cosmos1u3cjgn7t7v6edpy2szvydxucarzkyjj26az3k8",
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
						types.CommentsStateBlocked,
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
				types.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Invalid user answer",
			genesis: types.NewGenesisState(
				nil,
				[]types.UserAnswer{
					types.NewUserAnswer("post_id", "", []string{}),
				},
				nil,
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
				[]types.PostReaction{
					types.NewPostReaction("post_id", "", "", ""),
				},
				nil,
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
				nil,
				types.DefaultParams(),
			),
			expError: true,
		},
		{
			name: "Double reports",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				nil,
				[]types.Report{
					types.NewReport("post_id", "type", "message", "user"),
					types.NewReport("post_id", "type", "message", "user"),
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
				suite.Require().Panics(func() { suite.k.InitGenesis(suite.ctx, *test.genesis) })
			} else {
				suite.k.InitGenesis(suite.ctx, *test.genesis)

				suite.Require().Equal(test.expState.posts, suite.k.GetPosts(suite.ctx))
				suite.Require().Equal(test.expState.registeredReactions, suite.k.GetRegisteredReactions(suite.ctx))
				suite.Require().Equal(test.expState.postReactions, suite.k.GetAllPostReactions(suite.ctx))
				suite.Require().Equal(test.expState.userAnswers, suite.k.GetAllUserAnswers(suite.ctx))
				suite.Require().Equal(test.expState.params, suite.k.GetParams(suite.ctx))
			}
		})
	}
}
