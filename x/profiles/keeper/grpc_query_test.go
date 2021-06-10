package keeper_test

import (
	"encoding/hex"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

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
