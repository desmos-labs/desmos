package types_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"github.com/desmos-labs/desmos/x/profiles/types/common"
	"github.com/stretchr/testify/require"
)

func TestNewGenesis(t *testing.T) {
	profiles := types.Profiles{}
	nameSurnameParams := types.MonikerParams{}
	monikerParams := types.DtagParams{}
	bioParams := sdk.Int{}
	params := types.NewParams(nameSurnameParams, monikerParams, bioParams)

	relationships := types.Relationships{}
	usersRelationshipIDs := map[string]types.RelationshipIDs{}

	expGenState := types.GenesisState{
		Profiles:           profiles,
		Params:             params,
		Relationships:      relationships,
		UsersRelationships: usersRelationshipIDs,
	}

	actualGenState := types.NewGenesisState(profiles, params, relationships, usersRelationshipIDs)
	require.Equal(t, expGenState, actualGenState)
}

func TestValidateGenesis(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	timeZone, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	date := time.Date(2010, 10, 02, 12, 10, 00, 00, timeZone)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

	monoRelationship := types.NewMonodirectionalRelationship(user, otherUser)
	biRelationship := types.NewBiDirectionalRelationship(user, otherUser, types.Sent)

	relationshipIDsMap := map[string]types.RelationshipIDs{
		user.String():      {monoRelationship.ID, biRelationship.ID},
		otherUser.String(): {biRelationship.ID},
	}

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
			name: "Genesis with invalid relationship return error",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("custom_dtag1", user, date).
						WithBio(common.NewStrPtr("biography")).
						WithPictures(
							common.NewStrPtr("https://test.com/profile-pic"),
							common.NewStrPtr("https://test.com/cover-pic"),
						),
				),
				Params: types.DefaultParams(),
				Relationships: types.Relationships{
					monoRelationship,
					types.NewBiDirectionalRelationship(sdk.AccAddress{}, monoRelationship.Receiver, types.Sent),
				},
				UsersRelationships: relationshipIDsMap,
			},
			shouldError: true,
		},
		{
			name: "Genesis with invalid relationshipIDs return error",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("custom_dtag1", user, date).
						WithBio(common.NewStrPtr("biography")).
						WithPictures(
							common.NewStrPtr("https://test.com/profile-pic"),
							common.NewStrPtr("https://test.com/cover-pic"),
						),
				),
				Params: types.DefaultParams(),
				Relationships: types.Relationships{
					monoRelationship,
					biRelationship,
				},
				UsersRelationships: map[string]types.RelationshipIDs{
					user.String():      {monoRelationship.ID, types.RelationshipID("")},
					otherUser.String(): {biRelationship.ID},
				},
			},
			shouldError: true,
		},
		{
			name: "Invalid params returns error",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("custom_dtag1", user, date).
						WithBio(common.NewStrPtr("biography")).
						WithPictures(
							common.NewStrPtr("https://test.com/profile-pic"),
							common.NewStrPtr("https://test.com/cover-pic"),
						),
				),
				Params: types.NewParams(types.NewMonikerParams(sdk.NewInt(-1), sdk.NewInt(10)), types.DefaultDtagParams(), types.DefaultMaxBioLength),
			},
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.GenesisState{
				Profiles: types.NewProfiles(
					types.NewProfile("custom_dtag1", user, date).
						WithBio(common.NewStrPtr("biography")).
						WithPictures(
							common.NewStrPtr("https://test.com/profile-pic"),
							common.NewStrPtr("https://test.com/cover-pic"),
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
