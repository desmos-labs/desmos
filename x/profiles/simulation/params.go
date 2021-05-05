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
		simulation.NewSimParamChange(types.ModuleName, string(types.UsernameLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomUsernameParams(r)
				return fmt.Sprintf(`{"min_username_len":"%s","max_username_len":"%s"}`,
					params.MinUsernameLength, params.MaxUsernameLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DTagLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomDTagParams(r)
				return fmt.Sprintf(`{"min_dtag_len":"%s","max_dtag_len":"%s"}`,
					params.MinDTagLength, params.MaxDTagLength)
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
