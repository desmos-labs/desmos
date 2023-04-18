package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	errors "cosmossdk.io/errors"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// LinkApplication defines a rpc handler method for MsgLinkApplication.
func (k Keeper) LinkApplication(
	goCtx context.Context, msg *types.MsgLinkApplication,
) (*types.MsgLinkApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// See if the link already exists
	link, found, err := k.GetApplicationLink(ctx, msg.Sender, msg.LinkData.Application, msg.LinkData.Username)
	if err != nil {
		return nil, err
	}

	if found && link.IsVerificationOngoing() {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest,
			"verification process of link for application %s and username %s is already happening",
			msg.LinkData.Application, msg.LinkData.Username)
	}

	if err := k.StartProfileConnection(
		ctx, msg.LinkData, msg.CallData, user,
		msg.SourcePort, msg.SourceChannel, msg.TimeoutHeight, msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("Application link created",
		"application", msg.LinkData.Application,
		"username", msg.LinkData.Username,
		"account", msg.Sender)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypesApplicationLinkCreated,
			sdk.NewAttribute(types.AttributeKeyUser, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyApplicationName, msg.LinkData.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, msg.LinkData.Username),
			sdk.NewAttribute(types.AttributeKeyApplicationLinkCreationTime, ctx.BlockTime().Format(time.RFC3339)),
		),
	})

	return &types.MsgLinkApplicationResponse{}, nil
}

// UnlinkApplication defines a rpc method for MsgUnlinkApplication
func (k msgServer) UnlinkApplication(
	goCtx context.Context, msg *types.MsgUnlinkApplication,
) (*types.MsgUnlinkApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the link
	link, found, err := k.GetApplicationLink(ctx, msg.Signer, msg.Application, msg.Username)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, errors.Wrap(sdkerrors.ErrNotFound, "application link not found")
	}

	// Delete the link
	k.DeleteApplicationLink(ctx, link)

	k.Logger(ctx).Info("Application link removed",
		"application", msg.Application,
		"username", msg.Username,
		"account", msg.Signer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeApplicationLinkDeleted,
			sdk.NewAttribute(types.AttributeKeyUser, msg.Signer),
			sdk.NewAttribute(types.AttributeKeyApplicationName, msg.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, msg.Username),
		),
	})

	return &types.MsgUnlinkApplicationResponse{}, nil
}
