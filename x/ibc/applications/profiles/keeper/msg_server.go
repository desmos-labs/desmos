package keeper

import (
	"context"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ types.MsgServer = Keeper{}

// ConnectProfile defines a rpc handler method for MsgConnectProfile.
func (k Keeper) ConnectProfile(goCtx context.Context, msg *types.MsgConnectProfile) (*types.MsgConnectProfileResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	user, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	verificationData := types.NewVerificationData(msg.VerificationMethod, msg.VerificationValue)

	if err := k.StartProfileConnection(
		ctx, msg.SourcePort, msg.SourceChannel,
		verificationData, user, msg.FeePayer,
		msg.TimeoutHeight, msg.TimeoutTimestamp,
	); err != nil {
		return nil, err
	}

	k.Logger(ctx).Info("IBC profile connection",
		"application", msg.Application,
		"username", msg.Username,
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

	return &types.MsgConnectProfileResponse{}, nil
}
