package app_test

import (
	"testing"

	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"

	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	defaultState := stakingtypes.DefaultGenesisState()
	require.Equal(t, "desmos", defaultState.Params.BondDenom)
}
