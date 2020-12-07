package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the profiles MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

func (k msgServer) SaveProfile(goCtx context.Context, msg *types.MsgSaveProfile) (*types.MsgSaveProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	profile, found := k.GetProfile(ctx, msg.Creator)

	// Create a new profile if not found
	if !found {
		profile = types.NewProfile(
			msg.Dtag,
			"",
			"",
			types.NewPictures("", ""),
			ctx.BlockTime(),
			msg.Creator,
		)
	}

	// If the DTag changes, delete all the previous DTag transfer requests
	if profile.Dtag != msg.Dtag {
		k.DeleteAllDTagTransferRequests(ctx, msg.Creator)
	}

	// Update the existing profile with the values provided from the user
	updated, err := profile.Update(types.NewProfile(
		msg.Dtag,
		msg.Moniker,
		msg.Bio,
		types.NewPictures(msg.ProfilePicture, msg.CoverPicture),
		time.Time{},
		types.DoNotModify,
	))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Validate the profile
	err = k.ValidateProfile(ctx, updated)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the profile
	if err := k.StoreProfile(ctx, updated); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeProfileSaved,
		sdk.NewAttribute(types.AttributeProfileDtag, updated.Dtag),
		sdk.NewAttribute(types.AttributeProfileCreator, updated.Creator),
		sdk.NewAttribute(types.AttributeProfileCreationTime, updated.CreationDate.Format(time.RFC3339Nano)),
	))

	return &types.MsgSaveProfileResponse{}, nil
}

func (k msgServer) DeleteProfile(goCtx context.Context, msg *types.MsgDeleteProfile) (*types.MsgDeleteProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveProfile(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeProfileDeleted,
		sdk.NewAttribute(types.AttributeProfileCreator, msg.Creator),
	))

	return &types.MsgDeleteProfileResponse{}, nil
}

func (k msgServer) RequestDTagTransfer(goCtx context.Context, msg *types.MsgRequestDTagTransfer) (*types.MsgRequestDTagTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the request's receiver has blocked the sender before
	if k.IsUserBlocked(ctx, msg.Receiver, msg.Sender) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the user with address %s has blocked you", msg.Receiver)
	}

	dtagToTrade := k.GetDtagFromAddress(ctx, msg.Receiver)
	if len(dtagToTrade) == 0 {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrInvalidRequest,
			"the user with address %s doesn't have a profile yet so their DTag cannot be transferred",
			msg.Receiver,
		)
	}

	transferRequest := types.NewDTagTransferRequest(dtagToTrade, msg.Sender, msg.Receiver)
	err := k.SaveDTagTransferRequest(ctx, transferRequest)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferRequest,
		sdk.NewAttribute(types.AttributeDTagToTrade, dtagToTrade),
		sdk.NewAttribute(types.AttributeRequestSender, transferRequest.Sender),
		sdk.NewAttribute(types.AttributeRequestReceiver, transferRequest.Receiver),
	))

	return &types.MsgRequestDTagTransferResponse{}, nil
}

func (k msgServer) CancelDTagTransfer(goCtx context.Context, msg *types.MsgCancelDTagTransfer) (*types.MsgCancelDTagTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteDTagTransferRequest(ctx, msg.Sender, msg.Receiver)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferCancel,
		sdk.NewAttribute(types.AttributeRequestSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRequestReceiver, msg.Receiver),
	))

	return &types.MsgCancelDTagTransferResponse{}, nil
}

func (k msgServer) AcceptDTagTransfer(goCtx context.Context, msg *types.MsgAcceptDTagTransfer) (*types.MsgAcceptDTagTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	requests := k.GetUserIncomingDTagTransferRequests(ctx, msg.Receiver)

	// Check if the receiving user request is present, if not return error
	found := false
	var dTagWanted string
	for _, req := range requests {
		if req.Sender == msg.Sender {
			dTagWanted = req.DtagToTrade
			found = true
			break
		}
	}

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no request made from %s", msg.Sender)
	}

	// Get the current owner profile
	currentOwnerProfile, exist := k.GetProfile(ctx, msg.Receiver)
	if !exist {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "profile of %s doesn't exist", msg.Receiver)
	}

	// Get the DTag to trade and make sure its correct
	dTagToTrade := currentOwnerProfile.Dtag
	if dTagWanted != dTagToTrade {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged")
	}

	// Change the DTag and validate the profile
	currentOwnerProfile.Dtag = msg.NewDtag
	err := k.ValidateProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the profile
	err = k.StoreProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, err
	}

	// Check for an existent profile of the receiving user
	receiverProfile, exist := k.GetProfile(ctx, msg.Sender)
	if !exist {
		receiverProfile = types.NewProfile(
			dTagToTrade,
			"",
			"",
			types.NewPictures("", ""),
			ctx.BlockTime(),
			msg.Sender,
		)
	} else {
		receiverProfile.Dtag = dTagToTrade
	}

	// Validate the receiver profile
	err = k.ValidateProfile(ctx, receiverProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the receiver profile
	err = k.StoreProfile(ctx, receiverProfile)
	if err != nil {
		return nil, err
	}

	k.DeleteAllDTagTransferRequests(ctx, msg.Receiver)
	k.DeleteAllDTagTransferRequests(ctx, msg.Sender)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferAccept,
		sdk.NewAttribute(types.AttributeDTagToTrade, dTagToTrade),
		sdk.NewAttribute(types.AttributeNewDTag, msg.NewDtag),
		sdk.NewAttribute(types.AttributeRequestSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRequestReceiver, msg.Receiver),
	))

	return &types.MsgAcceptDTagTransferResponse{}, nil
}

func (k msgServer) RefuseDTagTransfer(goCtx context.Context, msg *types.MsgRefuseDTagTransfer) (*types.MsgRefuseDTagTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteDTagTransferRequest(ctx, msg.Sender, msg.Receiver)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferRefuse,
		sdk.NewAttribute(types.AttributeRequestSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRequestReceiver, msg.Receiver),
	))

	return &types.MsgRefuseDTagTransferResponse{}, nil
}
