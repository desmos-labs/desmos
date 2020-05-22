package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/reports/internal/keeper"
)

const (
	OpWeightMsgReportPost = "op_weight_msg_report_post"

	DefaultGasValue = 2000000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper) sim.WeightedOperations {
	var weightMsgReportPost int
	appParams.GetOrGenerate(cdc, OpWeightMsgReportPost, &weightMsgReportPost, nil,
		func(_ *rand.Rand) {
			weightMsgReportPost = params.DefaultWeightMsgReportPost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgReportPost,
			SimulateMsgReportPost(k, ak),
		),
	}
}
