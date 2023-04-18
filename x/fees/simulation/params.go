package simulation

// DONTCOVER

import (
	"encoding/json"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/x/fees/types"
)

func ParamChanges(r *rand.Rand) []simtypes.LegacyParamChange {
	params := types.NewParams(randomMinFees(r))
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.MinFeesStoreKey),
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
