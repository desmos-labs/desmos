package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/desmos-labs/desmos/testutil"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	testCases := []struct {
		name       string
		store      func(ctx sdk.Context)
		expGenesis *types.GenesisState
	}{
		{
			name: "empty state",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, types.DefaultParams())
			},
			expGenesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				types.DefaultParams(),
				"",
				nil,
				nil,
			),
		},
		{
			name: "non-empty state",
			store: func(ctx sdk.Context) {

				profile := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				otherProfile := testutil.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

				err := suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)

				err = suite.k.StoreProfile(suite.ctx, otherProfile)
				suite.Require().NoError(err)

				dTagRequests := []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
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
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						"desmos-ibc-profiles",
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
				)
				suite.k.SetParams(ctx, params)
				suite.k.SetPort(ctx, "port-id")

				chainLinks := []types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1nc54z3kzyal57w6wcf5khmwrxx5rafnwvu0m5z", "cosmos"),
						types.NewProof(
							testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
							"909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b",
							"text",
						),
						types.NewChainConfig("cosmos"),
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				for _, link := range chainLinks {
					suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(link.User))
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
					suite.ak.SetAccount(ctx, testutil.ProfileFromAddr(link.User))
					suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
				}
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
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
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						"desmos-ibc-profiles",
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
				),
				"port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address("cosmos1nc54z3kzyal57w6wcf5khmwrxx5rafnwvu0m5z", "cosmos"),
						types.NewProof(
							testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
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

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			exported := suite.k.ExportGenesis(ctx)
			suite.Require().Equal(tc.expGenesis, exported)
		})
	}
}

func (suite *KeeperTestSuite) Test_InitGenesis() {
	ext := suite.GetRandomProfile()

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		genesis   *types.GenesisState
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "empty genesis",
			genesis: types.NewGenesisState(
				nil,
				nil,
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
			),
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
			shouldErr: true,
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
			shouldErr: true,
		},
		{
			name: "double chain link panics",
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
			shouldErr: true,
		},
		{
			name: "valid genesis does not panic",
			store: func(ctx sdk.Context) {
				profile1 := testutil.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile1)

				profile2 := testutil.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
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
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
				},
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					sdk.NewInt(1000),
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						"desmos-ibc-profiles",
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
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
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
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
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						"desmos-ibc-profiles",
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
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

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			if tc.shouldErr {
				suite.Require().Panics(func() { suite.k.InitGenesis(ctx, *tc.genesis) })
			} else {
				suite.Require().NotPanics(func() { suite.k.InitGenesis(ctx, *tc.genesis) })
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
