package profiles_test

import (
	"testing"
	"time"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v4/testutil/profilestesting"
	"github.com/desmos-labs/desmos/v4/x/profiles"

	"github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"

	relationshipskeeper "github.com/desmos-labs/desmos/v4/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v4/x/relationships/types"

	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/desmos-labs/desmos/v4/app"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestBeginBlocker(t *testing.T) {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		relationshipstypes.StoreKey, types.StoreKey,
	)
	tKeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)

	// Create an in-memory db
	memDB := db.NewMemDB()
	ms := store.NewCommitMultiStore(memDB)
	for _, key := range keys {
		ms.MountStoreWithDB(key, sdk.StoreTypeIAVL, memDB)
	}
	for _, tKey := range tKeys {
		ms.MountStoreWithDB(tKey, sdk.StoreTypeTransient, memDB)
	}

	err := ms.LoadLatestVersion()
	require.NoError(t, err)

	ctx := sdk.NewContext(ms, tmproto.Header{ChainID: "test-chain"}, false, log.NewNopLogger())
	cdc, legacyAmino := app.MakeCodecs()
	pk := paramskeeper.NewKeeper(cdc, legacyAmino, keys[paramstypes.StoreKey], tKeys[paramstypes.TStoreKey])
	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey])
	rk := relationshipskeeper.NewKeeper(cdc, keys[relationshipstypes.StoreKey], sk)
	ak := authkeeper.NewAccountKeeper(cdc, keys[authtypes.StoreKey], pk.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, app.GetMaccPerms())
	k := keeper.NewKeeper(cdc, legacyAmino, keys[types.StoreKey], pk.Subspace(types.DefaultParamsSpace), ak, rk, nil, nil, nil)

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
