package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
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
			types.EventTypesApplicationLinkCreated,
			sdk.NewAttribute(types.AttributeKeyUser, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyApplicationName, msg.LinkData.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, msg.LinkData.Username),
			sdk.NewAttribute(types.AttributeKeyApplicationLinkCreationTime, ctx.BlockTime().Format(time.RFC3339)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgLinkApplicationResponse{}, nil
}

func (k msgServer) UnlinkApplication(
	goCtx context.Context, msg *types.MsgUnlinkApplication,
) (*types.MsgUnlinkApplicationResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteApplicationLink(ctx, msg.Signer, msg.Application, msg.Username)
	if err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("Application link removed",
		"application", msg.Application,
		"username", msg.Username,
		"account", msg.Signer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeApplicationLinkDeleted,
			sdk.NewAttribute(types.AttributeKeyUser, msg.Signer),
			sdk.NewAttribute(types.AttributeKeyApplicationName, msg.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, msg.Username),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgUnlinkApplicationResponse{}, nil
}
