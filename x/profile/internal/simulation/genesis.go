package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

// RandomizedGenState generates a random GenesisState for auth
func RandomizedGenState(simsState *module.SimulationState) {

	profileGenesis := types.NewGenesisState(accs)
	simsState.GenState[types.ModuleName] = simsState.Cdc.MustMarshalJSON(profileGenesis)
}
