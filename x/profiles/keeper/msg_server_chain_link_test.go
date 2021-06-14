package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgLinkChainAccount() {
	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()

	destPriv := secp256k1.GenPrivKey()
	destPubKey := destPriv.PubKey()

	// Get Bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)
	srcSigHex := hex.EncodeToString(srcSig)

	invalidAny, err := codectypes.NewAnyWithValue(srcPriv)
	suite.Require().NoError(err)

	blockTime := time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC)

	tests := []struct {
		name      string
		store     func()
		msg       *types.MsgLinkChainAccount
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "Store chain link failed returns error",
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				destAddr,
			),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Invalid chain address packed value returns error",
			msg: &types.MsgLinkChainAccount{
				ChainAddress: invalidAny,
				Proof:        types.NewProof(srcPubKey, srcSigHex, srcAddr),
				ChainConfig:  types.NewChainConfig("cosmos"),
				Signer:       destAddr,
			},
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Create link successfully",
			store: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				suite.Require().NoError(srcBaseAcc.SetPubKey(srcPubKey))
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				suite.Require().NoError(destBaseAcc.SetPubKey(destPubKey))
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				profile, err := types.NewProfile(
					"custom_dtag",
					"my-nickname",
					"my-bio",
					types.NewPictures(
						"https://test.com/profile-picture",
						"https://test.com/cover-pic",
					),
					blockTime,
					destBaseAcc,
					nil,
				)
				suite.Require().NoError(err)
				suite.Require().NoError(suite.k.StoreProfile(suite.ctx, profile))
			},
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				destAddr,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeLinkChainAccount,
					sdk.NewAttribute(types.AttributeChainLinkSourceAddress, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, destAddr),
					sdk.NewAttribute(types.AttributeChainLinkCreationTime, blockTime.Format(time.RFC3339Nano)),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			suite.ctx = suite.ctx.WithBlockTime(blockTime)
			if test.store != nil {
				test.store()
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err = server.LinkChainAccount(sdk.WrapSDKContext(suite.ctx), test.msg)

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

				addrData := test.msg.ChainAddress.GetCachedValue().(types.AddressData)
				_, found := suite.k.GetAccountByChainLink(suite.ctx, test.msg.ChainConfig.Name, addrData.GetAddress())
				suite.Require().True(found)

				_, found, err := suite.k.GetProfile(suite.ctx, destAddr)
				suite.Require().NoError(err)
				suite.Require().True(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) Test_handleMsgUnlinkChainAccount() {
	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)

	link := types.NewChainLink(
		types.NewBech32Address(srcAddr, "cosmos"),
		types.NewProof(srcPubKey, srcSigHex, srcAddr),
		types.NewChainConfig("cosmos"),
		time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
	)

	validProfile := *suite.testData.profile

	tests := []struct {
		name            string
		msg             *types.MsgUnlinkChainAccount
		shouldErr       bool
		expEvents       sdk.Events
		existentProfile *types.Profile
		existentLinks   []types.ChainLink
	}{
		{
			name:            "Non-existent link exists returns error",
			msg:             types.NewMsgUnlinkChainAccount(validProfile.GetAddress().String(), "cosmos", srcAddr),
			shouldErr:       true,
			expEvents:       sdk.EmptyEvents(),
			existentProfile: &validProfile,
			existentLinks:   nil,
		},
		{
			name:      "No error message",
			msg:       types.NewMsgUnlinkChainAccount(validProfile.GetAddress().String(), "cosmos", srcAddr),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnlinkChainAccount,
					sdk.NewAttribute(types.AttributeChainLinkSourceAddress, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkDestinationAddress, suite.testData.profile.GetAddress().String()),
				),
			},
			existentProfile: &validProfile,
			existentLinks:   []types.ChainLink{link},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			err := suite.k.StoreProfile(suite.ctx, test.existentProfile)
			suite.Require().NoError(err)

			for _, link := range test.existentLinks {
				err := suite.k.StoreChainLink(suite.ctx, test.existentProfile.GetAddress().String(), link)
				suite.Require().NoError(err)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err = server.UnlinkChainAccount(sdk.WrapSDKContext(suite.ctx), test.msg)
			suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				profile, found, err := suite.k.GetProfile(suite.ctx, suite.testData.profile.GetAddress().String())
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Empty(profile.ChainsLinks)
			}
		})
	}
}
