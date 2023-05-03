package v6_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/app"
	"github.com/desmos-labs/desmos/v5/testutil/storetesting"

	v6 "github.com/desmos-labs/desmos/v5/x/posts/legacy/v6"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

type mockSubspace struct {
	ps types.Params
}

func newMockSubspace(ps types.Params) mockSubspace {
	return mockSubspace{ps: ps}
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps paramstypes.ParamSet) {
	*ps.(*types.Params) = ms.ps
}

func (ms mockSubspace) SetParamSet(ctx sdk.Context, ps paramstypes.ParamSet) {
	panic("unimplemented")
}

func TestMigrate(t *testing.T) {
	cdc, _ := app.MakeCodecs()

	keys := sdk.NewKVStoreKeys(types.StoreKey)

	testCases := []struct {
		name          string
		setupSubspace func() mockSubspace
		shouldErr     bool
		check         func(ctx sdk.Context)
	}{
		{
			name: "params migrates properly",
			setupSubspace: func() mockSubspace {
				return newMockSubspace(types.DefaultParams())
			},
			shouldErr: false,
			check: func(ctx sdk.Context) {
				store := ctx.KVStore(keys[types.StoreKey])

				var params types.Params
				bz := store.Get(types.ParamsKey)
				require.NoError(t, cdc.Unmarshal(bz, &params))
				require.Equal(t, types.DefaultParams(), params)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			ctx := storetesting.BuildContext(keys, nil, nil)

			mockSubspace := tc.setupSubspace()

			err := v6.MigrateStore(ctx, keys[types.StoreKey], mockSubspace, cdc)
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
