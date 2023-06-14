package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"

	"github.com/desmos-labs/desmos/v5/testutil/simtesting"
	"github.com/desmos-labs/desmos/v5/x/posts/keeper"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
	subspacessim "github.com/desmos-labs/desmos/v5/x/subspaces/simulation"
)

// DONTCOVER

// SimulateMsgRequestPostOwnerTransfer tests and runs a single msg request post owner transfer request
func SimulateMsgRequestPostOwnerTransfer(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		data, signer, skip := randomRequestPostOwnerTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRequestPostOwnerTransfer", "skip"), nil, nil
		}

		msg := types.NewMsgRequestPostOwnerTransfer(
			data.SubspaceID,
			data.PostID,
			data.Receiver,
			data.Sender,
		)
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomRequestPostOwnerTransferFields returns the data needed to request a post owner transfer request
func randomRequestPostOwnerTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (request types.PostOwnerTransferRequest, signer simtypes.Account, skip bool) {

	// Get a random post
	posts := k.GetPosts(ctx)
	if len(posts) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}
	post := RandomPost(r, posts)

	if !k.HasProfile(ctx, post.Owner) {
		// Skip because the sender has no profile
		skip = true
		return
	}

	// Get a random receiver
	receiver, _ := simtypes.RandomAcc(r, accs)
	if receiver.Address.String() == post.Owner {
		// Skip because the receiver is already the owner
		skip = true
		return
	}

	if !k.HasProfile(ctx, receiver.Address.String()) {
		// Skip because the receiver has no profile
		skip = true
		return
	}

	if k.HasUserBlocked(ctx, receiver.Address.String(), post.Owner, post.SubspaceID) {
		// Skip because the receiver has blocked the sender
		skip = true
		return
	}

	if k.HasPostOwnerTransferRequest(ctx, post.SubspaceID, post.ID) {
		// Skip because the request has already existed
		skip = true
		return
	}

	acc := subspacessim.GetAccount(post.Owner, accs)
	if acc == nil {
		// Skip because the sender is not an account we have access to
		skip = true
		return
	}
	signer = *acc

	request = types.NewPostOwnerTransferRequest(post.SubspaceID, post.ID, receiver.Address.String(), post.Owner)
	return request, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgCancelPostOwnerTransferRequest tests and runs a single msg cancel post owner transfer request
func SimulateMsgCancelPostOwnerTransferRequest(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		data, signer, skip := randomCancelPostOwnerTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgCancelPostOwnerTransferRequest", "skip"), nil, nil
		}

		msg := types.NewMsgCancelPostOwnerTransferRequest(
			data.SubspaceID,
			data.PostID,
			data.Sender,
		)
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomCancelPostOwnerTransferFields returns the data needed to cancel a post owner transfer request
func randomCancelPostOwnerTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (request types.PostOwnerTransferRequest, signer simtypes.Account, skip bool) {

	// Get a random post owner transfer request
	requests := k.GetAllPostOwnerTransferRequests(ctx)
	if len(requests) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}

	request = RandomPostOwnerTransferRequest(r, requests)

	acc := subspacessim.GetAccount(request.Sender, accs)
	if acc == nil {
		// Skip because the sender is not an account we have access to
		skip = true
		return
	}
	signer = *acc

	return request, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgAcceptPostOwnerTransferRequest tests and runs a single msg accept post owner transfer request
func SimulateMsgAcceptPostOwnerTransferRequest(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		data, signer, skip := randomAcceptPostOwnerTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgAcceptPostOwnerTransferRequest", "skip"), nil, nil
		}

		msg := types.NewMsgAcceptPostOwnerTransferRequest(
			data.SubspaceID,
			data.PostID,
			data.Receiver,
		)
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomAcceptPostOwnerTransferFields returns the data needed to accept a post owner transfer request
func randomAcceptPostOwnerTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (request types.PostOwnerTransferRequest, signer simtypes.Account, skip bool) {

	// Get a random post owner transfer request
	requests := k.GetAllPostOwnerTransferRequests(ctx)
	if len(requests) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}
	request = RandomPostOwnerTransferRequest(r, requests)

	if !k.HasProfile(ctx, request.Receiver) {
		// Skip because the sender has no profile
		skip = true
		return
	}

	acc := subspacessim.GetAccount(request.Receiver, accs)
	if acc == nil {
		// Skip because the sender is not an account we have access to
		skip = true
		return
	}
	signer = *acc

	return request, signer, false
}

// --------------------------------------------------------------------------------------------------------------------

// SimulateMsgRefusePostOwnerTransferRequest tests and runs a single msg refuse post owner transfer request
func SimulateMsgRefusePostOwnerTransferRequest(
	k keeper.Keeper, ak authkeeper.AccountKeeper, bk bankkeeper.Keeper,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simtypes.Account, chainID string,
	) (OperationMsg simtypes.OperationMsg, futureOps []simtypes.FutureOperation, err error) {
		data, signer, skip := randomRefusePostOwnerTransferFields(r, ctx, accs, k)
		if skip {
			return simtypes.NoOpMsg(types.RouterKey, "MsgRefusePostOwnerTransferRequest", "skip"), nil, nil
		}

		msg := types.NewMsgRefusePostOwnerTransferRequest(
			data.SubspaceID,
			data.PostID,
			data.Receiver,
		)
		return simtesting.SendMsg(r, app, ak, bk, msg, ctx, signer)
	}
}

// randomRefusePostOwnerTransferFields returns the data needed to refuse a post owner transfer request
func randomRefusePostOwnerTransferFields(
	r *rand.Rand, ctx sdk.Context, accs []simtypes.Account, k keeper.Keeper,
) (request types.PostOwnerTransferRequest, signer simtypes.Account, skip bool) {

	// Get a random post owner transfer request
	requests := k.GetAllPostOwnerTransferRequests(ctx)
	if len(requests) == 0 {
		// Skip because there are no posts
		skip = true
		return
	}

	request = RandomPostOwnerTransferRequest(r, requests)

	acc := subspacessim.GetAccount(request.Receiver, accs)
	if acc == nil {
		// Skip because the sender is not an account we have access to
		skip = true
		return
	}
	signer = *acc

	return request, signer, false
}
