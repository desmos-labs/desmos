package v5_test

import (
	"encoding/hex"
	"testing"
	"time"

	v5types "github.com/desmos-labs/desmos/v6/x/profiles/legacy/v5/types"

	v5 "github.com/desmos-labs/desmos/v6/x/profiles/legacy/v5"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	"github.com/desmos-labs/desmos/v6/testutil/storetesting"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
)

func TestMigrateStore(t *testing.T) {
	cdc, legacyAmino := app.MakeCodecs()

	// Build all the necessary keys
	keys := storetypes.NewKVStoreKeys(authtypes.StoreKey, types.StoreKey)
	tKeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	account := profilestesting.GetChainLinkAccount("cosmos", "cosmos")
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "application links owners are added properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store an application link
				linkKey := types.UserApplicationLinkKey("cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773", "twitter", "twitteruser")
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
				kvStore.Set(types.ApplicationLinkClientIDKey("client_id"), linkKey)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Check the application link owner
				key := types.ApplicationLinkOwnerKey(
					"twitter",
					"twitteruser",
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
				)
				require.Equal(t, []byte{0x01}, kvStore.Get(key))
			},
		},
		{
			name: "invalid chain links are deleted",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store the chain link
				signatureValue := []byte("custom value")
				signature := v5types.SingleSignatureData{
					Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
					Signature: signatureValue,
				}
				signatureAny := profilestesting.NewAny(&signature)

				chainLink := v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address("cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f", "cosmos"),
					v5types.Proof{
						PubKey:    account.PubKeyAny(),
						Signature: signatureAny,
						PlainText: hex.EncodeToString(signatureValue),
					},
					v5types.ChainConfig{Name: "cosmos"},
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)

				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the chain link is deleted and the owner key is not added
				require.False(t, kvStore.Has(types.ChainLinkOwnerKey(
					"cosmos",
					"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				)))
				require.False(t, kvStore.Has(types.ChainLinksStoreKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
					"cosmos10clxpupsmddtj7wu7g0wdysajqwp890mva046f",
				)))
			},
		},
		{
			name: "valid chain link owners are added properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store the chain link
				addr, _ := sdk.Bech32ifyAddressBytes("cosmos", account.PubKey().Address())
				chainLink := v5types.NewChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					v5types.NewBech32Address(addr, "cosmos"),
					v5types.NewProof(
						account.PubKey(),
						&v5types.SingleSignatureData{
							Mode:      signing.SignMode_SIGN_MODE_TEXTUAL,
							Signature: account.Sign("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47"),
						},
						hex.EncodeToString([]byte("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")),
					),
					v5types.NewChainConfig(account.ChainName()),
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)

				kvStore.Set(
					types.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						account.Bech32Address().GetValue(),
					),
					cdc.MustMarshal(&chainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				key := types.ChainLinkOwnerKey(
					"cosmos",
					account.Bech32Address().GetValue(),
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
				)
				require.Equal(t, []byte{0x01}, kvStore.Get(key))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v5.MigrateStore(ctx, keys[types.StoreKey], cdc, legacyAmino)
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
