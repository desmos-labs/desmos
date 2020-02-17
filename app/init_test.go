package app_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	defaultState := staking.DefaultGenesisState()
	assert.Equal(t, "desmos", defaultState.Params.BondDenom)
}
