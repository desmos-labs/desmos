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
			"desmos",
		),
	}
}
