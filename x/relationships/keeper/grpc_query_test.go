package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Relationships() {
	testCases := []struct {
		name             string
		store            func(ctx sdk.Context)
		req              *types.QueryRelationshipsRequest
		expRelationships []types.Relationship
	}{
		{
			name: "query relationships with specific user and counterparty",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				))
			},
			req: types.NewQueryRelationshipsRequest(
				0,
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				nil,
			),
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				),
			},
		},
		{
			name: "query relationships with specific user",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					0,
				))
			},
			req: types.NewQueryRelationshipsRequest(
				0,
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				nil,
			),
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				),
			},
		},
		{
			name: "query relationships without pagination",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				))
			},
			req: types.NewQueryRelationshipsRequest(0, "", "", nil),
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				),
			},
		},
		{
			name: "query relationsips with pagination",
			store: func(ctx sdk.Context) {
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				))
				suite.k.SaveRelationship(ctx, types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				))
			},
			req: types.NewQueryRelationshipsRequest(0, "", "", &query.PageRequest{Limit: 1}),
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
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

			res, err := suite.k.Relationships(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			suite.Require().Equal(tc.expRelationships, res.Relationships)
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Blocks() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryBlocksRequest
		expBlocks []types.UserBlock
	}{
		{
			name: "query blocks for specific blocker and blocked",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				))
			},
			req: types.NewQueryBlocksRequest(
				0,
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				nil,
			),
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				),
			},
		},
		{
			name: "query blocks for specific user",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"",
					0,
				))
			},
			req: types.NewQueryBlocksRequest(
				0,
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"",
				nil,
			),
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				),
			},
		},
		{
			name: "query blocks without pagination",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				))
			},
			req: types.NewQueryBlocksRequest(0, "", "", nil),
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				),
			},
		},
		{
			name: "query blocks with pagination",
			store: func(ctx sdk.Context) {
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"",
					0,
				))
				suite.k.SaveUserBlock(ctx, types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
				))
			},
			req: types.NewQueryBlocksRequest(0, "", "", &query.PageRequest{Limit: 1}),
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"",
					0,
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

			res, err := suite.k.Blocks(sdk.WrapSDKContext(ctx), tc.req)
			suite.Require().NoError(err)
			suite.Require().NotNil(res)
			suite.Require().Equal(tc.expBlocks, res.Blocks)
		})
	}
}
