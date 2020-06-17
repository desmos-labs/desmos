package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"math/rand"
)

const (
	NSParamsKey      = "nameSurnameLenParams"
	monikerParamsKey = "monikerLenParams"
	bioParamsKey     = "bioLenParams"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {
	accs := randomAccounts(simsState)

	var nsParams models.NameSurnameLenParams
	simsState.AppParams.GetOrGenerate(simsState.Cdc, NSParamsKey, &nsParams, simsState.Rand,
		func(r *rand.Rand) { nsParams = RandomNameSurnameParams(r) })

	var monikerParams models.MonikerLenParams
	simsState.AppParams.GetOrGenerate(simsState.Cdc, monikerParamsKey, &monikerParams, simsState.Rand,
		func(r *rand.Rand) { monikerParams = RandomMonikerParams(r) })

	var bioParams models.BioLenParams
	simsState.AppParams.GetOrGenerate(simsState.Cdc, bioParamsKey, &bioParams, simsState.Rand,
		func(r *rand.Rand) { bioParams = RandomBioParams(r) })

	profileGenesis := types.NewGenesisState(
		accs,
		nsParams,
		monikerParams,
		bioParams,
	)

	fmt.Printf("Selected randomly generated profile parameters:\n%s\n%s\n%s\n",
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.NameSurnameLenParams),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.MonikerLenParams),
		codec.MustMarshalJSONIndent(simsState.Cdc, profileGenesis.BioLenParams),
	)

	simsState.GenState[models.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomAccounts returns randomly generated genesis accounts
func randomAccounts(simState *module.SimulationState) (accounts models.Profiles) {
	accountsNumber := simState.Rand.Intn(50)

	accounts = make(models.Profiles, accountsNumber)
	for i := 0; i < accountsNumber; i++ {
		accountData := RandomProfileData(simState.Rand, simState.Accounts)
		account := models.Profile{
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
