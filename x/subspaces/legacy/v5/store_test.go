package v4_test

import (
	"testing"

	v5 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v5"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"
)

func TestMigrateStore(t *testing.T) {
	cdc, legacyAminoCdc := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(authtypes.StoreKey, paramstypes.StoreKey, types.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Build the params keeper
	paramsKeeper := paramskeeper.NewKeeper(
		cdc, legacyAminoCdc, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey],
	)
	// Build the auth keeper
	authKeeper := authkeeper.NewAccountKeeper(
		cdc,
		keys[authtypes.StoreKey],
		paramsKeeper.Subspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
	)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "accounts of all the users inside groups are created properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.GroupMemberStoreKey(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"), []byte{0x01})
			},
			check: func(ctx sdk.Context) {
				userAcc, _ := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				require.True(t, authKeeper.HasAccount(ctx, userAcc))
			},
		},
		{
			name: "accounts of all the users having subspace permissions are created properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				permission := types.NewUserPermission(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm", types.NewPermissions("edit subspace"))
				store.Set(types.UserPermissionStoreKey(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"), cdc.MustMarshal(&permission))
			},
			check: func(ctx sdk.Context) {
				userAcc, _ := sdk.AccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
				require.True(t, authKeeper.HasAccount(ctx, userAcc))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, tKeys, memKeys)
			if tc.store != nil {
				tc.store(ctx)
			}

			err := v5.MigrateStore(ctx, keys[types.StoreKey], authKeeper)
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
