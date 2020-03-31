package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
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
		accountData := RandomAccountData(simState.Rand, simState.Accounts)
		account := types.Profile{
			Moniker: accountData.Moniker,
			Name:    accountData.Name,
			Surname: accountData.Surname,
			Bio:     accountData.Bio,
			Creator: accountData.Creator.Address,
		}

		accounts[i] = account
	}

	return accounts
}
