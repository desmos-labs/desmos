package v2_test

import (
	"testing"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v2 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v2"
	"github.com/desmos-labs/desmos/v2/x/relationships/types"
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
			name: "valid data is migrated properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				blockBz := cdc.MustMarshal(&v2.UserBlock{
					Blocker:    "blocker",
					Blocked:    "blocked",
					Reason:     "reason",
					SubspaceID: "",
				})
				store.Set(v2.UserBlockStoreKey("blocker", "", "blocked"), blockBz)

				relBz := cdc.MustMarshal(&v2.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "1",
				})
				store.Set(v2.RelationshipsStoreKey("user", "1", "recipient"), relBz)

				relBz = cdc.MustMarshal(&v2.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "2",
				})
				store.Set(v2.RelationshipsStoreKey("user", "2", "recipient"), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				// Make sure all blocks are deleted
				require.False(t, store.Has(v2.UserBlockStoreKey("blocker", "", "blocked")))

				// Make sure all relationships are deleted
				require.False(t, store.Has(v2.RelationshipsStoreKey("user", "1", "recipient")))
				require.False(t, store.Has(v2.RelationshipsStoreKey("user", "1", "recipient")))
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

			// TODO: Get the real account keeper and amino here
			err := v2.MigrateStore(ctx, authkeeper.AccountKeeper{}, storeKey, nil, cdc)
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
