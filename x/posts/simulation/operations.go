package simulation

// DONTCOVER

import (
	"math/rand"

	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/cosmos/cosmos-sdk/codec"
	sim "github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// Simulation operation weights constants
//
//nolint:gosec // These are not hardcoded credentials
const (
	DefaultWeightMsgCreatePost               int = 80
	DefaultWeightMsgEditPost                 int = 40
	DefaultWeightMsgDeletePost               int = 20
	DefaultWeightMsgAddPostAttachment        int = 50
	DefaultWeightMsgRemovePostAttachment     int = 50
	DefaultWeightMsgAnswerPoll               int = 50
	DefaultWeightMsgMovePost                 int = 10
	DefaultWeightMsgRequestPostOwnerTransfer int = 10
	DefaultWeightMsgCancelPostOwnerTransfer  int = 10
	DefaultWeightMsgAcceptPostOwnerTransfer  int = 10
	DefaultWeightMsgRefusePostOwnerTransfer  int = 10

	OpWeightMsgCreatePost               = "op_weight_msg_create_post"
	OpWeightMsgEditPost                 = "op_weight_msg_edit_post"
	OpWeightMsgDeletePost               = "op_weight_msg_delete_post"
	OpWeightMsgAddPostAttachment        = "op_weight_msg_add_post_attachment"
	OpWeightMsgRemovePostAttachment     = "op_weight_msg_remove_post_attachment"
	OpWeightMsgAnswerPoll               = "op_weight_msg_answer_poll"
	OpWeightMsgMovePost                 = "op_weight_msg_move_post"
	OpWeightMsgRequestPostOwnerTransfer = "op_weight_msg_request_post_owner_transfer"
	OpWeightMsgCancelPostOwnerTransfer  = "op_weight_msg_cancel_post_owner_transfer"
	OpWeightMsgAcceptPostOwnerTransfer  = "op_weight_msg_accept_post_owner_transfer"
	OpWeightMsgRefusePostOwnerTransfer  = "op_weight_msg_refuse_post_owner_transfer"

	DefaultGasValue = 200000
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	k keeper.Keeper, sk types.SubspacesKeeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) sim.WeightedOperations {
	var weightMsgCreatePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgCreatePost, &weightMsgCreatePost, nil,
		func(_ *rand.Rand) {
			weightMsgCreatePost = DefaultWeightMsgCreatePost
		},
	)

	var weightMsgEditPost int
	appParams.GetOrGenerate(cdc, OpWeightMsgEditPost, &weightMsgEditPost, nil,
		func(_ *rand.Rand) {
			weightMsgEditPost = DefaultWeightMsgEditPost
		},
	)

	var weightMsgDeletePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgDeletePost, &weightMsgDeletePost, nil,
		func(_ *rand.Rand) {
			weightMsgDeletePost = DefaultWeightMsgDeletePost
		},
	)

	var weightMsgAddPostAttachment int
	appParams.GetOrGenerate(cdc, OpWeightMsgAddPostAttachment, &weightMsgAddPostAttachment, nil,
		func(_ *rand.Rand) {
			weightMsgAddPostAttachment = DefaultWeightMsgAddPostAttachment
		},
	)

	var weightMsgRemovePostAttachment int
	appParams.GetOrGenerate(cdc, OpWeightMsgRemovePostAttachment, &weightMsgRemovePostAttachment, nil,
		func(_ *rand.Rand) {
			weightMsgRemovePostAttachment = DefaultWeightMsgRemovePostAttachment
		},
	)

	var weightMsgAnswerPoll int
	appParams.GetOrGenerate(cdc, OpWeightMsgAnswerPoll, &weightMsgAnswerPoll, nil,
		func(_ *rand.Rand) {
			weightMsgAnswerPoll = DefaultWeightMsgAnswerPoll
		},
	)

	var weightMsgMovePost int
	appParams.GetOrGenerate(cdc, OpWeightMsgMovePost, &weightMsgMovePost, nil,
		func(_ *rand.Rand) {
			weightMsgMovePost = DefaultWeightMsgMovePost
		},
	)

	var weightMsgRequestPostOwnerTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgRequestPostOwnerTransfer, &weightMsgRequestPostOwnerTransfer, nil,
		func(r *rand.Rand) {
			weightMsgRequestPostOwnerTransfer = DefaultWeightMsgRequestPostOwnerTransfer
		},
	)

	var weightMsgCancelPostOwnerTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgCancelPostOwnerTransfer, &weightMsgCancelPostOwnerTransfer, nil,
		func(r *rand.Rand) {
			weightMsgCancelPostOwnerTransfer = DefaultWeightMsgCancelPostOwnerTransfer
		},
	)

	var weightMsgAcceptPostOwnerTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgAcceptPostOwnerTransfer, &weightMsgAcceptPostOwnerTransfer, nil,
		func(r *rand.Rand) {
			weightMsgAcceptPostOwnerTransfer = DefaultWeightMsgAcceptPostOwnerTransfer
		},
	)

	var weightMsgRefusePostOwnerTransfer int
	appParams.GetOrGenerate(cdc, OpWeightMsgRefusePostOwnerTransfer, &weightMsgRefusePostOwnerTransfer, nil,
		func(r *rand.Rand) {
			weightMsgRefusePostOwnerTransfer = DefaultWeightMsgRefusePostOwnerTransfer
		},
	)

	return sim.WeightedOperations{
		sim.NewWeightedOperation(
			weightMsgCreatePost,
			SimulateMsgCreatePost(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgEditPost,
			SimulateMsgEditPost(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgDeletePost,
			SimulateMsgDeletePost(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAddPostAttachment,
			SimulateMsgAddPostAttachment(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRemovePostAttachment,
			SimulateMsgRemovePostAttachment(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAnswerPoll,
			SimulateMsgAnswerPoll(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgMovePost,
			SimulateMsgMovePost(k, sk, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRequestPostOwnerTransfer,
			SimulateMsgRequestPostOwnerTransfer(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgCancelPostOwnerTransfer,
			SimulateMsgCancelPostOwnerTransfer(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAcceptPostOwnerTransfer,
			SimulateMsgAcceptPostOwnerTransfer(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRefusePostOwnerTransfer,
			SimulateMsgRefusePostOwnerTransfer(k, ak, bk),
		),
	}
}
