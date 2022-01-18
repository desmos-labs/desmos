package keeper_test

import (
	"time"

	"github.com/desmos-labs/desmos/v2/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (suite *KeeperTestSuite) TestQueryServer_Profile() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryProfileRequest
		shouldErr   bool
		expResponse *types.QueryProfileResponse
	}{
		{
			name:      "empty user returns error",
			req:       types.NewQueryProfileRequest(""),
			shouldErr: true,
		},
		{
			name:      "non existing DTag returns error",
			req:       types.NewQueryProfileRequest("invalid-dtag"),
			shouldErr: true,
		},
		{
			name:        "profile not found",
			req:         types.NewQueryProfileRequest("cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa"),
			shouldErr:   false,
			expResponse: &types.QueryProfileResponse{Profile: nil},
		},
		{
			name: "found profile using DTag",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.StoreProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			req:       types.NewQueryProfileRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x-dtag"),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: testutil.NewAny(testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
			},
		},
		{
			name: "found profile using address",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.StoreProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			req:       types.NewQueryProfileRequest("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: testutil.NewAny(testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
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

			res, err := suite.k.Profile(sdk.WrapSDKContext(ctx), tc.req)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expResponse, res)
				if tc.expResponse.Profile != nil {
					suite.Require().Equal(tc.expResponse, res)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_IncomingDTagTransferRequests() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryIncomingDTagTransferRequestsRequest
		shouldErr   bool
		expRequests []types.DTagTransferRequest
	}{
		{
			name: "valid request without pagination",
			store: func(ctx sdk.Context) {
				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				request = types.NewDTagTransferRequest(
					"dtag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(request.Receiver)))
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				nil,
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				),
			},
		},
		{
			name: "valid request with pagination",
			store: func(ctx sdk.Context) {
				receiver := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(receiver)))

				request := types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					receiver,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))

				request = types.NewDTagTransferRequest(
					"dtag",
					"cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn",
					receiver,
				)
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, request))
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				&query.PageRequest{Limit: 1},
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
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

			res, err := suite.k.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx), tc.req)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expRequests, res.Requests)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_Params() {
	ctx, _ := suite.ctx.CacheContext()

	suite.k.SetParams(ctx, types.DefaultParams())

	res, err := suite.k.Params(sdk.WrapSDKContext(ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	suite.Require().Equal(types.DefaultParams(), res.Params)
}

func (suite *KeeperTestSuite) TestQueryServer_ChainLinks() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryChainLinksRequest
		shouldErr bool
		expLinks  []types.ChainLink
	}{
		{
			name: "valid request without pagination",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)
			},
			req: &types.QueryChainLinksRequest{
				User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
		{
			name: "valid request with pagination",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)

				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)
			},
			req: &types.QueryChainLinksRequest{
				User:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLinks: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.ChainLinks(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expLinks, res.Links)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserChainLink() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryUserChainLinkRequest
		shouldErr bool
		expRes    *types.QueryUserChainLinkResponse
	}{
		{
			name: "not found link returns error",
			req: &types.QueryUserChainLinkRequest{
				User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				ChainName: "cosmos",
				Target:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			shouldErr: true,
		},
		{
			name: "existing chain link returns proper response",
			store: func(ctx sdk.Context) {
				address := "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"
				suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(address))

				link := types.NewChainLink(
					address,
					types.NewBech32Address("cosmos1nc54z3kzyal57w6wcf5khmwrxx5rafnwvu0m5z", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
			},
			req: &types.QueryUserChainLinkRequest{
				User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				ChainName: "cosmos",
				Target:    "cosmos1nc54z3kzyal57w6wcf5khmwrxx5rafnwvu0m5z",
			},
			shouldErr: false,
			expRes: &types.QueryUserChainLinkResponse{
				Link: types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1nc54z3kzyal57w6wcf5khmwrxx5rafnwvu0m5z", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.UserChainLink(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(tc.expRes, res)
			}
		})
	}
}

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
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))
			},
			req:       &types.QueryRelationshipsRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
			expRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
		},
		{
			name: "query relationsips with pagination",
			store: func(ctx sdk.Context) {
				relationship := types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
				suite.Require().NoError(suite.k.SaveRelationship(ctx, relationship))

				relationship = types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(relationship.Creator)))
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
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
			},
			req:       &types.QueryBlocksRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
			expBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason1",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
				suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))

				block = types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"reason2",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				)
				suite.Require().NoError(suite.k.StoreProfile(ctx, testutil.ProfileFromAddr(block.Blocker)))
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
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
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

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinks() {
	testCases := []struct {
		name                string
		store               func(ctx sdk.Context)
		req                 *types.QueryApplicationLinksRequest
		shouldErr           bool
		expApplicationLinks []types.ApplicationLink
	}{
		{
			name:                "empty requests return empty result",
			req:                 types.NewQueryApplicationLinksRequest("user", nil),
			shouldErr:           false,
			expApplicationLinks: nil,
		},
		{
			name: "valid paginated request returns proper response",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("github", "githubuser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				))
			},
			req: types.NewQueryApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				&query.PageRequest{Limit: 1, Offset: 1, CountTotal: true},
			),
			shouldErr: false,
			expApplicationLinks: []types.ApplicationLink{
				types.NewApplicationLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData(
							"twitter",
							"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
						),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.ApplicationLinks(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expApplicationLinks, res.Links)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserApplicationLink() {
	testCases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryUserApplicationLinkRequest
		shouldErr   bool
		expResponse *types.QueryUserApplicationLinkResponse
	}{
		{
			name:      "not found link returns error",
			req:       types.NewQueryUserApplicationLinkRequest("user", "application", "username"),
			shouldErr: true,
		},
		{
			name: "valid request returns proper response",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			req: types.NewQueryUserApplicationLinkRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"twitter",
				"twitteruser",
			),
			shouldErr: false,
			expResponse: &types.QueryUserApplicationLinkResponse{
				Link: types.NewApplicationLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData(
							"twitter",
							"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
						),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.UserApplicationLink(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinkByClientID() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		req       *types.QueryApplicationLinkByClientIDRequest
		shouldErr bool
		expLink   types.ApplicationLink
	}{
		{
			name:      "not found link returns error",
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: true,
		},
		{
			name: "valid request returns proper response",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: false,
			expLink: types.NewApplicationLink(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				types.NewData("twitter", "twitteruser"),
				types.ApplicationLinkStateInitialized,
				types.NewOracleRequest(
					0,
					1,
					types.NewOracleRequestCallData(
						"twitter",
						"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
					),
					"client_id",
				),
				nil,
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

			res, err := suite.k.ApplicationLinkByClientID(sdk.WrapSDKContext(ctx), tc.req)
			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expLink, res.Link)
			}
		})
	}
}
