package app

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
)

// Init initializes the application, overriding the default genesis states that should be changed
func Init() {
	staking.DefaultGenesisState = stakingGenesisState
}

// stakingGenesisState returns the default genesis state for the staking module, replacing the
// bond denom from stake to desmos
func stakingGenesisState() staking.GenesisState {
	return staking.GenesisState{
		Params: staking.NewParams(
			staking.DefaultUnbondingTime,
			staking.DefaultMaxValidators,
			staking.DefaultMaxEntries,
			"4e188d9c17150037d5199bbdb91ae1eb2a78a15aca04cb35530cccb81494b36e",
		),
	}
}
