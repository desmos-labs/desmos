package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v5/x/posts/types"
)

// RequestPostOwnerTransfer defines the rpc method for Msg/RequestPostOwnerTransfer
func (k msgServer) RequestPostOwnerTransfer(goCtx context.Context, msg *types.MsgRequestPostOwnerTransfer) (*types.MsgRequestPostOwnerTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender has profile
	if !k.HasProfile(ctx, msg.Sender) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post without having a profile")
	}

	// Check if the receiver has profile
	if !k.HasProfile(ctx, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post to the user having no profile")
	}

	// Make sure the receiver has not blocked the sender
	if k.HasUserBlocked(ctx, msg.Receiver, msg.Sender, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "receiver has blocked you")
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Make sure the sender matches the owner
	if post.Owner != msg.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post that you do not own")
	}

	// Check if the post owner transfer request exists
	if k.HasPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request already exists")
	}

	// Save the post owner transfer request
	k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
		msg.SubspaceID,
		msg.PostID,
		msg.Receiver,
		msg.Sender,
	))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeRequestPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgRequestPostOwnerTransferResponse{}, nil
}

// CancelPostOwnerTransfer defines the rpc method for Msg/CancelPostOwnerTransfer
func (k msgServer) CancelPostOwnerTransferRequest(goCtx context.Context, msg *types.MsgCancelPostOwnerTransferRequest) (*types.MsgCancelPostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request sender matches the sender
	if request.Sender != msg.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot cancel a post owner transfer request that you are not the sender")
	}

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeCancelPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgCancelPostOwnerTransferRequestResponse{}, nil
}

// AcceptPostOwnerTransfer defines the rpc method for Msg/AcceptPostOwnerTransfer
func (k msgServer) AcceptPostOwnerTransferRequest(goCtx context.Context, msg *types.MsgAcceptPostOwnerTransferRequest) (*types.MsgAcceptPostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the receiver has profile
	if !k.HasProfile(ctx, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot accept a post owner transfer request without having a profile")
	}

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request receiver matches the receiver
	if request.Receiver != msg.Receiver {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot accept a post owner transfer request that you are not the receiver")
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Make sure the post owner matches the sender
	if post.Owner != request.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "the sender of the post owner transfer request does not own the post")
	}

	// Update the post and validate it
	newPost := types.NewOwnerTransfer(msg.Receiver, ctx.BlockTime()).Update(post)
	err := newPost.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the update
	k.SavePost(ctx, newPost)

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeAcceptPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	})

	return &types.MsgAcceptPostOwnerTransferRequestResponse{}, nil
}

// RefusePostOwnerTransfer defines the rpc method for Msg/RefusePostOwnerTransfer
func (k msgServer) RefusePostOwnerTransferRequest(goCtx context.Context, msg *types.MsgRefusePostOwnerTransferRequest) (*types.MsgRefusePostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request receiver matches the receiver
	if request.Receiver != msg.Receiver {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot refuse a post owner transfer request that you are not the receiver")
	}

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeRefusePostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	})

	return &types.MsgRefusePostOwnerTransferRequestResponse{}, nil
}
