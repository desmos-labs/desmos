package simulation

// DONTCOVER

import (
	"math/rand"

	feeskeeper "github.com/desmos-labs/desmos/v3/x/fees/keeper"

	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v3/app/params"
	"github.com/desmos-labs/desmos/v3/x/relationships/keeper"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateRelationship = "op_weight_msg_create_relationship"
	OpWeightMsgDeleteRelationship = "op_weight_msg_delete_relationship"
	OpWeightMsgBlockUser          = "op_weight_msg_block_user"
	OpWeightMsgUnBlockUser        = "op_weight_msg_unblock_user"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) sim.WeightedOperations {
	var weightMsgCreateRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreateRelationship, &weightMsgCreateRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgCreateRelationship = params.DefaultWeightMsgCreateRelationship
		},
	)

	var weightMsgDeleteRelationship int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeleteRelationship, &weightMsgDeleteRelationship, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteRelationship = params.DefaultWeightMsgDeleteRelationship
		},
	)

	var weightMsgBlockUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgBlockUser, &weightMsgBlockUser, nil,
		func(_ *rand.Rand) {
			weightMsgBlockUser = params.DefaultWeightMsgBlockUser
		},
	)

	var weightMsgUnblockUser int
	appParams.GetOrGenerate(cdc, OpWeightMsgUnBlockUser, &weightMsgUnblockUser, nil,
		func(_ *rand.Rand) {
			weightMsgBlockUser = params.DefaultWeightMsgUnblockUser
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreateRelationship,
			SimulateMsgCreateRelationship(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteRelationship,
			SimulateMsgDeleteRelationship(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgBlockUser,
			SimulateMsgBlockUser(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgUnblockUser,
			SimulateMsgUnblockUser(k, ak, bk, fk),
		),
	}
}
