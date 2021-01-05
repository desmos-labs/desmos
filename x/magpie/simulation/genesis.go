package simulation

// DONTCOVER

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simState *module.SimulationState) {
	sessionLength := simState.Rand.Int63n(1000) + 1 // Session length cannot be zero
	sessions := randomSessions(simState)
	genState := types.NewGenesisState(uint64(sessionLength), sessions)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(genState)
}

// randomSessions returns a randomly generated list of sessions
func randomSessions(simState *module.SimulationState) []types.Session {
	sessionsLength := 0

	sessions := make([]types.Session, sessionsLength)
	for i := 0; i < sessionsLength; i++ {

		simAccount, _ := simtypes.RandomAcc(simState.Rand, simState.Accounts)
		created := simState.Rand.Int63n(20000)
		expiry := simState.Rand.Int63n(1000) + created

		data := RandomSessionData(simAccount, simState.Rand)

		// Create the session
		sessions[i] = types.NewSession(
			types.SessionID{Value: uint64(i)},
			data.Owner.Address.String(),
			uint64(created),
			uint64(expiry),
			data.Namespace,
			data.ExternalOwner,
			data.PubKey,
			data.Signature,
		)
	}

	return sessions
}
