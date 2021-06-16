package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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

func (suite *KeeperTestSuite) Test_ProfileByChainLink() {
	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)
	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)

	link := types.NewChainLink(
		types.NewBech32Address(srcAddr, "cosmos"),
		types.NewProof(srcPubKey, srcSigHex, srcAddr),
		types.NewChainConfig("cosmos"),
		suite.testData.profile.CreationDate,
	)

	usecases := []struct {
		name        string
		store       func()
		req         *types.QueryProfileByChainLinkRequest
		shouldErr   bool
		expResponse *types.QueryProfileByChainLinkResponse
	}{
		{
			name:  "empty request returns error",
			store: func() {},
			req: &types.QueryProfileByChainLinkRequest{
				ChainName:     "",
				TargetAddress: "",
			},
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name: "invalid linked address returns error",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey(link.ChainConfig.Name, srcAddr)
				store.Set(key, []byte("invalid"))
			},
			req: &types.QueryProfileByChainLinkRequest{
				ChainName:     "cosmos",
				TargetAddress: srcAddr,
			},
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name: "destination has no profile returns error",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey(link.ChainConfig.Name, srcAddr)
				acc, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)
				store.Set(key, acc)
			},
			req: &types.QueryProfileByChainLinkRequest{
				ChainName:     "cosmos",
				TargetAddress: srcAddr,
			},
			shouldErr:   true,
			expResponse: nil,
		},
		{
			name: "valid request",
			store: func() {
				err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
				suite.Require().NoError(err)

				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey(link.ChainConfig.Name, srcAddr)
				store.Set(key, []byte(suite.testData.profile.GetAddress()))
			},
			req: &types.QueryProfileByChainLinkRequest{
				ChainName:     "cosmos",
				TargetAddress: srcAddr,
			},
			shouldErr: false,
			expResponse: &types.QueryProfileByChainLinkResponse{
				Profile: suite.codeToAny(suite.testData.profile),
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			suite.SetupTest()
			uc.store()

			res, err := suite.k.ProfileByChainLink(sdk.WrapSDKContext(suite.ctx), uc.req)
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
