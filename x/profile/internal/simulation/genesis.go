package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simsState *module.SimulationState) {
	accs := randomAccounts(simsState)
	profileGenesis := types.NewGenesisState(accs)
	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomAccounts returns randomly generated genesis accounts
func randomAccounts(simState *module.SimulationState) (accounts types.Profiles) {
	accountsNumber := simState.Rand.Intn(50)

	accounts = make(types.Profiles, accountsNumber)
	for i := 0; i < accountsNumber; i++ {
		simAccount, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		accounts[i] = NewRandomProfile(simState.Rand, simAccount.Address)
	}

	return accounts
}
