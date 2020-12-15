package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

func TestValidateGenesis(t *testing.T) {
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
			name: "Genesis with invalid relationship returns error",
			genesis: types.NewGenesisState(
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"",
						"",
					),
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
			),
			shouldError: true,
		},
		{
			name: "Genesis with invalid users blocks return error",
			genesis: types.NewGenesisState(
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
			),
			shouldError: true,
		},
		{
			name: "Valid Genesis returns no errors",
			genesis: types.NewGenesisState(
				[]types.Relationship{
					types.NewRelationship(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewRelationship(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
				[]types.UserBlock{
					types.NewUserBlock(
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
					types.NewUserBlock(
						"cosmos1y54exmx84cqtasvjnskf9f63djuuj68p7hqf47",
						"cosmos1cjf97gpzwmaf30pzvaargfgr884mpp5ak8f7ns",
						"reason",
						"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
					),
				},
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
