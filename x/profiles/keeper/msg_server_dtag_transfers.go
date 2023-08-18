package keeper

import (
	"context"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/desmos-labs/desmos/v6/x/profiles/types"
)

// RequestDTagTransfer defines a rpc method for MsgRequestDTagTransfer
func (k msgServer) RequestDTagTransfer(goCtx context.Context, msg *types.MsgRequestDTagTransfer) (*types.MsgRequestDTagTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the request's receiver has blocked the sender before
	if k.IsUserBlocked(ctx, msg.Receiver, msg.Sender) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address %s has blocked you", msg.Receiver)
	}

	profile, found, err := k.GetProfile(ctx, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "the request recipient does not have a profile")
	}

	dTagToTrade := profile.DTag
	if len(dTagToTrade) == 0 {
		return nil, errors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"the user with address %s doesn't have a profile yet so their DTag cannot be transferred",
			msg.Receiver,
		)
	}

	transferRequest := types.NewDTagTransferRequest(dTagToTrade, msg.Sender, msg.Receiver)
	err = k.SaveDTagTransferRequest(ctx, transferRequest)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeDTagTransferRequest,
			sdk.NewAttribute(types.AttributeKeyDTagToTrade, dTagToTrade),
			sdk.NewAttribute(types.AttributeKeyRequestSender, transferRequest.Sender),
			sdk.NewAttribute(types.AttributeKeyRequestReceiver, transferRequest.Receiver),
		),
	})

	return &types.MsgRequestDTagTransferResponse{}, nil
}

// CancelDTagTransferRequest defines a rpc method for MsgCancelDTagTransferRequest
func (k msgServer) CancelDTagTransferRequest(goCtx context.Context, msg *types.MsgCancelDTagTransferRequest) (*types.MsgCancelDTagTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the request exists
	if !k.HasDTagTransferRequest(ctx, msg.Sender, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", msg.Sender, msg.Receiver)
	}

	// Delete the request
	k.DeleteDTagTransferRequest(ctx, msg.Sender, msg.Receiver)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeDTagTransferCancel,
			sdk.NewAttribute(types.AttributeKeyRequestSender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRequestReceiver, msg.Receiver),
		),
	})

	return &types.MsgCancelDTagTransferRequestResponse{}, nil
}

// AcceptDTagTransferRequest defines a rpc method for MsgAcceptDTagTransferRequest
func (k msgServer) AcceptDTagTransferRequest(goCtx context.Context, msg *types.MsgAcceptDTagTransferRequest) (*types.MsgAcceptDTagTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	request, found, err := k.GetDTagTransferRequest(ctx, msg.Sender, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "no request made from %s", msg.Sender)
	}

	// Get the current owner profile
	currentOwnerProfile, exist, err := k.GetProfile(ctx, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "profile of %s doesn't exist", msg.Receiver)
	}

	// Get the DTag to trade and make sure its correct
	dTagWanted := request.DTagToTrade
	dTagToTrade := currentOwnerProfile.DTag
	if dTagWanted != dTagToTrade {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged")
	}

	// Change the DTag and validate the profile
	currentOwnerProfile.DTag = msg.NewDTag
	err = k.ValidateProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Check for an existent profile of the receiving user
	receiverProfile, exist, err := k.GetProfile(ctx, msg.Sender)
	if err != nil {
		return nil, err
	}

	if exist && msg.NewDTag == receiverProfile.DTag {
		err = k.storeProfileWithoutDTagCheck(ctx, currentOwnerProfile)
		if err != nil {
			return nil, err
		}
	} else {
		err = k.Keeper.SaveProfile(ctx, currentOwnerProfile)
		if err != nil {
			return nil, err
		}
	}

	if !exist {
		add, err := sdk.AccAddressFromBech32(msg.Sender)
		if err != nil {
			return nil, err
		}

		senderAcc := k.ak.GetAccount(ctx, add)
		if senderAcc == nil {
			senderAcc = authtypes.NewBaseAccountWithAddress(add)
		}

		receiverProfile, err = types.NewProfileFromAccount(dTagToTrade, senderAcc, ctx.BlockTime())
		if err != nil {
			return nil, err
		}
	} else {
		receiverProfile.DTag = dTagToTrade
	}

	// Validate the receiver profile
	err = k.ValidateProfile(ctx, receiverProfile)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the receiver profile
	err = k.Keeper.SaveProfile(ctx, receiverProfile)
	if err != nil {
		return nil, err
	}

	k.DeleteAllUserIncomingDTagTransferRequests(ctx, msg.Receiver)
	k.DeleteAllUserIncomingDTagTransferRequests(ctx, msg.Sender)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeDTagTransferAccept,
			sdk.NewAttribute(types.AttributeKeyDTagToTrade, dTagToTrade),
			sdk.NewAttribute(types.AttributeKeyNewDTag, msg.NewDTag),
			sdk.NewAttribute(types.AttributeKeyRequestSender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRequestReceiver, msg.Receiver),
		),
	})

	return &types.MsgAcceptDTagTransferRequestResponse{}, nil
}

// RefuseDTagTransferRequest defines a rpc method for MsgRefuseDTagTransferRequest
func (k msgServer) RefuseDTagTransferRequest(goCtx context.Context, msg *types.MsgRefuseDTagTransferRequest) (*types.MsgRefuseDTagTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the request exists
	if !k.HasDTagTransferRequest(ctx, msg.Sender, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "request from %s to %s not found", msg.Sender, msg.Receiver)
	}

	// Delete the request
	k.DeleteDTagTransferRequest(ctx, msg.Sender, msg.Receiver)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeDTagTransferRefuse,
			sdk.NewAttribute(types.AttributeKeyRequestSender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyRequestReceiver, msg.Receiver),
		),
	})

	return &types.MsgRefuseDTagTransferRequestResponse{}, nil
}
