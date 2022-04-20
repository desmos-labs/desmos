package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v3/app/params"
	"github.com/desmos-labs/desmos/v3/x/profiles/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgSaveProfile         = "op_weight_msg_save_profile"
	OpWeightMsgDeleteProfile       = "op_weight_msg_delete_profile"
	OpWeightMsgRequestDTagTransfer = "op_weight_msg_request_dtag_transfer"
	OpWeightMsgAcceptDTagTransfer  = "op_weight_msg_accept_dtag_transfer_request"
	OpWeightMsgRefuseDTagTransfer  = "op_weight_msg_refuse_dtag_transfer_request"
	OpWeightMsgCancelDTagTransfer  = "op_weight_msg_cancel_dtag_transfer_request"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) sim.WeightedOperations {
	var weightMsgSaveProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgSaveProfile, &weightMsgSaveProfile, nil,
		func(_ *rand.Rand) {
			weightMsgSaveProfile = params.DefaultWeightMsgSaveProfile
		},
	)

	var weightMsgDeleteProfile int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteProfile, &weightMsgDeleteProfile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteProfile = params.DefaultWeightMsgDeleteProfile
		},
	)

	var weightMsgRequestDTagTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgRequestDTagTransfer, &weightMsgRequestDTagTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgRequestDTagTransfer = params.DefaultWeightMsgRequestDTagTransfer
		},
	)

	var weightMsgAcceptDTagTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgAcceptDTagTransfer, &weightMsgAcceptDTagTransfer, nil,
		func(_ *rand.Rand) {
			weightMsgAcceptDTagTransfer = params.DefaultWeightMsgAcceptDTagTransfer
		},
	)

	var weightMsgRefuseDTagTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgRefuseDTagTransfer, &weightMsgRefuseDTagTransfer, nil,
		func(r *rand.Rand) {
			weightMsgRefuseDTagTransfer = params.DefaultWeightMsgRefuseDTagTransfer
		},
	)

	var weightMsgCancelDTagTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgCancelDTagTransfer, &weightMsgCancelDTagTransfer, nil,
		func(r *rand.Rand) {
			weightMsgCancelDTagTransfer = params.DefaultWeightMsgCancelDTagTransfer
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgSaveProfile,
			SimulateMsgSaveProfile(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteProfile,
			SimulateMsgDeleteProfile(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRequestDTagTransfer,
			SimulateMsgRequestDTagTransfer(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgAcceptDTagTransfer,
			SimulateMsgAcceptDTagTransfer(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRefuseDTagTransfer,
			SimulateMsgRefuseDTagTransfer(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgCancelDTagTransfer,
			SimulateMsgCancelDTagTransfer(k, ak, bk, fk),
		),
	}
}
