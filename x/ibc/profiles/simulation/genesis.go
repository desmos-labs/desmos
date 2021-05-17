package simulation

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
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

	gs := types.GenesisState{
		PortID: portID,
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&gs)
}
