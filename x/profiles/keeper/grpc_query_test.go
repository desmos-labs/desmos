package keeper_test

import (
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
				suite.testData.profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.DTag),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile),
			},
		},
		{
			name: "found profile - using address",
			storedProfiles: []*types.Profile{
				suite.testData.profile,
			},
			req:       types.NewQueryProfileRequest(suite.testData.profile.GetAddress().String()),
			shouldErr: false,
			expResponse: &types.QueryProfileResponse{
				Profile: suite.codeToAny(suite.testData.profile),
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

func (suite *KeeperTestSuite) Test_DTagTransfers() {
	usecases := []struct {
		name           string
		storedRequests []types.DTagTransferRequest
		req            *types.QueryDTagTransfersRequest
		shouldErr      bool
		expResponse    *types.QueryDTagTransfersResponse
	}{
		{
			name:      "invalid user",
			req:       types.NewQueryDTagTransfersRequest("invalid-address"),
			shouldErr: true,
		},
		{
			name: "valid request",
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
			req:       types.NewQueryDTagTransfersRequest("cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa"),
			shouldErr: false,
			expResponse: &types.QueryDTagTransfersResponse{
				Requests: []types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos19mj6dkd85m84gxvf8d929w572z5h9q0u8d8wpa",
					),
				},
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

			res, err := suite.k.DTagTransfers(sdk.WrapSDKContext(suite.ctx), uc.req)

			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().NotNil(res)

				suite.Require().Equal(uc.expResponse, res)
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
