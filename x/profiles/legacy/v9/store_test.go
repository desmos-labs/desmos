package v9_test

import (
	"encoding/hex"
	"testing"
	"time"

	v9 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v9"
	v9types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v9/types"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/store"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
)

func buildContext(
	keys map[string]*sdk.KVStoreKey, tKeys map[string]*sdk.TransientStoreKey, memKeys map[string]*sdk.MemoryStoreKey,
) sdk.Context {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	for _, key := range keys {
		cms.MountStoreWithDB(key, sdk.StoreTypeIAVL, db)
	}
	for _, tKey := range tKeys {
		cms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, db)
	}
	for _, memKey := range memKeys {
		cms.MountStoreWithDB(memKey, sdk.StoreTypeMemory, nil)
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
	paramsKeeper := paramskeeper.NewKeeper(
		cdc,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)
	authKeeper := authkeeper.NewAccountKeeper(
		cdc,
		keys[authtypes.StoreKey],
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
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
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store a profile
				profile, err := v9types.NewProfile(
					"john_doe",
					"John Doe",
					"My name if John Doe",
					v9types.Pictures{
						Profile: "",
						Cover:   "",
					},
					time.Date(2020, 1, 1, 0, 0, 0, 00, time.UTC),
					profilestesting.AccountFromAddr("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
				)
				require.NoError(t, err)
				authKeeper.SetAccount(ctx, profile)

				// Store a DTag reference
				kvStore.Set(types.DTagStoreKey("john_doe"), []byte("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Check the profile to make sure it contains the same data
				v10Profile, err := types.NewProfile(
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
				require.Equal(t, v10Profile, profile)

				// Check the DTag association
				value := kvStore.Get(types.DTagStoreKey("john_doe"))
				require.Equal(t, []byte("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"), value)
			},
		},
		{
			name: "chain links are migrated properly - Bech32Address",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)

				chainLink := v9types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v9types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					v9types.NewProof(
						pubKey,
						&v9types.SingleSignature{
							ValueType: v9types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					v9types.NewChainConfig("cosmos"),
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

				// Make sure the signature type has been updated correctly
				var stored types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)

				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)
				require.Equal(t, types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewAddress("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", types.GENERATION_ALGORITHM_COSMOS, types.NewBech32Encoding("cosmos")),
					types.NewProof(
						pubKey,
						&types.SingleSignature{
							ValueType: types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					types.NewChainConfig("cosmos"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "chain links are migrated properly - Base58Address",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)

				chainLink := v9types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v9types.NewBase58Address("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN"),
					v9types.NewProof(
						pubKey,
						&v9types.SingleSignature{
							ValueType: v9types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					v9types.NewChainConfig("solana"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"solana",
						"5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the signature type has been updated correctly
				var stored types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"solana",
					"5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN",
				)), &stored)

				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)
				expected := types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewAddress("5AfetAwZzftP8i5JBNatzWeccfXd4KvKq6TRfAvacFaN", types.GENERATION_ALGORITHM_DO_NOTHING, types.NewBase58Encoding("")),
					types.NewProof(
						pubKey,
						&types.SingleSignature{
							ValueType: types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					types.NewChainConfig("solana"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				require.True(t, expected.Equal(stored))
			},
		},
		{
			name: "chain links are migrated properly - HexAddress",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)

				chainLink := v9types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v9types.NewHexAddress("0x941991947b6ec9f5537bcac30c1295e8154df4cc", "0x"),
					v9types.NewProof(
						pubKey,
						&v9types.SingleSignature{
							ValueType: v9types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					v9types.NewChainConfig("ethereum"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)
				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"ethereum",
						"0x941991947b6ec9f5537bcac30c1295e8154df4cc",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the signature type has been updated correctly
				var stored types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"ethereum",
					"0x941991947b6ec9f5537bcac30c1295e8154df4cc",
				)), &stored)

				signature, err := hex.DecodeString("1234")
				require.NoError(t, err)
				require.Equal(t, types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					types.NewAddress("0x941991947b6ec9f5537bcac30c1295e8154df4cc", types.GENERATION_ALGORITHM_EVM, types.NewHexEncoding("0x", false)),
					types.NewProof(
						pubKey,
						&types.SingleSignature{
							ValueType: types.SIGNATURE_VALUE_TYPE_RAW,
							Signature: signature,
						},
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e7631626574",
					),
					types.NewChainConfig("ethereum"),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
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

			err := v9.MigrateStore(ctx, authKeeper, keys[types.StoreKey], cdc)
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
