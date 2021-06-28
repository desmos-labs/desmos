package keeper_test

import (
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) codeToAny(profile *types.Profile) *codectypes.Any {
	accountAny, err := codectypes.NewAnyWithValue(profile)
	suite.Require().NoError(err)
	return accountAny
}

func (suite *KeeperTestSuite) Test_Profile() {
	usecases := []struct {
		name           string
		storedProfiles []*types.Profile
		req            *types.QueryProfileRequest
		shouldErr      bool
		expResponse    *types.QueryProfileResponse
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
			name: "found profile - using dtag",
			storedProfiles: []*types.Profile{
				suite.testData.profile.Profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.DTag),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile.Profile),
			},
		},
		{
			name: "found profile - using address",
			storedProfiles: []*types.Profile{
				suite.testData.profile.Profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.GetAddress().String()),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile.Profile),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, profile := range uc.storedProfiles {
				suite.Require().NoError(suite.k.StoreProfile(suite.ctx, profile))
			}

			res, err := suite.k.Profile(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expResponse, res)

				if uc.expResponse.Profile != nil {
					// Make sure the cached value is not nil (this is to grant that UnpackInterfaces work properly)
					suite.Require().NotNil(res.Profile.GetCachedValue())
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_IncomingDTagTransferRequests() {
	usecases := []struct {
		name           string
		storedRequests []types.DTagTransferRequest
		req            *types.QueryIncomingDTagTransferRequestsRequest
		shouldErr      bool
		expRequests    []types.DTagTransferRequest
	}{
		{
			name:      "invalid user",
			req:       types.NewQueryIncomingDTagTransferRequestsRequest("invalid-address", nil),
			shouldErr: true,
		},
		{
			name: "valid request without pagination",
			storedRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
				types.NewDTagTransferRequest(
					"dtag-2",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				),
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				nil,
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
			},
		},
		{
			name: "valid request with pagination",
			storedRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
				types.NewDTagTransferRequest(
					"dtag-2",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
			},
			req: types.NewQueryIncomingDTagTransferRequestsRequest(
				"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				&query.PageRequest{Limit: 1, Offset: 1, CountTotal: true},
			),
			shouldErr: false,
			expRequests: []types.DTagTransferRequest{
				types.NewDTagTransferRequest(
					"dtag",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, req := range uc.storedRequests {
				suite.Require().NoError(suite.k.SaveDTagTransferRequest(suite.ctx, req))
			}

			res, err := suite.k.IncomingDTagTransferRequests(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expRequests, res.Requests)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_Params() {
	suite.k.SetParams(suite.ctx, types.DefaultParams())

	res, err := suite.k.Params(sdk.WrapSDKContext(suite.ctx), &types.QueryParamsRequest{})
	suite.Require().NoError(err)
	suite.Require().NotNil(res)

	suite.Require().Equal(types.DefaultParams(), res.Params)
}

func (suite *KeeperTestSuite) Test_UserChainLinks() {
	var pubKey, err = sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	suite.Require().NoError(err)

	usecases := []struct {
		name      string
		store     func()
		req       *types.QueryUserChainLinksRequest
		shouldErr bool
		expLen    int
	}{
		{
			name: "query user chain links without pagination",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)

				link1 := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				)

				key := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link1))

				link2 := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				)
				key2 := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw")
				store.Set(key2, types.MustMarshalChainLink(suite.cdc, link2))
			},
			req: &types.QueryUserChainLinksRequest{
				User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "query user chain links with pagination",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)

				link1 := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				)

				key := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link1))

				link2 := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				)
				key2 := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw")
				store.Set(key2, types.MustMarshalChainLink(suite.cdc, link2))
			},
			req: &types.QueryUserChainLinksRequest{
				User:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				Pagination: &query.PageRequest{Limit: 1},
			},
			shouldErr: false,
			expLen:    1,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()
			uc.store()

			res, err := suite.k.UserChainLinks(sdk.WrapSDKContext(suite.ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expLen, len(res.Links))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_UserChainLink() {
	var pubKey, err = sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	suite.Require().NoError(err)

	usecases := []struct {
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
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))

				link := types.NewChainLink(
					address,
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				)
				suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
			},
			req: &types.QueryUserChainLinkRequest{
				User:      "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				ChainName: "cosmos",
				Target:    "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			},
			shouldErr: false,
			expRes: &types.QueryUserChainLinkResponse{
				Link: types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						pubKey,
						"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
						"text",
					),
					types.NewChainConfig("cosmos"),
					suite.testData.profile.CreationDate,
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			res, err := suite.k.UserChainLink(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expRes, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_UserRelationships() {
	usecases := []struct {
		name                string
		storedRelationships []types.Relationship
		req                 *types.QueryUserRelationshipsRequest
		shouldErr           bool
		expLen              int
	}{
		{
			name: "query relationsips without pagination",
			storedRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			req:       &types.QueryUserRelationshipsRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "query relationsips with pagination",
			storedRelationships: []types.Relationship{
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewRelationship(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			req:       &types.QueryUserRelationshipsRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", Pagination: &query.PageRequest{Limit: 1}},
			shouldErr: false,
			expLen:    1,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, relationship := range uc.storedRelationships {
				suite.k.SaveRelationship(suite.ctx, relationship)
			}

			res, err := suite.k.UserRelationships(sdk.WrapSDKContext(suite.ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Relationships))
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_UserBlocks() {
	usecases := []struct {
		name                string
		storedUserBlocks []types.UserBlock
		req                 *types.QueryUserBlocksRequest
		shouldErr           bool
		expLen              int
	}{
		{
			name: "query blocks without pagination",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"reason1",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason2",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			req:       &types.QueryUserBlocksRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"},
			shouldErr: false,
			expLen:    2,
		},
		{
			name: "query blocks with pagination",
			storedUserBlocks: []types.UserBlock{
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4",
					"reason1",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
				types.NewUserBlock(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					"reason2",
					"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
				),
			},
			req:       &types.QueryUserBlocksRequest{User: "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", Pagination: &query.PageRequest{Limit: 1}},
			shouldErr: false,
			expLen:    1,
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()

			for _, UserBlock := range uc.storedUserBlocks {
				suite.k.SaveUserBlock(suite.ctx, UserBlock)
			}

			res, err := suite.k.UserBlocks(sdk.WrapSDKContext(suite.ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)
				suite.Require().Equal(uc.expLen, len(res.Blocks))
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserApplicationLinks() {
	usecases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryUserApplicationLinksRequest
		shouldErr   bool
		expResponse *types.QueryUserApplicationLinksResponse
	}{
		{
			name:      "empty requests return empty result",
			req:       types.NewQueryUserApplicationLinksRequest("user", nil),
			shouldErr: false,
			expResponse: &types.QueryUserApplicationLinksResponse{
				Links:      nil,
				Pagination: &query.PageResponse{Total: 0, NextKey: nil},
			},
		},
		{
			name: "valid paginated request returns proper response",
			store: func(ctx sdk.Context) {
				profile := suite.CreateProfileFromAddress("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
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
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				))
			},
			req: types.NewQueryUserApplicationLinksRequest(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				&query.PageRequest{Limit: 1, Offset: 1, CountTotal: true},
			),
			shouldErr: false,
			expResponse: &types.QueryUserApplicationLinksResponse{
				Links: []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				Pagination: &query.PageResponse{Total: 2, NextKey: nil},
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			res, err := suite.k.UserApplicationLinks(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_UserApplicationLink() {
	usecases := []struct {
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
				profile := suite.CreateProfileFromAddress("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
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
						-1,
						1,
						types.NewOracleRequestCallData(
							"twitter",
							"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
						),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			res, err := suite.k.UserApplicationLink(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestQueryServer_ApplicationLinkByClientID() {
	usecases := []struct {
		name        string
		store       func(ctx sdk.Context)
		req         *types.QueryApplicationLinkByClientIDRequest
		shouldErr   bool
		expResponse *types.QueryApplicationLinkByClientIDResponse
	}{
		{
			name:      "not found link returns error",
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: true,
		},
		{
			name: "valid request returns proper response",
			store: func(ctx sdk.Context) {
				profile := suite.CreateProfileFromAddress("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile)

				suite.Require().NoError(suite.k.SaveApplicationLink(
					ctx,
					types.NewApplicationLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData(
								"twitter",
								"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
							),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					)),
				)
			},
			req:       types.NewQueryApplicationLinkByClientIDRequest("client_id"),
			shouldErr: false,
			expResponse: &types.QueryApplicationLinkByClientIDResponse{
				Link: types.NewApplicationLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						-1,
						1,
						types.NewOracleRequestCallData(
							"twitter",
							"7B22757365726E616D65223A22526963636172646F4D222C22676973745F6964223A223732306530303732333930613930316262383065353966643630643766646564227D",
						),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			res, err := suite.k.ApplicationLinkByClientID(sdk.WrapSDKContext(ctx), uc.req)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(uc.expResponse, res)
			}
		})
	}
}
