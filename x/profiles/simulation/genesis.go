package simulation

// DONTCOVER

import (
	"fmt"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// RandomizedGenState generates a random GenesisState for profile
func RandomizedGenState(simsState *module.SimulationState) {

	profileGenesis := types.NewGenesisState(
		randomDTagTransferRequests(simsState),
		types.NewParams(
			RandomMonikerParams(simsState.Rand),
			RandomDTagParams(simsState.Rand),
			RandomBioParams(simsState.Rand),
		),
	)

	bz, err := simsState.Cdc.MarshalJSON(profileGenesis)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated profile parameters:\n%s\n", bz)

	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}

// randomDTagTransferRequests returns randomly generated genesis dTag transfer requests
func randomDTagTransferRequests(simState *module.SimulationState) []types.DTagTransferRequest {
	dtagTransferRequestsNumber := simState.Rand.Intn(20)

	dtagTransferRequests := make([]types.DTagTransferRequest, dtagTransferRequestsNumber)
	for i := 0; i < dtagTransferRequestsNumber; i++ {
		simAccount, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		simAccount2, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		dtagTransferRequests[i] = types.NewDTagTransferRequest(
			RandomDTag(simState.Rand),
			simAccount.Address.String(),
			simAccount2.Address.String(),
		)
	}

	return dtagTransferRequests
}
