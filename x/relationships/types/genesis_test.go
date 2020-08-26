package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
	"github.com/stretchr/testify/require"
)

func TestNewGenesis(t *testing.T) {
	usersRelationships := map[string][]sdk.AccAddress{}

	expGenState := types.GenesisState{
		UsersRelationships: usersRelationships,
	}

	actualGenState := types.NewGenesisState(usersRelationships)
	require.Equal(t, expGenState, actualGenState)
}

func TestValidateGenesis(t *testing.T) {
	var user, err = sdk.AccAddressFromBech32("cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47")
	require.NoError(t, err)

	otherUser, err := sdk.AccAddressFromBech32("cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns")
	require.NoError(t, err)

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
			name: "Genesis with invalid relationship return error",
			genesis: types.GenesisState{
				UsersRelationships: map[string][]sdk.AccAddress{
					user.String():      {sdk.AccAddress{}},
					otherUser.String(): {user},
				},
			},
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.GenesisState{
				UsersRelationships: map[string][]sdk.AccAddress{
					user.String():      {otherUser},
					otherUser.String(): {user},
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
