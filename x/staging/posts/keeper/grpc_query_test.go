package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

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
