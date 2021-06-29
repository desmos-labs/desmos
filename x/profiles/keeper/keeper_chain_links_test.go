package keeper_test

import (
	"encoding/hex"
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_StoreChainLink() {
	// Generate source and destination key
	ext := suite.GetRandomProfile()
	sig := hex.EncodeToString(ext.Sign([]byte(ext.GetAddress().String())))

	invalidAny, err := codectypes.NewAnyWithValue(ext.privKey)
	suite.Require().NoError(err)

	useCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		link      types.ChainLink
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "Invalid chain address packed value returns error",
			link: types.ChainLink{
				Address:      invalidAny,
				Proof:        types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				ChainConfig:  types.NewChainConfig("cosmos"),
				CreationTime: time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			},
			shouldErr: true,
		},
		{
			name: "Invalid chain address returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address("", "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "Invalid proof returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, "wrong"),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "Link already existing returns error",
			store: func(ctx sdk.Context) {
				address := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				profile := suite.CreateProfileFromAddress(address)
				suite.ak.SetAccount(ctx, profile)

				link := types.NewChainLink(
					address,
					types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
					types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
			check: func(ctx sdk.Context) {
				links := suite.k.GetChainLinks(ctx)
				suite.Require().Len(links, 1)
				suite.Require().Contains(links, types.NewChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
					types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
			},
		},
		{
			name: "Invalid user returns error",
			link: types.NewChainLink(
				"",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "User with no profile returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "Valid conditions return no error",
			store: func(ctx sdk.Context) {
				profile := suite.CreateProfileFromAddress("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err = suite.k.StoreProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: false,
			check: func(ctx sdk.Context) {
				links := suite.k.GetChainLinks(ctx)
				suite.Require().Len(links, 1)
				suite.Require().Contains(links, types.NewChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
					types.NewProof(ext.GetPubKey(), sig, ext.GetAddress().String()),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
			},
		},
	}

	for _, uc := range useCases {
		suite.Run(uc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if uc.store != nil {
				uc.store(ctx)
			}

			err = suite.k.SaveChainLink(ctx, uc.link)
			if uc.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				if uc.check != nil {
					uc.check(ctx)
				}
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

func (suite *KeeperTestSuite) TestKeeper_DeleteChainLink() {
	profileAcc, err := sdk.AccAddressFromBech32(suite.testData.user)
	suite.Require().NoError(err)

	tests := []struct {
		name      string
		store     func(ctx sdk.Context)
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
			store: func(ctx sdk.Context) {
				err = suite.k.StoreProfile(ctx, suite.testData.profile.Profile)
				suite.Require().NoError(err)
			},
			owner:     suite.testData.profile.GetAddress().String(),
			chainName: "cosmos",
			address:   suite.testData.otherUser,
			shouldErr: true,
		},
		{
			name: "Valid condition returns no error",
			store: func(ctx sdk.Context) {
				// Store profile
				profile := suite.testData.profile
				err = suite.k.StoreProfile(ctx, profile.Profile)
				suite.Require().NoError(err)

				// Store link
				store := ctx.KVStore(suite.storeKey)
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
			ctx, _ := suite.ctx.CacheContext()
			if test.store != nil {
				test.store(ctx)
			}

			err = suite.k.DeleteChainLink(ctx, test.owner, test.chainName, test.address)
			if test.shouldErr {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)

				_, found := suite.k.GetChainLink(ctx, test.owner, test.chainName, test.address)
				suite.Require().False(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteAllUserChainLinks() {
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		user  string
		check func(ctx sdk.Context)
	}{
		{
			name: "empty links are deleted properly",
			user: "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
			check: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				var iterations = 0
				suite.k.IterateUserChainLinks(ctx, address, func(index int64, link types.ChainLink) (stop bool) {
					iterations++
					return false
				})
				suite.Require().Zero(iterations)
			},
		},
		{
			name: "existing chain links are deleted properly",
			store: func(ctx sdk.Context) {
				pubKey := `{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A6jN4EPjj8mHf722yjEaKaGdJpxnTR40pDvXlX1mni9C"}`

				var any codectypes.Any
				err := suite.cdc.UnmarshalJSON([]byte(pubKey), &any)
				suite.Require().NoError(err)

				var key cryptotypes.PubKey
				err = suite.cdc.UnpackAny(&any, &key)
				suite.Require().NoError(err)

				store := ctx.KVStore(suite.storeKey)
				user := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				link := types.NewChainLink(
					user,
					types.NewBech32Address("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "cosmos"),
					types.NewProof(key, "signature", "plain text"),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetAddress()),
					suite.cdc.MustMarshalBinaryBare(&link),
				)

				link = types.NewChainLink(
					user,
					types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
					types.NewProof(key, "signature", "plain text"),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetAddress()),
					suite.cdc.MustMarshalBinaryBare(&link),
				)
			},
			user: "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
			check: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				var iterations = 0
				suite.k.IterateUserChainLinks(ctx, address, func(index int64, link types.ChainLink) (stop bool) {
					iterations++
					return false
				})
				suite.Require().Zero(iterations)
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

			suite.k.DeleteAllUserChainLinks(ctx, tc.user)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
