package keeper_test

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v5/testutil/profilestesting"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"

	"github.com/desmos-labs/desmos/v5/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_ExportGenesis() {
	chainLinkAccount := profilestesting.GetChainLinkAccount("cosmos", "cosmos")
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
				types.DefaultParams(),
				"",
				nil,
				nil,
				nil,
			),
		},
		{
			name: "non-empty state",
			store: func(ctx sdk.Context) {

				profile := profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				otherProfile := profilestesting.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")

				err := suite.k.SaveProfile(suite.ctx, profile)
				suite.Require().NoError(err)

				err = suite.k.SaveProfile(suite.ctx, otherProfile)
				suite.Require().NoError(err)

				dTagRequests := []types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				}
				for _, req := range dTagRequests {
					suite.Require().NoError(suite.k.SaveDTagTransferRequest(ctx, req))
				}

				params := types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					types.NewBioParams(sdk.NewInt(1000)),
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
					types.NewAppLinksParams(types.DefaultAppLinksValidityDuration),
				)
				suite.k.SetParams(ctx, params)
				suite.k.SetPort(ctx, "port-id")

				chainLinks := []types.ChainLink{
					chainLinkAccount.GetBech32ChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				for _, link := range chainLinks {
					suite.ak.SetAccount(ctx, profilestesting.ProfileFromAddr(link.User))
					suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
				}

				defaultExternalAddressEntries := []types.DefaultExternalAddressEntry{{
					Owner:     "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					ChainName: "cosmos",
					Target:    chainLinkAccount.Bech32Address().GetValue(),
				}}
				for _, entry := range defaultExternalAddressEntries {
					suite.k.SaveDefaultExternalAddress(ctx, entry.Owner, entry.ChainName, entry.Target)
				}

				applicationLinks := []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				}
				for _, link := range applicationLinks {
					suite.ak.SetAccount(ctx, profilestesting.ProfileFromAddr(link.User))
					suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
				}
			},
			expGenesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest("dtag-2", "sender-2", "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					types.NewDTagTransferRequest("dtag-1", "sender-1", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				},
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					types.NewBioParams(sdk.NewInt(1000)),
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
					types.NewAppLinksParams(types.DefaultAppLinksValidityDuration),
				),
				"port-id",
				[]types.ChainLink{
					chainLinkAccount.GetBech32ChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types.DefaultExternalAddressEntry{
					types.NewDefaultExternalAddressEntry(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						chainLinkAccount.Bech32Address().GetValue(),
					),
				},
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
						time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
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
		name        string
		setup       func()
		store       func(ctx sdk.Context)
		genesis     *types.GenesisState
		shouldPanic bool
		check       func(ctx sdk.Context)
	}{
		{
			name: "bind port panics",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
				nil,
			),
			setup: func() {
				suite.scopedKeeper.EXPECT().GetCapability(gomock.Any(), host.PortPath(types.IBCPortID)).Return(capabilitytypes.NewCapability(1), false)
				suite.portKeeper.EXPECT().BindPort(gomock.Any(), types.IBCPortID).Return(capabilitytypes.NewCapability(1))
				suite.scopedKeeper.EXPECT().ClaimCapability(gomock.Any(), capabilitytypes.NewCapability(1), host.PortPath(types.IBCPortID)).Return(fmt.Errorf("failed to bind port"))
			},
			shouldPanic: true,
		},
		{
			name: "empty genesis - unbounded port",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
				nil,
			),
			setup: func() {
				suite.scopedKeeper.EXPECT().GetCapability(gomock.Any(), host.PortPath(types.IBCPortID)).Return(capabilitytypes.NewCapability(1), false)
				suite.portKeeper.EXPECT().BindPort(gomock.Any(), types.IBCPortID).Return(capabilitytypes.NewCapability(1))
				suite.scopedKeeper.EXPECT().ClaimCapability(gomock.Any(), capabilitytypes.NewCapability(1), host.PortPath(types.IBCPortID)).Return(nil)
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal([]types.DTagTransferRequest(nil), suite.k.GetDTagTransferRequests(ctx))
				suite.Require().Equal(types.DefaultParams(), suite.k.GetParams(ctx))
				suite.Require().Equal(types.IBCPortID, suite.k.GetPort(ctx))
				suite.Require().Equal([]types.ApplicationLink(nil), suite.k.GetApplicationLinks(ctx))
			},
		},
		{
			name: "empty genesis - bounded port",
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				types.IBCPortID,
				nil,
				nil,
				nil,
			),
			setup: func() {
				suite.scopedKeeper.EXPECT().GetCapability(gomock.Any(), host.PortPath(types.IBCPortID)).Return(capabilitytypes.NewCapability(1), true)
			},
			check: func(ctx sdk.Context) {
				suite.Require().Equal([]types.DTagTransferRequest(nil), suite.k.GetDTagTransferRequests(ctx))
				suite.Require().Equal(types.DefaultParams(), suite.k.GetParams(ctx))
				suite.Require().Equal(types.IBCPortID, suite.k.GetPort(ctx))
				suite.Require().Equal([]types.ApplicationLink(nil), suite.k.GetApplicationLinks(ctx))
			},
		},
		{
			name: "double chain link panics",
			setup: func() {
				suite.scopedKeeper.EXPECT().GetCapability(gomock.Any(), host.PortPath("profiles-port-id")).Return(capabilitytypes.NewCapability(1), true)
			},
			genesis: types.NewGenesisState(
				nil,
				types.DefaultParams(),
				"profiles-port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(ext.GetPubKey(), profilestesting.SingleSignatureFromHex(hex.EncodeToString(ext.Sign(ext.GetAddress()))), hex.EncodeToString([]byte(ext.GetAddress().String()))),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(ext.GetPubKey(), profilestesting.SingleSignatureFromHex(hex.EncodeToString(ext.Sign(ext.GetAddress()))), hex.EncodeToString([]byte(ext.GetAddress().String()))),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				nil,
				nil,
			),
			shouldPanic: true,
		},
		{
			name: "valid genesis does not panic",
			setup: func() {
				suite.scopedKeeper.EXPECT().GetCapability(gomock.Any(), host.PortPath("profiles-port-id")).Return(capabilitytypes.NewCapability(1), true)
			},
			store: func(ctx sdk.Context) {
				profile1 := profilestesting.ProfileFromAddr("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				suite.ak.SetAccount(ctx, profile1)

				profile2 := profilestesting.ProfileFromAddr("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
				suite.ak.SetAccount(ctx, profile2)

				err := suite.k.SaveProfile(suite.ctx, profile1)
				suite.Require().NoError(err)

				err = suite.k.SaveProfile(suite.ctx, profile2)
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
				types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					types.NewBioParams(sdk.NewInt(1000)),
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
					types.NewAppLinksParams(types.DefaultAppLinksValidityDuration),
				),
				"profiles-port-id",
				[]types.ChainLink{
					types.NewChainLink(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
						types.NewProof(
							ext.GetPubKey(),
							profilestesting.SingleSignatureFromHex(
								hex.EncodeToString(ext.Sign([]byte("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"))),
							),
							hex.EncodeToString([]byte("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")),
						),
						types.NewChainConfig("cosmos"),
						time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					),
				},
				[]types.DefaultExternalAddressEntry{
					types.NewDefaultExternalAddressEntry("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", ext.GetAddress().String()),
				},
				[]types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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

				params := types.NewParams(
					types.NewNicknameParams(sdk.NewInt(100), sdk.NewInt(200)),
					types.NewDTagParams("regex", sdk.NewInt(100), sdk.NewInt(200)),
					types.NewBioParams(sdk.NewInt(1000)),
					types.NewOracleParams(
						32,
						10,
						6,
						50_000,
						200_000,
						sdk.NewCoin("band", sdk.NewInt(10)),
					),
					types.NewAppLinksParams(types.DefaultAppLinksValidityDuration),
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
							profilestesting.SingleSignatureFromHex(
								hex.EncodeToString(ext.Sign([]byte("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"))),
							),
							hex.EncodeToString([]byte("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")),
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
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
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
			if tc.setup != nil {
				tc.setup()
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			if tc.shouldPanic {
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
