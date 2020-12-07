package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	posts "github.com/desmos-labs/desmos/x/posts/keeper"

	"github.com/desmos-labs/desmos/app/params"
)

const (
	OpWeightMsgReportPost = "op_weight_msg_report_post"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler,
	pk posts.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {
	var weightMsgReportPost int
	appParams.GetOrGenerate(cdc, OpWeightMsgReportPost, &weightMsgReportPost, nil,
		func(_ *rand.Rand) {
			weightMsgReportPost = params.DefaultWeightMsgReportPost
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgReportPost,
			SimulateMsgReportPost(pk, ak, bk),
		),
	}
}
