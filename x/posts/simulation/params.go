package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v3/x/posts/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxTextLengthKey),
			func(r *rand.Rand) string {
				maxTextLength := RandomMaxTextLength(r)
				return fmt.Sprintf(`%d`, maxTextLength)
			},
		),
	}
}
