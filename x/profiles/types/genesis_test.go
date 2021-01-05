package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func TestValidateGenesis(t *testing.T) {
	date, err := time.Parse(time.RFC3339, "2010-10-02T12:10:00.000Z")
	require.NoError(t, err)

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
			name: "Genesis with invalid profile returns error (empty DTag)",
			genesis: types.NewGenesisState(
				[]types.Profile{
					types.NewProfile(
						"",
						"",
						"",
						types.NewPictures("", ""),
						date,
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					), // An empty tag should return an error
				},
				nil,
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Invalid params returns error",
			genesis: types.NewGenesisState(
				[]types.Profile{
					types.NewProfile(
						"custom_dtag1",
						"",
						"biography",
						types.NewPictures("https://test.com/profile-pic", "https://test.com/cover-pic"),
						date,
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
				nil,
				types.NewParams(
					types.NewMonikerParams(sdk.NewInt(-1), sdk.NewInt(10)),
					types.DefaultDtagParams(),
					types.DefaultMaxBioLength,
				),
			),
			shouldError: true,
		},
		{
			name: "Invalid dTag requests returns error",
			genesis: types.NewGenesisState(
				[]types.Profile{
					types.NewProfile(
						"custom_dtag1",
						"",
						"biography",
						types.NewPictures(
							"https://test.com/profile-pic",
							"https://test.com/cover-pic",
						),
						date,
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
				types.DefaultParams(),
			),
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.NewGenesisState(
				[]types.Profile{
					types.NewProfile(
						"custom_dtag1",
						"",
						"biography",
						types.NewPictures(
							"https://test.com/profile-pic",
							"https://test.com/cover-pic",
						),
						date,
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
					),
				},
				[]types.DTagTransferRequest{
					types.NewDTagTransferRequest(
						"dtag",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
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
