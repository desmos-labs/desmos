package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) Test_handleMsgLinkChainAccount() {
	var existentProfiles []*types.Profile

	// Generate source and destination key
	srcPriv := secp256k1.GenPrivKey()
	destPriv := secp256k1.GenPrivKey()
	srcPubKey := srcPriv.PubKey()
	destPubKey := destPriv.PubKey()

	// Get bech32 encoded addresses
	srcAddr, err := bech32.ConvertAndEncode("cosmos", srcPubKey.Address().Bytes())
	suite.Require().NoError(err)
	destAddr, err := bech32.ConvertAndEncode("cosmos", destPubKey.Address().Bytes())
	suite.Require().NoError(err)

	// Get signature by signing with keys
	srcSig, err := srcPriv.Sign([]byte(srcAddr))
	suite.Require().NoError(err)
	destSig, err := destPriv.Sign([]byte(destAddr))
	suite.Require().NoError(err)

	srcSigHex := hex.EncodeToString(srcSig)
	destSigHex := hex.EncodeToString(destSig)

	blockTime := suite.testData.profile.CreationDate

	tests := []struct {
		name      string
		store     func()
		msg       *types.MsgLinkChainAccount
		shouldErr bool
		expEvents sdk.Events
	}{
		{
			name: "Invalid destination proof returns error",
			store: func() {
			},
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				destAddr,
				types.NewProof(destPubKey, destSigHex, "wrong"),
			),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Store chain link failed returns error",
			store: func() {
			},
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				destAddr,
				types.NewProof(destPubKey, destSigHex, destAddr),
			),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "Create link successfully",
			store: func() {
				srcAccAddr, err := sdk.AccAddressFromBech32(srcAddr)
				suite.Require().NoError(err)

				srcBaseAcc := authtypes.NewBaseAccountWithAddress(srcAccAddr)
				srcBaseAcc.SetPubKey(srcPubKey)
				suite.ak.SetAccount(suite.ctx, srcBaseAcc)

				destAccAddr, err := sdk.AccAddressFromBech32(destAddr)
				suite.Require().NoError(err)
				destBaseAcc := authtypes.NewBaseAccountWithAddress(destAccAddr)
				destBaseAcc.SetPubKey(destPubKey)
				suite.ak.SetAccount(suite.ctx, destBaseAcc)

				existentProfiles = []*types.Profile{
					suite.CheckProfileNoError(types.NewProfile(
						"custom_dtag",
						"my-nickname",
						"my-bio",
						types.NewPictures(
							"https://test.com/profile-picture",
							"https://test.com/cover-pic",
						),
						suite.testData.profile.CreationDate,
						destBaseAcc,
					)),
				}
			},
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				destAddr,
				types.NewProof(destPubKey, destSigHex, destAddr),
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeLinkChainAccount,
					sdk.NewAttribute(types.AttributeChainLinkAccountTarget, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkAccountOwner, destAddr),
					sdk.NewAttribute(types.AttributeChainLinkCreated, blockTime.Format(time.RFC3339Nano)),
				),
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()
			test.store()
			suite.ctx = suite.ctx.WithBlockTime(blockTime)
		})
		for _, acc := range existentProfiles {
			err := suite.k.StoreProfile(suite.ctx, acc)
			suite.Require().NoError(err)
		}

		server := keeper.NewMsgServerImpl(suite.k)
		_, err := server.LinkChainAccount(sdk.WrapSDKContext(suite.ctx), test.msg)
		suite.Require().Equal(test.expEvents, suite.ctx.EventManager().Events())

		if test.shouldErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)

			stored := suite.k.GetAllAccountsByChainLink(suite.ctx)
			suite.Require().NotEmpty(stored)

			_, found, err := suite.k.GetProfile(suite.ctx, destAddr)
			suite.Require().NoError(err)
			suite.Require().True(found)
		}
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
					sdk.NewAttribute(types.AttributeChainLinkAccountTarget, srcAddr),
					sdk.NewAttribute(types.AttributeChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeChainLinkAccountOwner, suite.testData.user),
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

				profile, found, err := suite.k.GetProfile(suite.ctx, suite.testData.user)
				suite.Require().NoError(err)
				suite.Require().True(found)
				suite.Require().Empty(profile.ChainsLinks)
			}
		})
	}
}
