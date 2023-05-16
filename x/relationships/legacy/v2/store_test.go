package v2_test

import (
	"testing"

	v2 "github.com/desmos-labs/desmos/v5/x/relationships/legacy/v2"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	v1 "github.com/desmos-labs/desmos/v5/x/relationships/legacy/v1"
	"github.com/desmos-labs/desmos/v5/x/relationships/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "relationships are migrated properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				relBz := cdc.MustMarshal(&types.Relationship{
					Creator:      "user",
					Counterparty: "counterparty",
					SubspaceID:   1,
				})
				store.Set(v1.RelationshipsStoreKey("user", "counterparty", 1), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				// Make sure the old keys are deleted properly
				v1Key := v1.RelationshipsStoreKey("user", "counterparty", 1)
				require.False(t, store.Has(v1Key))

				v2Key := types.RelationshipsStoreKey("user", "counterparty", 1)
				require.True(t, store.Has(v2Key))

				var stored types.Relationship
				err := cdc.Unmarshal(store.Get(v2Key), &stored)
				require.NoError(t, err)
				require.Equal(t, types.Relationship{
					Creator:      "user",
					Counterparty: "counterparty",
					SubspaceID:   1,
				}, stored)
			},
		},
		{
			name: "user blocks are migrated properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				blockBz := cdc.MustMarshal(&types.UserBlock{
					Blocker:    "blocker",
					Blocked:    "blocked",
					Reason:     "reason",
					SubspaceID: 1,
				})
				store.Set(v1.UserBlockStoreKey("blocker", "blocked", 1), blockBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				// Make sure the old keys are deleted properly
				v1Key := v1.UserBlockStoreKey("blocker", "blocked", 1)
				require.False(t, store.Has(v1Key))

				v2Key := types.UserBlockStoreKey("blocker", "blocked", 1)
				require.True(t, store.Has(v2Key))

				var stored types.UserBlock
				err := cdc.Unmarshal(store.Get(v2Key), &stored)
				require.NoError(t, err)
				require.Equal(t, types.UserBlock{
					Blocker:    "blocker",
					Blocked:    "blocked",
					Reason:     "reason",
					SubspaceID: 1,
				}, stored)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := testutil.DefaultContext(storeKey, sdk.NewTransientStoreKey("test"))
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v2.MigrateStore(ctx, storeKey, cdc)
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
