package keeper_test

import (
	"encoding/hex"
	"time"

	"github.com/desmos-labs/desmos/v3/testutil"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (suite *KeeperTestSuite) TestKeeper_SaveChainLink() {
	// Generate source and destination key
	ext := suite.GetRandomProfile()
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		link      types.ChainLink
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "invalid chain address packed value returns error",
			link: types.ChainLink{
				Address: testutil.NewAny(ext.privKey),
				Proof: types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("")))),
					hex.EncodeToString([]byte("")),
				),
				ChainConfig:  types.NewChainConfig("cosmos"),
				CreationTime: time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			},
			shouldErr: true,
		},
		{
			name: "invalid chain address returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address("", "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
				),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "invalid proof returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					"wrong",
				),
				types.NewChainConfig("cosmos"),
				time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "link already existing returns error",
			store: func(ctx sdk.Context) {
				address := "cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"
				profile := testutil.ProfileFromAddr(address)
				suite.ak.SetAccount(ctx, profile)

				link := types.NewChainLink(
					address,
					types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
					types.NewProof(
						ext.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
						hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				suite.Require().NoError(suite.k.SaveChainLink(ctx, link))
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
				),
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
					types.NewProof(
						ext.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
						hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
			},
		},
		{
			name: "invalid user returns error",
			link: types.NewChainLink(
				"",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
				),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "user with no profile returns error",
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
				),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			shouldErr: true,
		},
		{
			name: "valid conditions return no error",
			store: func(ctx sdk.Context) {
				profile := testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")
				err := suite.k.SaveProfile(ctx, profile)
				suite.Require().NoError(err)
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(ext.GetAddress().String(), "cosmos"),
				types.NewProof(
					ext.GetPubKey(),
					testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
					hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
				),
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
					types.NewProof(
						ext.GetPubKey(),
						testutil.SingleSignatureProtoFromHex(hex.EncodeToString(ext.Sign([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))),
						hex.EncodeToString([]byte("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")),
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))

				// Check the additional keys
				store := ctx.KVStore(suite.storeKey)
				suite.Require().True(store.Has(types.ChainLinkOwnerKey(
					"cosmos",
					ext.GetAddress().String(),
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				)))
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

			err := suite.k.SaveChainLink(ctx, tc.link)
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

func (suite *KeeperTestSuite) TestKeeper_GetChainLink() {
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		owner     string
		chainName string
		address   string
		expFound  bool
	}{
		{
			name:      "non existent link returns false",
			owner:     "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			chainName: "cosmos",
			address:   "cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
			expFound:  false,
		},
		{
			name: "existent link returns true",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(suite.storeKey)
				key := types.ChainLinksStoreKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos", "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
				link := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns", "cosmos"),
					types.NewProof(secp256k1.GenPrivKey().PubKey(), testutil.SingleSignatureProtoFromHex("1234"), "706c61696e5f74657874"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				store.Set(key, types.MustMarshalChainLink(suite.cdc, link))
			},
			owner:     "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			chainName: "cosmos",
			address:   "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
			expFound:  true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			ctx, _ := suite.ctx.CacheContext()
			if tc.store != nil {
				tc.store(ctx)
			}

			_, found := suite.k.GetChainLink(ctx, tc.owner, tc.chainName, tc.address)
			if tc.expFound {
				suite.Require().True(found)
			} else {
				suite.Require().False(found)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_DeleteChainLink() {
	account := testutil.GetChainLinkAccount("cosmos", "cosmos")
	testCases := []struct {
		name  string
		store func(ctx sdk.Context)
		link  types.ChainLink
		check func(ctx sdk.Context)
	}{
		{
			name: "different user does not delete link",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
				err := suite.k.SaveChainLink(ctx, account.GetBech32ChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
				suite.Require().NoError(err)
			},
			link: types.NewChainLink(
				"cosmos1xvvggrlgjkhu4rva9j500rc52za2smxhluvftc",
				account.Bech32Address(),
				account.Bech32Proof("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasChainLink(ctx,
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos",
					account.Bech32Address().GetValue(),
				))
			},
		},
		{
			name: "different chain name does not delete link",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
				err := suite.k.SaveChainLink(ctx, account.GetBech32ChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
				suite.Require().NoError(err)
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address(account.Bech32Address().GetValue(), "cosmos"),
				account.Bech32Proof("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
				types.NewChainConfig("likecoin"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasChainLink(ctx,
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos",
					account.Bech32Address().GetValue(),
				))
			},
		},
		{
			name: "different external address does not delete the link",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
				err := suite.k.SaveChainLink(ctx, account.GetBech32ChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
				suite.Require().NoError(err)
			},
			link: types.NewChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				types.NewBech32Address("cosmos1xvvggrlgjkhu4rva9j500rc52za2smxhluvftc", "cosmos"),
				account.Bech32Proof("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x"),
				types.NewChainConfig("cosmos"),
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().True(suite.k.HasChainLink(ctx,
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos",
					account.Bech32Address().GetValue(),
				))
			},
		},
		{
			name: "proper data delete the link",
			store: func(ctx sdk.Context) {
				suite.Require().NoError(suite.k.SaveProfile(ctx, testutil.ProfileFromAddr("cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x")))
				err := suite.k.SaveChainLink(ctx, account.GetBech32ChainLink(
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				))
				suite.Require().NoError(err)
			},
			link: account.GetBech32ChainLink(
				"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
				time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
			),
			check: func(ctx sdk.Context) {
				suite.Require().False(suite.k.HasChainLink(ctx,
					"cosmos19xz3mrvzvp9ymgmudhpukucg6668l5haakh04x",
					"cosmos",
					account.Bech32Address().GetValue(),
				))

				// Check the additional keys
				store := ctx.KVStore(suite.storeKey)
				suite.Require().False(store.Has(types.ChainLinkOwnerKey(
					"cosmos",
					account.Bech32Address().GetValue(),
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)))
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

			suite.k.DeleteChainLink(ctx, tc.link)
			if tc.check != nil {
				tc.check(ctx)
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
					types.NewProof(key, testutil.SingleSignatureProtoFromHex("1234"), "706c61696e74657874"),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
				)

				link = types.NewChainLink(
					user,
					types.NewBech32Address("cosmos1xcy3els9ua75kdm783c3qu0rfa2eplesldfevn", "cosmos"),
					types.NewProof(key, testutil.SingleSignatureProtoFromHex("1234"), "706c61696e74657874"),
					types.NewChainConfig("cosmos"),
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				store.Set(
					types.ChainLinksStoreKey(link.User, link.ChainConfig.Name, link.GetAddressData().GetValue()),
					suite.cdc.MustMarshal(&link),
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
