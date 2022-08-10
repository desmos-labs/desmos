package wasm_test

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v4/x/relationships/types"
	"github.com/desmos-labs/desmos/v4/x/relationships/wasm"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func (suite *TestSuite) TestProfilesWasmQuerier_QueryCustom() {
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
			name: "relationships query request is parsed correctly",
			request: buildRelationshipsQueryRequest(suite.cdc,
				types.NewQueryRelationshipsRequest(
					0,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					nil,
				),
			),
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryRelationshipsResponse{
					Relationships: []types.Relationship{
						types.NewRelationship(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
							0,
						),
					},
					Pagination: nil,
				},
			),
		},
		{
			name: "blocks query request is parsed correctly",
			request: buildBlocksQueryRequest(suite.cdc, types.NewQueryBlocksRequest(
				0,
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				nil),
			),
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				))
			},
			shouldErr: false,
			expResponse: suite.cdc.MustMarshalJSON(
				&types.QueryBlocksResponse{
					Blocks: []types.UserBlock{
						types.NewUserBlock(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
							"",
							0,
						),
					},
					Pagination: nil,
				},
			),
		},
	}

	querier := wasm.NewRelationshipsWasmQuerier(suite.k, suite.cdc)

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
