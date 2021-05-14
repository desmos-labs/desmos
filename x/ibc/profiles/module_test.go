package ibcprofiles_test

import (
	"encoding/hex"
	"math"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/desmos-labs/desmos/testutil/ibctesting"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/suite"
)

type IBCProfilesTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA      *ibctesting.TestChain
	chainB      *ibctesting.TestChain
	queryClient types.QueryClient
}

func (suite *IBCProfilesTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))

	queryHelper := baseapp.NewQueryServerTestHelper(suite.chainA.GetContext(), suite.chainA.App.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, suite.chainA.App.IBCProfilesKeeper)
	suite.queryClient = types.NewQueryClient(queryHelper)
}

func TestIBCProfilesTestSuite(t *testing.T) {
	suite.Run(t, new(IBCProfilesTestSuite))
}

func (suite *IBCProfilesTestSuite) TestOnChanOpenInit() {
	var (
		channel     *channeltypes.Channel
		testChannel ibctesting.TestChannel
		connA       *ibctesting.TestConnection
		chanCap     *capabilitytypes.Capability
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"success", func() {}, true,
		},
		{
			"max channels reached", func() {
				testChannel.ID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				testChannel = suite.chainA.NextTestChannel(connA, ibctesting.MockPort)
			}, false,
		},
		{
			"invalid version", func() {
				channel.Version = "version"
			}, false,
		},
		{
			"capability already claimed", func() {
				err := suite.chainA.App.ScopedIBCProfilesKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(testChannel.PortID, testChannel.ID))
				suite.Require().NoError(err)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			_, _, connA, _ = suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
			testChannel = suite.chainA.NextTestChannel(connA, ibctesting.IBCProfilesPort)
			counterparty := channeltypes.NewCounterparty(testChannel.PortID, testChannel.ID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{connA.ID},
				Version:        types.Version,
			}

			module, _, err := suite.chainA.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App.ScopedIBCKeeper.NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.IBCProfilesPort, testChannel.ID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenInit(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				testChannel.PortID, testChannel.ID, chanCap, channel.Counterparty, channel.GetVersion(),
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *IBCProfilesTestSuite) TestOnChanOpenTry() {
	var (
		channel             *channeltypes.Channel
		testChannel         ibctesting.TestChannel
		connA               *ibctesting.TestConnection
		chanCap             *capabilitytypes.Capability
		counterpartyVersion string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"Success", func() {}, true,
		},
		{
			"Max channels reached", func() {
				testChannel.ID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"Capability already claimed in INIT should pass", func() {
				err := suite.chainA.App.ScopedIBCProfilesKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(testChannel.PortID, testChannel.ID))
				suite.Require().NoError(err)
			}, true,
		},
		{
			"Invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"Invalid port ID", func() {
				testChannel = suite.chainA.NextTestChannel(connA, ibctesting.MockPort)
			}, false,
		},
		{
			"Invalid version", func() {
				channel.Version = "version"
			}, false,
		},
		{
			"Invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			_, _, connA, _ = suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
			testChannel = suite.chainA.NextTestChannel(connA, ibctesting.IBCProfilesPort)
			counterparty := channeltypes.NewCounterparty(testChannel.PortID, testChannel.ID)
			channel = &channeltypes.Channel{
				State:          channeltypes.TRYOPEN,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{connA.ID},
				Version:        types.Version,
			}
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App.ScopedIBCKeeper.NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.IBCProfilesPort, testChannel.ID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenTry(suite.chainA.GetContext(), channel.Ordering, channel.GetConnectionHops(),
				testChannel.PortID, testChannel.ID, chanCap, channel.Counterparty, channel.GetVersion(), counterpartyVersion,
			)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

func (suite *IBCProfilesTestSuite) TestOnChanOpenAck() {
	var (
		testChannel         ibctesting.TestChannel
		connA               *ibctesting.TestConnection
		counterpartyVersion string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			"Success", func() {}, true,
		},
		{
			"Invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			_, _, connA, _ = suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
			testChannel = suite.chainA.NextTestChannel(connA, ibctesting.IBCProfilesPort)
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenAck(suite.chainA.GetContext(), testChannel.PortID, testChannel.ID, counterpartyVersion)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *IBCProfilesTestSuite) TestOnRecvPacket() {
	var packet channeltypes.Packet
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			name: "Success IBCAccountLink packet",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr := suite.chainA.Account.GetAddress().String()
				pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := srcAddr
				packetProof := []byte(srcAddr + "-" + destAddr)
				sig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex := hex.EncodeToString(sig)

				packetData := types.NewIBCAccountLinkPacketData(
					"cosmos",
					srcAddr,
					pubKeyHex,
					sigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: true,
		},
		{
			name: "Success IBCAccountConnection packet",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

				srcAddr := suite.chainA.Account.GetAddress().String()
				srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr + "-" + destAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex := hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex := hex.EncodeToString(dstSig)

				// send link from chainA to chainB
				packetData := types.NewIBCAccountConnectionPacketData(
					"cosmos",
					srcAddr,
					srcPubKeyHex,
					destAddr,
					srcSigHex,
					dstSigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: true,
		},
		{
			name: "Invalid packet struct",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				bz := []byte{}
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: false,
		},
		{
			name: "Invalid packet type",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				var modulePacket types.IBCProfilesPacketData
				modulePacket.Packet = &types.IBCProfilesPacketData_NoData{}
				bz, err := sdk.SortJSON(types.ProtoCdc.MustMarshalJSON(&modulePacket))
				suite.Require().NoError(err)
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			tc.malleate()
			module, _, err := suite.chainB.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainB.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainB.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			_, _, err = cbs.OnRecvPacket(suite.chainB.GetContext(), packet)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *IBCProfilesTestSuite) TestOnAcknowledgement() {
	var packet channeltypes.Packet
	var ack channeltypes.Acknowledgement
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			name: "Success IBCAccountLink ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr := suite.chainA.Account.GetAddress().String()
				pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := srcAddr

				packetProof := []byte(srcAddr + "-" + destAddr)
				sig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex := hex.EncodeToString(sig)

				packetData := types.NewIBCAccountLinkPacketData(
					"cosmos",
					srcAddr,
					pubKeyHex,
					sigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

				ackData := types.IBCAccountLinkPacketAck{SourceAddress: srcAddr}
				bz, err = ackData.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			expPass: true,
		},
		{
			name: "Invalid IBCAccountLink ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr := suite.chainA.Account.GetAddress().String()
				pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := srcAddr

				packetProof := []byte(srcAddr + "-" + destAddr)
				sig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex := hex.EncodeToString(sig)

				packetData := types.NewIBCAccountLinkPacketData(
					"cosmos",
					srcAddr,
					pubKeyHex,
					sigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

				ackData := types.IBCAccountLinkPacketAck{SourceAddress: srcAddr}
				bz, err = ackData.Marshal()
				ack = channeltypes.Acknowledgement{}
			},
			expPass: false,
		},
		{
			name: "Success IBCAccountConnection ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

				srcAddr := suite.chainA.Account.GetAddress().String()
				srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr + "-" + destAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex := hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex := hex.EncodeToString(dstSig)

				// send link from chainA to chainB
				packetData := types.NewIBCAccountConnectionPacketData(
					"cosmos",
					srcAddr,
					srcPubKeyHex,
					destAddr,
					srcSigHex,
					dstSigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

				ackData := types.IBCAccountConnectionPacketAck{SourceAddress: srcAddr}
				bz, err = ackData.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			expPass: true,
		},
		{
			name: "Invalid IBCAccountLink ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

				srcAddr := suite.chainA.Account.GetAddress().String()
				srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr + "-" + destAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex := hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex := hex.EncodeToString(dstSig)

				// send link from chainA to chainB
				packetData := types.NewIBCAccountConnectionPacketData(
					"cosmos",
					srcAddr,
					srcPubKeyHex,
					destAddr,
					srcSigHex,
					dstSigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

				ackData := types.IBCAccountConnectionPacketAck{SourceAddress: srcAddr}
				bz, err = ackData.Marshal()
				ack = channeltypes.Acknowledgement{}
			},
			expPass: false,
		},
		{
			name: "Invalid packet struct",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				bz := []byte{}
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
				ack = channeltypes.NewErrorAcknowledgement("failed")
			},
			expPass: false,
		},
		{
			name: "Invalid packet type",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				var modulePacket types.IBCProfilesPacketData
				modulePacket.Packet = &types.IBCProfilesPacketData_NoData{}
				bz, err := sdk.SortJSON(types.ProtoCdc.MustMarshalJSON(&modulePacket))
				suite.Require().NoError(err)
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
				ack = channeltypes.NewErrorAcknowledgement("failed")
			},
			expPass: false,
		},
		{
			name: "Error ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr := suite.chainA.Account.GetAddress().String()
				srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr + "-" + destAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex := hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex := hex.EncodeToString(dstSig)

				// send link from chainA to chainB
				packetData := types.NewIBCAccountConnectionPacketData(
					"cosmos",
					srcAddr,
					srcPubKeyHex,
					destAddr,
					srcSigHex,
					dstSigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
				ack = channeltypes.NewErrorAcknowledgement("failed")
			},
			expPass: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			tc.malleate()
			module, _, err := suite.chainB.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainB.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainB.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			bz, err := sdk.SortJSON(types.ProtoCdc.MustMarshalJSON(&ack))
			_, err = cbs.OnAcknowledgementPacket(suite.chainB.GetContext(), packet, bz)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *IBCProfilesTestSuite) TestOnTimeoutPacket() {
	var packet channeltypes.Packet
	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{

		{
			name: "Success IBCAccountLink timout",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr := suite.chainA.Account.GetAddress().String()
				pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := srcAddr

				packetProof := []byte(srcAddr + "-" + destAddr)
				sig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex := hex.EncodeToString(sig)

				packetData := types.NewIBCAccountLinkPacketData(
					"cosmos",
					srcAddr,
					pubKeyHex,
					sigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: true,
		},
		{
			name: "Success IBCAccountConnection timout",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

				srcAddr := suite.chainA.Account.GetAddress().String()
				srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr := suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr + "-" + destAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex := hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex := hex.EncodeToString(dstSig)

				// send link from chainA to chainB
				packetData := types.NewIBCAccountConnectionPacketData(
					"cosmos",
					srcAddr,
					srcPubKeyHex,
					destAddr,
					srcSigHex,
					dstSigHex,
				)
				bz, err := packetData.GetBytes()
				suite.Require().NoError(err)

				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: true,
		},
		{
			name: "Invalid packet struct",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				bz := []byte{}
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: false,
		},
		{
			name: "Invalid packet type",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				var modulePacket types.IBCProfilesPacketData
				modulePacket.Packet = &types.IBCProfilesPacketData_NoData{}
				bz, err := sdk.SortJSON(types.ProtoCdc.MustMarshalJSON(&modulePacket))
				suite.Require().NoError(err)
				packet = channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			},
			expPass: false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			tc.malleate()
			module, _, err := suite.chainB.App.IBCKeeper.PortKeeper.LookupModuleByPort(suite.chainB.GetContext(), ibctesting.IBCProfilesPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainB.App.IBCKeeper.Router.GetRoute(module)
			suite.Require().True(ok)

			_, err = cbs.OnTimeoutPacket(suite.chainB.GetContext(), packet)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
