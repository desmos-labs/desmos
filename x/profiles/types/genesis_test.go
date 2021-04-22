package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestValidateGenesis(t *testing.T) {
	addr1, _ := sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")

	tests := []struct {
		name        string
		genesis     *types.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Invalid params returns error",
			genesis: types.NewGenesisState(
				nil,
				types.NewParams(
					types.NewMonikerParams(sdk.NewInt(-1), sdk.NewInt(10)),
					types.DefaultDTagParams(),
					types.DefaultMaxBioLength,
				),
			),
			shouldError: true,
		},
		{
			name: "Invalid DTag requests returns error",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"",
						addr1.String(),
					),
				},
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Valid genesis returns no errors",
			genesis: types.NewGenesisState(
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						addr1.String(),
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
					),
				},
				types.DefaultParams(),
			),
			shouldError: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if test.shouldError {
				require.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
