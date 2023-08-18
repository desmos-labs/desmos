package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
)

func (suite *KeeperTestSuite) TestQueryServer_SubspacePosts() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QuerySubspacePostsRequest
		shouldErr bool
		expPosts  []types.Post
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQuerySubspacePostsRequest(0, nil),
			shouldErr: true,
		},
		{
			name: "valid request without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			request:   types.NewQuerySubspacePostsRequest(1, nil),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					0,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				),
				types.NewPost(
					1,
					0,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				),
			},
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			request: types.NewQuerySubspacePostsRequest(1, &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					0,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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

			res, err := suite.k.SubspacePosts(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPosts, res.Posts)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_SectionPosts() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		request   *types.QuerySectionPostsRequest
		shouldErr bool
		expPosts  []types.Post
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQuerySectionPostsRequest(0, 1, nil),
			shouldErr: true,
		},
		{
			name: "valid request without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			request:   types.NewQuerySectionPostsRequest(1, 1, nil),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				),
				types.NewPost(
					1,
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				),
			},
		},
		{
			name: "valid request with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					2,
					"",
					"Second post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 13, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			request: types.NewQuerySectionPostsRequest(1, 1, &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expPosts: []types.Post{
				types.NewPost(
					1,
					1,
					1,
					"",
					"First post!",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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

			res, err := suite.k.SectionPosts(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expPosts, res.Posts)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Post() {
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
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
				))
			},
			request:   types.NewQueryPostRequest(1, 1),
			shouldErr: false,
			expPost: types.NewPost(
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
				types.REPLY_SETTING_EVERYONE,
				time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				nil,
				"cosmos1r9jamre0x0qqy562rhhckt6sryztwhnvhafyz4",
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

func (suite *KeeperTestSuite) TestQueryServer_PostAttachments() {
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

func (suite *KeeperTestSuite) TestQueryServer_PollAnswers() {
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
			name: "valid request without user and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
			},
			request:   types.NewQueryPollAnswersRequest(1, 1, 1, "", nil),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"),
			},
		},
		{
			name: "valid request without user and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
			},
			request: types.NewQueryPollAnswersRequest(1, 1, 1, "", &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			},
		},
		{
			name: "valid request with user and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
			},
			request:   types.NewQueryPollAnswersRequest(1, 1, 1, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st", nil),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
			},
		},
		{
			name: "valid request with user and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(1, 1, 2, []uint32{1}, "cosmos1xnjndlr28kuymqexyk9m9m3kqyvm8fge0edxea"))
			},
			request: types.NewQueryPollAnswersRequest(1, 1, 1, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st", &query.PageRequest{
				Limit: 1,
			}),
			shouldErr: false,
			expAnswers: []types.UserAnswer{
				types.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1vs8dps0ktst5ekynmszxuxphfq08rhmepsn8st"),
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

func (suite *KeeperTestSuite) TestQueryServer_Params() {
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

func (suite *KeeperTestSuite) TestQueryServer_IncomingPostOwnerTransferRequests() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		request     *types.QueryIncomingPostOwnerTransferRequestsRequest
		shouldErr   bool
		expRequests []types.PostOwnerTransferRequest
	}{
		{
			name:      "invalid subspace id returns error",
			request:   types.NewQueryIncomingPostOwnerTransferRequestsRequest(0, "", nil),
			shouldErr: true,
		},
		{
			name: "valid request without receiver and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 3, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"))
			},
			request:   types.NewQueryIncomingPostOwnerTransferRequestsRequest(1, "", nil),
			shouldErr: false,
			expRequests: []types.PostOwnerTransferRequest{
				types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"),
				types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"),
				types.NewPostOwnerTransferRequest(1, 3, "other_receiver", "sender"),
				types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"),
			},
		},
		{
			name: "valid request with receiver and without pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 3, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"))
			},
			request:   types.NewQueryIncomingPostOwnerTransferRequestsRequest(1, "receiver", nil),
			shouldErr: false,
			expRequests: []types.PostOwnerTransferRequest{
				types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"),
				types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"),
			},
		},
		{
			name: "valid request without receiver and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 3, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"))
			},
			request:   types.NewQueryIncomingPostOwnerTransferRequestsRequest(1, "", &query.PageRequest{Limit: 1}),
			shouldErr: false,
			expRequests: []types.PostOwnerTransferRequest{
				types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"),
			},
		},
		{
			name: "valid request with receiver and with pagination returns properly",
			store: func(ctx sdk.Context) {
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 1, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 3, "other_receiver", "sender"))
				suite.k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(1, 4, "receiver", "sender"))
			},
			request:   types.NewQueryIncomingPostOwnerTransferRequestsRequest(1, "receiver", &query.PageRequest{Limit: 1}),
			shouldErr: false,
			expRequests: []types.PostOwnerTransferRequest{
				types.NewPostOwnerTransferRequest(1, 2, "receiver", "sender"),
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

			res, err := suite.k.IncomingPostOwnerTransferRequests(sdk.WrapSDKContext(ctx), tc.request)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expRequests, res.Requests)
			}
		})
	}
}
