package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	params := RandomParams(r)
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxPostMessageLengthKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%q",
					params.MaxPostMessageLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxOptionalDataFieldsNumberKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%q",
					params.MaxOptionalDataFieldsNumber)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxOptionalDataFieldValueLengthKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf("%q", params.MaxOptionalDataFieldValueLength)
			},
		),
	}
}
