package wasm_test

import (
	"encoding/json"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
	"github.com/desmos-labs/desmos/v3/x/posts/wasm"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func (suite *TestSuite) TestQuerier_QueryCustom() {
	subspacesQuery := subspacestypes.QuerySubspacesRequest{Pagination: nil}
	subspacesQueryBz, err := subspacesQuery.Marshal()
	suite.NoError(err)
	wrongQueryBz, err := json.Marshal(subspacesQueryBz)
	suite.NoError(err)

	testCases := []struct {
		name        string
		request     json.RawMessage
		store       func(ctx sdk.Context)
		shouldErr   bool
		expResponse []byte
	}{
		{
			name:        "wrong request type returns error",
			request:     wrongQueryBz,
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name: "subspace postss request is parsed correctly",
			request: buildSubspacePostsQueryRequest(
				suite.cdc,
				types.NewQuerySubspacePostsRequest(1, nil),
			),
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QuerySubspacePostsResponse{
					Posts: []types.Post{types.NewPost(
						1,
						0,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						[]types.PostReference{},
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					)},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				}),
		},
		{
			name: "sections posts request is parsed correctly",
			request: buildIncomingDtagTransferQueryRequest(
				suite.cdc,
				types.NewQuerySectionPostsRequest(1, 1, nil),
			),
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QuerySectionPostsResponse{
					Posts: []types.Post{types.NewPost(
						1,
						1,
						1,
						"External ID",
						"This is a text",
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
						0,
						nil,
						[]types.PostReference{},
						types.REPLY_SETTING_EVERYONE,
						time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
						nil,
					)},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name:    "post request request is parsed correctly",
			request: buildPostQueryRequest(suite.cdc, types.NewQueryPostRequest(1, 1)),
			store: func(ctx sdk.Context) {
				suite.k.SavePost(ctx, types.NewPost(
					1,
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryPostResponse{Post: types.NewPost(
					1,
					1,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					0,
					nil,
					[]types.PostReference{},
					types.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				)},
			),
		},
		{
			name: "post attachments request is parsed properly",
			request: buildAppLinksQueryRequest(suite.cdc, types.NewQueryPostAttachmentsRequest(
				1,
				1,
				nil,
			)),
			store: func(ctx sdk.Context) {
				suite.k.SaveAttachment(ctx, types.NewAttachment(
					1,
					1,
					1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(&types.QueryPostAttachmentsResponse{
				Attachments: []types.Attachment{types.NewAttachment(
					1,
					1,
					1,
					types.NewMedia("ftp://user:password@host:post/media.png", "media/png"),
				)},
				Pagination: &query.PageResponse{NextKey: nil, Total: 1}},
			),
		},
		{
			name:    "poll answers request is parsed correctly",
			request: buildPollAnswersQueryRequest(suite.cdc, types.NewQueryPollAnswersRequest(1, 1, 1, "", nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveUserAnswer(ctx, types.NewUserAnswer(
					1,
					1,
					1,
					[]uint32{1},
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryPollAnswersResponse{
					Answers: []types.UserAnswer{types.NewUserAnswer(
						1,
						1,
						1,
						[]uint32{1},
						"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					)},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
	}

	querier := wasm.NewPostsWasmQuerier(suite.k, suite.cdc)

	for _, tc := range testCases {
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}
			query, err := querier.QueryCustom(ctx, tc.request)
			if tc.shouldErr {
				suite.Error(err)
			} else {
				suite.NoError(err)
			}
			suite.Equal(tc.expResponse, query)
		})
	}
}
