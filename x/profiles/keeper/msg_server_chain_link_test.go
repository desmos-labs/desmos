package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/desmos-labs/desmos/v3/testutil"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestMsgServer_LinkChainAccount() {
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
	srcSig, err := srcPriv.Sign([]byte(destAddr))
	suite.Require().NoError(err)
	srcSigHex := hex.EncodeToString(srcSig)

	blockTime := time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgLinkChainAccount
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid chain link returns error",
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, testutil.SingleSignatureProtoFromHex(srcSigHex), hex.EncodeToString([]byte(srcAddr))),
				types.NewChainConfig("cosmos"),
				destAddr,
			),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "invalid chain address packed value returns error",
			msg: &types.MsgLinkChainAccount{
				ChainAddress: testutil.NewAny(srcPriv),
				Proof:        types.NewProof(srcPubKey, testutil.SingleSignatureProtoFromHex(srcSigHex), hex.EncodeToString([]byte(srcAddr))),
				ChainConfig:  types.NewChainConfig("cosmos"),
				Signer:       destAddr,
			},
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "valid request stores the link correctly",
			store: func(ctx sdk.Context) {
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
						"https://tc.com/profile-picture",
						"https://tc.com/cover-pic",
					),
					ctx.BlockTime(),
					destBaseAcc,
				)
				suite.Require().NoError(err)
				suite.Require().NoError(suite.k.SaveProfile(suite.ctx, profile))
			},
			msg: types.NewMsgLinkChainAccount(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, testutil.SingleSignatureProtoFromHex(srcSigHex), hex.EncodeToString([]byte(destAddr))),
				types.NewChainConfig("cosmos"),
				destAddr,
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeLinkChainAccount,
					sdk.NewAttribute(types.AttributeKeyChainLinkSourceAddress, srcAddr),
					sdk.NewAttribute(types.AttributeKeyChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeKeyChainLinkDestinationAddress, destAddr),
					sdk.NewAttribute(types.AttributeKeyChainLinkCreationTime, blockTime.Format(time.RFC3339Nano)),
				),
			},
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetChainLink(ctx, destAddr, "cosmos", srcAddr)
				suite.Require().True(found)

				_, found, err := suite.k.GetProfile(ctx, destAddr)
				suite.Require().NoError(err)
				suite.Require().True(found)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			ctx = ctx.WithBlockTime(blockTime)
			if tc.store != nil {
				tc.store(ctx)
			}

			server := keeper.NewMsgServerImpl(suite.k)
			_, err = server.LinkChainAccount(sdk.WrapSDKContext(ctx), tc.msg)

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

func (suite *KeeperTestSuite) TestMsgServer_UnlinkChainAccount() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		msg       *types.MsgUnlinkChainAccount
		shouldErr bool
		expEvents sdk.Events
		check     func(ctx sdk.Context)
	}{
		{
			name: "non existent link exists returns error",
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: true,
			expEvents: sdk.EmptyEvents(),
		},
		{
			name: "found link returns no error",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(
						testutil.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d"),
						testutil.SingleSignatureProtoFromHex("909e38994b1583d3f14384c2e9a03c90064e8fd8e19b780bb0ba303dfe671a27287da04d0ce096ce9a140bd070ee36818f5519eb2070a16971efd8143855524b"),
						"74657874",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)
			},
			msg: types.NewMsgUnlinkChainAccount(
				"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				"cosmos",
				"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			),
			shouldErr: false,
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeUnlinkChainAccount,
					sdk.NewAttribute(types.AttributeKeyChainLinkSourceAddress, "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns"),
					sdk.NewAttribute(types.AttributeKeyChainLinkSourceChainName, "cosmos"),
					sdk.NewAttribute(types.AttributeKeyChainLinkDestinationAddress, "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
				),
			},
			check: func(ctx sdk.Context) {
				_, found := suite.k.GetChainLink(
					ctx,
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
				)
				suite.Require().False(found)
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

			server := keeper.NewMsgServerImpl(suite.k)
			_, err := server.UnlinkChainAccount(sdk.WrapSDKContext(ctx), tc.msg)
			suite.Require().Equal(tc.expEvents, ctx.EventManager().Events())

			if tc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
