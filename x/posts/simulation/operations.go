package simulation

// DONTCOVER

import (
	"math/rand"

	subspaceskeeper "github.com/desmos-labs/desmos/v4/x/subspaces/keeper"

	feeskeeper "github.com/desmos-labs/desmos/v4/x/fees/keeper"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v4/app/params"
	"github.com/desmos-labs/desmos/v4/x/posts/keeper"
)

// Simulation operation weights constants
//
//nolint:gosec // These are not hardcoded credentials
const (
	OpWeightMsgCreatePost           = "op_weight_msg_create_post"
	OpWeightMsgEditPost             = "op_weight_msg_edit_post"
	OpWeightMsgDeletePost           = "op_weight_msg_delete_post"
	OpWeightMsgAddPostAttachment    = "op_weight_msg_add_post_attachment"
	OpWeightMsgRemovePostAttachment = "op_weight_msg_remove_post_attachment"
	OpWeightMsgAnswerPoll           = "op_weight_msg_answer_poll"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, sk subspaceskeeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper, fk feeskeeper.Keeper,
) sim.WeightedOperations {
	var weightMsgCreatePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = params.DefaultWeightMsgCreatePost
		},
	)

	var weightMsgEditPost int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditPost, &weightMsgEditPost, nil,
		func(_ *rand.Rand) {
			weightMsgEditPost = params.DefaultWeightMsgEditPost
		},
	)

	var weightMsgDeletePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = params.DefaultWeightMsgDeletePost
		},
	)

	var weightMsgAddPostAttachment int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddPostAttachment, &weightMsgAddPostAttachment, nil,
		func(_ *rand.Rand) {
			weightMsgAddPostAttachment = params.DefaultWeightMsgAddPostAttachment
		},
	)

	var weightMsgRemovePostAttachment int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemovePostAttachment, &weightMsgRemovePostAttachment, nil,
		func(r *rand.Rand) {
			weightMsgRemovePostAttachment = params.DefaultWeightMsgRemovePostAttachment
		},
	)

	var weightMsgAnswerPoll int
	appParams.GetOrGenerate(cdc, OpWeightMsgAnswerPoll, &weightMsgAnswerPoll, nil,
		func(r *rand.Rand) {
			weightMsgAnswerPoll = params.DefaultWeightMsgAnswerPoll
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreatePost,
			SimulateMsgCreatePost(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgEditPost,
			SimulateMsgEditPost(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgDeletePost,
			SimulateMsgDeletePost(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgAddPostAttachment,
			SimulateMsgAddPostAttachment(k, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgRemovePostAttachment,
			SimulateMsgRemovePostAttachment(k, sk, ak, bk, fk),
		),
		sim.NewWeightedOperation(
			weightMsgAnswerPoll,
			SimulateMsgAnswerPoll(k, sk, ak, bk, fk),
		),
	}
}
