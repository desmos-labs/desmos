package keeper

import (
	"context"
	"time"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

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

	profile, found, err := k.GetProfile(ctx, msg.Creator)
	if err != nil {
		return nil, err
	}

	// Create a new profile if not found
	if !found {
		addr, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, err
		}

		profile, err = types.NewProfile(
			msg.DTag,
			"",
			"",
			types.NewPictures("", ""),
			ctx.BlockTime(),
			k.ak.GetAccount(ctx, addr),
		)
		if err != nil {
			return nil, err
		}
	}

	// If the DTag changes, delete all the previous DTag transfer requests
	if profile.DTag != msg.DTag {
		k.DeleteAllDTagTransferRequests(ctx, msg.Creator)
	}

	// Update the existing profile with the values provided from the user
	updated, err := profile.Update(types.NewProfileUpdate(
		msg.DTag,
		msg.Username,
		msg.Bio,
		types.NewPictures(msg.ProfilePicture, msg.CoverPicture),
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
		sdk.NewAttribute(types.AttributeProfileDTag, updated.DTag),
		sdk.NewAttribute(types.AttributeProfileCreator, updated.GetAddress().String()),
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

	profile, found, err := k.GetProfile(ctx, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the request recipient does not have a profile yet")
	}

	dTagToTrade := profile.DTag
	if len(dTagToTrade) == 0 {
		return nil, sdkerrors.Wrapf(
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

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeDTagTransferRequest,
		sdk.NewAttribute(types.AttributeDTagToTrade, dTagToTrade),
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
			dTagWanted = req.DTagToTrade
			found = true
			break
		}
	}

	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "no request made from %s", msg.Sender)
	}

	// Get the current owner profile
	currentOwnerProfile, exist, err := k.GetProfile(ctx, msg.Receiver)
	if err != nil {
		return nil, err
	}

	if !exist {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "profile of %s doesn't exist", msg.Receiver)
	}

	// Get the DTag to trade and make sure its correct
	dTagToTrade := currentOwnerProfile.DTag
	if dTagWanted != dTagToTrade {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "the owner's DTag is different from the one to be exchanged")
	}

	// Change the DTag and validate the profile
	currentOwnerProfile.DTag = msg.NewDTag
	err = k.ValidateProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the profile
	err = k.StoreProfile(ctx, currentOwnerProfile)
	if err != nil {
		return nil, err
	}

	// Check for an existent profile of the receiving user
	receiverProfile, exist, err := k.GetProfile(ctx, msg.Sender)
	if err != nil {
		return nil, err
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

		receiverProfile, err = types.NewProfile(
			dTagToTrade,
			"",
			"",
			types.NewPictures("", ""),
			ctx.BlockTime(),
			senderAcc,
		)
		if err != nil {
			return nil, err
		}
	} else {
		receiverProfile.DTag = dTagToTrade
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
		sdk.NewAttribute(types.AttributeNewDTag, msg.NewDTag),
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

func (k msgServer) CreateRelationship(goCtx context.Context, msg *types.MsgCreateRelationship) (*types.CreateRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the receiver has blocked the sender before
	if k.HasUserBlocked(ctx, msg.Receiver, msg.Sender, msg.Subspace) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "The user with address %s has blocked you", msg.Receiver)
	}

	// Save the relationship
	err := k.SaveRelationship(ctx, types.NewRelationship(msg.Sender, msg.Receiver, msg.Subspace))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	return &types.CreateRelationshipResponse{}, nil
}

func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.DeleteRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveRelationship(ctx, types.NewRelationship(msg.User, msg.Counterparty, msg.Subspace))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.User),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	return &types.DeleteRelationshipResponse{}, nil
}

func (k msgServer) BlockUser(goCtx context.Context, msg *types.MsgBlockUser) (*types.BlockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userBlock := types.NewUserBlock(msg.Blocker, msg.Blocked, msg.Reason, msg.Subspace)
	err := k.SaveUserBlock(ctx, userBlock)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
		sdk.NewAttribute(types.AttributeUserBlockReason, msg.Reason),
	))

	return &types.BlockUserResponse{}, nil
}

func (k msgServer) UnblockUser(goCtx context.Context, msg *types.MsgUnblockUser) (*types.UnblockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteUserBlock(ctx, msg.Blocker, msg.Blocked, msg.Subspace)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
	))

	return &types.UnblockUserResponse{}, nil
}
