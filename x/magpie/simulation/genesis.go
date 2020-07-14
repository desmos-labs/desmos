package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState) {
	sessionLength := simState.Rand.Int63n(1000) + 1 // Session length cannot be zero
	sessions := randomSessions(simState)
	genState := types.NewGenesisState(sessionLength, sessions)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genState)
}

// randomSessions returns a randomly generated list of sessions
func randomSessions(simState *module.SimulationState) types.Sessions {
	sessionsLength := 0

	sessions := make(types.Sessions, sessionsLength)
	for i := 0; i < sessionsLength; i++ {

		simAccount, _ := sim.RandomAcc(simState.Rand, simState.Accounts)
		created := simState.Rand.Int63n(20000)
		data := RandomSessionData(simAccount, simState.Rand)

		// Create the session
		sessions[i] = types.NewSession(
			types.SessionID(i),
			data.Owner.Address,
			created,
			simState.Rand.Int63n(1000)+created,
			data.Namespace,
			data.ExternalOwner,
			data.PubKey,
			data.Signature,
		)
	}

	return sessions
}
