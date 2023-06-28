package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

func TestParams_Validate(t *testing.T) {
	testCases := []struct {
		name      string
		params    types.Params
		shouldErr bool
	}{
		{
			name:      "invalid coins return error",
			params:    types.NewParams(sdk.Coins{sdk.Coin{Denom: "%invalid%", Amount: sdk.NewInt(100)}}),
			shouldErr: true,
		},
		{
			name:      "valid params return no error",
			params:    types.NewParams(sdk.NewCoins(sdk.Coin{Denom: "udsm", Amount: sdk.NewInt(100)})),
			shouldErr: false,
		},
		{
			name:      "default params return no error",
			params:    types.DefaultParams(),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := tc.params.Validate()
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
