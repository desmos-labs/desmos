package v300_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v2/app"
	v230 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v230"
	v300 "github.com/desmos-labs/desmos/v2/x/profiles/legacy/v300"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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
			name: "valid data returns no error",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				block := v230.UserBlock{
					Blocker:  "blocker",
					Blocked:  "blocked",
					Reason:   "reason",
					Subspace: "",
				}
				blockBz := cdc.MustMarshal(&block)
				store.Set(v230.UserBlockStoreKey(block.Blocker, block.Subspace, block.Blocked), blockBz)

				relationship := v230.Relationship{
					Creator:   "user",
					Recipient: "recipient",
					Subspace:  "2",
				}
				relBz := cdc.MustMarshal(&relationship)
				store.Set(append(types.RelationshipsStorePrefix, 0x01), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				oldBlockKey := v230.UserBlockStoreKey("blocker", "", "blocked")
				require.False(t, store.Has(oldBlockKey))

				expectedBlock := types.NewUserBlock("blocker", "blocked", "reason", 0)
				expectedBlockKey := types.UserBlockStoreKey(expectedBlock.Blocker, expectedBlock.Subspace, expectedBlock.Blocked)
				require.True(t, store.Has(expectedBlockKey))

				var storedBlock types.UserBlock
				cdc.MustUnmarshal(store.Get(expectedBlockKey), &storedBlock)
				require.Equal(t, expectedBlock, storedBlock)

				oldRelationshipKey := v230.RelationshipsStoreKey("user", "2", "recipient")
				require.False(t, store.Has(oldRelationshipKey))

				expectedRelationship := types.NewRelationship("user", "recipient", 2)
				expectedRelationshipKey := types.RelationshipsStoreKey(expectedRelationship.Creator, expectedRelationship.Subspace, expectedRelationship.Recipient)
				require.True(t, store.Has(expectedRelationshipKey))

				var storedRelationship types.Relationship
				cdc.MustUnmarshal(store.Get(expectedRelationshipKey), &storedRelationship)
				require.Equal(t, expectedRelationship, storedRelationship)
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

			err := v300.MigrateStore(ctx, storeKey, cdc)
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
