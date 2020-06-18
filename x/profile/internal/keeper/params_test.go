package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeeper_SetParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := models.NewNameSurnameLenParams(&min, &max)
	monikerParams := models.NewMonikerLenParams(&min, &max)
	bioParams := models.NewBioLenParams(max)

	params := models.NewParams(nsParams, monikerParams, bioParams)

	k.SetParams(ctx, params)

	actualParams := k.GetParams(ctx)

	require.Equal(t, params, actualParams)
}

func TestKeeper_GetParams(t *testing.T) {
	min := sdk.NewInt(2)
	max := sdk.NewInt(1000)
	ctx, k := SetupTestInput()
	nsParams := models.NewNameSurnameLenParams(&min, &max)
	monikerParams := models.NewMonikerLenParams(&min, &max)
	bioParams := models.NewBioLenParams(max)
	params := models.NewParams(nsParams, monikerParams, bioParams)

	tests := []struct {
		name      string
		params    *models.Params
		expParams *models.Params
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
