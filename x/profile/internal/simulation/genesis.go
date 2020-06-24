package simulation

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

const (
	ParamsKey = "profileParams"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	profileGenesis := types.NewGenesisState(
		randomProfiles(simsState),
		types.NewParams(RandomNameSurnameParams(simsState.Rand), RandomMonikerParams(simsState.Rand), RandomBioParams(simsState.Rand)),
	)

	fmt.Printf("Selected randomly generated profile parameters:\n%s\n%s\n%s\n",
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.NameSurnameLengths),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.MonikerLengths),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.Params.MaxBioLen),
	)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomProfiles returns randomly generated genesis accounts
func randomProfiles(simState *module.SimulationState) (accounts types.Profiles) {
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
