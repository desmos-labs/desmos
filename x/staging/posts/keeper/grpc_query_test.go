package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

func (suite *KeeperTestSuite) Test_Posts() {
	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.Require().NoError(err)
	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.Require().NoError(err)

	posts := []types.Post{
		{
			PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message #desmos",
			Created:              creationDate,
			LastEdited:           creationDate.Add(1),
			CommentsState:        types.CommentsStateAllowed,
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			Attachments: types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
				),
			),
			Poll: &types.Poll{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewProvidedAnswer("1", "Yes"),
					types.NewProvidedAnswer("2", "No"),
				),
				EndDate:               pollEndDate,
				AllowsMultipleAnswers: true,
				AllowsAnswerEdits:     true,
			},
		},
		{
			PostID:               "29de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message",
			Created:              creationDate.Add(2),
			LastEdited:           creationDate.Add(2),
			CommentsState:        types.CommentsStateAllowed,
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
			Attachments: types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
				),
			),
			Poll: &types.Poll{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewProvidedAnswer("1", "Yes"),
					types.NewProvidedAnswer("2", "No"),
				),
				EndDate:               pollEndDate,
				AllowsMultipleAnswers: true,
				AllowsAnswerEdits:     true,
			},
		},
	}

	usecases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryPostsRequest
		shouldErr   bool
		expResponse *types.QueryPostsResponse
	}{
		{
			name:      "invalid request without subspace id returns error",
			store:     func(ctx sdk.Context) {},
			req:       &types.QueryPostsRequest{},
			shouldErr: true,
		},
		{
			name: "request with subspace id returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req:       &types.QueryPostsRequest{SubspaceId: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			shouldErr: false,
			expResponse: &types.QueryPostsResponse{
				Posts:      posts,
				Pagination: &query.PageResponse{Total: 2},
			},
		},
		{
			name: "request with pagination returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{
				SubspaceId: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				Pagination: &query.PageRequest{Limit: 1, Offset: 0},
			},
			shouldErr: false,
			expResponse: &types.QueryPostsResponse{
				Posts: []types.Post{posts[0]},
				Pagination: &query.PageResponse{
					NextKey: []byte("29de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"),
				},
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()
			if uc.store != nil {
				uc.store(suite.ctx)
			}

			res, err := suite.k.Posts(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_Post() {
	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.Require().NoError(err)
	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.Require().NoError(err)

	post := types.Post{
		PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		Message:              "Post message",
		Created:              creationDate,
		LastEdited:           creationDate.Add(1),
		Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		AdditionalAttributes: nil,
		Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		Attachments: types.NewAttachments(
			types.NewAttachment(
				"https://uri.com",
				"text/plain",
				[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			),
		),
		Poll: &types.Poll{
			Question: "poll?",
			ProvidedAnswers: types.NewPollAnswers(
				types.NewProvidedAnswer("1", "Yes"),
				types.NewProvidedAnswer("2", "No"),
			),
			EndDate:               pollEndDate,
			AllowsMultipleAnswers: true,
			AllowsAnswerEdits:     true,
		},
	}

	usecases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryPostRequest
		shouldErr   bool
		expResponse *types.QueryPostResponse
	}{
		{
			name:      "request with invalid post id returns error",
			req:       &types.QueryPostRequest{PostId: ""},
			shouldErr: true,
		},
		{
			name:      "request non existent post returns error",
			req:       &types.QueryPostRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "valid request returns data properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, post)
			},
			req:       &types.QueryPostRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: false,
			expResponse: &types.QueryPostResponse{
				Post: post,
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()
			if uc.store != nil {
				uc.store(suite.ctx)
			}
			res, err := suite.k.Post(sdk.WrapSDKContext(suite.ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_RegisteredReactions() {
	usecases := []struct {
		name            string
		storedReactions []types.RegisteredReaction
		req             *types.QueryRegisteredReactionsRequest
		expLen          int
	}{
		{
			name: "query registered reactions without subspace and pagination",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{},
			expLen: 2,
		},
		{
			name: "query registered reactions with a subspace",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{SubspaceId: "subspace1"},
			expLen: 1,
		},
		{
			name: "query registered reactions with pagination",
			storedReactions: []types.RegisteredReaction{
				types.NewRegisteredReaction(
					"creator",
					":smile:",
					"smile",
					"subspace1",
				),
				types.NewRegisteredReaction(
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					":fire:",
					"fire",
					"subspace2",
				),
			},
			req:    &types.QueryRegisteredReactionsRequest{Pagination: &query.PageRequest{Limit: 1}},
			expLen: 1,
		},
	}

	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			suite.SetupTest()
			for _, reaction := range uc.storedReactions {
				suite.k.SaveRegisteredReaction(suite.ctx, reaction)
			}

			res, err := suite.k.RegisteredReactions(sdk.WrapSDKContext(suite.ctx), uc.req)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			suite.Require().Equal(uc.expLen, len(res.Reactions))
		})
	}
}

func (suite *KeeperTestSuite) Test_UserAnswers() {
	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserAnswersRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid post id returns error",
			req:       &types.QueryUserAnswersRequest{},
			shouldErr: true,
		},
		{
			name:      "non existent post returns error",
			req:       &types.QueryUserAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "non existent poll data in the post returns error",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					Poll:                 nil,
				}
				suite.k.SavePost(ctx, post)
			},
			req:       &types.QueryUserAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					Poll:                 suite.testData.post.Poll,
				}
				suite.k.SavePost(ctx, post)

				answers := []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				}

				for _, answer := range answers {
					suite.k.SaveUserAnswer(suite.ctx, answer)
				}

			},
			req:       &types.QueryUserAnswersRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid request with user returns properly",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					Poll:                 suite.testData.post.Poll,
				}
				suite.k.SavePost(ctx, post)

				answers := []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				}

				for _, answer := range answers {
					suite.k.SaveUserAnswer(suite.ctx, answer)
				}

			},
			req: &types.QueryUserAnswersRequest{
				PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				User:   "cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
			},
			shouldErr: false,
			expLen:    1,
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				post := types.Post{
					PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
					Message:              "Post with poll data",
					Created:              suite.testData.post.Created,
					LastEdited:           suite.testData.post.LastEdited,
					Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					AdditionalAttributes: nil,
					Creator:              suite.testData.postOwner,
					Poll:                 suite.testData.post.Poll,
				}
				suite.k.SavePost(ctx, post)

				answers := []types.UserAnswer{
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1unacjuhyamzks5yu7qwlfuahdedd838e6fmdta",
						[]string{"1"},
					),
					types.NewUserAnswer(
						"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						[]string{"1"},
					),
				}

				for _, answer := range answers {
					suite.k.SaveUserAnswer(suite.ctx, answer)
				}

			},
			req: &types.QueryUserAnswersRequest{
				PostId:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}

	suite.SetupTest()
	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}
			res, err := suite.k.UserAnswers(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Answers))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_PostReactions() {

	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.Require().NoError(err)

	post := types.Post{
		PostID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		Message:    "Post message",
		Created:    creationDate,
		LastEdited: creationDate.Add(1),
		Subspace:   "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		Creator:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	}

	reactions := []types.PostReaction{
		types.NewPostReaction(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			":smile:",
			"reaction",
			"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		),
		types.NewPostReaction(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			":smile:",
			"reaction",
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		),
	}

	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryPostReactionsRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid post id returns error",
			req:       &types.QueryPostReactionsRequest{},
			shouldErr: true,
		},
		{
			name:      "non existent post return error",
			req:       &types.QueryPostReactionsRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, post)
				for _, reaction := range reactions {
					suite.k.SavePostReaction(ctx, reaction)
				}
			},
			req:       &types.QueryPostReactionsRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, post)
				for _, reaction := range reactions {
					suite.k.SavePostReaction(ctx, reaction)
				}
			},
			req: &types.QueryPostReactionsRequest{
				PostId:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}
	suite.SetupTest()
	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}
			res, err := suite.k.PostReactions(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Reactions))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_PostComments() {
	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.Require().NoError(err)
	pollEndDate, err := time.Parse(time.RFC3339, "2050-01-01T15:15:00.000Z")
	suite.Require().NoError(err)
	posts := []types.Post{
		{
			PostID:               "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message #desmos",
			Created:              creationDate,
			LastEdited:           creationDate.Add(1),
			CommentsState:        types.CommentsStateAllowed,
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			Attachments: types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
				),
			),
			Poll: &types.Poll{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewProvidedAnswer("1", "Yes"),
					types.NewProvidedAnswer("2", "No"),
				),
				EndDate:               pollEndDate,
				AllowsMultipleAnswers: true,
				AllowsAnswerEdits:     true,
			},
		},
		{
			PostID:               "29de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message",
			Created:              creationDate.Add(2),
			LastEdited:           creationDate.Add(2),
			CommentsState:        types.CommentsStateAllowed,
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		},
		{
			PostID:               "39de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			ParentID:             "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			Message:              "Post message",
			Created:              creationDate.Add(2),
			LastEdited:           creationDate.Add(2),
			CommentsState:        types.CommentsStateAllowed,
			Subspace:             "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
		},
	}

	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryPostCommentsRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid request return error",
			req:       &types.QueryPostCommentsRequest{},
			shouldErr: true,
		},
		{
			name: "non existent post id returns error",
			req: &types.QueryPostCommentsRequest{
				PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			},
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostCommentsRequest{
				PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostCommentsRequest{
				PostId:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}
	suite.SetupTest()
	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}
			res, err := suite.k.PostComments(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Comments))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_Reports() {

	creationDate, err := time.Parse(time.RFC3339, "2020-01-01T15:15:00.000Z")
	suite.Require().NoError(err)

	post := types.Post{
		PostID:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
		Message:    "Post message",
		Created:    creationDate,
		LastEdited: creationDate.Add(1),
		Subspace:   "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		Creator:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
	}

	reports := []types.Report{
		types.NewReport(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			[]string{"scam"},
			"this is a test",
			"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
		),
		types.NewReport(
			"19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
			[]string{"scam"},
			"this is a test",
			"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
		),
	}

	usecases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryReportsRequest
		shouldErr bool
		expLen    int
	}{
		{
			name:      "invalid post id returns error",
			req:       &types.QueryReportsRequest{},
			shouldErr: true,
		},
		{
			name:      "non existent post returns error",
			req:       &types.QueryReportsRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: true,
		},
		{
			name: "valid request returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(suite.ctx, types.DefaultParams())
				suite.k.SavePost(ctx, post)
				for _, report := range reports {
					err := suite.k.SaveReport(ctx, report)
					suite.Require().NoError(err)
				}
			},
			req:       &types.QueryReportsRequest{PostId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, post)
				suite.k.SetParams(suite.ctx, types.DefaultParams())
				for _, report := range reports {
					err := suite.k.SaveReport(ctx, report)
					suite.Require().NoError(err)
				}
			},
			req: &types.QueryReportsRequest{
				PostId:     "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}
	suite.SetupTest()
	for _, uc := range usecases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}
			res, err := suite.k.Reports(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Reports))
			}
		})
	}
}
