package subspaces_test

import (
	"testing"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"cosmossdk.io/store"
	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/feegrant"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	"github.com/desmos-labs/desmos/v6/x/subspaces"
	"github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestBeginBlocker(t *testing.T) {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		types.StoreKey,
	)

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	err := ms.LoadLatestVersion()
	require.NoError(t, err)

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	cdc, _ := app.MakeCodecs()

	keeper := keeper.NewKeeper(cdc, keys[types.StoreKey], nil, nil, "authority")

	testCases := []struct {
		name     string
		setupCtx func(ctx sdk.Context) sdk.Context
		store    func(ctx sdk.Context)
		check    func(ctx sdk.Context)
	}{
		{
			name: "allowance is not expired before time",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				// Just 1 nanosecond after time
				expiration := time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC)

				keeper.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", math.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				expiration := time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC)
				key := types.GroupAllowanceKey(1, 1)
				require.True(t, kvStore.Has(key))
				require.True(t, kvStore.Has(types.ExpiringAllowanceKey(&expiration, key)))
			},
		},
		{
			name: "allowance is expired after time",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC))
			},
			store: func(ctx sdk.Context) {
				// Just 1 nanosecond before time
				expiration := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)

				keeper.SaveGrant(ctx, types.NewGrant(
					1,
					"cosmos1x5pjlvufs4znnhhkwe8v4tw3kz30f3lxgwza53",
					types.NewGroupGrantee(1),
					&feegrant.BasicAllowance{
						SpendLimit: sdk.NewCoins(sdk.NewCoin("test", math.NewInt(1))),
						Expiration: &expiration,
					},
				))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[types.StoreKey])
				expiration := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
				key := types.GroupAllowanceKey(1, 1)
				require.False(t, kvStore.Has(key))
				require.False(t, kvStore.Has(types.ExpiringAllowanceKey(&expiration, key)))
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx, _ := ctx.CacheContext()
			if tc.setupCtx != nil {
				ctx = tc.setupCtx(ctx)
			}
			if tc.store != nil {
				tc.store(ctx)
			}

			subspaces.BeginBlocker(ctx, keeper)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
