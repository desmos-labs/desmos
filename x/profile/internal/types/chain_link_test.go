package types_test

import (
	"fmt"
	"testing"

	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/stretchr/testify/require"
)

func TestChainLink_Validate(t *testing.T) {
	tests := []struct {
		name      string
		chainLink types.ChainLink
		expErr    error
	}{
		{
			name:      "empty chain link name returns error",
			chainLink: types.NewChainLink("", "4B06BEE128323FF26A75821637541F3640C280286312C638CAF07399C3961BFE"),
			expErr:    fmt.Errorf("chain name cannot be empty or blank"),
		},
		{
			name:      "invalid chain link hash returns error",
			chainLink: types.NewChainLink("cosmos", "txhash"),
			expErr:    fmt.Errorf("transaction hash of cosmos chain must be a valid sha-256 hash"),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.chainLink.Validate()
			require.Equal(t, test.expErr, actual)
		})
	}
}

func TestChainLink_Equals(t *testing.T) {
	tests := []struct {
		name      string
		chainLink types.ChainLink
		otherCl   types.ChainLink
		expBool   bool
	}{
		{
			name:      "different chainLink returns false",
			chainLink: types.NewChainLink("cosmos", "4B06BEE128323FF26A75821637541F3640C280286312C638CAF07399C3961BFE"),
			otherCl:   types.NewChainLink("desmos", "4B06BEE128323FF26A75821637541F3640C280286312C638CAF07399C3961BFE"),
			expBool:   false,
		},
		{
			name:      "equal chainLink returns true",
			chainLink: types.NewChainLink("desmos", "4B06BEE128323FF26A75821637541F3640C280286312C638CAF07399C3961BFE"),
			otherCl:   types.NewChainLink("desmos", "4B06BEE128323FF26A75821637541F3640C280286312C638CAF07399C3961BFE"),
			expBool:   true,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			actual := test.chainLink.Equals(test.otherCl)
			require.Equal(t, actual, test.expBool)
		})
	}
}
