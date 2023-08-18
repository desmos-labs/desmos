package v7_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/feegrant"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	v7 "github.com/desmos-labs/desmos/v6/x/subspaces/legacy/v7"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func isExpiringQueueEmpty(ctx sdk.Context, key storetypes.StoreKey) bool {
	store := ctx.KVStore(key)
	expiringStore := prefix.NewStore(store, types.ExpiringAllowanceQueuePrefix)
	iterator := expiringStore.Iterator(nil, nil)
	defer iterator.Close()

	count := 0
	for ; iterator.Valid(); iterator.Next() {
		count++
	}

	return count == 0
}

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	expiration := time.Date(2100, 7, 7, 0, 0, 0, 0, time.UTC)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "user allowance added to expiring queue properly",
			store: func(ctx sdk.Context) {
				grant := types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				)
				ctx.KVStore(keys[types.StoreKey]).Set(types.UserAllowanceKey(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"), cdc.MustMarshal(&grant))
			},
			check: func(ctx sdk.Context) {
				require.True(
					t,
					ctx.KVStore(keys[types.StoreKey]).Has(types.ExpiringAllowanceKey(&expiration, types.UserAllowanceKey(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"))),
				)
			},
		},
		{
			name: "user allowance without expiration skipped properly",
			store: func(ctx sdk.Context) {
				grant := types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewUserGrantee("cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
					},
				)
				ctx.KVStore(keys[types.StoreKey]).Set(types.UserAllowanceKey(1, "cosmos1m0czrla04f7rp3zg7dsgc4kla54q7pc4xt00l5"), cdc.MustMarshal(&grant))

			},
			check: func(ctx sdk.Context) {
				require.True(t, isExpiringQueueEmpty(ctx, keys[types.StoreKey]))
			},
		},
		{
			name: "group allowance added to expiring queue properly",
			store: func(ctx sdk.Context) {
				grant := types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
						Expiration: &expiration,
					},
				)
				ctx.KVStore(keys[types.StoreKey]).Set(types.GroupAllowanceKey(1, 1), cdc.MustMarshal(&grant))
			},
			check: func(ctx sdk.Context) {
				require.True(
					t,
					ctx.KVStore(keys[types.StoreKey]).Has(types.ExpiringAllowanceKey(&expiration, types.GroupAllowanceKey(1, 1))),
				)
			},
		},
		{
			name: "group allowance without expiration skipped properly",
			store: func(ctx sdk.Context) {
				grant := types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", sdk.NewInt(1))),
					},
				)
				ctx.KVStore(keys[types.StoreKey]).Set(types.GroupAllowanceKey(1, 1), cdc.MustMarshal(&grant))
			},
			check: func(ctx sdk.Context) {
				require.True(t, isExpiringQueueEmpty(ctx, keys[types.StoreKey]))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, nil, nil)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v7.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
