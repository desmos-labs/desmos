package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func (suite *KeeperTestsuite) TestQueryServer_Posts() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryPostsRequest
		shouldErr bool
		expPosts  []types.Post
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryPostsRequest(0, nil),
			shouldErr: true,
		},
		{
			name: "valid request without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
				))
			},
			request:   types.NewQueryPostsRequest(1, nil),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
				types.NewPost(
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
				),
			},
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
				))
			},
			request: types.NewQueryPostsRequest(1, &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				),
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

			res, err := suite.k.Posts(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPosts, res.Posts)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_Post() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryPostRequest
		shouldErr bool
		expPost   types.Post
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryPostRequest(0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryPostRequest(1, 0),
			shouldErr: true,
		},
		{
			name:      "not found post returns error",
			request:   types.NewQueryPostRequest(1, 1),
			shouldErr: true,
		},
		{
			name: "existing post is returned properly",
			store: func(ctx sdk.Context) {
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
			},
			request:   types.NewQueryPostRequest(1, 1),
			shouldErr: false,
			expPost: types.NewPost(
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
			),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Post(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPost, res.Post)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_PostAttachments() {
	testCases := []struct {
		name           string
		store          func(ctx sdk.Context)
		request        *types.QueryPostAttachmentsRequest
		shouldErr      bool
		expAttachments []types.Attachment
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryPostAttachmentsRequest(0, 1, nil),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryPostAttachmentsRequest(1, 0, nil),
			shouldErr: true,
		},
		{
			name: "valid request without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/second-image.png",
					"image/png",
				)))
			},
			request: types.NewQueryPostAttachmentsRequest(1, 1, nil),
			expAttachments: []types.Attachment{
				types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
				types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/second-image.png",
					"image/png",
				)),
			},
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)))
				suite.k.SaveAttachment(ctx, types.NewAttachment(1, 1, 2, types.NewMedia(
					"ftp://user:password@example.com/second-image.png",
					"image/png",
				)))
			},
			request: types.NewQueryPostAttachmentsRequest(1, 1, &query.PageRequest{
				Limit: 1,
			}),
			expAttachments: []types.Attachment{
				types.NewAttachment(1, 1, 1, types.NewMedia(
					"ftp://user:password@example.com/image.png",
					"image/png",
				)),
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

			res, err := suite.k.PostAttachments(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expAttachments, res.Attachments)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_PollAnswers() {
	firstUser, err := sdk.AccAddressFromBech32("cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st")
	suite.Require().NoError(err)

	secondUser, err := sdk.AccAddressFromBech32("cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea")
	suite.Require().NoError(err)

	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		request    *types.QueryPollAnswersRequest
		shouldErr  bool
		expAnswers []types.UserAnswer
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryPollAnswersRequest(0, 1, 1, "", nil),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryPollAnswersRequest(1, 0, 1, "", nil),
			shouldErr: true,
		},
		{
			name:      "invalid poll id returns error",
			request:   types.NewQueryPollAnswersRequest(1, 1, 0, "", nil),
			shouldErr: true,
		},
		{
			name:      "invalid user returns error",
			request:   types.NewQueryPollAnswersRequest(1, 1, 1, "user", nil),
			shouldErr: true,
		},
		{
			name: "valid request without user and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, secondUser))
			},
			request:   types.NewQueryPollAnswersRequest(1, 1, 1, "", nil),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser),
				types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser),
			},
		},
		{
			name: "valid request without user and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, secondUser))
			},
			request: types.NewQueryPollAnswersRequest(1, 1, 1, "", &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser),
			},
		},
		{
			name: "valid request with user and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, secondUser))
			},
			request:   types.NewQueryPollAnswersRequest(1, 1, 1, firstUser.String(), nil),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser),
			},
		},
		{
			name: "valid request with user and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, firstUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, secondUser))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, secondUser))
			},
			request: types.NewQueryPollAnswersRequest(1, 1, 1, firstUser.String(), &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, firstUser),
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

			res, err := suite.k.PollAnswers(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expAnswers, res.Answers)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_PollTallyResults() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		request    *types.QueryPollTallyResultsRequest
		shouldErr  bool
		expResults types.PollTallyResults
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryPollTallyResultsRequest(0, 1, 1),
			shouldErr: true,
		},
		{
			name:      "invalid post id returns error",
			request:   types.NewQueryPollTallyResultsRequest(1, 0, 1),
			shouldErr: true,
		},
		{
			name:      "invalid poll id returns error",
			request:   types.NewQueryPollTallyResultsRequest(1, 1, 0),
			shouldErr: true,
		},
		{
			name:      "not found results return error",
			request:   types.NewQueryPollTallyResultsRequest(1, 1, 1),
			shouldErr: true,
		},
		{
			name:    "results are returned properly",
			request: types.NewQueryPollTallyResultsRequest(1, 1, 1),
			store: func(ctx sdk.Context) {
				suite.k.SavePollTallyResults(ctx, types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
					types.NewAnswerResult(0, 10),
					types.NewAnswerResult(1, 14),
				}))
			},
			shouldErr: false,
			expResults: types.NewPollTallyResults(1, 1, 1, []types.PollTallyResults_AnswerResult{
				types.NewAnswerResult(0, 10),
				types.NewAnswerResult(1, 14),
			}),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.PollTallyResults(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResults, res.Results)
			}
		})
	}
}

func (suite *KeeperTestsuite) TestQueryServer_Params() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QueryParamsRequest
		expParams types.Params
	}{
		{
			name: "params are returned properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.NewParams(200))
			},
			request:   types.NewQueryParamsRequest(),
			expParams: types.NewParams(200),
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			res, err := suite.k.Params(sdk.WrapSDKContext(ctx), tc.request)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expParams, res.Params)
		})
	}
}
