package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

func ParamChanges(r *rand.Rand) []simtypes.ParamChange {
	return []simtypes.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, string(types.NicknameParamsKey),
			func(r *rand.Rand) string {
				params := RandomNicknameParams(r)
				return fmt.Sprintf(`{"min_length":"%s","max_length":"%s"}`,
					params.MinLength, params.MaxLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.DTagParamsKey),
			func(r *rand.Rand) string {
				params := RandomDTagParams(r)
				return fmt.Sprintf(`{"min_length":"%s","max_length":"%s"}`,
					params.MinLength, params.MaxLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.BioParamsKey),
			func(r *rand.Rand) string {
				params := RandomBioParams(r)
				return fmt.Sprintf(`{"max_length":"%s"}`, params.MaxLength)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.OracleParamsKey),
			func(r *rand.Rand) string {
				params := RandomOracleParams(r)
				feeAmountBz, err := json.Marshal(params.FeeAmount)
				if err != nil {
					panic(err)
				}
				return fmt.Sprintf(
					`{"script_id":"%d", "ask_count":"%d", "min_count":"%d", "prepare_gas":"%d", "execute_gas":"%d", "fee_amount":%s}`,
					params.ScriptID, params.AskCount, params.MinCount, params.PrepareGas, params.ExecuteGas, string(feeAmountBz),
				)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, string(types.AppLinksParamsKey),
			func(r *rand.Rand) string {
				params := RandomAppLinksParams(r)
				return fmt.Sprintf(`{"validity_duration":"%d"}`, params.ValidityDuration)
			},
		),
	}
}
