package simulation

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"math/rand"
)

const (
	keyNameSurnameParams = "nameSurnameLenParams"
	keyMonikerParams     = "monikerLenParams"
	keyBioParams         = "bioLenParams"
)

func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keyNameSurnameParams,
			func(r *rand.Rand) string {
				params := RandomNameSurnameParams(r)
				return fmt.Sprintf(`{"min_name_surname_len":"%s","max_name_surname_len":"%s"}`,
					params.MinNameSurnameLen, params.MaxNameSurnameLen)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyMonikerParams,
			func(r *rand.Rand) string {
				params := RandomMonikerParams(r)
				return fmt.Sprintf(`{"min_moniker_len":"%s","max_moniker_len":"%s"}`,
					params.MinMonikerLen, params.MaxMonikerLen)
			},
		),
		simulation.NewSimParamChange(types.ModuleName, keyBioParams,
			func(r *rand.Rand) string {
				params := RandomBioParams(r)
				return fmt.Sprintf(`{"max_bio_len":"%s"}`, params.MaxBioLen)
			},
		),
	}
}
