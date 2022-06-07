package simulation

// DONTCOVER

import (
	"encoding/json"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

func ParamChanges(_ *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.ReasonsKey),
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
