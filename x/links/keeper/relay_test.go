package keeper_test

import (
	"encoding/hex"

	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	ibctesting "github.com/desmos-labs/desmos/testing"
	"github.com/desmos-labs/desmos/x/links/types"
)

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
			},
			expPass: true,
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

			// send coin from chainA to chainB
			transferMsg := types.NewMsgCreateIBCAccountConnection(
				channelA.PortID, channelA.ID, 0,
				"cosmos", srcAddress, srcPubKeyHex, dstAddress, srcSigHex, dstSigHex,
			)
			err := suite.coordinator.SendMsg(suite.chainA, suite.chainB, channelB.ClientID, transferMsg)
			suite.Require().NoError(err) // message committed

			data := types.NewIBCAccountConnectionPacketData(
				"cosmos", srcAddress, srcPubKeyHex, dstAddress, srcSigHex, dstSigHex,
			)
			bz, _ := data.GetBytes()
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
