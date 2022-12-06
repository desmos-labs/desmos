package v6_test

import (
	"encoding/hex"
	"testing"
	"time"

	v8 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v8"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/cosmos/cosmos-sdk/store"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	v9types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v9/types"
	types "github.com/desmos-labs/desmos/v4/x/profiles/types"
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

	pubKey := secp256k1.GenPrivKey().PubKey()
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "chain links are migrated properly - SIGNATURE_VALUE_TYPE_COSMOS_AMINO",
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
						"7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a226a756e6f2d31222c22666565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2230222c2264656e6f6d223a226a756e6f227d5d2c22676173223a2231227d2c226d656d6f223a226465736d6f7331366336307938743876726132377a6a673261726c6364353864636b3963776e37703666777464222c226d736773223a5b5d2c2273657175656e6365223a2230227d",
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
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the signature type has been updated correctly
				var stored v9types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				signature, err := stored.Proof.GetSignature()
				require.NoError(t, err)

				valueType, err := signature.GetValueType()
				require.NoError(t, err)
				require.Equal(t, v9types.SIGNATURE_VALUE_TYPE_COSMOS_AMINO, valueType)
			},
		},
		{
			name: "chain links are migrated properly - SIGNATURE_VALUE_TYPE_COSMOS_DIRECT",
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
						"0aba010a88010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e6412680a2c74657272613166387736767a6835376c357a6532347871386378677432663873753672733064793578716863122c74657272613166387736767a6835376c357a65323478713863786774326638737536727330647935787168631a0a0a05756c756e61120130122d6465736d6f7331677570676e73666776733038776174777466646c34613572393538396375733368756a30617312600a4e0a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2102ac68a389793aef8b4121268090a0822afc62f0de082c54c370feb97952ff6ec212040a020801120e0a0a0a05756c756e6112013010011a09626f6d6261792d313220ab8d0c",
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
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the signature type has been updated correctly
				var stored v9types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				signature, err := stored.Proof.GetSignature()
				require.NoError(t, err)

				valueType, err := signature.GetValueType()
				require.NoError(t, err)
				require.Equal(t, v9types.SIGNATURE_VALUE_TYPE_COSMOS_DIRECT, valueType)
			},
		},
		{
			name: "chain links are migrated properly - SIGNATURE_VALUE_TYPE_RAW",
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
							ValueType: v9types.SIGNATURE_VALUE_TYPE_COSMOS_DIRECT,
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
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the signature type has been updated correctly
				var stored v9types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				signature, err := stored.Proof.GetSignature()
				require.NoError(t, err)

				valueType, err := signature.GetValueType()
				require.NoError(t, err)
				require.Equal(t, v9types.SIGNATURE_VALUE_TYPE_RAW, valueType)
			},
		},
		{
			name: "chain links are migrated properly - multi signature",
			store: func(ctx sdk.Context) {
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
					  "value_type": "SIGNATURE_VALUE_TYPE_COSMOS_DIRECT",
					  "signature": "o/8Z4KeRtihOrG0pyW+0xQXUrRDq2kcA1FhvImGFgjtUG8Nxb10e6kx9m8pHCU6KZcwb0vaBso7jhTHlIy5zfA=="
					},
					{
					  "@type": "/desmos.profiles.v3.SingleSignature",
					  "value_type": "SIGNATURE_VALUE_TYPE_COSMOS_DIRECT",
					  "signature": "o/8Z4KeRtihOrG0pyW+0xQXUrRDq2kcA1FhvImGFgjtUG8Nxb10e6kx9m8pHCU6KZcwb0vaBso7jhTHlIy5zfA=="
					}
				  ]
				}`
				var signatureData v9types.Signature
				err := cdc.UnmarshalInterfaceJSON([]byte(signatureJSON), &signatureData)
				require.NoError(t, err)

				chainLink := v9types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v9types.NewBech32Address("cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs", "cosmos"),
					v9types.NewProof(pubKey, signatureData, "7b226163636f756e745f6e756d626572223a2230222c22636861696e5f6964223a226a756e6f2d31222c22666565223a7b22616d6f756e74223a5b7b22616d6f756e74223a2230222c2264656e6f6d223a226a756e6f227d5d2c22676173223a2231227d2c226d656d6f223a226465736d6f7331366336307938743876726132377a6a673261726c6364353864636b3963776e37703666777464222c226d736773223a5b5d2c2273657175656e6365223a2230227d"),
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
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				var stored v9types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				)), &stored)
				signature, err := stored.Proof.GetSignature()
				require.NoError(t, err)

				valueType, err := signature.GetValueType()
				require.NoError(t, err)
				require.Equal(t, v9types.SIGNATURE_VALUE_TYPE_COSMOS_AMINO, valueType)
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

			err := v8.MigrateStore(ctx, keys[types.StoreKey], cdc, legacyAmino)
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
