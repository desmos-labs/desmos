package v5_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	v6 "github.com/desmos-labs/desmos/v6/x/subspaces/legacy/v6"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
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
			name: "non default group is skipped properly",
			store: func(ctx sdk.Context) {
				oldGroup := types.NewUserGroup(
					1,
					1,
					1,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				)
				ctx.KVStore(keys[types.StoreKey]).Set(
					types.GroupStoreKey(1, oldGroup.SectionID, oldGroup.ID),
					cdc.MustMarshal(&oldGroup),
				)
			},
			check: func(ctx sdk.Context) {
				var newGroup types.UserGroup
				cdc.MustUnmarshal(ctx.KVStore(keys[types.StoreKey]).
					Get(types.GroupStoreKey(1, 1, 1)), &newGroup)

				require.Equal(t, types.NewUserGroup(
					1,
					1,
					1,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				), newGroup)
			},
		},
		{
			name: "non-moved default group is migrated properly",
			store: func(ctx sdk.Context) {
				oldGroup := types.NewUserGroup(
					1,
					types.RootSectionID,
					types.DefaultGroupID,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				)
				ctx.KVStore(keys[types.StoreKey]).Set(
					types.GroupStoreKey(1, oldGroup.SectionID, oldGroup.ID),
					cdc.MustMarshal(&oldGroup),
				)
			},
			check: func(ctx sdk.Context) {
				var newGroup types.UserGroup
				cdc.MustUnmarshal(ctx.KVStore(keys[types.StoreKey]).
					Get(types.GroupStoreKey(1, types.RootSectionID, types.DefaultGroupID)), &newGroup)

				require.Equal(t, types.NewUserGroup(
					1,
					types.RootSectionID,
					types.DefaultGroupID,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				), newGroup)
			},
		},
		{
			name: "moved default group is migrated properly",
			store: func(ctx sdk.Context) {
				oldGroup := types.NewUserGroup(
					1,
					1,
					types.DefaultGroupID,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				)
				ctx.KVStore(keys[types.StoreKey]).Set(
					types.GroupStoreKey(1, oldGroup.SectionID, oldGroup.ID),
					cdc.MustMarshal(&oldGroup),
				)
			},
			check: func(ctx sdk.Context) {
				var newGroup types.UserGroup
				cdc.MustUnmarshal(ctx.KVStore(keys[types.StoreKey]).
					Get(types.GroupStoreKey(1, types.RootSectionID, types.DefaultGroupID)), &newGroup)

				require.Equal(t, types.NewUserGroup(
					1,
					types.RootSectionID,
					types.DefaultGroupID,
					"default group",
					"description",
					types.NewPermissions(types.NewPermission(types.PermissionEverything)),
				), newGroup)
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
