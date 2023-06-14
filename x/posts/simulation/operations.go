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
	DefaultWeightMsgCreatePost                     int = 80
	DefaultWeightMsgEditPost                       int = 40
	DefaultWeightMsgDeletePost                     int = 20
	DefaultWeightMsgAddPostAttachment              int = 50
	DefaultWeightMsgRemovePostAttachment           int = 50
	DefaultWeightMsgAnswerPoll                     int = 50
	DefaultWeightMsgMovePost                       int = 10
	DefaultWeightMsgRequestPostOwnerTransfer       int = 10
	DefaultWeightMsgCancelPostOwnerTransferRequest int = 10
	DefaultWeightMsgAcceptPostOwnerTransferRequest int = 10
	DefaultWeightMsgRefusePostOwnerTransferRequest int = 10

	OpWeightMsgCreatePost                     = "op_weight_msg_create_post"
	OpWeightMsgEditPost                       = "op_weight_msg_edit_post"
	OpWeightMsgDeletePost                     = "op_weight_msg_delete_post"
	OpWeightMsgAddPostAttachment              = "op_weight_msg_add_post_attachment"
	OpWeightMsgRemovePostAttachment           = "op_weight_msg_remove_post_attachment"
	OpWeightMsgAnswerPoll                     = "op_weight_msg_answer_poll"
	OpWeightMsgMovePost                       = "op_weight_msg_move_post"
	OpWeightMsgRequestPostOwnerTransfer       = "op_weight_msg_request_post_owner_transfer"
	OpWeightMsgCancelPostOwnerTransferRequest = "op_weight_msg_cancel_post_owner_transfer_request"
	OpWeightMsgAcceptPostOwnerTransferRequest = "op_weight_msg_accept_post_owner_transfer_request"
	OpWeightMsgRefusePostOwnerTransferRequest = "op_weight_msg_refuse_post_owner_transfer_request"

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

	var weightMsgCancelPostOwnerTransferRequest int
	appParams.GetOrGenerate(cdc, OpWeightMsgCancelPostOwnerTransferRequest, &weightMsgCancelPostOwnerTransferRequest, nil,
		func(r *rand.Rand) {
			weightMsgCancelPostOwnerTransferRequest = DefaultWeightMsgCancelPostOwnerTransferRequest
		},
	)

	var weightMsgAcceptPostOwnerTransferRequest int
	appParams.GetOrGenerate(cdc, OpWeightMsgAcceptPostOwnerTransferRequest, &weightMsgAcceptPostOwnerTransferRequest, nil,
		func(r *rand.Rand) {
			weightMsgAcceptPostOwnerTransferRequest = DefaultWeightMsgAcceptPostOwnerTransferRequest
		},
	)

	var weightMsgRefusePostOwnerTransferRequest int
	appParams.GetOrGenerate(cdc, OpWeightMsgRefusePostOwnerTransferRequest, &weightMsgRefusePostOwnerTransferRequest, nil,
		func(r *rand.Rand) {
			weightMsgRefusePostOwnerTransferRequest = DefaultWeightMsgRefusePostOwnerTransferRequest
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
			weightMsgCancelPostOwnerTransferRequest,
			SimulateMsgCancelPostOwnerTransferRequest(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgAcceptPostOwnerTransferRequest,
			SimulateMsgAcceptPostOwnerTransferRequest(k, ak, bk),
		),
		sim.NewWeightedOperation(
			weightMsgRefusePostOwnerTransferRequest,
			SimulateMsgRefusePostOwnerTransferRequest(k, ak, bk),
		),
	}
}
