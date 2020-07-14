package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/x/posts/types"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	params := RandomParams(r)
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxPostMessageLengthKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"max_post_message_length":"%s"`,
					params.MaxPostMessageLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxOptionalDataFieldsNumberKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"max_optional_data_fields_number":"%s"`,
					params.MaxOptionalDataFieldsNumber)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.MaxOptionalDataFieldValueLengthKey),
			func(r *rand.Rand) string {
				return fmt.Sprintf(`{"max_optional_data_field_value_length":"%s"}`, params.MaxOptionalDataFieldValueLength)
			},
		),
	}
}
