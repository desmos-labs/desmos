package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v5/app/params"
	"github.com/desmos-labs/desmos/v5/x/reports/keeper"
	"github.com/desmos-labs/desmos/v5/x/reports/types"
)

// Simulation operation weights constants
// #nosec G101 -- This is a false positive
const (
	OpWeightMsgCreateReport          = "op_weight_msg_create_report"
	OpWeightMsgDeleteReport          = "op_weight_msg_delete_report"
	OpWeightMsgSupportStandardReason = "op_weight_msg_support_standard_reason"
	OpWeightMsgAddReason             = "op_weight_msg_add_reason"
	OpWeightMsgRemoveReason          = "op_weight_msg_remove_reason"

	DefaultGasValue = 200_000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, sk types.SubspacesKeeper, pk types.PostsKeeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgCreateReport int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateReport, &weightMsgCreateReport, nil,
		func(_ *rand.Rand) {
			weightMsgCreateReport = params.DefaultWeightMsgCreateReport
		},
	)

	var weightMsgDeleteReport int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteReport, &weightMsgDeleteReport, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteReport = params.DefaultWeightMsgDeleteReport
		},
	)

	var weightMsgSupportStandardReason int
	appParams.GetOrGenerate(cdc, OpWeightMsgSupportStandardReason, &weightMsgSupportStandardReason, nil,
		func(_ *rand.Rand) {
			weightMsgSupportStandardReason = params.DefaultWeightMsgSupportStandardReason
		},
	)

	var weightMsgAddReason int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddReason, &weightMsgAddReason, nil,
		func(_ *rand.Rand) {
			weightMsgAddReason = params.DefaultWeightMsgAddReason
		},
	)

	var weightMsgRemoveReason int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveReason, &weightMsgRemoveReason, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveReason = params.DefaultWeightMsgRemoveReason
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateReport,
			SimulateMsgCreateReport(k, sk, pk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteReport,
			SimulateMsgDeleteReport(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgSupportStandardReason,
			SimulateMsgSupportStandardReason(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAddReason,
			SimulateMsgAddReason(sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveReason,
			SimulateMsgRemoveReason(k, sk, ak, bk),
		),
	}
}
