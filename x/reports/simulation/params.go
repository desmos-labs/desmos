package simulation

// DONTCOVER

import (
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v4/x/reports/types"
)

func ParamChanges(_ *rand.Rand) []simtypes.LegacyParamChange {
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.ReasonsKey),
			func(r *rand.Rand) string {
				standardReasons := GetRandomStandardReasons(r, 10)
				bz, err := json.Marshal(&standardReasons)
				if err != nil {
					panic(err)
				}
				return string(bz)
			},
		),
	}
}
