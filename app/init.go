package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

const (
	DefaultBondDenom = "desmos"

	// gov
	DefaultStartingProposalID uint64 = 1
)

var (
	// gov
	DefaultMinDepositTokens = sdk.TokensFromConsensusPower(10)
	DefaultQuorum           = sdk.NewDecWithPrec(334, 3)
	DefaultThreshold        = sdk.NewDecWithPrec(5, 1)
	DefaultVeto             = sdk.NewDecWithPrec(334, 3)
)

// Init initializes the application, overriding the default genesis states that should be changed
func Init() {
	staking.DefaultGenesisState = stakingGenesisState
	gov.DefaultGenesisState = govGenesisState
}

// stakingGenesisState returns the default genesis state for the staking module, replacing the
// bond denom from stake to desmos
func stakingGenesisState() staking.GenesisState {
	return staking.GenesisState{
		Params: staking.NewParams(
			staking.DefaultUnbondingTime,
			staking.DefaultMaxValidators,
			staking.DefaultMaxEntries,
			0,
			DefaultBondDenom,
		),
	}
}

func govGenesisState() gov.GenesisState {
	return gov.NewGenesisState(
		DefaultStartingProposalID,
		gov.NewDepositParams(
			sdk.NewCoins(sdk.NewCoin(DefaultBondDenom, DefaultMinDepositTokens)),
			gov.DefaultPeriod,
		),
		gov.NewVotingParams(gov.DefaultPeriod),
		gov.NewTallyParams(DefaultQuorum, DefaultThreshold, DefaultVeto),
	)
}
