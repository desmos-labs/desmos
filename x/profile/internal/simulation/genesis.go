package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	accs := randomAccounts(simsState)
	profileParams := randomProfileParams(simsState)
	profileGenesis := types.NewGenesisState(
		accs,
		profileParams.NameSurnameParams,
		profileParams.MonikerParams,
		profileParams.BioParams,
	)
	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomAccounts returns randomly generated genesis accounts
func randomAccounts(simState *module.SimulationState) (accounts types.Profiles) {
	accountsNumber := simState.Rand.Intn(50)

	accounts = make(types.Profiles, accountsNumber)
	for i := 0; i < accountsNumber; i++ {
		accountData := RandomProfileData(simState.Rand, simState.Accounts)
		account := types.Profile{
			Moniker: accountData.Moniker,
			Name:    &accountData.Name,
			Surname: &accountData.Surname,
			Bio:     &accountData.Bio,
			Creator: accountData.Creator.Address,
		}

		accounts[i] = account
	}

	return accounts
}

// randomProfileParams returns randomly generated genesis parameters
func randomProfileParams(simState *module.SimulationState) ProfileParams {
	return RandomProfileParams(simState.Rand)
}
