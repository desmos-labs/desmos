package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantUserAllowance defines a rpc method for MsgGrantUserAllowance
func (k msgServer) GrantUserAllowance(goCtx context.Context, msg *types.MsgGrantUserAllowance) (*types.MsgGrantUserAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}
	// Checking for duplicate entry
	_, found, err := k.GetUserGrant(ctx, msg.SubspaceID, msg.Granter, msg.Grantee)
	if found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}
	if err != nil {
		return nil, err
	}

	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return nil, err
	}
	grant, err := types.NewUserGrant(msg.SubspaceID, msg.Granter, msg.Grantee, allowance)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.SaveUserGrant(ctx, grant)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeGrantUserAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGrantee, msg.Grantee),
		),
	})
	return &types.MsgGrantUserAllowanceResponse{}, nil
}

// RevokeUserAllowance defines a rpc method for MsgRevokeUserAllowance
func (k msgServer) RevokeUserAllowance(goCtx context.Context, msg *types.MsgRevokeUserAllowance) (*types.MsgRevokeUserAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.RemoveUserAllowance(ctx, msg.SubspaceID, msg.Granter, msg.Grantee)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeRevokeUserAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGrantee, msg.Grantee),
		),
	})
	return &types.MsgRevokeUserAllowanceResponse{}, nil
}

// GrantGroupAllowance defines a rpc method for MsgGrantGroupAllowance
func (k msgServer) GrantGroupAllowance(goCtx context.Context, msg *types.MsgGrantGroupAllowance) (*types.MsgGrantGroupAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}
	if !k.HasUserGroup(ctx, msg.SubspaceID, msg.GroupID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}
	// Checking for duplicate entry
	if _, found, _ := k.GetGroupGrant(ctx, msg.SubspaceID, msg.Granter, msg.GroupID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
	}

	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return nil, err
	}
	grant, err := types.NewGroupGrant(msg.SubspaceID, msg.Granter, msg.GroupID, allowance)
	if err != nil {
		return nil, err
	}
	err = k.Keeper.SaveGroupGrant(ctx, grant)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeGrantGroupAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})
	return &types.MsgGrantGroupAllowanceResponse{}, nil
}

// RevokeGroupAllowance defines a rpc method for MsgRevokeGroupAllowance
func (k msgServer) RevokeGroupAllowance(goCtx context.Context, msg *types.MsgRevokeGroupAllowance) (*types.MsgRevokeGroupAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.RemoveGroupGrant(ctx, msg.SubspaceID, msg.Granter, msg.GroupID)
	if err != nil {
		return nil, err
	}
	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
		sdk.NewEvent(
			types.EventTypeRevokeGroupAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})
	return &types.MsgRevokeGroupAllowanceResponse{}, nil
}
