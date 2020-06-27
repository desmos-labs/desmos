package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestNewGenesis(t *testing.T) {
	profiles := types.Profiles{}
	nameSurnameParams := types.MonikerParams{}
	monikerParams := types.DtagParams{}
	bioParams := sdk.Int{}
	params := types.NewParams(nameSurnameParams, monikerParams, bioParams)

	expGenState := types.GenesisState{
		Profiles: profiles,
		Params:   params,
	}

	actualGenState := types.NewGenesisState(profiles, params)
	require.Equal(t, expGenState, actualGenState)
}

func TestValidateGenesis(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	tests := []struct {
		name        string
		genesis     types.GenesisState
		shouldError bool
	}{
		{
			name:        "DefaultGenesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid profile errors",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("", user, date), // An empty tag should return an error
				),
				Params: types.DefaultParams(),
			},
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("custom_dtag1", user, date).
						WithBio(newStrPtr("biography")).
						WithPictures(
							newStrPtr("https://test.com/profile-pic"),
							newStrPtr("https://test.com/cover-pic"),
						),
				),
				Params: types.DefaultParams(),
			},
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
