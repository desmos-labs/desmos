package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

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

// CreateSubspace defines a rpc method for MsgCreateSubspace
func (k msgServer) CreateSubspace(goCtx context.Context, msg *types.MsgCreateSubspace) (*types.MsgCreateSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the next subspace ID
	subspaceID, err := k.GetSubspaceID(ctx)
	if err != nil {
		return nil, err
	}

	// Create and validate the subspace
	subspace := types.NewSubspace(subspaceID, msg.Name, msg.Description, msg.Treasury, msg.Owner, msg.Creator, ctx.BlockTime())
	if err := subspace.Validate(); err != nil {
		return nil, err
	}

	// Save the subspace
	k.SaveSubspace(ctx, subspace)

	// Update the id for the next subspace
	k.SetSubspaceID(ctx, subspace.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, subspace.Creator),
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
		SubspaceID: subspace.ID,
	}, nil
}

// EditSubspace defines a rpc method for MsgEditSubspace
func (k msgServer) EditSubspace(goCtx context.Context, msg *types.MsgEditSubspace) (*types.MsgEditSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the if the subspace exists
	subspace, exists := k.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to edit
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionEditSubspace) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage this subspace")
	}

	// Update the subspace and validate it
	updated := subspace.Update(types.NewSubspaceUpdate(msg.Name, msg.Description, msg.Owner, msg.Treasury))
	err = updated.Validate()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the subspace
	k.SaveSubspace(ctx, updated)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeEditSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", updated.ID)),
		),
	})

	return &types.MsgEditSubspaceResponse{}, nil
}

// DeleteSubspace defines a rpc method for MsgDeleteSubspace
func (k msgServer) DeleteSubspace(goCtx context.Context, msg *types.MsgDeleteSubspace) (*types.MsgDeleteSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to edit
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionDeleteSubspace) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage this subspace")
	}

	// Delete the subspace
	k.Keeper.DeleteSubspace(ctx, msg.SubspaceID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeleteSubspace,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgDeleteSubspaceResponse{}, nil
}

// CreateUserGroup defines a rpc method for MsgCreateUserGroup
func (k msgServer) CreateUserGroup(goCtx context.Context, msg *types.MsgCreateUserGroup) (*types.MsgCreateUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.SubspaceID)
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", msg.Creator)
	}

	// Check the permissions to create a group
	if !k.HasPermission(ctx, msg.SubspaceID, creator, types.PermissionManageGroups) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in this subspace")
	}
	if !k.HasPermission(ctx, msg.SubspaceID, creator, types.PermissionSetPermissions) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the default permissions are valid
	if !types.ArePermissionsValid(msg.DefaultPermissions) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Get the next group ID
	groupID, err := k.GetGroupID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create and validate the group
	group := types.NewUserGroup(msg.SubspaceID, groupID, msg.Name, msg.Description, msg.DefaultPermissions)
	if err := group.Validate(); err != nil {
		return nil, err
	}

	// Save the group
	k.SaveUserGroup(ctx, group)

	// Update the id for the next group
	k.SetGroupID(ctx, msg.SubspaceID, group.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeCreateUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", group.ID)),
		),
	})

	return &types.MsgCreateUserGroupResponse{GroupID: groupID}, nil
}

// EditUserGroup defines a rpc method for MsgEditUserGroup
func (k msgServer) EditUserGroup(goCtx context.Context, msg *types.MsgEditUserGroup) (*types.MsgEditUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to create a group
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionManageGroups) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in this subspace")
	}

	// Update the group and validate it
	updated := group.Update(types.NewGroupUpdate(msg.Name, msg.Description))
	err = updated.Validate()
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the updated group
	k.SaveUserGroup(ctx, updated)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeEditUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})

	return &types.MsgEditUserGroupResponse{}, nil
}

// SetUserGroupPermissions defines a rpc method for MsgSetUserGroupPermissions
func (k msgServer) SetUserGroupPermissions(goCtx context.Context, msg *types.MsgSetUserGroupPermissions) (*types.MsgSetUserGroupPermissionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, found := k.GetSubspace(ctx, msg.SubspaceID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionSetPermissions) {
		return nil, sdkerrors.Wrapf(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the permission is valid
	if !types.ArePermissionsValid(msg.Permissions) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Make sure that the user is not part of the group they want to change the permissions for, unless they are the owner
	if subspace.Owner != msg.Signer && k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, signer) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "cannot set the permissions for a group you are part of")
	}

	// Set the group permissions and store the group
	group.Permissions = msg.Permissions
	k.SaveUserGroup(ctx, group)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
		),
		sdk.NewEvent(
			types.EventTypeSetUserGroupPermissions,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})

	return &types.MsgSetUserGroupPermissionsResponse{}, nil
}

// DeleteUserGroup defines a rpc method for MsgDeleteUserGroup
func (k msgServer) DeleteUserGroup(goCtx context.Context, msg *types.MsgDeleteUserGroup) (*types.MsgDeleteUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	if !k.HasUserGroup(ctx, msg.SubspaceID, msg.GroupID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check for permissions
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionManageGroups) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot delete user groups in this subspace")
	}

	// Delete the group
	k.Keeper.DeleteUserGroup(ctx, msg.SubspaceID, msg.GroupID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeleteUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})

	return &types.MsgDeleteUserGroupResponse{}, nil
}

// AddUserToUserGroup defines a rpc method for MsgAddUserToUserGroup
func (k msgServer) AddUserToUserGroup(goCtx context.Context, msg *types.MsgAddUserToUserGroup) (*types.MsgAddUserToUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	if !k.HasUserGroup(ctx, msg.SubspaceID, msg.GroupID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionSetPermissions) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, err
	}

	// Check if the user is already part of the group
	if k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, user) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user is already part of group %d", msg.GroupID)
	}

	// Set the user group
	err = k.AddUserToGroup(ctx, msg.SubspaceID, msg.GroupID, user)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeAddUserToGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgAddUserToUserGroupResponse{}, nil
}

// RemoveUserFromUserGroup defines a rpc method for MsgRemoveUserFromUserGroup
func (k msgServer) RemoveUserFromUserGroup(goCtx context.Context, msg *types.MsgRemoveUserFromUserGroup) (*types.MsgRemoveUserFromUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	if !k.HasUserGroup(ctx, msg.SubspaceID, msg.GroupID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionSetPermissions) {
		return nil, sdkerrors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, err
	}

	// Check if the user is already part of the group
	if !k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, user) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "user is not part of group %d", msg.GroupID)
	}

	// Remove the user group
	k.RemoveUserFromGroup(ctx, msg.SubspaceID, msg.GroupID, user)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeRemoveUserFromGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgRemoveUserFromUserGroupResponse{}, nil
}

// SetUserPermissions defines a rpc method for MsgSetUserPermissions
func (k msgServer) SetUserPermissions(goCtx context.Context, msg *types.MsgSetUserPermissions) (*types.MsgSetUserPermissionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address: %s", msg.User)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, msg.SubspaceID, signer, types.PermissionSetPermissions) {
		return nil, sdkerrors.Wrapf(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the permission is valid
	if !types.ArePermissionsValid(msg.Permissions) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Set the permissions
	if msg.Permissions == nil {
		// Remove the permission to clear the store if empty permissions are provided
		k.RemoveUserPermissions(ctx, msg.SubspaceID, user)
	} else {
		k.Keeper.SetUserPermissions(ctx, msg.SubspaceID, user, msg.Permissions)
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, signer.String()),
		),
		sdk.NewEvent(
			types.EventTypeSetUserPermissions,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUser, user.String()),
		),
	})

	return &types.MsgSetUserPermissionsResponse{}, nil
}
