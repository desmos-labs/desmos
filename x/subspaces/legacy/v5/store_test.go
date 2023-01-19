package v5_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"
	v5 "github.com/desmos-labs/desmos/v4/x/subspaces/legacy/v5"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, legacyAmino := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(types.StoreKey, authtypes.StoreKey, paramstypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	// Build the x/auth keeper
	paramsKeeper := paramskeeper.NewKeeper(
		cdc,
		legacyAmino,
		keys[paramstypes.StoreKey],
		tKeys[paramstypes.TStoreKey],
	)
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
			name: "subspace is migrated properly",
			store: func(ctx sdk.Context) {
				oldSubspace := types.NewSubspace(
					1,
					"name",
					"description",
					"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
					"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
					"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
					time.Date(2023, 1, 11, 1, 1, 1, 1, time.UTC),
				)
				ctx.KVStore(keys[types.StoreKey]).Set(types.SubspaceStoreKey(1), cdc.MustMarshal(&oldSubspace))
			},
			check: func(ctx sdk.Context) {
				var newSubspace types.Subspace
				err := cdc.Unmarshal(ctx.KVStore(keys[types.StoreKey]).Get(types.SubspaceStoreKey(1)), &newSubspace)
				require.NoError(t, err)
				require.Equal(t, types.NewSubspace(
					1,
					"name",
					"description",
					"cosmos1cyjzgj9j7d2gdqk78pa0fgvfnlzradat97aek9",
					"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
					"cosmos10ya9y35qkf4puaklx5fs07sxfxqncx9usgsnz6",
					time.Date(2023, 1, 11, 1, 1, 1, 1, time.UTC),
				), newSubspace)
			},
		},
		{
			name: "accounts of all the users inside groups are created properly",
			store: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				store.Set(types.GroupMemberStoreKey(1, 1, "cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm"), []byte{0x01})
			},
			check: func(ctx sdk.Context) {
				userAcc := sdk.MustAccAddressFromBech32("cosmos1a0cj0j6ujn2xap8p40y6648d0w2npytw3xvenm")
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

			err := v5.MigrateStore(ctx, keys[types.StoreKey], cdc, authKeeper)
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
