package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/fees/types"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	params := types.DefaultParams()
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MinFeesStoreKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"min_fees":"%s"`,
					params.MinFees)
			},
		),
	}
}
