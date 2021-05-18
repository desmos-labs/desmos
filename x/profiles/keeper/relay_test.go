package keeper_test

import (
	"encoding/hex"
	"time"

	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/desmos-labs/desmos/testutil/ibctesting"
	ibcprofilestypes "github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestIBCAccountConnectionPacket() {

	suite.Run("Packet transformation test", func() {
		suite.SetupIBCTest()

		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
		channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

		height := uint64(suite.chainA.GetContext().BlockHeight())

		srcAddr := suite.chainA.Account.GetAddress().String()
		srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
		destAddr := suite.chainB.Account.GetAddress().String()
		packetProof := []byte(srcAddr)
		srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
		srcSigHex := hex.EncodeToString(srcSig)
		dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
		dstSigHex := hex.EncodeToString(dstSig)

		// send link from chainA to chainB
		packetData := ibcprofilestypes.NewIBCAccountConnectionPacketData(
			"cosmos",
			"test-net",
			srcAddr,
			srcPubKeyHex,
			destAddr,
			srcSigHex,
			dstSigHex,
		)

		msg := ibcprofilestypes.NewMsgCreateIBCAccountConnection(channelA.PortID, channelA.ID, packetData, 0)
		err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
		suite.Require().NoError(err) // message committed

		bz, _ := packetData.GetBytes()
		suite.Require().NoError(err)

		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(height, height+100), 0)
		packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
		proof, proofHeight := suite.chainA.QueryProof(packetKey)

		recvMsg := channeltypes.NewMsgRecvPacket(packet, proof, proofHeight, suite.chainB.Account.GetAddress())
		err = suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, recvMsg)
		suite.Require().NoError(err) // message committed
	})
}

func (suite *KeeperTestSuite) TestOnRecvIBCAccountConnectionPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		srcAddr            string
		srcPubKeyHex       string
		destAddr           string
		srcSigHex          string
		dstSigHex          string
	)

	tests := []struct {
		name        string
		malleate    func()
		stubPacket  func(*ibcprofilestypes.IBCAccountConnectionPacketData)
		doubleStore bool
		expPass     bool
	}{
		{
			name: "Create link from source chain successfully",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr = suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: true,
		},
		{
			name: "Non existent destination address on destination chain",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr = suite.chainA.Account.GetAddress().String()

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
		{
			name: "Invalid packet",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr = suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex = hex.EncodeToString(srcSig)
			},
			stubPacket: func(p *ibcprofilestypes.IBCAccountConnectionPacketData) {
				p.DestinationSignature = "---"
			},
			expPass: false,
		},
		{
			name: "Invalid signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr = suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign([]byte{0})
				dstSigHex = hex.EncodeToString(dstSig)

			},
			expPass: false,
		},
		{
			name: "Link already exists",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
				destAddr = suite.chainB.Account.GetAddress().String()

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				srcSigHex = hex.EncodeToString(srcSig)
				dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
				dstSigHex = hex.EncodeToString(dstSig)
			},
			doubleStore: true,
			expPass:     false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			packetData := ibcprofilestypes.NewIBCAccountConnectionPacketData(
				"cosmos",
				"test-net", srcAddr, srcPubKeyHex, destAddr, srcSigHex, dstSigHex,
			)
			// create account connection from chainA to chainB
			msg := ibcprofilestypes.NewMsgCreateIBCAccountConnection(
				channelA.PortID, channelA.ID, packetData, 0,
			)
			err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
			suite.Require().NoError(err) // message committed

			if test.stubPacket != nil {
				test.stubPacket(&packetData)
			}

			bz, _ := packetData.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			if test.doubleStore == true {
				proof := types.NewProof(packetData.SourcePubKey, packetData.SourceSignature)
				chainConfig := types.NewChainConfig(packetData.SourceChainID, packetData.SourceChainPrefix)
				link := types.NewLink(packetData.SourceAddress, proof, chainConfig, time.Now())
				suite.chainB.App.ProfileKeeper.StoreLink(suite.chainB.GetContext(), link)
			}

			_, err = suite.chainB.App.ProfileKeeper.OnRecvIBCAccountConnectionPacket(
				suite.chainB.GetContext(),
				packet,
				packetData,
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
			name: "Receive success ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				packetAck := ibcprofilestypes.IBCAccountConnectionPacketAck{SourceAddress: suite.chainA.Account.GetAddress().String()}
				bz, _ := packetAck.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			success: true,
		},
		{
			name: "Receive unsuccess ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.NewErrorAcknowledgement("failed ibc porifles packet")
			},
			success: false,
		},
		{
			name: "Receive invalid ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.Acknowledgement{}
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
			destAddr := suite.chainA.Account.GetAddress().String()

			packetProof := []byte(srcAddr)
			srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
			srcSigHex := hex.EncodeToString(srcSig)
			dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
			dstSigHex := hex.EncodeToString(dstSig)

			data := ibcprofilestypes.NewIBCAccountConnectionPacketData(
				"cosmos",
				"test-net",
				srcAddr,
				srcPubKeyHex,
				destAddr,
				srcSigHex,
				dstSigHex,
			)
			bz, _ := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			err := suite.chainA.App.ProfileKeeper.OnAcknowledgementIBCAccountConnectionPacket(suite.chainA.GetContext(), packet, data, ack)
			if test.success {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnTimeoutIBCAccountConnectionPacket() {
	suite.Run("Receive timeout packet and returns nil", func() {
		suite.SetupIBCTest()
		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
		channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
		srcAddr := suite.chainA.Account.GetAddress().String()
		srcPubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
		destAddr := suite.chainA.Account.GetAddress().String()

		packetProof := []byte(srcAddr)
		srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
		srcSigHex := hex.EncodeToString(srcSig)
		dstSig, _ := suite.chainB.PrivKey.Sign(srcSig)
		dstSigHex := hex.EncodeToString(dstSig)

		data := ibcprofilestypes.NewIBCAccountConnectionPacketData(
			"cosmos",
			"test-net",
			srcAddr,
			srcPubKeyHex,
			destAddr,
			srcSigHex,
			dstSigHex,
		)
		bz, _ := data.GetBytes()
		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
		err := suite.chainA.App.ProfileKeeper.OnTimeoutIBCAccountConnectionPacket(suite.chainA.GetContext(), packet, data)
		suite.Require().NoError(err)
	})
}

// ___________________________________________________________________________________________________________________

func (suite *KeeperTestSuite) TestIBCAccountLinkPacket() {

	suite.Run("Packet transformation test", func() {
		suite.SetupIBCTest()

		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
		channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)

		height := uint64(suite.chainA.GetContext().BlockHeight())

		srcAddr := suite.chainA.Account.GetAddress().String()
		pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
		packetProof := []byte(srcAddr)
		sig, _ := suite.chainA.PrivKey.Sign(packetProof)
		sigHex := hex.EncodeToString(sig)

		packetData := ibcprofilestypes.NewIBCAccountLinkPacketData(
			"cosmos",
			"test-net",
			srcAddr,
			pubKeyHex,
			sigHex,
		)

		msg := ibcprofilestypes.NewMsgCreateIBCAccountLink(channelA.PortID, channelA.ID, packetData, 0)
		err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
		suite.Require().NoError(err) // message committed

		// send link from chainA to chainB
		bz, err := packetData.GetBytes()
		suite.Require().NoError(err)

		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(height, height+100), 0)
		packetKey := host.PacketCommitmentKey(packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())
		proof, proofHeight := suite.chainA.QueryProof(packetKey)
		suite.T().Log(proofHeight)
		suite.T().Log(height)

		recvMsg := channeltypes.NewMsgRecvPacket(packet, proof, proofHeight, suite.chainB.Account.GetAddress())
		err = suite.coordinator.SendMsg(suite.chainB, suite.chainA, channelA.ClientID, recvMsg)
		suite.Require().NoError(err) // message committed
	})
}

func (suite *KeeperTestSuite) TestOnRecvIBCAccountLinkPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		srcAddr            string
		srcPubKeyHex       string
		sigHex             string
	)

	tests := []struct {
		name        string
		malleate    func()
		stubPacket  func(*ibcprofilestypes.IBCAccountLinkPacketData)
		doubleStore bool
		expPass     bool
	}{
		{
			name: "Create link from source chain successfully",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex = hex.EncodeToString(srcSig)

			},
			expPass: true,
		},
		{
			name: "Invalid packet",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex = hex.EncodeToString(srcSig)
			},
			stubPacket: func(p *ibcprofilestypes.IBCAccountLinkPacketData) {
				p.Signature = "="
			},
			expPass: false,
		},
		{
			name: "Invalid source signature",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())

				srcSig, _ := suite.chainA.PrivKey.Sign([]byte{0})
				sigHex = hex.EncodeToString(srcSig)

			},
			expPass: false,
		},
		{
			name: "Link already exists",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				srcAddr = suite.chainA.Account.GetAddress().String()
				srcPubKeyHex = hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())

				packetProof := []byte(srcAddr)
				srcSig, _ := suite.chainA.PrivKey.Sign(packetProof)
				sigHex = hex.EncodeToString(srcSig)

			},
			doubleStore: true,
			expPass:     false,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			test.malleate()

			packetData := ibcprofilestypes.NewIBCAccountLinkPacketData(
				"cosmos",
				"test-net", srcAddr, srcPubKeyHex, sigHex,
			)

			// send coin from chainA to chainB
			msg := ibcprofilestypes.NewMsgCreateIBCAccountLink(
				channelA.PortID, channelA.ID, packetData, 0,
			)
			err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, msg)
			suite.Require().NoError(err) // message committed

			if test.stubPacket != nil {
				test.stubPacket(&packetData)
			}

			bz, _ := packetData.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			if test.doubleStore == true {
				proof := types.NewProof(packetData.SourcePubKey, packetData.Signature)
				chainConfig := types.NewChainConfig(packetData.SourceChainID, packetData.SourceChainPrefix)
				link := types.NewLink(packetData.SourceAddress, proof, chainConfig, time.Now())
				suite.chainB.App.ProfileKeeper.StoreLink(suite.chainB.GetContext(), link)
			}

			_, err = suite.chainB.App.ProfileKeeper.OnRecvIBCAccountLinkPacket(
				suite.chainB.GetContext(),
				packet,
				packetData,
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
			name: "Receive success ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				packetAck := ibcprofilestypes.IBCAccountLinkPacketAck{SourceAddress: suite.chainA.Account.GetAddress().String()}
				bz, _ := packetAck.Marshal()
				ack = channeltypes.NewResultAcknowledgement(bz)
			},
			success: true,
		},
		{
			name: "Receive unsuccess ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.NewErrorAcknowledgement("failed ibc profiles packet")
			},
			success: false,
		},
		{
			name: "Receive invalid ack",
			malleate: func() {
				_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
				channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
				ack = channeltypes.Acknowledgement{}
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
			packetProof := []byte(srcAddr)
			sig, _ := suite.chainA.PrivKey.Sign(packetProof)
			sigHex := hex.EncodeToString(sig)

			data := ibcprofilestypes.NewIBCAccountLinkPacketData(
				"cosmos",
				"test-net",
				srcAddr,
				pubKeyHex,
				sigHex,
			)
			bz, _ := data.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
			err := suite.chainA.App.ProfileKeeper.OnAcknowledgementIBCAccountLinkPacket(suite.chainA.GetContext(), packet, data, ack)
			if test.success {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestOnTimeoutIBCAccountLinkPacket() {
	suite.Run("Receive timeout packet and returns nil", func() {
		suite.SetupIBCTest()
		_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)
		channelA, channelB := suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
		srcAddr := suite.chainA.Account.GetAddress().String()
		pubKeyHex := hex.EncodeToString(suite.chainA.Account.GetPubKey().Bytes())
		packetProof := []byte(srcAddr)
		sig, _ := suite.chainA.PrivKey.Sign(packetProof)
		sigHex := hex.EncodeToString(sig)

		data := ibcprofilestypes.NewIBCAccountLinkPacketData(
			"cosmos",
			"test-net",
			srcAddr,
			pubKeyHex,
			sigHex,
		)
		bz, _ := data.GetBytes()
		packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)
		err := suite.chainA.App.ProfileKeeper.OnTimeoutIBCAccountLinkPacket(suite.chainA.GetContext(), packet, data)
		suite.Require().NoError(err)
	})
}
