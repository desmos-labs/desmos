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
			PollData: &types.PollData{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
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
			Subspace:             "5e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
			AdditionalAttributes: nil,
			Creator:              "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
			Attachments: types.NewAttachments(
				types.NewAttachment(
					"https://uri.com",
					"text/plain",
					[]string{"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
				),
			),
			PollData: &types.PollData{
				Question: "poll?",
				ProvidedAnswers: types.NewPollAnswers(
					types.NewPollAnswer("1", "Yes"),
					types.NewPollAnswer("2", "No"),
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
		expResponse *types.QueryPostsResponse
	}{
		{
			name: "request with parent id returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{ParentId: "19de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af"},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[1]},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
		{
			name: "request with creation time returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{CreationTime: &creationDate},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[0]},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
		{
			name: "request with parent id returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{Subspace: "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e"},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[0]},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
		{
			name: "request with creator returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{Creator: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[0]},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
		{
			name: "request with hashtags returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{Hashtags: []string{"desmos"}},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[0]},
				Pagination: &query.PageResponse{Total: 1},
			},
		},
		{
			name: "request with pagination returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{Pagination: &query.PageRequest{Limit: 1, Offset: 0}},
			expResponse: &types.QueryPostsResponse{
				Posts:      []types.Post{posts[0]},
				Pagination: &query.PageResponse{NextKey: []byte("29de02e105c68a60e45c289bff19fde745bca9c63c38f2095b59e8e8090ae1af")},
			},
		},
		{
			name: "request without properties returns properly",
			store: func(ctx sdk.Context) {
				for _, post := range posts {
					suite.k.SavePost(ctx, post)
				}
			},
			req: &types.QueryPostsRequest{},
			expResponse: &types.QueryPostsResponse{
				Posts:      posts,
				Pagination: &query.PageResponse{Total: 2},
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
			suite.Require().NoError(err)
			suite.Require().Equal(uc.expResponse, res)

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
		PollData: &types.PollData{
			Question: "poll?",
			ProvidedAnswers: types.NewPollAnswers(
				types.NewPollAnswer("1", "Yes"),
				types.NewPollAnswer("2", "No"),
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
			req:    &types.QueryRegisteredReactionsRequest{Subspace: "subspace1"},
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
			suite.Require().Equal(uc.expLen, len(res.RegisteredReactions))
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
					PollData:             nil,
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
					PollData:             suite.testData.post.PollData,
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
					PollData:             suite.testData.post.PollData,
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
					PollData:             suite.testData.post.PollData,
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
