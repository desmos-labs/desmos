package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/posts/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_SetParams(t *testing.T) {
	ctx, k := SetupTestInput()

	params := types.DefaultParams()
	k.SetParams(ctx, params)

	actualParams := k.GetParams(ctx)

	require.Equal(t, params, actualParams)
}

func TestKeeper_GetParams(t *testing.T) {
	ctx, k := SetupTestInput()

	params := types.DefaultParams()
	k.SetParams(ctx, params)

	actualParams := k.GetParams(ctx)

	require.Equal(t, params, actualParams)

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
