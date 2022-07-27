package v4_test

import (
	"encoding/hex"
	"testing"
	"time"

	v4 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4"
	types2 "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4/types"
	v5types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v5/types"

	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

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
	v4types "github.com/desmos-labs/desmos/v4/x/profiles/legacy/v4/types"
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
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, profilestypes.StoreKey)
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

	// Build common data
	pubKey := profilestesting.PubKeyFromBech32("cosmospub1addwnpepqvryxhhqhw52c4ny5twtfzf3fsrjqhx0x5cuya0fylw0wu0eqptykeqhr4d")
	pubKeyAny := profilestesting.NewAny(pubKey)

	addressAny := profilestesting.NewAny(&v4types.Bech32Address{
		Value:  "cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
		Prefix: "cosmos",
	})

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "profiles are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Store a profile
				profile, err := types2.NewProfile(
					"john_doe",
					"John Doe",
					"My name if John Doe",
					v4types.Pictures{
						Profile: "",
						Cover:   "",
					},
					time.Date(2020, 1, 1, 0, 0, 0, 00, time.UTC),
					profilestesting.AccountFromAddr("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
				)
				require.NoError(t, err)
				authKeeper.SetAccount(ctx, profile)

				// Store a DTag reference
				kvStore.Set(profilestypes.DTagStoreKey("john_doe"), []byte("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Check the profile to make sure it contains the same data
				v2Profile, err := v5types.NewProfile(
					"john_doe",
					"John Doe",
					"My name if John Doe",
					v5types.NewPictures("", ""),
					time.Date(2020, 1, 1, 0, 0, 0, 00, time.UTC),
					profilestesting.AccountFromAddr("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
				)
				require.NoError(t, err)

				sdkAddr, err := sdk.AccAddressFromBech32("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf")
				require.NoError(t, err)

				account := authKeeper.GetAccount(ctx, sdkAddr)
				profile, ok := account.(*v5types.Profile)
				require.True(t, ok)
				require.Equal(t, v2Profile, profile)

				// Check the DTag association
				value := kvStore.Get(profilestypes.DTagStoreKey("john_doe"))
				require.Equal(t, []byte("cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"), value)
			},
		},
		{
			name: "DTag transfer requests are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Store a DTag transfer request
				kvStore.Set(
					profilestypes.DTagTransferRequestStoreKey("cosmos13vsgmgs9tjktnnc6pkln7pm4jswxmeajrqc4xd", "cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf"),
					cdc.MustMarshal(&v4types.DTagTransferRequest{
						DTagToTrade: "john_doe",
						Sender:      "cosmos13vsgmgs9tjktnnc6pkln7pm4jswxmeajrqc4xd",
						Receiver:    "cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf",
					}),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Check the migrated DTag transfer request
				var stored v5types.DTagTransferRequest
				cdc.MustUnmarshal(kvStore.Get(profilestypes.DTagTransferRequestStoreKey(
					"cosmos13vsgmgs9tjktnnc6pkln7pm4jswxmeajrqc4xd",
					"cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf",
				)), &stored)

				require.Equal(t, v5types.NewDTagTransferRequest(
					"john_doe",
					"cosmos13vsgmgs9tjktnnc6pkln7pm4jswxmeajrqc4xd",
					"cosmos1nejmx335u222dj6lg7qjqrufchkpazu8e0semf",
				), stored)
			},
		},
		{
			name: "application links are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Store an application link
				linkKey := profilestypes.UserApplicationLinkKey("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser")
				kvStore.Set(
					linkKey,
					cdc.MustMarshal(&v5types.ApplicationLink{
						User:  "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
						Data:  v5types.NewData("twitter", "twitteruser"),
						State: v5types.ApplicationLinkStateInitialized,
						OracleRequest: v5types.NewOracleRequest(
							0,
							1,
							v5types.NewOracleRequestCallData("twitter", "calldata"),
							"client_id",
						),
						Result:       nil,
						CreationTime: time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					}),
				)

				// Store an application link client id
				kvStore.Set(profilestypes.ApplicationLinkClientIDKey("client_id"), linkKey)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Check the application links
				linkKey := profilestypes.UserApplicationLinkKey(
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				)

				var stored v5types.ApplicationLink
				cdc.MustUnmarshal(kvStore.Get(linkKey), &stored)
				require.Equal(t, v5types.NewApplicationLink(
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
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				), stored)

				// Check the application link client id
				require.Equal(t, linkKey, kvStore.Get(profilestypes.ApplicationLinkClientIDKey("client_id")))
			},
		},
		{
			name: "leftover application client id keys are deleted properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])
				kvStore.Set(v4types.ApplicationLinkClientIDKey("client_id"), []byte("client_id_value"))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])
				require.False(t, kvStore.Has(v4types.ApplicationLinkClientIDKey("client_id")))
				require.False(t, kvStore.Has(profilestypes.ApplicationLinkClientIDKey("client_id")))
			},
		},
		{
			name: "chain link is migrated properly (text signature)",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Store the chain link
				kvStore.Set(
					v4types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
					),
					cdc.MustMarshal(&v4types.ChainLink{
						User:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						Address: addressAny,
						Proof: v4types.Proof{
							PubKey:    pubKeyAny,
							Signature: "7369676E6174757265",
							PlainText: "74657874",
						},
						ChainConfig:  v4types.ChainConfig{Name: "cosmos"},
						CreationTime: time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					}),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				signatureValue, err := hex.DecodeString("7369676E6174757265")
				require.NoError(t, err)
				signature := v5types.SingleSignatureData{
					Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
					Signature: signatureValue,
				}
				signatureAny := profilestesting.NewAny(&signature)

				var stored v5types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(profilestypes.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				)), &stored)
				require.Equal(t, v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					v5types.Proof{
						PubKey:    pubKeyAny,
						Signature: signatureAny,
						PlainText: "74657874",
					},
					v5types.ChainConfig{Name: "cosmos"},
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "chain link is migrated properly (direct sign tx signature)",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Build the signature
				bz, err := cdc.Marshal(&tx.SignDoc{
					BodyBytes:     nil,
					AuthInfoBytes: nil,
					ChainId:       "test-chain",
					AccountNumber: 1,
				})
				require.NoError(t, err)
				plainTextValue := hex.EncodeToString(bz)

				// Store the chain link
				kvStore.Set(
					v4types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
					),
					cdc.MustMarshal(&v4types.ChainLink{
						User:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						Address: addressAny,
						Proof: v4types.Proof{
							PubKey:    pubKeyAny,
							Signature: "7369676E6174757265",
							PlainText: plainTextValue,
						},
						ChainConfig:  v4types.ChainConfig{Name: "cosmos"},
						CreationTime: time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					}),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				bz, err := cdc.Marshal(&tx.SignDoc{
					BodyBytes:     nil,
					AuthInfoBytes: nil,
					ChainId:       "test-chain",
					AccountNumber: 1,
				})
				require.NoError(t, err)
				plainTextValue := hex.EncodeToString(bz)

				signatureValue, err := hex.DecodeString("7369676E6174757265")
				require.NoError(t, err)

				signatureAny := profilestesting.NewAny(&v5types.SingleSignatureData{
					Mode:      signing.SignMode_SIGN_MODE_DIRECT,
					Signature: signatureValue,
				})

				var stored v5types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(profilestypes.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				)), &stored)
				require.Equal(t, v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					v5types.Proof{
						PubKey:    pubKeyAny,
						Signature: signatureAny,
						PlainText: plainTextValue,
					},
					v5types.ChainConfig{Name: "cosmos"},
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "chain link is migrated properly (amino sign tx signature)",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Build the signature
				bz, err := legacyAmino.MarshalJSON(&legacytx.StdSignDoc{
					AccountNumber: 1,
					Sequence:      1,
					TimeoutHeight: 0,
					ChainID:       "test-chain",
					Memo:          "This is my memo",
					Fee:           nil,
					Msgs:          nil,
				})
				require.NoError(t, err)
				plainTextValue := hex.EncodeToString(bz)

				// Store the chain link
				kvStore.Set(
					v4types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
					),
					cdc.MustMarshal(&v4types.ChainLink{
						User:    "cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						Address: addressAny,
						Proof: v4types.Proof{
							PubKey:    pubKeyAny,
							Signature: "7369676E6174757265",
							PlainText: plainTextValue,
						},
						ChainConfig:  v4types.ChainConfig{Name: "cosmos"},
						CreationTime: time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
					}),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				bz, err := legacyAmino.MarshalJSON(&legacytx.StdSignDoc{
					AccountNumber: 1,
					Sequence:      1,
					TimeoutHeight: 0,
					ChainID:       "test-chain",
					Memo:          "This is my memo",
					Fee:           nil,
					Msgs:          nil,
				})
				require.NoError(t, err)
				plainTextValue := hex.EncodeToString(bz)

				signatureValue, err := hex.DecodeString("7369676E6174757265")
				require.NoError(t, err)

				signature := v5types.SingleSignatureData{
					Mode:      signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON,
					Signature: signatureValue,
				}
				signatureAny := profilestesting.NewAny(&signature)

				var stored v5types.ChainLink
				cdc.MustUnmarshal(kvStore.Get(profilestypes.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				)), &stored)
				require.Equal(t, v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					v5types.Proof{
						PubKey:    pubKeyAny,
						Signature: signatureAny,
						PlainText: plainTextValue,
					},
					v5types.ChainConfig{Name: "cosmos"},
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				), stored)
			},
		},
		{
			name: "user blocks and relationships are deleted properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Store a user block
				kvStore.Set(
					v4types.UserBlockStoreKey("blocker", "", "blocked"),
					cdc.MustMarshal(&v4types.UserBlock{
						Blocker:    "blocker",
						Blocked:    "blocked",
						Reason:     "reason",
						SubspaceID: "",
					}),
				)

				// Store some relationships
				kvStore.Set(
					v4types.RelationshipsStoreKey("user", "1", "recipient"),
					cdc.MustMarshal(&v4types.Relationship{
						Creator:    "user",
						Recipient:  "recipient",
						SubspaceID: "1",
					}),
				)
				kvStore.Set(
					v4types.RelationshipsStoreKey("user", "2", "recipient"),
					cdc.MustMarshal(&v4types.Relationship{
						Creator:    "user",
						Recipient:  "recipient",
						SubspaceID: "2",
					}),
				)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[profilestypes.StoreKey])

				// Make sure all blocks are deleted
				require.False(t, kvStore.Has(v4types.UserBlockStoreKey("blocker", "", "blocked")))

				// Make sure all relationships are deleted
				require.False(t, kvStore.Has(v4types.RelationshipsStoreKey("user", "1", "recipient")))
				require.False(t, kvStore.Has(v4types.RelationshipsStoreKey("user", "1", "recipient")))
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

			err := v4.MigrateStore(ctx, authKeeper, keys[profilestypes.StoreKey], legacyAmino, cdc)
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
