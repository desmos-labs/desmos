package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v6/app/params"
	"github.com/desmos-labs/desmos/v6/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

// Simulation operation weights constants
//
//nolint:gosec // This is a false positive
const (
	OpWeightMsgAddReaction              = "op_weight_msg_add_reactions"
	OpWeightMsgRemoveReaction           = "op_weight_msg_remove_reaction"
	OpWeightMsgAddRegisteredReaction    = "op_weight_msg_add_registered_reaction"
	OpWeightMsgEditRegisteredReaction   = "op_weight_msg_edit_registered_reaction"
	OpWeightMsgRemoveRegisteredReaction = "op_weight_msg_remove_registered_reaction"
	OpWeightMsgSetReactionsParams       = "op_weight_msg_set_reactions_params"

	DefaultGasValue = 200_000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, profilesKeeper types.ProfilesKeeper, sk types.SubspacesKeeper, pk types.PostsKeeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgAddReaction int
	appParams.GetOrGenerate(OpWeightMsgAddReaction, &weightMsgAddReaction, nil,
		func(_ *rand.Rand) {
			weightMsgAddReaction = params.DefaultWeightMsgAddReaction
		},
	)

	var weightMsgRemoveReaction int
	appParams.GetOrGenerate(OpWeightMsgRemoveReaction, &weightMsgRemoveReaction, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveReaction = params.DefaultWeightMsgRemoveReaction
		},
	)

	var weightMsgAddRegisteredReaction int
	appParams.GetOrGenerate(OpWeightMsgAddRegisteredReaction, &weightMsgAddRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgAddRegisteredReaction = params.DefaultWeightMsgAddRegisteredReaction
		},
	)

	var weightMsgEditRegisteredReaction int
	appParams.GetOrGenerate(OpWeightMsgEditRegisteredReaction, &weightMsgEditRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgEditRegisteredReaction = params.DefaultWeightMsgEditRegisteredReaction
		},
	)

	var weightMsgRemoveRegisteredReaction int
	appParams.GetOrGenerate(OpWeightMsgRemoveRegisteredReaction, &weightMsgRemoveRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveRegisteredReaction = params.DefaultWeightMsgRemoveRegisteredReaction
		},
	)

	var weightMsgSetReactionsParams int
	appParams.GetOrGenerate(OpWeightMsgSetReactionsParams, &weightMsgSetReactionsParams, nil,
		func(_ *rand.Rand) {
			weightMsgSetReactionsParams = params.DefaultWeightMsgSetReactionsParams
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgAddReaction,
			SimulateMsgAddReaction(k, profilesKeeper, sk, pk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveReaction,
			SimulateMsgRemoveReaction(k, sk, pk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAddRegisteredReaction,
			SimulateMsgAddRegisteredReaction(sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgEditRegisteredReaction,
			SimulateMsgEditRegisteredReaction(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveRegisteredReaction,
			SimulateMsgRemoveRegisteredReaction(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgSetReactionsParams,
			SimulateMsgSetReactionsParams(sk, ak, bk),
		),
	}
}
