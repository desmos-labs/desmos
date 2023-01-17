package posts_test

import (
	"testing"
	"time"

	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	profileskeeper "github.com/desmos-labs/desmos/v4/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v4/x/profiles/types"

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
	"github.com/desmos-labs/desmos/v4/x/posts"
	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"
	"github.com/desmos-labs/desmos/v4/x/posts/types"
	poststypes "github.com/desmos-labs/desmos/v4/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

func TestEndBlocker(t *testing.T) {
	// Define store keys
	keys := sdk.NewMemoryStoreKeys(
		paramstypes.StoreKey, authtypes.StoreKey,
		profilestypes.StoreKey, relationshipstypes.StoreKey,
		subspacestypes.StoreKey, types.StoreKey,
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
	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey], nil, nil)
	rk := relationshipskeeper.NewKeeper(cdc, keys[relationshipstypes.StoreKey], sk)
	authKeeper := authkeeper.NewAccountKeeper(cdc, keys[authtypes.StoreKey], pk.Subspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, app.GetMaccPerms())
	profilesKeeper := profileskeeper.NewKeeper(cdc, legacyAmino, keys[profilestypes.StoreKey], pk.Subspace(profilestypes.DefaultParamsSpace), authKeeper, rk, nil, nil, nil)
	keeper := postskeeper.NewKeeper(cdc, keys[poststypes.StoreKey], pk.Subspace(types.DefaultParamsSpace), profilesKeeper, sk, rk)

	testCases := []struct {
		name     string
		setupCtx func(ctx sdk.Context) sdk.Context
		store    func(ctx sdk.Context)
		check    func(ctx sdk.Context)
	}{
		{
			name: "active poll is not tallied before time",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC))
			},
			store: func(ctx sdk.Context) {
				attachment := poststypes.NewAttachment(1, 1, 1, poststypes.NewPoll(
					"What animal is best?",
					[]poststypes.Poll_ProvidedAnswer{
						poststypes.NewProvidedAnswer("Cat", nil),
						poststypes.NewProvidedAnswer("Dog", nil),
						poststypes.NewProvidedAnswer("No one of the above", nil),
					},

					// Just 1 nanosecond before time
					time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC),

					true,
					false,
					nil,
				))
				keeper.SaveAttachment(ctx, attachment)
				keeper.InsertActivePollQueue(ctx, attachment)

				keeper.SaveUserAnswer(ctx, poststypes.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				keeper.SaveUserAnswer(ctx, poststypes.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
			},
			check: func(ctx sdk.Context) {
				kvStore := ctx.KVStore(keys[poststypes.StoreKey])
				endTime := time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC)
				require.True(t, kvStore.Has(poststypes.ActivePollQueueKey(1, 1, 1, endTime)))
			},
		},
		{
			name: "active poll is tallied after time",
			setupCtx: func(ctx sdk.Context) sdk.Context {
				return ctx.WithBlockTime(time.Date(2020, 1, 1, 12, 00, 00, 001, time.UTC))
			},
			store: func(ctx sdk.Context) {
				attachment := poststypes.NewAttachment(1, 1, 1, poststypes.NewPoll(
					"What animal is best?",
					[]poststypes.Poll_ProvidedAnswer{
						poststypes.NewProvidedAnswer("Cat", nil),
						poststypes.NewProvidedAnswer("Dog", nil),
						poststypes.NewProvidedAnswer("No one of the above", nil),
					},

					// Just 1 nanosecond before time
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),

					true,
					false,
					nil,
				))
				keeper.SaveAttachment(ctx, attachment)
				keeper.InsertActivePollQueue(ctx, attachment)

				keeper.SaveUserAnswer(ctx, poststypes.NewUserAnswer(1, 1, 1, []uint32{0, 1}, "cosmos1pmklwgqjqmgc4ynevmtset85uwm0uau90jdtfn"))
				keeper.SaveUserAnswer(ctx, poststypes.NewUserAnswer(1, 1, 1, []uint32{1}, "cosmos1zmqjufkg44ngswgf4vmn7evp8k6h07erdyxefd"))
			},
			check: func(ctx sdk.Context) {
				poll, found := keeper.GetPoll(ctx, 1, 1, 1)
				require.True(t, found)
				require.Equal(t, poststypes.NewPollTallyResults([]poststypes.PollTallyResults_AnswerResult{
					poststypes.NewAnswerResult(0, 1),
					poststypes.NewAnswerResult(1, 2),
					poststypes.NewAnswerResult(2, 0),
				}), poll.FinalTallyResults)

				kvStore := ctx.KVStore(keys[poststypes.StoreKey])
				endTime := time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC)
				require.False(t, kvStore.Has(poststypes.ActivePollQueueKey(1, 1, 1, endTime)))
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

			posts.EndBlocker(ctx, keeper)

			if tc.check != nil {
				tc.check(ctx)
			}
		})
	}
}
