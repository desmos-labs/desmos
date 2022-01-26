package keeper_test

import (
	"github.com/desmos-labs/desmos/v2/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Relationships() {
	testCases := []struct {
		name             string
		store            func(ctx sdk.Context)
		req              *types.QueryRelationshipsRequest
		shouldErr        bool
		expRelationships []types.Relationship
	}{
		{
			name: "query relationships without pagination",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			req:       &types.QueryRelationshipsRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
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
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			req: &types.QueryRelationshipsRequest{
				User:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
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
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(tc.expRelationships, res.Relationships)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Blocks() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryBlocksRequest
		shouldErr bool
		expBlocks []types.UserBlock
	}{
		{
			name: "query blocks without pagination",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason1",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			req:       &types.QueryBlocksRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					0,
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason1",
					0,
				),
			},
		},
		{
			name: "query blocks with pagination",
			store: func(ctx sdk.Context) {
				block := types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason1",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					0,
				)
				suite.Require().NoError(suite.pk.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			req: &types.QueryBlocksRequest{
				User:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
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
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(tc.expBlocks, res.Blocks)
			}
		})
	}
}
