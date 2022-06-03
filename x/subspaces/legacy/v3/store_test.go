package v3_test

import (
	"testing"
	"time"

	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"

	v3 "github.com/desmos-labs/desmos/v3/x/subspaces/legacy/v3"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/app"
	"github.com/desmos-labs/desmos/v3/testutil"
	v2 "github.com/desmos-labs/desmos/v3/x/subspaces/legacy/v2"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "section data are set properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				subspace := types.NewSubspace(
					1,
					"Test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				)
				kvStore.Set(types.SubspaceStoreKey(1), cdc.MustMarshal(&subspace))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Check the next section id
				initialSectionID := types.GetSectionIDFromBytes(kvStore.Get(types.NextSectionIDStoreKey(1)))
				require.Equal(t, uint32(1), initialSectionID)

				var section types.Section
				cdc.MustUnmarshal(kvStore.Get(types.SectionStoreKey(1, 0)), &section)
				require.Equal(t, types.DefaultSection(1), section)
			},
		},
		{
			name: "groups permissions are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				group := v2.NewUserGroup(
					1,
					1,
					"Test group",
					"",
					v2.PermissionWrite|v2.PermissionModerateContent,
				)
				kvStore.Set(v2.GroupStoreKey(group.SubspaceID, group.ID), cdc.MustMarshal(&group))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the old key does not exist
				require.False(t, kvStore.Has(v2.GroupStoreKey(1, 1)))

				// Check the permissions
				var group types.UserGroup
				cdc.MustUnmarshal(kvStore.Get(types.GroupStoreKey(1, types.RootSectionID, 1)), &group)
				require.Equal(t, types.NewUserGroup(
					1,
					types.RootSectionID,
					1,
					"Test group",
					"",
					types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionModerateContent),
				), group)
			},
		},
		{
			name: "user permissions are migrated properly",
			store: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				addr, err := sdk.AccAddressFromBech32("cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")
				require.NoError(t, err)

				kvStore.Set(v2.UserPermissionStoreKey(1, addr), v2.MarshalPermission(v2.PermissionEverything))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])

				// Make sure the old key does not exist
				user, err := sdk.AccAddressFromBech32("cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")
				require.NoError(t, err)
				require.False(t, kvStore.Has(v2.UserPermissionStoreKey(1, user)))

				// Check the permissions
				var stored types.UserPermission
				cdc.MustUnmarshal(kvStore.Get(types.UserPermissionStoreKey(1, 0, "cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")), &stored)
				require.Equal(t, types.NewUserPermission(
					1,
					types.RootSectionID,
					"cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv",
					types.NewPermissions(types.PermissionEverything),
				), stored)
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
