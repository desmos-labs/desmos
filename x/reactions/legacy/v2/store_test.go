package v2_test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/address"
	"github.com/cosmos/cosmos-sdk/runtime"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/testutil/storetesting"
	profileskeeper "github.com/desmos-labs/desmos/v6/x/profiles/keeper"
	profilestypes "github.com/desmos-labs/desmos/v6/x/profiles/types"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v6/app"
	postskeeper "github.com/desmos-labs/desmos/v6/x/posts/keeper"
	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	v2 "github.com/desmos-labs/desmos/v6/x/reactions/legacy/v2"
	"github.com/desmos-labs/desmos/v6/x/reactions/types"
	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	relationshipstypes "github.com/desmos-labs/desmos/v6/x/relationships/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v6/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

func TestMigrateStore(t *testing.T) {
	cdc, legacyAminoCdc := app.MakeCodecs()

	// Build all the necessary keys
	keys := storetypes.NewKVStoreKeys(paramstypes.StoreKey, relationshipstypes.StoreKey, subspacestypes.StoreKey, poststypes.StoreKey, types.StoreKey)
	tKeys := storetypes.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := storetypes.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	authKeeper := authkeeper.NewAccountKeeper(
		cdc,
		runtime.NewKVStoreService(keys[authtypes.StoreKey]),
		authtypes.ProtoBaseAccount,
		app.GetMaccPerms(),
		address.NewBech32Codec("cosmos"),
		"cosmos",
		authtypes.NewModuleAddress("gov").String(),
	)

	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey], nil, nil, "authority")
	rk := relationshipskeeper.NewKeeper(cdc, keys[relationshipstypes.StoreKey], sk)
	ak := profileskeeper.NewKeeper(cdc, legacyAminoCdc, keys[profilestypes.StoreKey], authKeeper, rk, nil, nil, nil, authtypes.NewModuleAddress("gov").String())
	pk := postskeeper.NewKeeper(cdc, keys[poststypes.StoreKey], ak, sk, rk, authtypes.NewModuleAddress("gov").String())

	testCases := []struct {
		name      string
		store     func(ctx sdk.Context)
		shouldErr bool
		check     func(ctx sdk.Context)
	}{
		{
			name: "next registered reaction ids are set for existing subspaces",
			store: func(ctx sdk.Context) {
				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"This is a test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					2,
					"This is another test subspace",
					"This is anoter test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.Equal(t, uint32(1), types.GetRegisteredReactionIDFromBytes(store.Get(types.NextRegisteredReactionIDStoreKey(1))))
				require.Equal(t, uint32(1), types.GetRegisteredReactionIDFromBytes(store.Get(types.NextRegisteredReactionIDStoreKey(2))))
			},
		},
		{
			name: "params are set for existing subspaces",
			store: func(ctx sdk.Context) {
				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					1,
					"This is a test subspace",
					"This is a test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))

				sk.SaveSubspace(ctx, subspacestypes.NewSubspace(
					2,
					"This is another test subspace",
					"This is anoter test subspace",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					"cosmos1s0he0z3g92zwsxdj83h0ky9w463sx7gq9mqtgn",
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
				))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var params types.SubspaceReactionsParams

				cdc.MustUnmarshal(store.Get(types.SubspaceReactionsParamsStoreKey(1)), &params)
				require.Equal(t, types.DefaultReactionsParams(1), params)

				cdc.MustUnmarshal(store.Get(types.SubspaceReactionsParamsStoreKey(2)), &params)
				require.Equal(t, types.DefaultReactionsParams(2), params)
			},
		},
		{
			name: "next reaction ids are set properly for existing posts",
			store: func(ctx sdk.Context) {
				pk.SavePost(ctx, poststypes.NewPost(
					1,
					0,
					1,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))

				pk.SavePost(ctx, poststypes.NewPost(
					2,
					0,
					2,
					"External ID",
					"This is a text",
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
					1,
					nil,
					nil,
					nil,
					poststypes.REPLY_SETTING_EVERYONE,
					time.Date(2020, 1, 1, 12, 00, 00, 000, time.UTC),
					nil,
					"cosmos13t6y2nnugtshwuy0zkrq287a95lyy8vzleaxmd",
				))
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])
				require.Equal(t, uint32(1), types.GetRegisteredReactionIDFromBytes(store.Get(types.NextReactionIDStoreKey(1, 1))))
				require.Equal(t, uint32(1), types.GetRegisteredReactionIDFromBytes(store.Get(types.NextReactionIDStoreKey(2, 2))))
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

			err := v2.MigrateStore(ctx, keys[types.StoreKey], sk, pk, cdc)
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
