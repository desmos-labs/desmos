package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/subspaces/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateSubspace = "op_weight_msg_create_subspace"
	OpWeightMsgEditSubspace   = "op_weight_msg_edit_subspace"
	OpWeightMsgAddAdmin       = "op_weight_msg_add_admin"
	OpWeightMsgRemoveAdmin    = "op_weight_msg_remove_admin"
	OpWeightMsgRegisterUser   = "op_weight_msg_register_user"
	OpWeightMsgUnregisterUser = "op_weight_msg_unregister_user"
	OpWeightMsgBanUser        = "op_weight_msg_ban_user"
	OpWeightMsgUnbanUser      = "op_weight_msg_unban_user"

	DefaultGasValue = 500_000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONMarshaler,
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgCreateSubspace int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateSubspace, &weightMsgCreateSubspace, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSubspace = params.DefaultWeightMsgCreateSubspace
		},
	)

	var weightMsgEditSubspace int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditSubspace, &weightMsgEditSubspace, nil,
		func(_ *rand.Rand) {
			weightMsgEditSubspace = params.DefaultWeightMsgEditSubspace
		},
	)

	var weightMsgAddAdmin int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddAdmin, &weightMsgAddAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgAddAdmin = params.DefaultWeightMsgAddAmin
		},
	)

	var weightMsgRemoveAdmin int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveAdmin, &weightMsgRemoveAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveAdmin = params.DefaultWeightMsgRemoveAdmin
		},
	)

	var weightMsgRegisterUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterUser, &weightMsgRegisterUser, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterUser = params.DefaultWeightMsgRegisterUser
		},
	)

	var weightMsgUnregisterUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgUnregisterUser, &weightMsgUnregisterUser, nil,
		func(_ *rand.Rand) {
			weightMsgUnregisterUser = params.DefaultWeightMsgUnregisterUser
		},
	)

	var weightMsgBanUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgBanUser, &weightMsgBanUser, nil,
		func(_ *rand.Rand) {
			weightMsgBanUser = params.DefaultWeightMsgBanUser
		},
	)

	var weightMsgUnbanUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgUnbanUser, &weightMsgUnbanUser, nil,
		func(_ *rand.Rand) {
			weightMsgUnbanUser = params.DefaultWeightMsgUnbanUser
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateSubspace,
			SimulateMsgCreateSubspace(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgEditSubspace,
			SimulateMsgEditSubspace(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAddAdmin,
			SimulateMsgRegisterUser(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAddAdmin,
			SimulateMsgAddAdmin(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveAdmin,
			SimulateMsgRemoveAdmin(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRegisterUser,
			SimulateMsgRegisterUser(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgUnregisterUser,
			SimulateMsgUnregisterUser(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgBanUser,
			SimulateMsgBanUser(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgUnbanUser,
			SimulateMsgUnbanUser(k, ak, bk),
		),
	}
}
