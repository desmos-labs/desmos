package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/desmos-labs/desmos/app/params"
	"github.com/desmos-labs/desmos/x/fees"
	"github.com/desmos-labs/desmos/x/relationships/keeper"
)

const (
	OpWeightMsgCreateRelationship = "op_weight_msg_create_relationship"
	OpWeightMsgDeleteRelationship = "op_weight_msg_delete_relationship"
	OpWeightMsgBlockUser          = "op_weight_msg_block_user"
	OpWeightMsgUnBlockUser        = "op_weight_msg_unblock_user"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(appParams sim.AppParams, cdc *codec.Codec, k keeper.Keeper, ak auth.AccountKeeper,
	fk fees.Keeper) sim.WeightedOperations {
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
			SimulateMsgCreateRelationship(k, ak, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeleteRelationship,
			SimulateMsgDeleteRelationship(k, ak, fk),
		),
		sim.NewWeightedOperation(
			weightMsgBlockUser,
			SimulateMsgBlockUser(k, ak, fk),
		),
		sim.NewWeightedOperation(
			weightMsgUnblockUser,
			SimulateMsgUnblockUser(k, ak, fk),
		),
	}
}
