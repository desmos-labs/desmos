package v300_test

//func TestMigrateStore(t *testing.T) {
//	cdc, _ := app.MakeCodecs()
//	storeKey := sdk.NewKVStoreKey(types.StoreKey)
//	testCases := []struct {
//		name      string
//		store     func(ctx sdk.Context)
//		shouldErr bool
//		check     func(ctx sdk.Context)
//	}{
//		{
//			name: "valid data returns no error",
//			store: func(ctx sdk.Context) {
//				store := ctx.KVStore(storeKey)
//
//				block := v2.UserBlock{
//					Blocker:    "blocker",
//					Blocked:    "blocked",
//					Reason:     "reason",
//					SubspaceID: "",
//				}
//				blockBz := cdc.MustMarshal(&block)
//				store.Set(v2.UserBlockStoreKey(block.Blocker, block.SubspaceID, block.Blocked), blockBz)
//
//				relationship := v2.Relationship{
//					Creator:    "user",
//					Recipient:  "recipient",
//					SubspaceID: "2",
//				}
//				relBz := cdc.MustMarshal(&relationship)
//				store.Set(append(types.RelationshipsStorePrefix, 0x01), relBz)
//			},
//			shouldErr: false,
//			check: func(ctx sdk.Context) {
//				store := ctx.KVStore(storeKey)
//
//				oldBlockKey := v2.UserBlockStoreKey("blocker", "", "blocked")
//				require.False(t, store.Has(oldBlockKey))
//
//				expectedBlock := types.NewUserBlock("blocker", "blocked", "reason", 0)
//				expectedBlockKey := types.UserBlockStoreKey(expectedBlock.Blocker, expectedBlock.SubspaceID, expectedBlock.Blocked)
//				require.True(t, store.Has(expectedBlockKey))
//
//				var storedBlock types.UserBlock
//				cdc.MustUnmarshal(store.Get(expectedBlockKey), &storedBlock)
//				require.Equal(t, expectedBlock, storedBlock)
//
//				oldRelationshipKey := v2.RelationshipsStoreKey("user", "2", "recipient")
//				require.False(t, store.Has(oldRelationshipKey))
//
//				expectedRelationship := types.NewRelationship("user", "recipient", 2)
//				expectedRelationshipKey := types.RelationshipsStoreKey(expectedRelationship.Creator, expectedRelationship.Counterparty, expectedRelationship.SubspaceID)
//				require.True(t, store.Has(expectedRelationshipKey))
//
//				var storedRelationship types.Relationship
//				cdc.MustUnmarshal(store.Get(expectedRelationshipKey), &storedRelationship)
//				require.Equal(t, expectedRelationship, storedRelationship)
//			},
//		},
//	}
//
//	for _, tc := range testCases {
//		tc := tc
//		t.Run(tc.name, func(t *testing.T) {
//			ctx := testutil.DefaultContext(storeKey, sdk.NewTransientStoreKey("test"))
//			if tc.store != nil {
//				tc.store(ctx)
//			}
//
//			err := v3.MigrateStore(ctx, storeKey, cdc)
//			if tc.shouldErr {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//				if tc.check != nil {
//					tc.check(ctx)
//				}
//			}
//		})
//	}
//}
