package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v3/testutil"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestOnRecvPacket() {
	var (
		packetData types.LinkChainAccountPacketData
		srcAddr    string
		destAddr   string
	)

	testCases := []struct {
		name        string
		malleate    func(srcAddr, srcSigHex, destAddr, destSigHex string)
		store       func()
		doubleStore bool
		expPass     bool
	}{
		{
			name: "invalid packet returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.LinkChainAccountPacketData{
					SourceAddress: nil,
					SourceProof: types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					SourceChainConfig:  types.NewChainConfig("cosmos"),
					DestinationAddress: destAddr,
					DestinationProof: types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				}
			},
			expPass: false,
		},
		{
			name: "failed unpack source address returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				invalidAny, err := codectypes.NewAnyWithValue(secp256k1.GenPrivKey())
				suite.Require().NoError(err)
				packetData = types.LinkChainAccountPacketData{
					SourceAddress: invalidAny,
					SourceProof: types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					SourceChainConfig:  types.NewChainConfig("cosmos"),
					DestinationAddress: destAddr,
					DestinationProof: types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				}
			},
			expPass: false,
		},
		{
			name: "invalid destination address returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig("cosmos"),
					"cosmos1asdjlansdjhasd",
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			expPass: false,
		},
		{
			name: "destination address without profile returns error",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig("cosmos"),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			expPass: false,
		},
		{
			name: "returns error if the profile public key does not match provided public key",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig("cosmos"),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainA.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"tc-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.SaveProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: false,
		},
		{
			name: "returns error when source proof verification fails",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						"696e76616c6964",
					),
					types.NewChainConfig(
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"tc-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.SaveProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: false,
		},
		{
			name: "returns error when destination proof verification fails",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig(
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						"696e76616c6964",
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"tc-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.SaveProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: false,
		},
		{
			name: "returns error when failed to store chain link",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig(
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"tc-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.SaveProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)

				// Store link
				store := suite.chainB.GetContext().KVStore(suite.chainB.App.GetKey(types.StoreKey))
				key := types.ChainLinksStoreKey(baseAcc.GetAddress().String(), "cosmos", srcAddr)
				link := types.NewChainLink(
					addr.String(),
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(suite.chainA.Account.GetPubKey(), testutil.SingleSignatureProtoFromHex("1234"), hex.EncodeToString([]byte(srcAddr))),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link))
			},
			expPass: false,
		},
		{
			name: "valid link is created successfully",
			malleate: func(srcAddr, srcSigHex, destAddr, destSigHex string) {
				packetData = types.NewLinkChainAccountPacketData(
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(
						suite.chainA.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(srcSigHex),
						hex.EncodeToString([]byte(destAddr)),
					),
					types.NewChainConfig(
						"cosmos",
					),
					destAddr,
					types.NewProof(
						suite.chainB.Account.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(destSigHex),
						hex.EncodeToString([]byte(srcAddr)),
					),
				)
			},
			store: func() {
				addr := suite.chainB.Account.GetAddress()
				baseAcc := authtypes.NewBaseAccountWithAddress(addr)
				baseAcc.SetPubKey(suite.chainB.Account.GetPubKey())

				profile, err := types.NewProfile(
					"dtag",
					"tc-user",
					"biography",
					types.NewPictures(
						"https://shorturl.at/adnX3",
						"https://shorturl.at/cgpyF",
					),
					time.Time{},
					baseAcc,
				)
				suite.Require().NoError(err)
				err = suite.chainB.App.ProfileKeeper.SaveProfile(suite.chainB.GetContext(), profile)
				suite.Require().NoError(err)
			},
			expPass: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupIBCTest()
			srcAddr = suite.chainA.Account.GetAddress().String()
			destAddr = suite.chainB.Account.GetAddress().String()

			srcSig, err := suite.chainA.PrivKey.Sign([]byte(destAddr))
			suite.NoError(err)
			srcSigHex := hex.EncodeToString(srcSig)

			dstSig, err := suite.chainB.PrivKey.Sign([]byte(srcAddr))
			suite.NoError(err)
			destSigHex := hex.EncodeToString(dstSig)

			tc.malleate(srcAddr, srcSigHex, destAddr, destSigHex)
			if tc.store != nil {
				tc.store()
			}

			_, err = suite.chainB.App.ProfileKeeper.OnRecvLinkChainAccountPacket(suite.chainB.GetContext(), packetData)
			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}
