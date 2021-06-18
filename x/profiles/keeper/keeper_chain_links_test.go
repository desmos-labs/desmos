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

	tests := []struct {
		name      string
		store     func()
		user      string
		link      types.ChainLink
		shouldErr bool
		expStored []types.ChainLink
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
				profileAcc.String(),
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
				profileAcc.String(),
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
				key := types.ChainLinksStoreKey(profileAcc.String(), "cosmos", srcAddr)
				link := types.NewChainLink(
					profileAcc.String(),
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(secp256k1.GenPrivKey().PubKey(), "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link))
				key := types.ChainsLinksStoreKey("cosmos", srcAddr)
				store.Set(key, suite.testData.profile.GetAddress())
			},
			user: suite.testData.user,
			link: types.NewChainLink(
				profileAcc.String(),
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []types.ChainLink{
				types.NewChainLink(
					profileAcc.String(),
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(secp256k1.GenPrivKey().PubKey(), "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				),
			},
			expStored: []sdk.AccAddress{suite.testData.profile.GetAddress()},
		},
		{
			name: "Invalid user returns error",
			user: "",
			link: types.NewChainLink(
				profileAcc.String(),
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []types.ChainLink{},
		},
		{
			name:  "User with no profile returns error",
			store: func() {},
			user:  suite.testData.user,
			link: types.NewChainLink(
				profileAcc.String(),
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			expStored: []types.ChainLink{},
		},
		{
			name: "Valid conditions return no error",
			store: func() {
				err := suite.k.StoreProfile(suite.ctx, suite.testData.profile)
				suite.Require().NoError(err)
			},
			user: suite.testData.profile.GetAddress().String(),
			link: types.NewChainLink(
				profileAcc.String(),
				types.NewBech32Address(srcAddr, "cosmos"),
				types.NewProof(srcPubKey, srcSigHex, srcAddr),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
			expStored: []sdk.AccAddress{suite.testData.profile.GetAddress()},
			expStored: []types.ChainLink{
				types.NewChainLink(
					profileAcc.String(),
					types.NewBech32Address(srcAddr, "cosmos"),
					types.NewProof(srcPubKey, srcSigHex, srcAddr),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store()
			}

			err = suite.k.StoreChainLink(suite.ctx, test.link)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				address := test.link.Address.GetCachedValue().(types.AddressData)
				link, found := suite.k.GetChainLink(suite.ctx, profileAcc.String(), test.link.ChainConfig.Name, address.GetAddress())
				suite.Require().True(found)
				suite.Require().Contains(test.expStored, link)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteChainLink() {

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
			owner:     suite.testData.profile.GetAddress().String(),
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name: "Valid condition returns no error",
			store: func() {
				// Store profile
				profile := suite.testData.profile
				err = suite.k.StoreProfile(suite.ctx, profile)
				suite.Require().NoError(err)

				// Store link
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainLinksStoreKey(profile.GetAddress().String(), "cosmos", suite.testData.otherUser)
				store.Set(key, profileAcc)
			},
			owner:     suite.testData.profile.GetAddress().String(),
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

				_, found := suite.k.GetChainLink(suite.ctx, test.owner, test.chainName, test.address)
				suite.Require().False(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetChainLink() {
	tests := []struct {
		name        string
		store       func()
		owner       string
		chainName   string
		address     string
		shouldFound bool
	}{
		{
			name:        "Non existent link returns anything",
			owner:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			chainName:   "cosmos",
			address:     "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			shouldFound: false,
		},
		{
			name: "Existent link returns no error",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				key := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", suite.testData.user)
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(secp256k1.GenPrivKey().PubKey(), "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link))
			},
			owner:       "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			chainName:   "cosmos",
			address:     suite.testData.user,
			shouldFound: true,
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store()
			}

			_, found := suite.k.GetChainLink(suite.ctx, test.owner, test.chainName, test.address)
			if test.shouldFound {
				suite.Require().True(found)
			} else {
				suite.Require().False(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetAllChainLink() {
	pub1 := secp256k1.GenPrivKey().PubKey()
	pub2 := secp256k1.GenPrivKey().PubKey()

	tests := []struct {
		name      string
		store     func()
		expStored []types.ChainLink
	}{
		{
			name:      "Non existent link returns empty array",
			expStored: []types.ChainLink{},
		},
		{
			name: "Existent links returns all links",
			store: func() {
				store := suite.ctx.KVStore(suite.storeKey)
				store.Set(
					types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f"),
					types.MustMarshalChainLink(
						suite.cdc,
						types.NewChainLink(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
							types.NewProof(pub1, "signature", "plain_text"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
					),
				)
				store.Set(
					types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs"),
					types.MustMarshalChainLink(
						suite.cdc,
						types.NewChainLink(
							"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
							types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
							types.NewProof(pub2, "signature", "plain_text"),
							types.NewChainConfig("cosmos"),
							time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
						),
					),
				)
			},
			expStored: []types.ChainLink{
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					types.NewProof(pub1, "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				),
				types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					types.NewProof(pub2, "signature", "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				),
			},
		},
	}

	for _, test := range tests {
		suite.Run(test.name, func() {
			suite.SetupTest()
			if test.store != nil {
				test.store()
			}
			links := suite.k.GetAllChainLinks(suite.ctx)
			suite.Require().Equal(len(test.expStored), len(links))
			for _, link := range links {
				suite.Require().Contains(test.expStored, link)
			}
		})
	}
}
