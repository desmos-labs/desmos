package types_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

func TestValidateGenesis(t *testing.T) {
	tests := []struct {
		Genesis  *types.GenesisState
		ExpError error
	}{
		{
			Genesis:  types.NewGenesisState(0, nil),
			ExpError: fmt.Errorf("invalid default session length: 0"),
		},
		{
			Genesis:  types.NewGenesisState(1, nil),
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
