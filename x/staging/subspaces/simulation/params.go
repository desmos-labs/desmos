package simulation

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

// DONTCOVER

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.NameLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomNameParams(r)
				return fmt.Sprintf(`{"min_name_len":"%s","max_name_len":"%s"}`,
					params.MinNameLength, params.MaxNameLength)
			},
		),
	}
}
