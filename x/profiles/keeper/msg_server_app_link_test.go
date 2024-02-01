package keeper_test

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/cometbft/cometbft/libs/log"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
	"github.com/golang/mock/gomock"

	"github.com/desmos-labs/desmos/v6/pkg/obi"
	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"
	oracletypes "github.com/desmos-labs/desmos/v6/x/oracle/types"
	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_LinkApplication() {
	blockTime := time.Date(2023, 9, 6, 0, 0, 0, 0, time.UTC)
	params := types.DefaultParams()

	testCases := []struct {
		name      string
		setup     func()
		store     func(ctx sdk.Context)
		msg       *types.MsgLinkApplication
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid user address returns error",
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"invalid",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "ongoing application link exists returns error",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					types.NewData("twitter", "twitteruser"),
					types.AppLinkStateVerificationStarted,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 3, 1, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "ibc channel does not exist returns error",
			setup: func() {
				suite.channelKeeper.EXPECT().
					GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
					Return(channeltypes.Channel{}, false)
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "ibc capability does not exist returns error",
			setup: func() {
				suite.channelKeeper.
					EXPECT().
					GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
					Return(channeltypes.Channel{
						Counterparty: channeltypes.Counterparty{
							PortId:    "oracle",
							ChannelId: "channel-0",
						},
					}, true)

				suite.scopedKeeper.EXPECT().
					GetCapability(gomock.Any(), host.ChannelCapabilityPath(types.IBCPortID, "channel-0")).
					Return(nil, false)
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "failed to send ibc packet returns error",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, params)
			},
			setup: func() {
				suite.channelKeeper.EXPECT().
					GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
					Return(channeltypes.Channel{
						Counterparty: channeltypes.Counterparty{
							PortId:    "oracle",
							ChannelId: "channel-0",
						},
					}, true)

				suite.scopedKeeper.EXPECT().
					GetCapability(gomock.Any(), host.ChannelCapabilityPath(types.IBCPortID, "channel-0")).
					Return(&capabilitytypes.Capability{Index: 1}, true)

				suite.channelKeeper.EXPECT().
					SendPacket(
						gomock.Any(),
						&capabilitytypes.Capability{Index: 1},
						types.IBCPortID,
						"channel-0",
						clienttypes.NewHeight(0, 1000),
						uint64(0),
						oracletypes.NewOracleRequestPacketData(
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
							oracletypes.OracleScriptID(params.Oracle.ScriptID),
							obi.MustEncode(keeper.OracleScriptCallData{"twitter", "call_data"}),
							params.Oracle.AskCount,
							params.Oracle.MinCount,
							params.Oracle.FeeAmount,
							params.Oracle.PrepareGas,
							params.Oracle.ExecuteGas,
						).GetBytes(),
					).
					Return(uint64(0), fmt.Errorf("invalid"))
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			shouldErr: true,
		},
		{
			name: "link application starts properly",
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, params)

				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.ak.SetAccount(ctx, profile)
			},
			setup: func() {
				suite.channelKeeper.EXPECT().
					GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
					Return(channeltypes.Channel{
						Counterparty: channeltypes.Counterparty{
							PortId:    "oracle",
							ChannelId: "channel-0",
						},
					}, true)

				suite.scopedKeeper.EXPECT().
					GetCapability(gomock.Any(), host.ChannelCapabilityPath(types.IBCPortID, "channel-0")).
					Return(&capabilitytypes.Capability{Index: 1}, true)

				suite.channelKeeper.EXPECT().
					SendPacket(
						gomock.Any(),
						&capabilitytypes.Capability{Index: 1},
						types.IBCPortID,
						"channel-0",
						clienttypes.NewHeight(0, 1000),
						uint64(0),
						oracletypes.NewOracleRequestPacketData(
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
							oracletypes.OracleScriptID(params.Oracle.ScriptID),
							obi.MustEncode(keeper.OracleScriptCallData{"twitter", "call_data"}),
							params.Oracle.AskCount,
							params.Oracle.MinCount,
							params.Oracle.FeeAmount,
							params.Oracle.PrepareGas,
							params.Oracle.ExecuteGas,
						).GetBytes(),
					).
					Return(uint64(0), nil)
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeSavedApplicationLink,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
				),
				sdk.NewEvent(
					types.EventTypeCreatedApplicationLink,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
					sdk.NewAttribute(types.AttributeKeyApplicationLinkCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				link, found, err := suite.k.GetApplicationLink(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal(
					types.NewApplicationLink(
						"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							0,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
						),
						nil,
						time.Date(2023, 9, 6, 00, 00, 00, 000, time.UTC),
						time.Date(2024, 9, 5, 00, 00, 00, 000, time.UTC),
					),
					link,
				)
			},
		},
		{
			name: "link application starts properly - replace non ongoing application link",
			setup: func() {
				suite.channelKeeper.EXPECT().
					GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
					Return(channeltypes.Channel{
						Counterparty: channeltypes.Counterparty{
							PortId:    "oracle",
							ChannelId: "channel-0",
						},
					}, true)

				suite.scopedKeeper.EXPECT().
					GetCapability(gomock.Any(), host.ChannelCapabilityPath(types.IBCPortID, "channel-0")).
					Return(&capabilitytypes.Capability{Index: 1}, true)

				suite.channelKeeper.EXPECT().
					SendPacket(
						gomock.Any(),
						&capabilitytypes.Capability{Index: 1},
						types.IBCPortID,
						"channel-0",
						clienttypes.NewHeight(0, 1000),
						uint64(0),
						oracletypes.NewOracleRequestPacketData(
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
							oracletypes.OracleScriptID(params.Oracle.ScriptID),
							obi.MustEncode(keeper.OracleScriptCallData{"twitter", "call_data"}),
							params.Oracle.AskCount,
							params.Oracle.MinCount,
							params.Oracle.FeeAmount,
							params.Oracle.PrepareGas,
							params.Oracle.ExecuteGas,
						).GetBytes(),
					).
					Return(uint64(0), nil)
			},
			store: func(ctx sdk.Context) {
				suite.k.SetParams(ctx, params)

				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					types.NewData("twitter", "twitteruser"),
					types.AppLinkStateVerificationTimedOut,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 3, 1, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
			},
			msg: types.NewMsgLinkApplication(
				types.NewData("twitter", "twitteruser"),
				"call_data",
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				types.IBCPortID,
				"channel-0",
				clienttypes.NewHeight(0, 1000),
				0,
			),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeSavedApplicationLink,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
				),
				sdk.NewEvent(
					types.EventTypeCreatedApplicationLink,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
					sdk.NewAttribute(types.AttributeKeyApplicationLinkCreationTime, blockTime.Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				link, found, err := suite.k.GetApplicationLink(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser")
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Equal(
					types.NewApplicationLink(
						"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
						types.NewData("twitter", "twitteruser"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							0,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
						),
						nil,
						time.Date(2023, 9, 6, 00, 00, 00, 000, time.UTC),
						time.Date(2024, 9, 5, 00, 00, 00, 000, time.UTC),
					),
					link,
				)
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

			ctx = ctx.WithBlockTime(blockTime).WithEventManager(sdk.NewEventManager())
			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.LinkApplication(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_LinkApplication_Logger() {
	// Setup profiles and keepers
	params := types.DefaultParams()
	suite.k.SetParams(suite.ctx, params)
	profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
	suite.ak.SetAccount(suite.ctx, profile)

	suite.channelKeeper.EXPECT().
		GetChannel(gomock.Any(), types.IBCPortID, "channel-0").
		Return(channeltypes.Channel{
			Counterparty: channeltypes.Counterparty{
				PortId:    "oracle",
				ChannelId: "channel-0",
			},
		}, true)

	suite.scopedKeeper.EXPECT().
		GetCapability(gomock.Any(), host.ChannelCapabilityPath(types.IBCPortID, "channel-0")).
		Return(&capabilitytypes.Capability{Index: 1}, true)

	suite.channelKeeper.EXPECT().
		SendPacket(
			gomock.Any(),
			&capabilitytypes.Capability{Index: 1},
			types.IBCPortID,
			"channel-0",
			clienttypes.NewHeight(0, 1000),
			uint64(0),
			oracletypes.NewOracleRequestPacketData(
				"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773-twitter-twitteruser",
				oracletypes.OracleScriptID(params.Oracle.ScriptID),
				obi.MustEncode(keeper.OracleScriptCallData{"twitter", "call_data"}),
				params.Oracle.AskCount,
				params.Oracle.MinCount,
				params.Oracle.FeeAmount,
				params.Oracle.PrepareGas,
				params.Oracle.ExecuteGas,
			).GetBytes(),
		).
		Return(uint64(0), nil)

	// Setup Logger
	var buf bytes.Buffer
	ctx, _ := suite.ctx.CacheContext()
	ctx = ctx.WithLogger(log.NewTMLogger(&buf))

	// Execute
	server := keeper.NewMsgServerImpl(suite.k)
	_, err := server.LinkApplication(ctx, types.NewMsgLinkApplication(
		types.NewData("twitter", "twitteruser"),
		"call_data",
		"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
		types.IBCPortID,
		"channel-0",
		clienttypes.NewHeight(0, 1000),
		0,
	))
	suite.Require().NoError(err)

	// Check logs
	msg := strings.TrimSpace(buf.String())
	suite.Require().Contains(msg, "Application link created")
	suite.Require().Contains(msg, fmt.Sprintf("application=%s", "twitter"))
	suite.Require().Contains(msg, fmt.Sprintf("username=%s", "twitteruser"))
	suite.Require().Contains(msg, fmt.Sprintf("account=%s", "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"))
}

func (suite *KeeperTestSuite) TestMsgServer_UnlinkApplication() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgUnlinkApplication
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name:      "link not found returns error",
			msg:       types.NewMsgUnlinkApplication("twitter", "twitteruser", "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			shouldErr: true,
		},
		{
			name: "link deleted properly",
			store: func(ctx sdk.Context) {
				profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
				suite.ak.SetAccount(ctx, profile)

				link := types.NewApplicationLink(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					types.NewData("twitter", "twitteruser"),
					types.AppLinkStateVerificationTimedOut,
					types.NewOracleRequest(
						1,
						1,
						types.NewOracleRequestCallData("twitter", "tweet-123456789"),
						"client_id",
					),
					nil,
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 3, 1, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveApplicationLink(ctx, link))
			},
			msg: types.NewMsgUnlinkApplication("twitter", "twitteruser", "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeDeletedApplicationLink,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
					sdk.NewAttribute(types.AttributeKeyApplicationLinkExpirationTime, time.Date(2020, 3, 1, 00, 00, 00, 000, time.UTC).Format(time.RFC3339)),
				),
			},
			check: func(ctx sdk.Context) {
				suite.Require().False(
					suite.k.HasApplicationLink(ctx, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser"),
				)
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

			ctx = ctx.WithEventManager(sdk.NewEventManager())
			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.UnlinkApplication(sdk.WrapSDKContext(ctx), tc.msg)

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestMsgServer_UnlinkApplication_Logger() {
	// Setup profiles and link
	profile := profilestesting.ProfileFromAddr("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773")
	suite.ak.SetAccount(suite.ctx, profile)

	suite.Require().NoError(suite.k.SaveApplicationLink(
		suite.ctx,
		types.NewApplicationLink(
			"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			types.NewData("twitter", "twitteruser"),
			types.AppLinkStateVerificationTimedOut,
			types.NewOracleRequest(
				1,
				1,
				types.NewOracleRequestCallData("twitter", "tweet-123456789"),
				"client_id",
			),
			nil,
			time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
			time.Date(2020, 3, 1, 00, 00, 00, 000, time.UTC),
		)))

	// Setup Logger
	var buf bytes.Buffer
	ctx, _ := suite.ctx.CacheContext()
	ctx = ctx.WithLogger(log.NewTMLogger(&buf))

	// Execute
	server := keeper.NewMsgServerImpl(suite.k)
	_, err := server.UnlinkApplication(ctx, types.NewMsgUnlinkApplication(
		"twitter",
		"twitteruser",
		"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
	))
	suite.Require().NoError(err)

	// Check logs
	msg := strings.TrimSpace(buf.String())
	suite.Require().Contains(msg, "Application link removed")
	suite.Require().Contains(msg, fmt.Sprintf("application=%s", "twitter"))
	suite.Require().Contains(msg, fmt.Sprintf("username=%s", "twitteruser"))
	suite.Require().Contains(msg, fmt.Sprintf("account=%s", "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"))
}
