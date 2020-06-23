package types_test

import (
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestNewGenesis(t *testing.T) {
	profiles := models.Profiles{}
	nameSurnameParams := models.NameSurnameLengths{}
	monikerParams := models.MonikerLengths{}
	bioParams := sdk.Int{}
	params := models.NewParams(nameSurnameParams, monikerParams, bioParams)

	expGenState := types.GenesisState{
		Profiles: profiles,
		Params:   params,
	}

	actualGenState := types.NewGenesisState(profiles, params)
	require.Equal(t, expGenState, actualGenState)
}

func TestValidateGenesis(t *testing.T) {
	var testPostOwner, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	var name = "name"
	var surname = "surname"
	var bio = "biography"

	var testProfilePic = "https://shorturl.at/adnX3"
	var testCoverPic = "https://shorturl.at/cgpyF"
	var testPictures = models.NewPictures(&testProfilePic, &testCoverPic)

	validNameMinParam := sdk.NewInt(3)
	validNameMaxParam := sdk.NewInt(800)
	validMinMonikerParam := sdk.NewInt(5)
	invalidMaxMonikerParam := sdk.NewInt(-1)

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
				Profiles: models.Profiles{
					models.Profile{
						Name:     &name,
						Surname:  &surname,
						Moniker:  "",
						Bio:      &bio,
						Pictures: testPictures,
						Creator:  testPostOwner,
					},
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid params errors",
			genesis: types.GenesisState{
				Profiles: models.Profiles{
					models.Profile{
						Name:     &name,
						Surname:  &surname,
						Moniker:  "moniker",
						Bio:      &bio,
						Pictures: testPictures,
						Creator:  testPostOwner,
					},
				},
				Params: models.Params{
					NameSurnameLengths: models.NameSurnameLengths{
						MinNameSurnameLen: validNameMinParam,
						MaxNameSurnameLen: validNameMaxParam,
					},
					MonikerLengths: models.MonikerLengths{
						MinMonikerLen: validMinMonikerParam,
						MaxMonikerLen: invalidMaxMonikerParam,
					},
					MaxBioLen: sdk.NewInt(30),
				},
			},
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.GenesisState{
				Profiles: models.Profiles{
					models.Profile{
						Name:     &name,
						Surname:  &surname,
						Moniker:  "moniker",
						Bio:      &bio,
						Pictures: testPictures,
						Creator:  testPostOwner,
					},
				},
				Params: models.Params{
					NameSurnameLengths: models.DefaultNameSurnameLenParams(),
					MonikerLengths:     models.DefaultMonikerLenParams(),
					MaxBioLen:          models.DefaultMaxBioLength,
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
