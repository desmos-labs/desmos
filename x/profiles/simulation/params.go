package simulation

// DONTCOVER

import (
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
				return fmt.Sprintf(`{"min_moniker_len":"%s","max_moniker_len":"%s"}`,
					params.MinMonikerLength, params.MaxMonikerLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DtagLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomDTagParams(r)
				return fmt.Sprintf(`{"min_dtag_len":"%s","max_dtag_len":"%s"}`,
					params.MinDtagLength, params.MaxDtagLength)
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
