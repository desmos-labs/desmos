package v2_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/testutil/storetesting"
	v2 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v2"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
	subspaceskeeper "github.com/desmos-labs/desmos/v5/x/subspaces/keeper"
	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) *mockSubspace {
	return &mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func (ms *mockSubspace) SetParamSet(ctx sdk.Context, ps paramstypes.ParamSet) {
	ms.ps = *ps.(*types.Params)
}

func TestMigrateStore(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	// Build all the necessary keys
	keys := sdk.NewKVStoreKeys(paramstypes.StoreKey, subspacestypes.StoreKey, types.StoreKey)

	paramsSubspace := newMockSubspace(types.Params{})

	sk := subspaceskeeper.NewKeeper(cdc, keys[subspacestypes.StoreKey], nil, nil, "authority")

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
			ctx := storetesting.BuildContext(keys, nil, nil)
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
