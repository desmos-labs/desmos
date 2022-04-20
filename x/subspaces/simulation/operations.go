package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v3/app/params"
	"github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
)

// Simulation operation weights constants
// #nosec G101 -- This is a false positive
const (
	OpWeightMsgCreateSubspace          = "op_weight_msg_create_subspace"
	OpWeightMsgEditSubspace            = "op_weight_msg_edit_subspace"
	OpWeightMsgDeleteSubspace          = "op_weight_msg_delete_subspace"
	OpWeightMsgCreateUserGroup         = "op_weight_msg_create_user_group"
	OpWeightMsgEditUserGroup           = "op_weight_msg_edit_user_group"
	OpWeightMsgSetUserGroupPermissions = "op_weight_msg_set_user_group_permissions"
	OpWeightMsgDeleteUserGroup         = "op_weight_msg_delete_user_group"
	OpWeightMsgAddUserToUserGroup      = "op_weight_msg_add_user_to_user_group"
	OpWeightMsgRemoveUserFromUserGroup = "op_weight_msg_remove_user_from_user_group"
	OpWeightMsgSetUserPermissions      = "op_weight_msg_set_user_permissions"

	DefaultGasValue = 200_000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
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

	var weightMsgDeleteSubspace int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteSubspace, &weightMsgDeleteSubspace, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSubspace = params.DefaultWeightMsgDeleteSubspace
		},
	)

	var weightMsgCreateUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateUserGroup, &weightMsgCreateUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgCreateUserGroup = params.DefaultWeightMsgCreateUserGroup
		},
	)

	var weightMsgEditUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditUserGroup, &weightMsgEditUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgEditUserGroup = params.DefaultWeightMsgEditUserGroup
		},
	)

	var weightMsgSetUserGroupPermissions int
	appParams.GetOrGenerate(cdc, OpWeightMsgSetUserGroupPermissions, &weightMsgSetUserGroupPermissions, nil,
		func(_ *rand.Rand) {
			weightMsgSetUserGroupPermissions = params.DefaultWeightMsgSetUserGroupPermissions
		},
	)

	var weightMsgDeleteUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteUserGroup, &weightMsgDeleteUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteUserGroup = params.DefaultWeightMsgDeleteUserGroup
		},
	)

	var weightMsgAddUserToUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddUserToUserGroup, &weightMsgAddUserToUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgAddUserToUserGroup = params.DefaultWeightMsgAddUserToUserGroup
		},
	)

	var weightMsgRemoveUserFromUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveUserFromUserGroup, &weightMsgRemoveUserFromUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveUserFromUserGroup = params.DefaultWeightMsgRemoveUserFromUserGroup
		},
	)

	var weightMsgSetUserPermissions int
	appParams.GetOrGenerate(cdc, OpWeightMsgSetUserPermissions, &weightMsgSetUserPermissions, nil,
		func(_ *rand.Rand) {
			weightMsgSetUserPermissions = params.DefaultWeightMsgSetUserPermissions
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateSubspace,
			SimulateMsgCreateSubspace(ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgEditSubspace,
			SimulateMsgEditSubspace(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteSubspace,
			SimulateMsgDeleteSubspace(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgCreateUserGroup,
			SimulateMsgCreateUserGroup(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgEditUserGroup,
			SimulateMsgEditUserGroup(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgSetUserGroupPermissions,
			SimulateMsgSetUserGroupPermissions(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteUserGroup,
			SimulateMsgDeleteUserGroup(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgAddUserToUserGroup,
			SimulateMsgAddUserToUserGroup(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveUserFromUserGroup,
			SimulateMsgRemoveUserFromUserGroup(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgSetUserPermissions,
			SimulateMsgSetUserPermissions(k, ak, bk, fk),
		),
	}
}
