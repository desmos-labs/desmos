package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MonikerLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomMonikerParams(r)
				return fmt.Sprintf(`{"min_moniker_len":"%s","max_moniker_len":"%s"}`,
					params.MinMonikerLen, params.MaxMonikerLen)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DtagLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomDTagParams(r)
				return fmt.Sprintf(`{"min_dtag_len":"%s","max_dtag_len":"%s"}`,
					params.MinDtagLen, params.MaxDtagLen)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxBioLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomBioParams(r)
				return fmt.Sprintf(`{"max_bio_len":"%s"}`, params)
			},
		),
	}
}
