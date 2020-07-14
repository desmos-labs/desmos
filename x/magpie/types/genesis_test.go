package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/stretchr/testify/require"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		Genesis  types.GenesisState
		ExpError error
	}{
		{
			Genesis:  types.GenesisState{DefaultSessionLength: -1},
			ExpError: fmt.Errorf("invalid default session length: -1"),
		},
		{
			Genesis:  types.GenesisState{DefaultSessionLength: 0},
			ExpError: fmt.Errorf("invalid default session length: 0"),
		},
		{
			Genesis:  types.GenesisState{DefaultSessionLength: 1},
			ExpError: nil,
		},
		{
			Genesis:  types.DefaultGenesisState(),
			ExpError: nil,
		},
	}

	for index, test := range tests {
		test := test
		t.Run(fmt.Sprintf("Test %d", index), func(t *testing.T) {
			require.Equal(t, test.ExpError, types.ValidateGenesis(test.Genesis))
		})
	}
}
