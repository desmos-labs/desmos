package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MonikerLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomMonikerParams(r)
				jsonBz, _ := json.Marshal(params)
				return string(jsonBz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DtagLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomDTagParams(r)
				jsonBz, _ := json.Marshal(params)
				return string(jsonBz)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxBioLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomBioParams(r)
				return fmt.Sprintf("%q", params)
			},
		),
	}
}
