package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
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
			expGenesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), "", nil),
		},
		{
			name: "non-empty state",
			store: func(ctx sdk.Context) {
				dTagRequests := []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				}
				for _, req := range dTagRequests {
					suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, req))
				}

				relationships := []types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				for _, rel := range relationships {
					suite.Require().NoError(suite.k.SaveRelationship(ctx, rel))
				}

				blocks := []types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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

				address := "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"
				suite.ak.SetAccount(ctx, suite.CreateProfileFromAddress(address))
				err := suite.k.SaveApplicationLink(ctx,
					types.NewApplicationLink(
						address,
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
				)
				suite.Require().NoError(err)
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
	profile1 := suite.CreateProfileFromAddress("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	profile2 := suite.CreateProfileFromAddress("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

	addr3, err := sdk.AccAddressFromBech32("cosmos1s3nh6tafl4amaxkke9kdejhp09lk93g9ev39r4")
	suite.Require().NoError(err)

	addr4, err := sdk.AccAddressFromBech32("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0")
	suite.Require().NoError(err)

	pubKey4, err := sdk.GetPubKeyFromBech32(
		sdk.Bech32PubKeyTypeAccPub,
		"cosmospub1addwnpepq0j8zw4t6tg3v8gh7d2d799gjhue7ewwmpg2hwr77f9kuuyzgqtrw5r6wec",
	)
	suite.Require().NoError(err)

	usecases := []struct {
		name         string
		authAccounts []authtypes.AccountI
		genesis      *types.GenesisState
		expErr       bool
		check        func(ctx sdk.Context)
	}{
		{
			name:    "empty genesis",
			genesis: types.NewGenesisState(nil, nil, nil, types.DefaultParams(), types.IBCPortID, nil),
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
			),
			expErr: true,
		},
		{
			name: "double chain link panics",
			authAccounts: []authtypes.AccountI{
				suite.CheckProfileNoError(types.NewProfile(
					"dtag-3",
					"nickname-3",
					"bio-3",
					types.NewPictures("profile-3", "cover-3"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					authtypes.NewBaseAccountWithAddress(addr4),
					[]types.ChainLink{
						types.NewChainLink(
							types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
							types.NewProof(pubKey4, "sig_hex", "addr"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
						types.NewChainLink(
							types.NewBech32Address("cosmos1xmquc944hzu6n6qtljcexkuhhz76mucxtgm5x0", "cosmos"),
							types.NewProof(pubKey4, "sig_hex", "addr"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
					},
				)),
			},
			genesis: types.NewGenesisState(
				nil,
				[]types.Relationship{},
				[]types.UserBlock{},
				types.DefaultParams(),
				"profiles-port-id",
				nil,
			),
			expErr: true,
		},
		{
			name: "valid genesis does not panic",
			authAccounts: []authtypes.AccountI{
				profile1,
				profile2,
				authtypes.NewBaseAccountWithAddress(addr3),
			},
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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
					types.NewDTagTransferRequest("dtag-1", "sender-1", "receiver-1"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "receiver-2"),
				}
				suite.Require().Equal(requests, suite.k.GetDTagTransferRequests(ctx))

				relationships := []types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				}
				suite.Require().Equal(relationships, suite.k.GetAllRelationships(ctx))

				blocks := []types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
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

				linksEntries := []types.ApplicationLink{
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
				suite.Require().Equal(linksEntries, suite.k.GetApplicationLinks(ctx))
			},
		},
	}

	for _, uc := range usecases {
		uc := uc
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			for _, acc := range uc.authAccounts {
				suite.ak.SetAccount(ctx, acc)
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
