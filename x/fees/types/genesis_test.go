package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

func TestValidateGenesis(t *testing.T) {
	testCases := []struct {
		name      string
		genesis   *types.GenesisState
		shouldErr bool
	}{
		{
			name: "invalid params return error",
			genesis: types.NewGenesisState(types.NewParams(
				[]types.MinFee{
					types.NewMinFee("", sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(1)))),
				},
			)),
			shouldErr: true,
		},
		{
			name:      "default genesis returns no error",
			genesis:   types.DefaultGenesisState(),
			shouldErr: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			err := types.ValidateGenesis(tc.genesis)
			if tc.shouldErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
