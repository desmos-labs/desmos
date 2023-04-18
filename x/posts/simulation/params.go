package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/desmos-labs/desmos/v4/x/posts/types"
)

func ParamChanges(r *rand.Rand) []simtypes.LegacyParamChange {
	return []simtypes.LegacyParamChange{
		simulation.NewSimLegacyParamChange(types.ModuleName, string(types.MaxTextLengthKey),
			func(r *rand.Rand) string {
				maxTextLength := RandomMaxTextLength(r)
				return fmt.Sprintf(`%d`, maxTextLength)
			},
		),
	}
}
