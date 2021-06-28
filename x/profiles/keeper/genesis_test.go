package keeper_test

import (
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	var pubKey, err = sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d",
	)
	suite.Require().NoError(err)

	usecases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "empty state",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			expGenesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), "", nil, nil),
		},
		{
			name: "non-empty state",
			store: func(ctx sdk.Context) {

				profile := suite.CreateProfileFromAddress(suite.testData.user)
				otherProfile := suite.CreateProfileFromAddress(suite.testData.otherUser)

				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)

				err = suite.k.StoreProfile(suite.ctx, otherProfile)
				suite.Require().NoError(err)

				dTagRequests := []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", suite.testData.otherUser),
					types.NewDTagTransferRequest("dtag-1", "sender-1", suite.testData.user),
				}
				for _, req := range dTagRequests {
					suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, req))
				}

				relationships := []types.Relationship{
					types.NewRelationship(
						suite.testData.otherUser,
						suite.testData.user,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						suite.testData.user,
						suite.testData.otherUser,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				for _, rel := range relationships {
					suite.Require().NoError(suite.k.SaveRelationship(ctx, rel))
				}

				blocks := []types.UserBlock{
					types.NewUserBlock(
						suite.testData.otherUser,
						suite.testData.user,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						suite.testData.user,
						suite.testData.otherUser,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				for _, block := range blocks {
					suite.Require().NoError(suite.k.SaveUserBlock(ctx, block))
				}

				params := types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				)
				suite.k.SetParams(ctx, params)
				suite.k.SetPort(ctx, "port-id")

				chainLinks := []types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				for _, link := range chainLinks {
					suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(link.User))
					suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
				}

				applicationLinks := []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				for _, link := range applicationLinks {
					suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(link.User))
					suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
				}
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", suite.testData.otherUser),
					types.NewDTagTransferRequest("dtag-1", "sender-1", suite.testData.user),
				},
				[]types.Relationship{
					types.NewRelationship(
						suite.testData.otherUser,
						suite.testData.user,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						suite.testData.user,
						suite.testData.otherUser,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						suite.testData.otherUser,
						suite.testData.user,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						suite.testData.user,
						suite.testData.otherUser,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
				"port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
						types.NewProof(
							pubKey,
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			),
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			exported := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(uc.expGenesis, exported)
		})
	}
}

func (suite *KeeperTestSuite) Test_InitGenesis() {
	ext := suite.GetRandomProfile()

	usecases := []struct {
		name    string
		store   func(ctx sdk.Context)
		genesis *types.GenesisState
		expErr  bool
		check   func(ctx sdk.Context)
	}{
		{
			name:    "empty genesis",
			genesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), types.IBCPortID, nil, nil),
			check: func(ctx sdk.Context) {
				suite.Require().Equal([]types.DTagTransferRequest(nil), suite.k.GetDTagTransferRequests(ctx))
				suite.Require().Equal([]types.Relationship(nil), suite.k.GetAllRelationships(ctx))
				suite.Require().Equal([]types.UserBlock(nil), suite.k.GetAllUsersBlocks(ctx))
				suite.Require().Equal(types.DefaultParams(), suite.k.GetParams(ctx))
				suite.Require().Equal(types.IBCPortID, suite.k.GetPort(ctx))
				suite.Require().Equal([]types.ApplicationLink(nil), suite.k.GetApplicationLinks(ctx))
			},
		},
		{
			name: "double relationships panics",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{
					types.NewRelationship("creator", "recipient", "subspace"),
					types.NewRelationship("creator", "recipient", "subspace"),
				},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
				nil,
			),
			expErr: true,
		},
		{
			name: "double user block panics",
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{
					types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
					types.NewUserBlock("blocker", "blocked", "reason", "subspace"),
				},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
				nil,
			),
			expErr: true,
		},
		{
			name: "double chain link panics",
			store: func(ctx sdk.Context) {

			},
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(ext.GetPubKey(), hex.EncodeToString(ext.Sign(ext.GetAddress())), ext.GetAddress().String()),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(ext.GetPubKey(), hex.EncodeToString(ext.Sign(ext.GetAddress())), ext.GetAddress().String()),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
			),
			expErr: true,
		},
		{
			name: "valid genesis does not panic",
			store: func(ctx sdk.Context) {
				profile1 := suite.CreateProfileFromAddress("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile1)

				profile2 := suite.CreateProfileFromAddress("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
				suite.ak.SetAccount(ctx, profile2)

				err := suite.k.StoreProfile(suite.ctx, profile1)
				suite.Require().NoError(err)

				err = suite.k.StoreProfile(suite.ctx, profile2)
				suite.Require().NoError(err)

				addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
				suite.Require().NoError(err)
				suite.ak.SetAccount(ctx, authtypes.NewBaseAccountWithAddress(addr3))
			},
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", suite.testData.user),
					types.NewDTagTransferRequest("dtag-2", "sender-2", suite.testData.otherUser),
				},
				[]types.Relationship{
					types.NewRelationship(
						suite.testData.user,
						suite.testData.otherUser,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						suite.testData.otherUser,
						suite.testData.user,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						suite.testData.user,
						suite.testData.otherUser,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						suite.testData.otherUser,
						suite.testData.user,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				),
				"profiles-port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(
							ext.GetPubKey(),
							hex.EncodeToString(ext.Sign([]byte(ext.GetAddress().String()))),
							ext.GetAddress().String(),
						),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			),
			check: func(ctx sdk.Context) {
				requests := []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", suite.testData.otherUser),
					types.NewDTagTransferRequest("dtag-1", "sender-1", suite.testData.user),
				}
				suite.Require().Equal(requests, suite.k.GetDTagTransferRequests(ctx))

				relationships := []types.Relationship{
					types.NewRelationship(
						suite.testData.otherUser,
						suite.testData.user,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						suite.testData.user,
						suite.testData.otherUser,
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				suite.Require().Equal(relationships, suite.k.GetAllRelationships(ctx))

				blocks := []types.UserBlock{
					types.NewUserBlock(
						suite.testData.otherUser,
						suite.testData.user,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						suite.testData.user,
						suite.testData.otherUser,
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				suite.Require().Equal(blocks, suite.k.GetAllUsersBlocks(ctx))

				params := types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
				)
				suite.Require().Equal(params, suite.k.GetParams(ctx))

				portID := "profiles-port-id"
				suite.Require().Equal(portID, suite.k.GetPort(ctx))

				chainLinks := []types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(
							ext.GetPubKey(),
							hex.EncodeToString(ext.Sign([]byte(ext.GetAddress().String()))),
							ext.GetAddress().String(),
						),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				}
				suite.Require().Equal(chainLinks, suite.k.GetChainLinks(ctx))

				applicationLinks := []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							-1,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				suite.Require().Equal(applicationLinks, suite.k.GetApplicationLinks(ctx))
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

			if uc.expErr {
				suite.Require().Panics(func() { suite.k.InitGenesis(ctx, *uc.genesis) })
			} else {
				suite.Require().NotPanics(func() { suite.k.InitGenesis(ctx, *uc.genesis) })
				if uc.check != nil {
					uc.check(ctx)
				}
			}
		})
	}
}
