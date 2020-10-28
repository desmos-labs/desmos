package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

const (
	BondDenom = "desmos"
)

// Init initializes the application, overriding the default genesis states that should be changed
func Init() {
	// TODO: Check how this can be implemented again
	//stakingtypes.DefaultGenesisState = stakingGenesisState
	//govtypes.DefaultGenesisState = govGenesisState
}

// stakingGenesisState returns the default genesis state for the staking module, replacing the
// bond denom from stake to desmos
func stakingGenesisState() *stakingtypes.GenesisState {
	return stakingtypes.NewGenesisState(
		stakingtypes.NewParams(
			stakingtypes.DefaultUnbondingTime,
			stakingtypes.DefaultMaxValidators,
			stakingtypes.DefaultMaxEntries,
			stakingtypes.DefaultHistoricalEntries,
			BondDenom,
		),
		nil,
		nil,
	)
}

// govGenesisState returns the default genesis state for the gov module, replacing the
// bond denom from stake to desmos
func govGenesisState() *govtypes.GenesisState {
	return govtypes.NewGenesisState(
		govtypes.DefaultStartingProposalID,
		govtypes.NewDepositParams(
			sdk.NewCoins(sdk.NewCoin(BondDenom, govtypes.DefaultMinDepositTokens)),
			govtypes.DefaultPeriod,
		),
		govtypes.DefaultVotingParams(),
		govtypes.DefaultTallyParams(),
	)
}
