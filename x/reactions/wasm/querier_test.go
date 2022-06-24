package wasm_test

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
	"github.com/desmos-labs/desmos/v4/x/reactions/wasm"
)

func (suite *Testsuite) TestReactionsWasmQuerier_QueryCustom() {
	profilesQuery := profilestypes.QueryProfileRequest{User: ""}
	profilesQueryBz, err := profilesQuery.Marshal()
	suite.NoError(err)
	wrongQueryBz, err := json.Marshal(profilesQueryBz)
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
			name:    "reactions request is parsed correctly",
			request: buildReactionsQueryRequest(suite.cdc, types.NewQueryReactionsRequest(1, 1, "", nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryReactionsResponse{
					Reactions: []types.Reaction{types.NewReaction(
						1,
						1,
						1,
						types.NewRegisteredReactionValue(1),
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					)},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
		{
			name:    "reaction request is parsed correctly",
			request: buildReactionQueryRequest(suite.cdc, types.NewQueryReactionRequest(1, 1, 1)),
			store: func(ctx sdk.Context) {
				suite.k.SaveReaction(ctx, types.NewReaction(
					1,
					1,
					1,
					types.NewRegisteredReactionValue(1),
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryReactionResponse{
					Reaction: types.NewReaction(
						1,
						1,
						1,
						types.NewRegisteredReactionValue(1),
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
			),
		},
		{
			name:    "registered reasons request is parsed correctly",
			request: buildRegisteredReactionsQueryRequest(suite.cdc, types.NewQueryRegisteredReactionsRequest(1, nil)),
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(1, 1, "shorthand_code", "display_value"))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryRegisteredReactionsResponse{RegisteredReactions: []types.RegisteredReaction{types.NewRegisteredReaction(1, 1, "shorthand_code", "display_value")},
					Pagination: &query.PageResponse{NextKey: nil, Total: 1},
				},
			),
		},
		{
			name:    "registered reason request is parsed correctly",
			request: buildRegisteredReactionQueryRequest(suite.cdc, types.NewQueryRegisteredReactionRequest(1, 1)),
			store: func(ctx sdk.Context) {
				suite.k.SaveRegisteredReaction(ctx, types.NewRegisteredReaction(1, 1, "shorthand_code", "display_value"))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryRegisteredReactionResponse{
					RegisteredReaction: types.NewRegisteredReaction(1, 1, "shorthand_code", "display_value"),
				},
			),
		},
		{
			name:    "reactions parameters request is parsed correctly",
			request: buildReactionsParamsQueryRequest(suite.cdc, types.NewQueryReactionsParamsRequest(1)),
			store: func(ctx sdk.Context) {
				suite.k.SaveSubspaceReactionsParams(ctx, types.NewSubspaceReactionsParams(1, types.NewRegisteredReactionValueParams(true), types.NewFreeTextValueParams(true, 100, "")))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryReactionsParamsResponse{Params: types.NewSubspaceReactionsParams(1, types.NewRegisteredReactionValueParams(true), types.NewFreeTextValueParams(true, 100, ""))},
			),
		},
	}

	querier := wasm.NewReactionsWasmQuerier(suite.k, suite.cdc)

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
