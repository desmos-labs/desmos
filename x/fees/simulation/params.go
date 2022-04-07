package simulation

// DONTCOVER

import (
	"encoding/json"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	params := types.NewParams(randomMinFees(r))
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MinFeesStoreKey),
			func(r *rand.Rand) string {
				minFeesBytes, err := json.Marshal(params.MinFees)
				if err != nil {
					panic(err)
				}
				return string(minFeesBytes)
			},
		),
	}
}
