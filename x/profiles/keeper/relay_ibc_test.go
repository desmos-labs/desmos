package keeper_test

import (
	"encoding/hex"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	"github.com/cosmos/cosmos-sdk/x/ibc/core/exported"
	"github.com/desmos-labs/desmos/testutil/ibctesting"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	var (
		channelA, channelB ibctesting.TestChannel
		packetData         types.LinkChainAccountPacketData
	)

	tests := []struct {
		name        string
		malleate    func(srcAddr, srcSigHex, destAddr, destSigHex string)
		store       func()
		doubleStore bool
		expPass     bool
	}{
		{
			name: "Invalid packet returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address("", "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			store:   func() {},
			expPass: false,
		},
		{
			name: "Verify source proof failed returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						"invalid",
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			store:   func() {},
			expPass: false,
		},
		{
			name: "Verify destination proof failed returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						"invalid",
					),
				)
			},
			store:   func() {},
			expPass: false,
		},
		{
			name: "Destination has no profile returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			store:   func() {},
			expPass: false,
		},
		{
			name: "Duplicated links returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"test-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.StoreProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)

				err = suite.chainB.App.ProfileKeeper.StoreChainLink(
					suite.chainB.GetContext(),
					profile.GetAddress().String(),
					types.NewChainLink(
						types.NewBech32Address(suite.chainA.Account.GetAddress().String(), "cosmos"),
						types.NewProof(
							suite.chainA.Account.GetPubKey(),
							"signature",
							"plain_text",
						),
						types.NewChainConfig(
							"cosmos",
						),
						time.Time{},
					),
				)
				suite.Require().NoError(err)
			},
			expPass: true,
		},
		{
			name: "Create link from source chain successfully",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"test",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"test-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.StoreProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: true,
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupIBCTest()
			_, _, connA, connB := suite.coordinator.SetupClientConnections(suite.chainA, suite.chainB, exported.Tendermint)

			channelA, channelB = suite.coordinator.CreateIBCProfilesChannels(suite.chainA, suite.chainB, connA, connB, channeltypes.UNORDERED)
			srcAddr := suite.chainA.Account.GetAddress().String()

			srcSig, err := suite.chainA.PrivKey.Sign([]byte(srcAddr))
			suite.NoError(err)
			srcSigHex := hex.EncodeToString(srcSig)

			destAddr := suite.chainB.Account.GetAddress().String()
			dstSig, err := suite.chainB.PrivKey.Sign([]byte(destAddr))
			suite.NoError(err)
			destSigHex := hex.EncodeToString(dstSig)

			test.malleate(srcAddr, srcSigHex, destAddr, destSigHex)

			bz, _ := packetData.GetBytes()
			packet := channeltypes.NewPacket(bz, 1, channelA.PortID, channelA.ID, channelB.PortID, channelB.ID, clienttypes.NewHeight(0, 100), 0)

			test.store()
			_, err = suite.chainB.App.ProfileKeeper.OnRecvPacket(
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
