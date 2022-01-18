package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface
// for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) CreateSubspace(goCtx context.Context, msg *types.MsgCreateSubspace) (*types.MsgCreateSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the next subspace ID
	subspaceID, err := k.GetSubspaceID(ctx)
	if err != nil {
		return nil, err
	}

	// Create and validate the subspace
	subspace := types.NewSubspace(subspaceID, msg.Name, msg.Description, msg.Owner, msg.Creator, msg.Treasury, ctx.BlockTime())
	if err := subspace.Validate(); err != nil {
		return nil, err
	}

	// Save the subspace
	err = k.SaveSubspace(ctx, subspace)
	if err != nil {
		return nil, err
	}

	// Update the id for the next subspace
	k.SetSubspaceID(ctx, subspace.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeCreateSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", subspaceID)),
			sdk.NewAttribute(types.AttributeKeySubspaceName, subspace.Name),
			sdk.NewAttribute(types.AttributeKeySubspaceCreator, subspace.Creator),
			sdk.NewAttribute(types.AttributeKeyCreationTime, subspace.CreationTime.Format(time.RFC3339)),
		),
	})

	return &types.MsgCreateSubspaceResponse{
		SubspaceID: subspaceID,
	}, nil
}

func (k msgServer) EditSubspace(goCtx context.Context, msg *types.MsgEditSubspace) (*types.MsgEditSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the permission to edit
	hasPermissions, err := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionChangeInfo)
	if err != nil {
		return nil, err
	}
	if !hasPermissions {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot edit this subspace")
	}

	// Check the if the subspace exists
	subspace, exists := k.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %s not found", msg.SubspaceID)
	}

	// Update the subspace and validate it
	updated := subspace.Update(types.NewSubspaceUpdate(msg.Name, msg.Description, msg.Owner, msg.Treasury))
	err = updated.Validate()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the subspace
	err = k.SaveSubspace(ctx, updated)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeEditSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", updated.ID)),
		),
	})

	return &types.MsgEditSubspaceResponse{}, nil
}

func (k msgServer) CreateUserGroup(goCtx context.Context, msg *types.MsgCreateUserGroup) (*types.MsgCreateUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the permission first
	hasPermission, err := k.HasPermission(ctx, msg.SubspaceID, msg.Creator, types.PermissionSetPermissions)
	if err != nil {
		return nil, err
	}
	if !hasPermission {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot create user groups in this subspace")
	}

	// Check if there is another group
	if k.HasGroup(ctx, msg.SubspaceID, msg.GroupName) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with name %s already exists", msg.GroupName)
	}

	// Store the group
	k.SaveUserGroup(ctx, msg.SubspaceID, msg.GroupName, msg.DefaultPermissions)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeCreateUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupName, msg.GroupName),
		),
	})

	return &types.MsgCreateUserGroupResponse{}, nil
}

func (k msgServer) DeleteUserGroup(goCtx context.Context, msg *types.MsgDeleteUserGroup) (*types.MsgDeleteUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check for permissions
	hasPermission, err := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionSetPermissions)
	if err != nil {
		return nil, err
	}
	if !hasPermission {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot delete user groups in this subspace")
	}

	// Check if the group exists
	if !k.HasGroup(ctx, msg.SubspaceID, msg.GroupName) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %s could not be found", msg.GroupName)
	}

	// Delete the group
	k.Keeper.DeleteUserGroup(ctx, msg.SubspaceID, msg.GroupName)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeleteUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupName, msg.GroupName),
		),
	})

	return &types.MsgDeleteUserGroupResponse{}, nil
}

func (k msgServer) AddUserToUserGroup(goCtx context.Context, msg *types.MsgAddUserToUserGroup) (*types.MsgAddUserToUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the permissions
	hasPermission, err := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionSetPermissions)
	if err != nil {
		return nil, err
	}
	if !hasPermission {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, err
	}

	// Check if the user is already part of the group
	if k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupName, user) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user is already part of group %s", msg.GroupName)
	}

	// Set the user group
	err = k.AddUserToGroup(ctx, msg.SubspaceID, msg.GroupName, user)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeEditSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupName, msg.GroupName),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgAddUserToUserGroupResponse{}, nil
}

func (k msgServer) RemoveUserFromUserGroup(goCtx context.Context, msg *types.MsgRemoveUserFromUserGroup) (*types.MsgRemoveUserFromUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the permissions
	hasPermission, err := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionSetPermissions)
	if err != nil {
		return nil, err
	}
	if !hasPermission {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, err
	}

	// Check if the user is already part of the group
	if !k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupName, user) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user is not part of group %s", msg.GroupName)
	}

	// Set the user group
	err = k.RemoveUserFromGroup(ctx, msg.SubspaceID, msg.GroupName, user)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeEditSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupName, msg.GroupName),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgRemoveUserFromUserGroupResponse{}, nil
}

func (k msgServer) SetPermissions(goCtx context.Context, msg *types.MsgSetPermissions) (*types.MsgSetPermissionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the permissions
	hasPermission, err := k.HasPermission(ctx, msg.SubspaceID, msg.Signer, types.PermissionSetPermissions)
	if err != nil {
		return nil, err
	}
	if !hasPermission {
		return nil, sdkerrors.Wrapf(types.ErrPermissionDenied, "you cannot set other users permissions")
	}

	// Set the permissions
	k.Keeper.SetPermissions(ctx, msg.SubspaceID, msg.Target, msg.Permissions)

	return &types.MsgSetPermissionsResponse{}, nil
}
