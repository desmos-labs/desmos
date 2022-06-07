package v6_test

import (
	"testing"
	"time"

	"github.com/desmos-labs/desmos/v3/testutil"
	v6 "github.com/desmos-labs/desmos/v3/x/profiles/legacy/v6"
	profilestypes "github.com/desmos-labs/desmos/v3/x/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	olderChainAccount := testutil.GetChainLinkAccount("cosmos", "cosmos")
	newerChainAccount := testutil.GetChainLinkAccount("cosmos", "cosmos")

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "default external addresses are set properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Store the chain links
				newerChainLink := newerChainAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2020, 1, 3, 00, 00, 00, 000, time.UTC),
				)

				kvStore.Set(
					profilestypes.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						newerChainAccount.Bech32Address().Value,
					),
					cdc.MustMarshal(&newerChainLink),
				)

				olderChainLink := olderChainAccount.GetBech32ChainLink(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					time.Date(2020, 1, 2, 00, 00, 00, 000, time.UTC),
				)

				kvStore.Set(
					profilestypes.ChainLinksStoreKey(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos",
						olderChainAccount.Bech32Address().Value,
					),
					cdc.MustMarshal(&olderChainLink),
				)
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				key := profilestypes.DefaultExternalAddressKey(
					"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					"cosmos",
				)
				require.Equal(t, []byte(olderChainAccount.Bech32Address().Value), kvStore.Get(key))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := testutil.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v6.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
