package keeper_test

import (
	"encoding/hex"

	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/desmos-labs/desmos/testing"
	"github.com/desmos-labs/desmos/x/links/types"
)

// func (suite *KeeperTestSuite) TestIBCAccountConnectionPacket() {

// 	suite.Run("packet transformation test", func() {
// 		suite.SetupIBCTest()

// 		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
// 		channelA, channelB := suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

// 		srcAddress := suite.chainA.Account.GetAddress().String()
// 		srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
// 		dstAddress := suite.chainB.Account.GetAddress().String()
// 		link := types.NewLink(srcAddress, dstAddress)
// 		linkBz, _ := link.Marshal()
// 		srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
// 		srcSigHex := hex.EncodeToString(srcSig)
// 		dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
// 		dstSigHex := hex.EncodeToString(dstSig)

// 		// send link from chainA to chainB
// 		packetData := types.NewIBCAccountConnectionPacketData(
// 			"cosmos",
// 			srcAddress,
// 			srcPubKeyHex,
// 			dstAddress,
// 			srcSigHex,
// 			dstSigHex,
// 		)

// 		msg := types.NewMsgCreateIBCAccountConnection(channelA.PortID, channelA.ID, 0, srcAddress, srcPubKeyHex, dstAddress, srcSigHex, dstSigHex)
// 		err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
// 		suite.Require().NoError(err) // message committed

// 		bz, err := packetData.GetBytes()
// 		suite.Require().NoError(err)

// 		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
// 		packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
// 		proof, proofHeight := suite.chainA.QueryProof(packetKey)

// 		recvMsg := channeltypes.NewMsgRecvPacket(packet, proof, proofHeight, suite.chainB.Account.GetAddress())
// 		err = suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, recvMsg)
// 		suite.Require().NoError(err) // message committed
// 	})
// }

func (suite *KeeperTestSuite) TestTransmitIBCAccountConnectionPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "successful create link from source chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, _ = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

			},
			expPass: true,
		},
		{
			name: "source channel not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, _ = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				channelA.ID = ibctesting.InvalidID
			},
			expPass: false,
		},
		{
			name: "next seq send not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA = suite.chainA.NextTestChannel(connA, ibctesting.LinksPort)
				channelB = suite.chainB.NextTestChannel(connB, ibctesting.LinksPort)

				// manually create channel so next seq send is never set
				suite.chainA.App.IBCKeeper.ChannelKeeper.SetChannel(
					suite.chainA.GetContext(),
					channelA.PortID, channelA.ID,
					channeltypes.NewChannel(
						channeltypes.OPEN,
						channeltypes.ORDERED,
						channeltypes.NewCounterparty(channelB.PortID, channelB.ID),
						[]string{connA.ID},
						ibctesting.DefaultChannelVersion,
					),
				)
				suite.chainA.CreateChannelCapability(channelA.PortID, channelA.ID)
			},
			expPass: false,
		},
		{
			name: "channel capability not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				cap := suite.chainA.GetChannelCapability(channelA.PortID, channelA.ID)
				// Release channel capability
				suite.chainA.App.ScopedLinksKeeper.ReleaseCapability(suite.chainA.GetContext(), cap)
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			srcAddress := suite.chainA.Account.GetAddress().String()
			srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
			dstAddress := suite.chainB.Account.GetAddress().String()
			link := types.NewLink(srcAddress, dstAddress)
			linkBz, _ := link.Marshal()
			srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
			srcSigHex := hex.EncodeToString(srcSig)
			dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
			dstSigHex := hex.EncodeToString(dstSig)

			// send link from chainA to chainB
			packet := types.NewIBCAccountConnectionPacketData(
				"cosmos",
				srcAddress,
				srcPubKeyHex,
				dstAddress,
				srcSigHex,
				dstSigHex,
			)
			err := suite.chainA.App.LinksKeeper.TransmitIBCAccountConnectionPacket(
				suite.chainA.GetContext(),
				packet,
				channelA.PortID,
				channelA.ID,
				clienttypes.NewHeight(0, 100),
				0,
			)
			if test.expPass {
				suite.Require().NoError(err) // message committed
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnRecvIBCAccountConnectionPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		srcAddr            string
		srcPubKeyHex       string
		dstAddr            string
		srcSigHex          string
		dstSigHex          string
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "successful create link from source chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = suite.chainB.Account.GetAddress().String()

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: true,
		},
		{
			name: "non exist destination address on destination chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = suite.chainA.Account.GetAddress().String()

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
		{
			name: "invalid pubkey for source signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainB.Account.GetPubKey().Bytes())
				dstAddr = suite.chainA.Account.GetAddress().String()

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
		{
			name: "invalid source signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = suite.chainA.Account.GetAddress().String()

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainB.PrivKey.Sign(linkBz)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
		{
			name: "invalid destination signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = suite.chainA.Account.GetAddress().String()

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainA.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			// send coin from chainA to chainB
			msg := types.NewMsgCreateIBCAccountConnection(
				channelA.PortID, channelA.ID, 0,
				srcAddr, srcPubKeyHex, dstAddr, srcSigHex, dstSigHex,
			)
			err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
			suite.Require().NoError(err) // message committed

			data := types.NewIBCAccountConnectionPacketData(
				"cosmos", srcAddr, srcPubKeyHex, dstAddr, srcSigHex, dstSigHex,
			)
			bz := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			_, err = suite.chainB.App.LinksKeeper.OnRecvIBCAccountConnectionPacket(
				suite.chainB.GetContext(),
				packet,
				data,
			)
			if test.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnAcknowledgementIBCAccountConnectionPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		ack                channeltypes.Acknowledgement
	)
	tests := []struct {
		name     string
		malleate func()
		success  bool // success of ack
	}{
		{
			name: "success ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				packetAck := types.IBCAccountConnectionPacketAck{SourceAddress: suite.chainA.Account.GetAddress().String()}
				bz, _ := packetAck.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			success: true,
		},
		{
			name: "unsuccess ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.NewErrorAcknowledgement("failed links packet")
			},
			success: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			srcAddr := suite.chainA.Account.GetAddress().String()
			srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
			dstAddr := suite.chainA.Account.GetAddress().String()

			link := types.NewLink(srcAddr, dstAddr)
			linkBz, _ := link.Marshal()
			srcSig, _ := suite.chainB.PrivKey.Sign(linkBz)
			srcSigHex := hex.EncodeToString(srcSig)
			dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
			dstSigHex := hex.EncodeToString(dstSig)

			data := types.NewIBCAccountConnectionPacketData(
				"cosmos",
				srcAddr,
				srcPubKeyHex,
				dstAddr,
				srcSigHex,
				dstSigHex,
			)
			bz := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			err := suite.chainA.App.LinksKeeper.OnAcknowledgementIBCAccountConnectionPacket(suite.chainA.GetContext(), packet, data, ack)
			if test.success {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestIBCAccountLinkPacket() {

	suite.Run("packet transformation test", func() {
		suite.SetupIBCTest()

		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
		channelA, channelB := suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

		srcAddr := suite.chainA.Account.GetAddress().String()
		pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
		dstAddress := srcAddr
		link := types.NewLink(srcAddr, dstAddress)
		linkBz, _ := link.Marshal()
		sig, _ := suite.chainA.PrivKey.Sign(linkBz)
		sigHex := hex.EncodeToString(sig)

		msg := types.NewMsgCreateIBCAccountLink(channelA.PortID, channelA.ID, 0, srcAddr, pubKeyHex, sigHex)
		err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
		suite.Require().NoError(err) // message committed

		// send link from chainA to chainB
		packetData := types.NewIBCAccountLinkPacketData(
			"cosmos",
			srcAddr,
			pubKeyHex,
			sigHex,
		)
		bz, err := packetData.GetBytes()
		suite.Require().NoError(err)

		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
		packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
		proof, proofHeight := suite.chainA.QueryProof(packetKey)

		recvMsg := channeltypes.NewMsgRecvPacket(packet, proof, proofHeight, suite.chainB.Account.GetAddress())
		err = suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, recvMsg)
		suite.Require().NoError(err) // message committed
	})
}

func (suite *KeeperTestSuite) TestTransmitIBCAccountLinkPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "successful create link from source chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, _ = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

			},
			expPass: true,
		},
		{
			name: "source channel not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, _ = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				channelA.ID = ibctesting.InvalidID
			},
			expPass: false,
		},
		{
			name: "next seq send not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA = suite.chainA.NextTestChannel(connA, ibctesting.LinksPort)
				channelB = suite.chainB.NextTestChannel(connB, ibctesting.LinksPort)

				// manually create channel so next seq send is never set
				suite.chainA.App.IBCKeeper.ChannelKeeper.SetChannel(
					suite.chainA.GetContext(),
					channelA.PortID, channelA.ID,
					channeltypes.NewChannel(
						channeltypes.OPEN,
						channeltypes.ORDERED,
						channeltypes.NewCounterparty(channelB.PortID, channelB.ID),
						[]string{connA.ID},
						ibctesting.DefaultChannelVersion,
					),
				)
				suite.chainA.CreateChannelCapability(channelA.PortID, channelA.ID)
			},
			expPass: false,
		},
		{
			name: "channel capability not found",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				cap := suite.chainA.GetChannelCapability(channelA.PortID, channelA.ID)
				// Release channel capability
				suite.chainA.App.ScopedLinksKeeper.ReleaseCapability(suite.chainA.GetContext(), cap)
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			srcAddr := suite.chainA.Account.GetAddress().String()
			pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
			dstAddress := srcAddr
			link := types.NewLink(srcAddr, dstAddress)
			linkBz, _ := link.Marshal()
			sig, _ := suite.chainA.PrivKey.Sign(linkBz)
			sigHex := hex.EncodeToString(sig)

			// send link from chainA to chainB
			packet := types.NewIBCAccountLinkPacketData(
				"cosmos",
				srcAddr,
				pubKeyHex,
				sigHex,
			)
			err := suite.chainA.App.LinksKeeper.TransmitIBCAccountLinkPacket(
				suite.chainA.GetContext(),
				packet,
				channelA.PortID,
				channelA.ID,
				clienttypes.NewHeight(0, 100),
				0,
			)
			if test.expPass {
				suite.Require().NoError(err) // message committed
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnRecvIBCAccountLinkPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		srcAddr            string
		srcPubKeyHex       string
		dstAddr            string
		sigHex             string
	)

	tests := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			name: "successful create link from source chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = srcAddr

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				sigHex = hex.EncodeToString(srcSig)

			},
			expPass: true,
		},
		{
			name: "invalid source pubkey",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainB.Account.GetPubKey().Bytes())
				dstAddr = srcAddr

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainA.PrivKey.Sign(linkBz)
				sigHex = hex.EncodeToString(srcSig)

			},
			expPass: false,
		},
		{
			name: "invalid source signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				dstAddr = srcAddr

				link := types.NewLink(srcAddr, dstAddr)
				linkBz, _ := link.Marshal()
				srcSig, _ := suite.chainB.PrivKey.Sign(linkBz)
				sigHex = hex.EncodeToString(srcSig)
			},
			expPass: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			// send coin from chainA to chainB
			msg := types.NewMsgCreateIBCAccountLink(
				channelA.PortID, channelA.ID, 0,
				srcAddr, srcPubKeyHex, sigHex,
			)
			err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
			suite.Require().NoError(err) // message committed

			data := types.NewIBCAccountLinkPacketData(
				"cosmos", srcAddr, srcPubKeyHex, sigHex,
			)
			bz, _ := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			_, err = suite.chainB.App.LinksKeeper.OnRecvIBCAccountLinkPacket(
				suite.chainB.GetContext(),
				packet,
				data,
			)
			if test.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnAcknowledgementIBCAccountLinkPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		ack                channeltypes.Acknowledgement
	)
	tests := []struct {
		name     string
		malleate func()
		success  bool // success of ack
	}{
		{
			name: "success ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				packetAck := types.IBCAccountLinkPacketAck{SourceAddress: suite.chainA.Account.GetAddress().String()}
				bz, _ := packetAck.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			success: true,
		},
		{
			name: "unsuccess ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateLinksChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.NewErrorAcknowledgement("failed links packet")
			},
			success: false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			srcAddr := suite.chainA.Account.GetAddress().String()
			pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
			dstAddress := srcAddr
			link := types.NewLink(srcAddr, dstAddress)
			linkBz, _ := link.Marshal()
			sig, _ := suite.chainA.PrivKey.Sign(linkBz)
			sigHex := hex.EncodeToString(sig)

			data := types.NewIBCAccountLinkPacketData(
				"cosmos",
				srcAddr,
				pubKeyHex,
				sigHex,
			)
			bz, _ := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			err := suite.chainA.App.LinksKeeper.OnAcknowledgementIBCAccountLinkPacket(suite.chainA.GetContext(), packet, data, ack)
			if test.success {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
