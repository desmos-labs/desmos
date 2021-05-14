package types_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		name    string
		genesis *types.GenesisState
		expPass bool
	}{
		{
			name:    "Default Genesis does not error",
			genesis: types.DefaultGenesisState(),
			expPass: true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			if !test.expPass {
				require.Error(t, types.ValidateGenesis(test.genesis))
			} else {
				require.NoError(t, types.ValidateGenesis(test.genesis))
			}
		})
	}
}
