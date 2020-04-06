package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestNewGenesis(t *testing.T) {
	profiles := types.Profiles{}
	expGenState := types.GenesisState{Profiles: profiles}

	actualGenState := types.NewGenesisState(profiles)

	require.Equal(t, expGenState, actualGenState)
}

func TestValidateGenesis(t *testing.T) {
	var testPostOwner, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	var name = "name"
	var surname = "surname"
	var bio = "biography"

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
			name: "Genesis with invalid account errors",
			genesis: types.GenesisState{
				Profiles: types.Profiles{
					types.Profile{
						Name:     &name,
						Surname:  &surname,
						Moniker:  "",
						Bio:      &bio,
						Pictures: &testPictures,
						Creator:  testPostOwner,
					},
				},
			},
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.GenesisState{
				Profiles: types.Profiles{
					types.Profile{
						Name:     &name,
						Surname:  &surname,
						Moniker:  "moniker",
						Bio:      &bio,
						Pictures: &testPictures,
						Creator:  testPostOwner,
					},
				},
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
