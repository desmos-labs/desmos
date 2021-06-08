package keeper_test

import (
	"encoding/hex"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreChainLink() {
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

	invalidAny, err := codectypes.NewAnyWithValue(srcPriv)
	suite.Require().NoError(err)

	profileAcc, err := sdk.AccAddressFromBech32(suite.testData.user)
	suite.Require().NoError(err)

	tests := []struct {
		name      string
		store     func()
		user      string
		link      types.ChainLink
		shouldErr bool
		expStored []sdk.AccAddress
	}{
		{
			name: "Invalid chain address packed value returns error",
			user: suite.testData.user,
			link: types.ChainLink{
				Address:      invalidAny,
				Proof:        types.NewProof(srcPubKey, srcSigHex, srcAddr),
				ChainConfig:  types.NewChainConfig("cosmos"),
				CreationTime: time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			},
			shouldErr: true,
		},
		{
			name: "Invalid chain address returns error",
			user: suite.testData.user,
			link: types.NewChainLink(
				types.NewBech32Address("", "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "Invalid proof returns error",
			user: suite.testData.user,
			link: types.NewChainLink(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, "wrong"),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "Link already existing returns error",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey("cosmos", srcAddr)
				store.Set(key, profileAcc)
			},
			user: suite.testData.user,
			link: types.NewChainLink(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []sdk.AccAddress{profileAcc},
		},
		{
			name: "Invalid user returns error",
			user: "",
			link: types.NewChainLink(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []sdk.AccAddress{},
		},
		{
			name:  "User with no profile returns error",
			store: func() {},
			user:  suite.testData.user,
			link: types.NewChainLink(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []sdk.AccAddress{},
		},
		{
			name: "Valid conditions return no error",
			store: func() {
				err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
				suite.Require().NoError(err)
			},
			user: suite.testData.user,
			link: types.NewChainLink(
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
			expStored: []sdk.AccAddress{profileAcc},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store()
			}

			err = suite.k.StoreChainLink(suite.ctx, test.user, test.link)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				address := test.link.Address.GetCachedValue().(types.AddressData)
				addr, found := suite.k.GetAccountByChainLink(suite.ctx, test.link.ChainConfig.Name, address.GetAddress())
				suite.Require().True(found)
				suite.Require().Contains(test.expStored, addr)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteChainLink() {

	invalidAny, err := codectypes.NewAnyWithValue(secp256k1.GenPrivKey())
	suite.Require().NoError(err)

	profileAcc, err := sdk.AccAddressFromBech32(suite.testData.user)
	suite.Require().NoError(err)

	tests := []struct {
		name      string
		store     func()
		owner     string
		chainName string
		address   string
		shouldErr bool
	}{
		{
			name:      "Invalid owner address returns error",
			owner:     "",
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name:      "Owner without profile returns error",
			owner:     suite.testData.otherUser,
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name: "Target address not linked to the profile returns error",
			store: func() {
				err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
				suite.Require().NoError(err)
			},
			owner:     suite.testData.user,
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name: "Address data unpack failed returns error",
			store: func() {
				// Store profile
				profile := *suite.testData.profile
				link := types.ChainLink{
					Address: invalidAny,
					Proof: types.NewProof(
						secp256k1.GenPrivKey().PubKey(),
						"signature",
						"plain_text",
					),
					ChainConfig:  types.NewChainConfig("cosmos"),
					CreationTime: suite.testData.profile.CreationDate,
				}
				profile.ChainsLinks = []types.ChainLink{link}
				err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
				suite.Require().NoError(err)

				// Store link
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey("cosmos", suite.testData.otherUser)
				store.Set(key, profileAcc)
			},
			owner:     suite.testData.user,
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name: "Valid condition returns no error",
			store: func() {
				// Store profile
				profile := suite.testData.profile
				profile.ChainsLinks = []types.ChainLink{
					types.NewChainLink(
						types.NewBech32Address(suite.testData.otherUser, "cosmos"),
						types.NewProof(secp256k1.GenPrivKey().PubKey(), "signature", "plain_text"),
						types.NewChainConfig("cosmos"),
						suite.testData.profile.CreationDate,
					),
				}
				err = suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)

				// Store link
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey("cosmos", suite.testData.otherUser)
				store.Set(key, profileAcc)
			},
			owner:     suite.testData.user,
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: false,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store()
			}

			err = suite.k.DeleteChainLink(suite.ctx, test.owner, test.chainName, test.address)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				_, found := suite.k.GetAccountByChainLink(suite.ctx, test.chainName, test.address)
				suite.Require().False(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAccountByChainLink() {
	tests := []struct {
		name        string
		store       func()
		chainName   string
		address     string
		shouldFound bool
		expRes      string
	}{
		{
			name:        "Non existent link returns anything",
			store:       func() {},
			chainName:   "cosmos",
			address:     suite.testData.user,
			shouldFound: false,
			expRes:      "",
		},
		{
			name: "existent link returns no error",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainsLinksStoreKey("cosmos", suite.testData.user)
				acc, err := sdk.AccAddressFromBech32(suite.testData.otherUser)
				suite.Require().NoError(err)
				store.Set(key, acc)
			},
			chainName:   "cosmos",
			address:     suite.testData.user,
			shouldFound: true,
			expRes:      suite.testData.otherUser,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			test.store()
			acc, found := suite.k.GetAccountByChainLink(suite.ctx, test.chainName, test.address)
			if test.shouldFound {
				suite.Require().True(found)
				suite.Equal(test.expRes, acc.String())
			} else {
				suite.Require().False(found)
			}
		})
	}
}
