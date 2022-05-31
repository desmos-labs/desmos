package v3_test

import (
	"testing"

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

				// Check the permissions
				var group types.UserGroup
				cdc.MustUnmarshal(kvStore.Get(v2.GroupStoreKey(1, 1)), &group)
				require.Equal(t, types.CombinePermissions(poststypes.PermissionWrite, poststypes.PermissionModerateContent), group.Permissions)
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

				addr, err := sdk.AccAddressFromBech32("cosmos12e7ejq92sma437d3svemgfvl8sul8lxfs69mjv")
				require.NoError(t, err)

				// Check the permissions
				var stored types.UserPermission
				cdc.MustUnmarshal(kvStore.Get(types.UserPermissionStoreKey(1, addr)), &stored)
				require.Equal(t, types.NewUserPermission(
					1,
					addr.String(),
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
