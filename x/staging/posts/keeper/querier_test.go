package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/staging/posts/keeper"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) Test_queryPost() {
	tests := []struct {
		name                string
		path                []string
		storedPosts         []types.Post
		storedReactions     []types.PostReactionsEntry
		registeredReactions []types.RegisteredReaction
		storedAnswers       []types.UserAnswer
		expError            bool
		expResult           types.PostQueryResponse
	}{
		{
			name:                "Invalid query endpoint",
			path:                []string{"invalid", ""},
			registeredReactions: nil,
			expError:            true,
		},
		{
			name:                "Invalid ID returns error",
			path:                []string{types.QueryPost, ""},
			registeredReactions: nil,
			expError:            true,
		},
		{
			name:                "Post not found returns error",
			path:                []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			registeredReactions: nil,
			expError:            true,
		},
		{
			name: "Post without reactions is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			registeredReactions: nil,
			path:                []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expError:            false,
			expResult: types.NewPostResponse(
				types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				[]types.UserAnswer{
					types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				},
				[]types.PostReaction{},
				[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			),
		},
		{
			name: "Post without children is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			registeredReactions: nil,
			path:                []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expError:            false,
			expResult: types.NewPostResponse(
				types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				[]types.UserAnswer{
					types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				},
				[]types.PostReaction{},
				[]string{},
			),
		},
		{
			name: "Post without medias is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					PollData:             suite.testData.post.PollData,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			storedReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
						),
					},
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.postOwner,
					":like:",
					"https://smile.jpg",
					"",
				),
			},
			path:     []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expError: false,
			expResult: types.NewPostResponse(
				types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					PollData:             suite.testData.post.PollData,
				},
				[]types.UserAnswer{
					types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				},
				[]types.PostReaction{
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					),
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
					),
				},
				[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			),
		},
		{
			name: "Post without poll and poll answers is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
			},
			storedReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
						),
					},
				),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.postOwner,
					":like:",
					"https://smile.jpg",
					"",
				),
			},
			path:     []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expError: false,
			expResult: types.NewPostResponse(
				types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
				nil,
				[]types.PostReaction{
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					),
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
					),
				},
				[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			),
		},
		{
			name: "Post with all data is returned properly",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
			},
			storedReactions: []types.PostReactionsEntry{
				types.NewPostReactionsEntry(
					"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					[]types.PostReaction{
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						),
						types.NewPostReaction(
							":like:",
							"https://smile.jpg",
							"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
						),
					},
				),
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			registeredReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					suite.testData.postOwner,
					":like:",
					"https://smile.jpg",
					"",
				),
			},
			path:     []string{types.QueryPost, "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expError: false,
			expResult: types.NewPostResponse(
				types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				[]types.UserAnswer{
					types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				},
				[]types.PostReaction{
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					),
					types.NewPostReaction(
						":like:",
						"https://smile.jpg",
						"cosmos1r2plnngkwnahajl3d2a7fvzcsxf6djlt380f3l",
					),
				},
				[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, reaction := range test.registeredReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			for _, p := range test.storedPosts {
				suite.k.SavePost(suite.ctx, p)
			}

			for index, ans := range test.storedAnswers {
				suite.k.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			for _, entry := range test.storedReactions {
				for _, reaction := range entry.Reactions {
					err := suite.k.SavePostReaction(suite.ctx, entry.PostID, reaction)
					suite.Require().NoError(err)
				}
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Require().Equal(string(expected), string(result))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryPosts() {
	tests := []struct {
		name          string
		storedPosts   []types.Post
		storedAnswers []types.UserAnswer
		params        types.QueryPostsParams
		expResponse   []types.PostQueryResponse
	}{
		{
			name: "Empty params returns all",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					ParentID:             "",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					PollData:             suite.testData.post.PollData,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			params: types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.Post{
						PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						ParentID:             "",
						Message:              "Parent",
						Created:              suite.testData.post.Created,
						LastEdited:           suite.testData.post.LastEdited,
						AdditionalAttributes: nil,
						Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						Attachments:          suite.testData.post.Attachments,
						PollData:             suite.testData.post.PollData,
					},
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
					},
					[]types.PostReaction{},
					[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
				),
				types.NewPostResponse(
					types.Post{
						PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
						ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						Message:              "Child",
						Created:              suite.testData.post.Created,
						LastEdited:           suite.testData.post.LastEdited,
						AdditionalAttributes: nil,
						Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						PollData:             suite.testData.post.PollData,
					},
					nil,
					[]types.PostReaction{},
					[]string{},
				),
			},
		},
		{
			name: "Empty params returns all posts without medias",
			storedPosts: []types.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					PollData:             suite.testData.post.PollData,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			params: types.QueryPostsParams{},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.Post{
						PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
						ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						Message:              "Child",
						Created:              suite.testData.post.Created,
						LastEdited:           suite.testData.post.LastEdited,
						AdditionalAttributes: nil,
						Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						PollData:             suite.testData.post.PollData,
					},
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
					},
					[]types.PostReaction{},
					[]string{},
				),
			},
		},
		{
			name: "Empty params returns all posts without poll data and poll answers",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
			},
			params: types.QueryPostsParams{Page: 1, Limit: 1},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.Post{
						PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						Message:              "Parent",
						Created:              suite.testData.post.Created,
						LastEdited:           suite.testData.post.LastEdited,
						AdditionalAttributes: nil,
						Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						Attachments:          suite.testData.post.Attachments,
					},
					nil,
					[]types.PostReaction{},
					[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
				),
			},
		},
		{
			name: "Non empty params return proper posts",
			storedPosts: []types.Post{
				{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Parent",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Child",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					Attachments:          suite.testData.post.Attachments,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			params: types.QueryPostsParams{Page: 1, Limit: 1},
			expResponse: []types.PostQueryResponse{
				types.NewPostResponse(
					types.Post{
						PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						Message:              "Parent",
						Created:              suite.testData.post.Created,
						LastEdited:           suite.testData.post.LastEdited,
						AdditionalAttributes: nil,
						Creator:              "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
						Attachments:          suite.testData.post.Attachments,
						PollData:             suite.testData.post.PollData,
					},
					[]types.UserAnswer{
						types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
					},
					[]types.PostReaction{},
					[]string{"f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, p := range test.storedPosts {
				suite.k.SavePost(suite.ctx, p)
			}

			for index, ans := range test.storedAnswers {
				suite.k.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			request := abci.RequestQuery{Data: suite.legacyAminoCdc.MustMarshalJSON(&test.params)}

			result, err := querier(suite.ctx, []string{types.QueryPosts}, request)
			suite.Require().NoError(err)

			expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResponse)
			suite.Require().Equal(string(expected), string(result))
		})
	}
}

func (suite *KeeperTestSuite) Test_queryPollAnswers() {
	tests := []struct {
		name          string
		path          []string
		storedPosts   []types.Post
		storedAnswers []types.UserAnswer
		expError      bool
		expResult     types.QueryPollAnswersResponse
	}{
		{
			name:     "Invalid post id returns error",
			path:     []string{types.QueryPollAnswers, ""},
			expError: true,
		},
		{
			name:     "Post not found returns error",
			path:     []string{types.QueryPollAnswers, "1"},
			expError: true,
		},
		{
			name:     "No post associated with ID given",
			path:     []string{types.QueryPollAnswers, "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			expError: true,
		},
		{
			name: "Post without poll returns error",
			path: []string{types.QueryPollAnswers, "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			storedPosts: []types.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					Message:              "post with poll",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					Attachments:          suite.testData.post.Attachments,
				},
			},
			expError: true,
		},
		{
			name: "Returns answers details of the post correctly",
			path: []string{types.QueryPollAnswers, "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd"},
			storedPosts: []types.Post{
				{
					PostID:               "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
					Message:              "post with poll",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					AdditionalAttributes: nil,
					Creator:              suite.testData.post.Creator,
					Attachments:          suite.testData.post.Attachments,
					PollData:             suite.testData.post.PollData,
				},
			},
			storedAnswers: []types.UserAnswer{
				types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
			},
			expResult: types.QueryPollAnswersResponse{
				PostID: "f1b909289cd23188c19da17ae5d5a05ad65623b0fad756e5e03c8c936ca876fd",
				Answers: []types.UserAnswer{
					types.NewUserAnswer([]string{"1"}, "cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4"),
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, p := range test.storedPosts {
				suite.k.SavePost(suite.ctx, p)
			}

			for index, ans := range test.storedAnswers {
				suite.k.SavePollAnswers(suite.ctx, test.storedPosts[index].PostID, ans)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				expected := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Require().Equal(string(expected), string(result))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryRegisteredReactions() {
	tests := []struct {
		name            string
		path            []string
		storedReactions []types.RegisteredReaction
		expError        bool
		expResult       []types.RegisteredReaction
	}{
		{
			name: "PostReactions returned properly",
			path: []string{types.QueryRegisteredReactions},
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":smile:",
					"http://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":sad:",
					"http://sad.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			expError: false,
			expResult: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":sad:",
					"http://sad.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":smile:",
					"http://smile.jpg",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			for _, r := range test.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, r)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				actual := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
				suite.Require().Equal(string(actual), string(result))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_queryParams() {
	tests := []struct {
		name         string
		storedParams types.Params
		path         []string
		expResult    types.Params
	}{
		{
			name:         "Returning posts parameters correctly",
			storedParams: types.DefaultParams(),
			path:         []string{types.QueryParams},
			expResult:    types.DefaultParams(),
		},
		{
			name:         "Returning non default params",
			storedParams: types.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
			path:         []string{types.QueryParams},
			expResult:    types.NewParams(sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1), sdk.NewInt(1)),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			suite.k.SetParams(suite.ctx, test.storedParams)

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})
			suite.Require().Nil(err)

			actual := codec.MustMarshalJSONIndent(suite.legacyAminoCdc, &test.expResult)
			suite.Require().Equal(string(actual), string(result))
		})
	}
}

func (suite *KeeperTestSuite) Test_queryReports() {
	tests := []struct {
		name          string
		path          []string
		storedReports []types.Report
		expErr        bool
		expResponse   []types.Report
	}{
		{
			name:          "Invalid Post ID",
			path:          []string{types.QueryReports, "1234"},
			storedReports: nil,
			expErr:        true,
		},
		{
			name: "Valid request returns correctly",
			path: []string{types.QueryReports, suite.testData.postID},
			storedReports: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.postOwner,
				),
				types.NewReport(
					"other_post",
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
			expErr: false,
			expResponse: []types.Report{
				types.NewReport(
					suite.testData.postID,
					"type",
					"message",
					suite.testData.postOwner,
				),
			},
		},
		{
			name:          "Empty reports and valid id",
			path:          []string{types.QueryReports, suite.testData.postID},
			storedReports: nil,
			expErr:        false,
			expResponse:   nil,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, rep := range test.storedReports {
				err := suite.k.SaveReport(suite.ctx, rep)
				suite.Require().NoError(err)
			}

			querier := keeper.NewQuerier(suite.k, suite.legacyAminoCdc)
			result, err := querier(suite.ctx, test.path, abci.RequestQuery{})

			if test.expErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				var reports []types.Report
				suite.Require().NoError(suite.legacyAminoCdc.UnmarshalJSON(result, &reports))
				suite.Require().Equal(test.expResponse, reports)
			}
		})
	}
}
