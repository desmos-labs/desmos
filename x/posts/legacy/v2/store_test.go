package v2_test

import (
	"testing"
	"time"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v4/app"
	"github.com/desmos-labs/desmos/v4/testutil/storetesting"
	v2 "github.com/desmos-labs/desmos/v4/x/posts/legacy/v2"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, amino := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(paramstypes.StoreKey, subspacestypes.StoreKey, types.StoreKey, authtypes.StoreKey)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	pk := paramskeeper.NewKeeper(cdc, amino, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])
	paramsSubspace := pk.Subspace(types.ModuleName)
	paramsSubspace = paramsSubspace.WithKeyTable(types.ParamKeyTable())

	authKeeper := authkeeper.NewAccountKeeper(cdc, keys[authtypes.StoreKey], pk.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, app.GetMaccPerms())
	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey], authKeeper)

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "next post id is set for existing subspaces",
			store: func(ctx sdk.Context) {
				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"This is a test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))

				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					2,
					"This is another test subspace",
					"This is anoter test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
				))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				postID := types.GetPostIDFromBytes(store.Get(types.NextPostIDStoreKey(1)))
				require.Equal(t, uint64(1), postID)

				postID = types.GetPostIDFromBytes(store.Get(types.NextPostIDStoreKey(2)))
				require.Equal(t, uint64(1), postID)
			},
		},
		{
			name:      "module params are set properly",
			shouldErr: false,
			check: func(ctx sdk.Context) {
				var params types.Params
				paramsSubspace.GetParamSet(ctx, &params)
				require.Equal(t, types.DefaultParams(), params)
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

			err := v2.MigrateStore(ctx, keys[types.StoreKey], paramsSubspace, sk)
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
