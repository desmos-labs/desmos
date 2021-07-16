package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (k msgServer) BlockUser(goCtx context.Context, msg *types.MsgBlockUser) (*types.MsgBlockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckUserPermissionsInSubspace(ctx, msg.Subspace, msg.Blocker); err != nil {
		return nil, err
	}

	if err := k.CheckUserPermissionsInSubspace(ctx, msg.Subspace, msg.Blocked); err != nil {
		return nil, err
	}

	userBlock := types.NewUserBlock(msg.Blocker, msg.Blocked, msg.Reason, msg.Subspace)
	err := k.SaveUserBlock(ctx, userBlock)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUser,
		sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeKeySubspace, msg.Subspace),
		sdk.NewAttribute(types.AttributeKeyUserBlockReason, msg.Reason),
	))

	return &types.MsgBlockUserResponse{}, nil
}

func (k msgServer) UnblockUser(goCtx context.Context, msg *types.MsgUnblockUser) (*types.MsgUnblockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteUserBlock(ctx, msg.Blocker, msg.Blocked, msg.Subspace)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeKeyUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeKeyUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeKeySubspace, msg.Subspace),
	))

	return &types.MsgUnblockUserResponse{}, nil
}
