package v3_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"
	v3 "github.com/desmos-labs/desmos/v4/x/reactions/legacy/v3"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "duplicated reactions does not exist works properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
			},
		},
		{
			name: "delete duplicated reactions properly -- free text",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewFreeTextValue("test"), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.False(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
			},
		},
		{
			name: "delete duplicated reactions properly -- registered reaction",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				reaction := types.NewReaction(1, 1, 1, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 1), cdc.MustMarshal(&reaction))

				duplicatedReaction := types.NewReaction(1, 1, 2, types.NewRegisteredReactionValue(1), "author")
				store.Set(types.ReactionStoreKey(1, 1, 2), cdc.MustMarshal(&duplicatedReaction))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.True(t, store.Has(types.ReactionStoreKey(1, 1, 1)))
				require.False(t, store.Has(types.ReactionStoreKey(1, 1, 2)))
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

			err := v3.MigrateStore(ctx, keys[types.StoreKey], cdc)
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
