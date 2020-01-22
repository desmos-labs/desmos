package app_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/desmos-labs/desmos/app"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	defaultParams := staking.DefaultParams()
	defaultState := staking.DefaultGenesisState()

	app.Init()

	customParams := staking.NewParams(
		defaultParams.UnbondingTime,
		defaultParams.MaxValidators,
		defaultParams.MaxEntries,
		defaultParams.BondDenom,
	)
	expected := staking.NewGenesisState(customParams, defaultState.Validators, defaultState.Delegations)

	customDefaultState := staking.DefaultGenesisState()
	assert.NotEqual(t, expected, customDefaultState)
	assert.Equal(t, "stake", defaultState.Params.BondDenom)
	assert.Equal(t, "4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e", customDefaultState.Params.BondDenom)
}
