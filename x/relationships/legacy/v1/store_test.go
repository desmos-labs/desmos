package v1_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"
	profilesv1beta1 "github.com/desmos-labs/desmos/v3/x/profiles/legacy/v1beta1"
	v1 "github.com/desmos-labs/desmos/v3/x/relationships/legacy/v1"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
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

				blockBz := cdc.MustMarshal(&profilesv1beta1.UserBlock{
					Blocker:    "blocker",
					Blocked:    "blocked",
					Reason:     "reason",
					SubspaceID: "",
				})
				store.Set(profilesv1beta1.UserBlockStoreKey("blocker", "", "blocked"), blockBz)

				relBz := cdc.MustMarshal(&profilesv1beta1.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "",
				})
				store.Set(profilesv1beta1.RelationshipsStoreKey("user", "", "recipient"), relBz)

				relBz = cdc.MustMarshal(&profilesv1beta1.Relationship{
					Creator:    "user",
					Recipient:  "recipient",
					SubspaceID: "2",
				})
				store.Set(profilesv1beta1.RelationshipsStoreKey("user", "2", "recipient"), relBz)
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(storeKey)

				// Make sure user blocks are migrated properly
				expectedBlock := types.NewUserBlock("blocker", "blocked", "reason", 0)
				expectedBlockKey := types.UserBlockStoreKey(expectedBlock.Blocker, expectedBlock.Blocked, expectedBlock.SubspaceID)
				require.True(t, store.Has(expectedBlockKey))

				var storedBlock types.UserBlock
				cdc.MustUnmarshal(store.Get(expectedBlockKey), &storedBlock)
				require.Equal(t, expectedBlock, storedBlock)

				// Make sure relationships with subspace 0 are not migrated
				expectedRelationshipKey := types.RelationshipsStoreKey("user", "recipient", 0)
				require.False(t, store.Has(expectedRelationshipKey))

				// Make sure other relationships are migrated properly
				expectedRelationship := types.NewRelationship("user", "recipient", 2)
				expectedRelationshipKey = types.RelationshipsStoreKey(expectedRelationship.Creator, expectedRelationship.Counterparty, expectedRelationship.SubspaceID)
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

			pk := profilesv1beta1.NewKeeper(storeKey, cdc)
			err := v1.MigrateStore(ctx, pk, storeKey, cdc)
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
