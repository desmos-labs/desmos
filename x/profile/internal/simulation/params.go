package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MonikerLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomNameSurnameParams(r)
				return fmt.Sprintf(`{"min_name_surname_len":"%s","max_name_surname_len":"%s"}`,
					params.MinMonikerLen, params.MaxMonikerLen)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DtagLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomMonikerParams(r)
				return fmt.Sprintf(`{"min_moniker_len":"%s","max_moniker_len":"%s"}`,
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
