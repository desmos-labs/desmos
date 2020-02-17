package app_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	defaultState := staking.DefaultGenesisState()
	require.Equal(t, "desmos", defaultState.Params.BondDenom)
}
