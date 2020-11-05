package simulation

// DONTCOVER

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/tendermint/tendermint/crypto/ed25519"

	"github.com/desmos-labs/desmos/x/reports/types"
)

func RandomizedGenState(simState *module.SimulationState) {
	reports := randomReports(simState)
	reportsGenesis := types.NewGenesisState(reports)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(reportsGenesis)
}

func randomReports(simState *module.SimulationState) (reportsMap []types.Report) {
	reportsMapLen := simState.Rand.Intn(50)

	reports := make([]types.Report, reportsMapLen)
	for i := 0; i < reportsMapLen; i++ {
		privKey := ed25519.GenPrivKey().PubKey()
		reports[i] = types.NewReport(
			RandomPostID(simState.Rand),
			RandomReportTypes(simState.Rand),
			RandomReportMessage(simState.Rand),
			sdk.AccAddress(privKey.Address()).String(),
		)
	}

	return reports
}
