package v6_test

import (
	"encoding/hex"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	v5types "github.com/desmos-labs/desmos/v6/x/profiles/legacy/v5/types"
	v6 "github.com/desmos-labs/desmos/v6/x/profiles/legacy/v6"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	dbm "github.com/cosmos/cosmos-db"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

func buildContext(
	keys map[string]*storetypes.KVStoreKey, tKeys map[string]*storetypes.TransientStoreKey, memKeys map[string]*storetypes.MemoryStoreKey,
) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	for _, key := range keys {
		cms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		cms.MountStoreWithDB(tKey, storetypes.StoreTypeTransient, db)
	}
	for _, memKey := range memKeys {
		cms.MountStoreWithDB(memKey, storetypes.StoreTypeMemory, nil)
	}

	err := cms.LoadLatestVersion()
	if err != nil {
		panic(err)
	}

	return sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger())
}

func TestMigrateStore(t *testing.T) {
	cdc, legacyAmino := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Build the x/auth keeper
	authKeeper := authkeeper.NewAccountKeeper(
		cdc,
		keys[authtypes.StoreKey],
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	pubKey := secp256k1.GenPrivKey().PubKey()
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "profiles are migrated properly",
			store: func(ctx sdk.Context) {
				// Store a profile
				profile, err := v5types.NewProfile(
					"john_doe",
					"John Doe",
					"My name if John Doe",
					v5types.Pictures{
						Profile: "",
						Cover:   "",
					},
					time.Date(2020, 1, 1, 0, 0, 0, 00, time.UTC),
					profilestesting.AccountFromAddr("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
				)
				require.NoError(t, err)
				authKeeper.SetAccount(ctx, profile)
			},
			check: func(ctx sdk.Context) {
				// Check the profile to make sure it contains the same data
				v7Profile, err := types.NewProfile(
					"john_doe",
					"John Doe",
					"My name if John Doe",
					types.NewPictures("", ""),
					time.Date(2020, 1, 1, 0, 0, 0, 00, time.UTC),
					profilestesting.AccountFromAddr("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
				)
				require.NoError(t, err)

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf")
				require.NoError(t, err)

				account := authKeeper.GetAccount(ctx, sdkAddr)
				profile, ok := account.(*types.Profile)
				require.True(t, ok)
				require.Equal(t, v7Profile, profile)
			},
		},
		{
			name: "application links are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store an application link
				link := v5types.NewApplicationLink(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					v5types.NewData("twitter", "twitteruser"),
					v5types.ApplicationLinkStateInitialized,
					v5types.NewOracleRequest(
						0,
						1,
						v5types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(types.UserApplicationLinkKey("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser"), cdc.MustMarshal(&link))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Check the application links
				linkKey := types.UserApplicationLinkKey(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				)

				var stored types.ApplicationLink
				cdc.MustUnmarshal(kvStore.Get(linkKey), &stored)
				require.Equal(t, types.NewApplicationLink(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2021, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
				), stored)

				// Check the application link expiration time
				require.Equal(t, []byte("client_id"), kvStore.Get(types.ApplicationLinkExpiringTimeKey(
					time.Date(2022, 1, 1, 00, 00, 00, 000, time.UTC),
					"client_id",
				)))
			},
		},
		{
			name: "chain links are migrated properly - single signature",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				sig, err := hex.DecodeString("1234")
				require.NoError(t, err)

				chainLink := v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					v5types.NewProof(
						pubKey,
						&v5types.SingleSignatureData{
							Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
							Signature: sig,
						},
						"plain_text",
					),
					v5types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				var stored types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				require.Equal(t, types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					types.NewProof(pubKey, profilestesting.SingleSignatureFromHex("1234"), "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "chain links are migrated properly - multi signature",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				var signatureJSON = `{
				  "@type": "/desmos.profiles.v2.MultiSignatureData",
				  "bit_array": {
					"extra_bits_stored": 3,
					"elems": "wA=="
				  },
				  "signatures": [
					{
					  "@type": "/desmos.profiles.v2.SingleSignatureData",
					  "mode": "SIGN_MODE_LEGACY_AMINO_JSON",
					  "signature": "J/xFZ4GKKYA+wT9ClATHExyBiswZVPUS89caM3nn7HQdJd6LFC9hFRZSsG73iq7/1YcHAj5ujfvpjJkBhQFkdg=="
					},
					{
					  "@type": "/desmos.profiles.v2.SingleSignatureData",
					  "mode": "SIGN_MODE_LEGACY_AMINO_JSON",
					  "signature": "k5TIZjDnr7lhiZrdj8GiEdFLjMOHAsU8qnAYUVV/NYMsEeEVENpNZ2V4oZs0KGUxdUdUmytL14zfgJ2vpVBB9w=="
					}
				  ]
				}`
				var signatureData v5types.SignatureData
				err := cdc.UnmarshalInterfaceJSON([]byte(signatureJSON), &signatureData)
				require.NoError(t, err)

				chainLink := v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					v5types.NewProof(pubKey, signatureData, "plain_text"),
					v5types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				var signatureJSON = `{
				  "@type": "/desmos.profiles.v3.CosmosMultiSignature",
				  "bit_array": {
					"extra_bits_stored": 3,
					"elems": "wA=="
				  },
				  "signatures": [
					{
					  "@type": "/desmos.profiles.v3.SingleSignature",
					  "value_type": "SIGNATURE_VALUE_TYPE_COSMOS_AMINO",
					  "signature": "J/xFZ4GKKYA+wT9ClATHExyBiswZVPUS89caM3nn7HQdJd6LFC9hFRZSsG73iq7/1YcHAj5ujfvpjJkBhQFkdg=="
					},
					{
					  "@type": "/desmos.profiles.v3.SingleSignature",
					  "value_type": "SIGNATURE_VALUE_TYPE_COSMOS_AMINO",
					  "signature": "k5TIZjDnr7lhiZrdj8GiEdFLjMOHAsU8qnAYUVV/NYMsEeEVENpNZ2V4oZs0KGUxdUdUmytL14zfgJ2vpVBB9w=="
					}
				  ]
				}`
				var signature types.Signature
				err := cdc.UnmarshalInterfaceJSON([]byte(signatureJSON), &signature)
				require.NoError(t, err)

				var stored types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				require.Equal(t, types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					types.NewProof(pubKey, signature, "plain_text"),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "default external addresses are set properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				sig, err := hex.DecodeString("1234")
				require.NoError(t, err)

				chainLink := v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					v5types.NewProof(
						pubKey,
						&v5types.SingleSignatureData{
							Mode:      signing.SignMode_SIGN_MODE_DIRECT,
							Signature: sig,
						},
						"plain_text",
					),
					v5types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				key := types.DefaultExternalAddressKey("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47", "cosmos")
				require.Equal(t, []byte("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs"), kvStore.Get(key))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := buildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v6.MigrateStore(ctx, authKeeper, keys[types.StoreKey], legacyAmino, cdc)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				if tc.check != nil {
					tc.check(ctx)
				}
			}
		})
	}
}
