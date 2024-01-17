package profiles_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v6/x/profiles"

	"github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v6/x/profiles/types"

	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"

	"cosmossdk.io/log"
	"cosmossdk.io/store"
	"cosmossdk.io/store/metrics"
	storetypes "cosmossdk.io/store/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	db "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	subspaceskeeper "github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestBeginBlocker(t *testing.T) {
	// Define store keys
	keys := storetypes.NewKVStoreKeys(
		authtypes.StoreKey, relationshipstypes.StoreKey, types.StoreKey,
	)

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB, log.NewNopLogger(), metrics.NewNoOpMetrics())
	for _, key := range keys {
		ms.MountStoreWithDB(key, storetypes.StoreTypeIAVL, memDB)
	}

	err := ms.LoadLatestVersion()
	require.NoError(t, err)

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	cdc, legacyAmino := app.MakeCodecs()
	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey], nil, nil, "authority")
	rk := relationshipskeeper.NewKeeper(cdc, keys[relationshipstypes.StoreKey], sk)
	ak := authkeeper.NewAccountKeeper(cdc, runtime.NewKVStoreService(keys[authtypes.StoreKey]), authtypes.ProtoBaseAccount, app.GetMaccPerms(), address.NewBech32Codec("cosmos"), "cosmos", authtypes.NewModuleAddress("gov").String())
	k := keeper.NewKeeper(cdc, legacyAmino, keys[types.StoreKey], ak, rk, nil, nil, nil, authtypes.NewModuleAddress("gov").String())

	testCases := []struct {
		name      string
		setupCtx  func(ctx sdk.Context) sdk.Context
		store     func(ctx sdk.Context)
		check     func(ctx sdk.Context)
		expEvents sdk.Events
	}{
		{
			name: "expired links are deleted correctly",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 01, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				address := "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"
				link := types.NewApplicationLink(
					address,
					types.NewData("twitter", "twitteruser"),
					types.ApplicationLinkStateInitialized,
					types.NewOracleRequest(
						0,
						1,
						types.NewOracleRequestCallData("twitter", "calldata"),
						"client_id",
					),
					nil,
					time.Date(2019, 1, 1, 00, 00, 00, 000, time.UTC),
					time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
				)

				require.NoError(t, k.SaveProfile(ctx, profilestesting.ProfileFromAddr(address)))
				require.NoError(t, k.SaveApplicationLink(ctx, link))
			},
			expEvents: sdk.Events{
				sdk.NewEvent(
					types.EventTypeApplicationLinkDeleted,
					sdk.NewAttribute(types.AttributeKeyUser, "cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773"),
					sdk.NewAttribute(types.AttributeKeyApplicationName, "twitter"),
					sdk.NewAttribute(types.AttributeKeyApplicationUsername, "twitteruser"),
					sdk.NewAttribute(
						types.AttributeKeyApplicationLinkExpirationTime,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC).Format(time.RFC3339),
					),
				),
			},
			check: func(ctx sdk.Context) {
				require.False(t, k.HasApplicationLink(ctx,
					"cosmos10nsdxxdvy9qka3zv0lzw8z9cnu6kanld8jh773",
					"twitter",
					"twitteruser",
				))
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

			// Reset the events
			ctx = ctx.WithEventManager(sdk.NewEventManager())

			// Run the BeginBlocker
			profiles.BeginBlocker(ctx, k)

			// Check the events and storage
			require.Equal(t, tc.expEvents, ctx.EventManager().Events())
			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
