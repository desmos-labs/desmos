package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/links/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		name        string
		genesis     *types.GenesisState
		shouldError bool
	}{
		{
			name:        "Default Genesis does not error",
			genesis:     types.DefaultGenesisState(),
			shouldError: false,
		},
		{
			name: "Genesis with invalid links returns error",
			genesis: types.NewGenesisState(
				types.PortID,
				[]types.Link{
					types.NewLink(
						"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
						"cosmos1c07g02fjmsl6dcumfsgttjkvnk4n9lxzek0dvn",
					),
					types.NewLink(
						"desmos1tw3jl54lmwn3mq6hjfvl5nsk4q70v34wc9nsyk",
						"cosmos1wnv4pk0ueawnt06dsdpnqmhqrqpwll39ssx6kn",
					),
				},
			),
			shouldError: true,
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
