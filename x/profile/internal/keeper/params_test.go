package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_SetParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := types.NewNameSurnameLenParams(min, max)
	monikerParams := types.NewMonikerLenParams(min, max)

	params := types.NewParams(nsParams, monikerParams, max)

	k.SetParams(ctx, params)

	actualParams := k.GetParams(ctx)

	require.Equal(t, params, actualParams)
}

func TestKeeper_GetParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := types.NewNameSurnameLenParams(min, max)
	monikerParams := types.NewMonikerLenParams(min, max)
	params := types.NewParams(nsParams, monikerParams, max)

	tests := []struct {
		name      string
		params    *types.Params
		expParams *types.Params
	}{
		{
			name:      "Returning previously set params",
			params:    &params,
			expParams: &params,
		},
		{
			name:      "Returning nothing",
			params:    nil,
			expParams: nil,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.params != nil {
				k.SetParams(ctx, *test.params)
			}

			if test.expParams != nil {
				require.Equal(t, *test.expParams, k.GetParams(ctx))
			}
		})
	}
}
