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
	nameSurnameParams := models.NameSurnameLenParams{}
	monikerParams := models.MonikerLenParams{}
	bioParams := models.BioLenParams{}

	expGenState := types.GenesisState{
		Profiles:             profiles,
		NameSurnameLenParams: nameSurnameParams,
		MonikerLenParams:     monikerParams,
		BioLenParams:         bioParams,
	}

	actualGenState := types.NewGenesisState(profiles, nameSurnameParams, monikerParams, bioParams)

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
	invalidNameMinParam := sdk.NewInt(1)
	validNameMaxParam := sdk.NewInt(800)
	invalidNameMaxParam := sdk.NewInt(1800)
	validMinMonikerParam := sdk.NewInt(5)
	validMaxMonikerParam := sdk.NewInt(50)
	invalidMinMonikerParam := sdk.NewInt(1)
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
			name: "Genesis with invalid profile's name/surname min params",
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &invalidNameMinParam,
					MaxNameSurnameLen: &validNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &validMinMonikerParam,
					MaxMonikerLen: &validMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(30),
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid profile's name/surname max params",
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &validNameMinParam,
					MaxNameSurnameLen: &invalidNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &validMinMonikerParam,
					MaxMonikerLen: &validMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(30),
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid profile's moniker min params",
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &validNameMinParam,
					MaxNameSurnameLen: &validNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &invalidMinMonikerParam,
					MaxMonikerLen: &validMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(30),
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid profile's moniker max negative params",
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &validNameMinParam,
					MaxNameSurnameLen: &validNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &validNameMinParam,
					MaxMonikerLen: &invalidMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(30),
				},
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid profile's bio params",
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &validNameMinParam,
					MaxNameSurnameLen: &validNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &validNameMinParam,
					MaxMonikerLen: &validMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(-50),
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
				NameSurnameLenParams: models.NameSurnameLenParams{
					MinNameSurnameLen: &validNameMinParam,
					MaxNameSurnameLen: &validNameMaxParam,
				},
				MonikerLenParams: models.MonikerLenParams{
					MinMonikerLen: &validMinMonikerParam,
					MaxMonikerLen: &validMaxMonikerParam,
				},
				BioLenParams: models.BioLenParams{
					MaxBioLen: sdk.NewInt(30),
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
