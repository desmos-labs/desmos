package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/desmos-labs/desmos/x/links/types"
)

// Simulation parameter constants
const port = "port_id"

// RandomizedGenState generates a random GenesisState for transfer.
func RandomizedGenState(simState *module.SimulationState) {
	var portID string
	simState.AppParams.GetOrGenerate(
		simState.Cdc, port, &portID, simState.Rand,
		func(r *rand.Rand) { portID = strings.ToLower(simtypes.RandStringOfLength(r, 20)) },
	)

	linksGenesis := types.GenesisState{
		PortId: portID,
		Links:  randomLinks(simState),
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&linksGenesis)
}

func randomLinks(simState *module.SimulationState) (linksList []types.Link) {
	linksListLen := simState.Rand.Intn(50)

	links := make([]types.Link, linksListLen)
	for i := 0; i < linksListLen; i++ {
		src, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		dst, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)

		links[i] = types.NewLink(
			src.Address.String(),
			dst.Address.String(),
		)
	}

	return links
}
