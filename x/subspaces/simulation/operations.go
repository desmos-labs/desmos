package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/app/params"
	"github.com/desmos-labs/desmos/v4/x/subspaces/keeper"
)

// Simulation operation weights constants
// #nosec G101 -- This is a false positive
const (
	OpWeightMsgCreateSubspace              = "op_weight_msg_create_subspace"
	OpWeightMsgEditSubspace                = "op_weight_msg_edit_subspace"
	OpWeightMsgDeleteSubspace              = "op_weight_msg_delete_subspace"
	OpWeightMsgCreateSection               = "op_weight_msg_create_section"
	OpWeightMsgEditSection                 = "op_weight_msg_edit_section"
	OpWeightMsgMoveSection                 = "op_weight_msg_move_section"
	OpWeightMsgDeleteSection               = "op_weight_msg_delete_section"
	OpWeightMsgCreateUserGroup             = "op_weight_msg_create_user_group"
	OpWeightMsgEditUserGroup               = "op_weight_msg_edit_user_group"
	OpWeightMsgMoveUserGroup               = "op_weight_msg_move_user_group"
	OpWeightMsgSetUserGroupPermissions     = "op_weight_msg_set_user_group_permissions"
	OpWeightMsgDeleteUserGroup             = "op_weight_msg_delete_user_group"
	OpWeightMsgAddUserToUserGroup          = "op_weight_msg_add_user_to_user_group"
	OpWeightMsgRemoveUserFromUserGroup     = "op_weight_msg_remove_user_from_user_group"
	OpWeightMsgSetUserPermissions          = "op_weight_msg_set_user_permissions"
	OpWeightMsgGrantTreasuryAuthorization  = "op_weight_msg_grant_treasury_authorization"
	OpWeightMsgRevokeTreasuryAuthorization = "op_weight_msg_revoke_treasury_authorization"
	OpWeightMsgGrantAllowance              = "op_weight_msg_grant_allowance"
	OpWeightMsgRevokeAllowance             = "op_weight_msg_revoke_allowance"

	DefaultGasValue = 200_000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper, authzk authzkeeper.Keeper,
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

	var weightMsgCreateSection int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateSection, &weightMsgCreateSection, nil,
		func(_ *rand.Rand) {
			weightMsgCreateSection = params.DefaultWeightMsgCreateSection
		},
	)

	var weightMsgEditSection int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditSection, &weightMsgEditSection, nil,
		func(_ *rand.Rand) {
			weightMsgEditSection = params.DefaultWeightMsgEditSection
		},
	)

	var weightMsgMoveSection int
	appParams.GetOrGenerate(cdc, OpWeightMsgMoveSection, &weightMsgMoveSection, nil,
		func(_ *rand.Rand) {
			weightMsgMoveSection = params.DefaultWeightMsgMoveSection
		},
	)

	var weightMsgDeleteSection int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteSection, &weightMsgDeleteSection, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteSection = params.DefaultWeightMsgDeleteSection
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

	var weightMsgMoveUserGroup int
	appParams.GetOrGenerate(cdc, OpWeightMsgMoveUserGroup, &weightMsgMoveUserGroup, nil,
		func(_ *rand.Rand) {
			weightMsgMoveUserGroup = params.DefaultWeightMsgMoveUserGroup
		})

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

	var weightMsgGrantTreasuryAuthorization int
	appParams.GetOrGenerate(cdc, OpWeightMsgGrantTreasuryAuthorization, &weightMsgGrantTreasuryAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgGrantTreasuryAuthorization = params.DefaultWeightMsgGrantTreasuryAuthorization
		},
	)

	var weightMsgRevokeTreasuryAuthorization int
	appParams.GetOrGenerate(cdc, OpWeightMsgRevokeTreasuryAuthorization, &weightMsgRevokeTreasuryAuthorization, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeTreasuryAuthorization = params.DefaultWeightMsgRevokeTreasuryAuthorization
		},
	)

	var weightMsgGrantUserAllowance int
	appParams.GetOrGenerate(cdc, OpWeightMsgGrantAllowance, &weightMsgGrantUserAllowance, nil,
		func(_ *rand.Rand) {
			weightMsgGrantUserAllowance = params.DefaultWeightMsgGrantAllowance
		},
	)

	var weightMsgRevokeUserAllowance int
	appParams.GetOrGenerate(cdc, OpWeightMsgRevokeAllowance, &weightMsgRevokeUserAllowance, nil,
		func(_ *rand.Rand) {
			weightMsgRevokeUserAllowance = params.DefaultWeightMsgRevokeAllowance
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
			weightMsgCreateSection,
			SimulateMsgCreateSection(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgEditSection,
			SimulateMsgEditSection(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgMoveSection,
			SimulateMsgMoveSection(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteSection,
			SimulateMsgDeleteSection(k, ak, bk, fk),
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
			weightMsgMoveUserGroup,
			SimulateMsgMoveUserGroup(k, ak, bk, fk),
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
		sim.NewWeightedOperation(
			weightMsgGrantTreasuryAuthorization,
			SimulateMsgGrantTreasuryAuthorization(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRevokeTreasuryAuthorization,
			SimulateMsgRevokeTreasuryAuthorization(k, ak, bk, fk, authzk),
		),
		sim.NewWeightedOperation(
			weightMsgGrantUserAllowance,
			SimulateMsgGrantAllowance(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRevokeUserAllowance,
			SimulateMsgRevokeAllowance(k, ak, bk, fk),
		),
	}
}
