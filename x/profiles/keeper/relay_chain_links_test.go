package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	var (
		packetData types.LinkChainAccountPacketData
		srcAddr    string
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
				packetData = types.LinkChainAccountPacketData{
					SourceAddress: nil,
					SourceProof: types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					SourceChainConfig: types.NewChainConfig(
						"cosmos",
					),
					DestinationAddress: destAddr,
					DestinationProof: types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				}
			},
			expPass: false,
		},
		{
			name: "Unpack source address failed returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				invalidAny, err := codectypes.NewAnyWithValue(secp256k1.GenPrivKey())
				suite.Require().NoError(err)
				packetData = types.LinkChainAccountPacketData{
					SourceAddress: invalidAny,
					SourceProof: types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					SourceChainConfig: types.NewChainConfig(
						"cosmos",
					),
					DestinationAddress: destAddr,
					DestinationProof: types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				}
			},
			expPass: false,
		},
		{
			name: "Invalid destination address returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"cosmos",
					),
					"cosmos1asdjlansdjhasd",
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			expPass: false,
		},
		{
			name: "Destination address without profile returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						destAddr,
					),
				)
			},
			expPass: false,
		},
		{
			name: "Profile public key does not equal to provided public key returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"cosmos",
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
				baseAcc.SetPubKey(suite.chainA.Account.GetPubKey())

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
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						destSigHex,
						"invalid",
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
			expPass: false,
		},
		{
			name: "Failed to store chain link returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						srcSigHex,
						srcAddr,
					),
					types.NewChainConfig(
						"cosmos",
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

				// Store link
				store := suite.chainB.GetContext().KVStore(suite.chainB.App.GetKey(types.StoreKey))
				key := types.ChainLinksStoreKey(baseAcc.GetAddress().String(), "cosmos", srcAddr)
				link := types.NewChainLink(
					addr.String(),
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(suite.chainA.Account.GetPubKey(), "signature", srcAddr),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link))
			},
			expPass: false,
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
						"cosmos",
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
			srcAddr = suite.chainA.Account.GetAddress().String()

			srcSig, err := suite.chainA.PrivKey.Sign([]byte(srcAddr))
			suite.NoError(err)
			srcSigHex := hex.EncodeToString(srcSig)

			destAddr := suite.chainB.Account.GetAddress().String()
			dstSig, err := suite.chainB.PrivKey.Sign([]byte(destAddr))
			suite.NoError(err)
			destSigHex := hex.EncodeToString(dstSig)

			test.malleate(srcAddr, srcSigHex, destAddr, destSigHex)
			if test.store != nil {
				test.store()
			}

			_, err = suite.chainB.App.ProfileKeeper.OnRecvLinkChainAccountPacket(suite.chainB.GetContext(), packetData)
			if test.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
