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
		simulation.NewSimParamChange(types.ModuleName, string(types.NicknameLenParamsKey),
			func(r *rand.Rand) string {
				params := RandomNicknameParams(r)
				return fmt.Sprintf(`{"min_nickname_len":"%s","max_nickname_len":"%s"}`,
					params.MinNicknameLength, params.MaxNicknameLength)
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
		simulation.NewSimParamChange(types.ModuleName, string(types.OracleParamsKey),
			func(r *rand.Rand) string {
				params := RandomOracleParams(r)
				return fmt.Sprintf(
					`{"script_id":"%d", "ask_count":"%d", "min_count":"%d", "prepare_gas":"%d", "execute_gas":"%d", "fee_payer":"%s", fee_coins:"%s"}`,
					params.ScriptID, params.AskCount, params.MinCount, params.PrepareGas, params.ExecuteGas, params.FeePayer, params.FeeAmount.String(),
				)
			},
		),
	}
}
