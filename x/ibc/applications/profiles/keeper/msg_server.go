package keeper

import (
	"context"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

// CreateApplicationLink defines a rpc handler method for MsgConnectProfile.
func (k Keeper) CreateApplicationLink(
	goCtx context.Context, msg *types.MsgCreateApplicationLink,
) (*types.MsgCreateApplicationLinkResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	if err := k.StartProfileConnection(
		ctx, msg.Application, msg.VerificationData, user,
		msg.SourcePort, msg.SourceChannel, msg.TimeoutHeight, msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("IBC profile connection",
		"application", msg.Application.Name,
		"username", msg.Application.Username,
		"account", msg.Sender)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeConnectProfile,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
		),
	})

	return &types.MsgCreateApplicationLinkResponse{}, nil
}
