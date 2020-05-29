package simulation

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func RandomizedGenState(simState *module.SimulationState) {
	reports := randomReports(simState)
	reportsGenesis := types.NewGenesisState(reports)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(reportsGenesis)
}

func randomReports(simState *module.SimulationState) (reportsMap map[string]types.Reports) {
	reportsMapLen := simState.Rand.Intn(rand.Intn(50))

	reportsMap = make(map[string]types.Reports, reportsMapLen)
	for i := 0; i < reportsMapLen; i++ {
		reportsLen := simState.Rand.Intn(20)
		reports := make(types.Reports, reportsLen)
		for j := 0; j < reportsLen; j++ {
			privKey := ed25519.GenPrivKey().PubKey()
			reports[j] = types.NewReport(
				RandomReportTypes(simState.Rand),
				RandomReportMessage(simState.Rand),
				sdk.AccAddress(privKey.Address()),
			)
		}
		reportsMap[RandomPostID(simState.Rand, simState.Accounts).String()] = reports
	}

	return reportsMap
}
