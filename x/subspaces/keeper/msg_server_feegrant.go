package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

// GrantAllowance defines a rpc method for MsgGrantAllowance
func (k msgServer) GrantAllowance(goCtx context.Context, msg *types.MsgGrantAllowance) (*types.MsgGrantAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, msg.Granter, types.PermissionManageAllowances) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage allowances in this subspace")
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
	}

	switch grantee := msg.Grantee.GetCachedValue().(type) {
	case *types.UserGrantee:
		if k.HasUserGrant(ctx, msg.SubspaceID, grantee.User) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
		}
		events = events.AppendEvent(sdk.NewEvent(
			types.EventTypeGrantAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyUserGrantee, grantee.User),
		))

	case *types.GroupGrantee:
		if !k.HasUserGroup(ctx, msg.SubspaceID, grantee.GroupID) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", grantee.GroupID)
		}
		if k.HasGroupGrant(ctx, msg.SubspaceID, grantee.GroupID) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "fee allowance already exists")
		}
		events = events.AppendEvent(sdk.NewEvent(
			types.EventTypeGrantAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGroupGrantee, fmt.Sprintf("%d", grantee.GroupID)),
		))

	default:
		panic(fmt.Errorf("unsupported type %T", grantee))
	}

	allowance, err := msg.GetUnpackedAllowance()
	if err != nil {
		return nil, err
	}
	k.Keeper.SaveGrant(ctx, types.NewGrant(msg.SubspaceID, msg.Granter, msg.Grantee.GetCachedValue().(types.Grantee), allowance))
	ctx.EventManager().EmitEvents(events)
	return &types.MsgGrantAllowanceResponse{}, nil
}

// RevokeAllowance defines a rpc method for MsgRevokeAllowance
func (k msgServer) RevokeAllowance(goCtx context.Context, msg *types.MsgRevokeAllowance) (*types.MsgRevokeAllowanceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, msg.Granter, types.PermissionManageAllowances) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage allowances in this subspace")
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Granter),
		),
	}

	switch grantee := msg.Grantee.GetCachedValue().(type) {
	case *types.UserGrantee:
		if !k.HasUserGrant(ctx, msg.SubspaceID, grantee.User) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "fee allowance does not exist")
		}
		k.DeleteUserGrant(ctx, msg.SubspaceID, grantee.User)
		events = events.AppendEvent(sdk.NewEvent(
			types.EventTypeRevokeAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyUserGrantee, grantee.User),
		))

	case *types.GroupGrantee:
		if !k.HasGroupGrant(ctx, msg.SubspaceID, grantee.GroupID) {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "fee allowance does not exist")
		}
		k.DeleteGroupGrant(ctx, msg.SubspaceID, grantee.GroupID)
		events = events.AppendEvent(sdk.NewEvent(
			types.EventTypeRevokeAllowance,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyGranter, msg.Granter),
			sdk.NewAttribute(types.AttributeKeyGroupGrantee, fmt.Sprintf("%d", grantee.GroupID)),
		))
	default:
		panic(fmt.Errorf("unsupported type %T", grantee))
	}
	ctx.EventManager().EmitEvents(events)
	return &types.MsgRevokeAllowanceResponse{}, nil
}
