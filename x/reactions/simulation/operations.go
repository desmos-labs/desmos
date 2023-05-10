package simulation

// DONTCOVER

import (
	"math/rand"

	postskeeper "github.com/desmos-labs/desmos/v4/x/posts/keeper"
	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/app/params"
	"github.com/desmos-labs/desmos/v4/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v4/x/reactions/types"
)

// Simulation operation weights constants
// #nosec G101 -- This is a false positive
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
	k keeper.Keeper, profilesKeeper types.ProfilesKeeper, sk subspaceskeeper.Keeper, pk postskeeper.Keeper,
	ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) sim.WeightedOperations {

	var weightMsgAddReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddReaction, &weightMsgAddReaction, nil,
		func(_ *rand.Rand) {
			weightMsgAddReaction = params.DefaultWeightMsgAddReaction
		},
	)

	var weightMsgRemoveReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveReaction, &weightMsgRemoveReaction, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveReaction = params.DefaultWeightMsgRemoveReaction
		},
	)

	var weightMsgAddRegisteredReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddRegisteredReaction, &weightMsgAddRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgAddRegisteredReaction = params.DefaultWeightMsgAddRegisteredReaction
		},
	)

	var weightMsgEditRegisteredReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditRegisteredReaction, &weightMsgEditRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgEditRegisteredReaction = params.DefaultWeightMsgEditRegisteredReaction
		},
	)

	var weightMsgRemoveRegisteredReaction int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveRegisteredReaction, &weightMsgRemoveRegisteredReaction, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveRegisteredReaction = params.DefaultWeightMsgRemoveRegisteredReaction
		},
	)

	var weightMsgSetReactionsParams int
	appParams.GetOrGenerate(cdc, OpWeightMsgSetReactionsParams, &weightMsgSetReactionsParams, nil,
		func(_ *rand.Rand) {
			weightMsgSetReactionsParams = params.DefaultWeightMsgSetReactionsParams
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgAddReaction,
			SimulateMsgAddReaction(k, profilesKeeper, sk, pk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveReaction,
			SimulateMsgRemoveReaction(k, sk, pk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgAddRegisteredReaction,
			SimulateMsgAddRegisteredReaction(sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgEditRegisteredReaction,
			SimulateMsgEditRegisteredReaction(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRemoveRegisteredReaction,
			SimulateMsgRemoveRegisteredReaction(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgSetReactionsParams,
			SimulateMsgSetReactionsParams(sk, ak, bk, fk),
		),
	}
}
