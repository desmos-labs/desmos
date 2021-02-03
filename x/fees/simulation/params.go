package simulation

// DONTCOVER

import (
	"encoding/json"
	//"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/x/fees/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	params := types.NewParams(GenMinFees(r))
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MinFeesStoreKey),
			func(r *rand.Rand) string {
				jsonBz, _ := json.Marshal(params.MinFees)
				return string(jsonBz)
			},
		),
	}
}
