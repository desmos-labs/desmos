package keeper

import (
	"context"
	"fmt"
	"time"

	"github.com/desmos-labs/desmos/v5/x/subspaces/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
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
	subspace := types.NewSubspace(subspaceID, msg.Name, msg.Description, types.GetTreasuryAddress(subspaceID).String(), msg.Owner, msg.Creator, ctx.BlockTime(), nil)
	if err := subspace.Validate(); err != nil {
		return nil, err
	}
	// Create a treasury account for subspace
	k.createAccountIfNotExists(ctx, subspace.Treasury)

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

	// Check if the subspace exists
	subspace, exists := k.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to edit
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionEditSubspace) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage this subspace")
	}

	// Update the subspace and validate it
	updated := subspace.Update(types.NewSubspaceUpdate(msg.Name, msg.Description, msg.Owner))
	err = updated.Validate()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
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

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to edit
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionDeleteSubspace) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage this subspace")
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

// CreateSection defines a rpc method for MsgCreateSection
func (k msgServer) CreateSection(goCtx context.Context, msg *types.MsgCreateSection) (*types.MsgCreateSectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the parent section
	if !k.HasSection(ctx, msg.SubspaceID, msg.ParentID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.ParentID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", msg.Creator)
	}

	// Check the permission to manage sections
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionManageSections) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage sections within this subspace")
	}

	// Get the next section ID
	sectionID, err := k.GetNextSectionID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create and validate the section
	section := types.NewSection(msg.SubspaceID, sectionID, msg.ParentID, msg.Name, msg.Description)
	err = section.Validate()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the section
	k.SaveSection(ctx, section)

	// Update the section id for the next one
	k.SetNextSectionID(ctx, section.SubspaceID, section.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
		sdk.NewEvent(
			types.EventTypeCreateSection,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", section.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeySectionID, fmt.Sprintf("%d", section.ID)),
		),
	})

	return &types.MsgCreateSectionResponse{
		SectionID: section.ID,
	}, nil
}

// EditSection defines a rpc method for MsgEditSection
func (k msgServer) EditSection(goCtx context.Context, msg *types.MsgEditSection) (*types.MsgEditSectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	section, found := k.GetSection(ctx, msg.SubspaceID, msg.SectionID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.SectionID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Editor)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid editor address: %s", msg.Editor)
	}

	// Check the permission to manage sections
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionManageSections) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage sections within this subspace")
	}

	// Update the section and validate it
	update := types.NewSectionUpdate(msg.Name, msg.Description)
	updated := section.Update(update)
	err = updated.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the section
	k.SaveSection(ctx, updated)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Editor),
		),
		sdk.NewEvent(
			types.EventTypeEditSection,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", section.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeySectionID, fmt.Sprintf("%d", section.ID)),
		),
	})

	return &types.MsgEditSectionResponse{}, nil
}

// MoveSection defines a rpc method for MsgMoveSection
func (k msgServer) MoveSection(goCtx context.Context, msg *types.MsgMoveSection) (*types.MsgMoveSectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	section, found := k.GetSection(ctx, msg.SubspaceID, msg.SectionID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.SectionID, msg.SubspaceID)
	}

	// Check if the destination section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.NewParentID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d does not exist inside subspace %d", msg.NewParentID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to manage sections
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionManageSections) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage sections within this subspace")
	}

	// Update the section parent id and validate it
	section.ParentID = msg.NewParentID
	err = section.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the section
	k.SaveSection(ctx, section)

	// Make sure the section path is valid
	if !k.IsSectionPathValid(ctx, section.SubspaceID, section.ID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid section path")
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeMoveSection,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeySectionID, fmt.Sprintf("%d", msg.SectionID)),
		),
	})

	return &types.MsgMoveSectionResponse{}, nil
}

// DeleteSection defines a rpc method for MsgDeleteSection
func (k msgServer) DeleteSection(goCtx context.Context, msg *types.MsgDeleteSection) (*types.MsgDeleteSectionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.SectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.SectionID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to manage sections
	if !k.HasPermission(ctx, msg.SubspaceID, types.RootSectionID, signer.String(), types.PermissionManageSections) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage sections within this subspace")
	}

	// Delete the section
	k.Keeper.DeleteSection(ctx, msg.SubspaceID, msg.SectionID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeleteSection,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeySectionID, fmt.Sprintf("%d", msg.SectionID)),
		),
	})

	return &types.MsgDeleteSectionResponse{}, nil
}

// CreateUserGroup defines a rpc method for MsgCreateUserGroup
func (k msgServer) CreateUserGroup(goCtx context.Context, msg *types.MsgCreateUserGroup) (*types.MsgCreateUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.SectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.SectionID, msg.SubspaceID)
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address: %s", msg.Creator)
	}

	// Check the permissions to manage groups
	if !k.HasPermission(ctx, msg.SubspaceID, msg.SectionID, creator.String(), types.PermissionManageGroups) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in this subspace")
	}
	if !k.HasPermission(ctx, msg.SubspaceID, msg.SectionID, creator.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the default permissions are valid
	if !types.ArePermissionsValid(msg.DefaultPermissions) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Get the next group ID
	groupID, err := k.GetNextGroupID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create and validate the group
	group := types.NewUserGroup(msg.SubspaceID, msg.SectionID, groupID, msg.Name, msg.Description, msg.DefaultPermissions)
	if err := group.Validate(); err != nil {
		return nil, err
	}

	// Save the group
	k.SaveUserGroup(ctx, group)

	// Update the id for the next group
	k.SetNextGroupID(ctx, msg.SubspaceID, group.ID+1)

	// Add the initial members, if any
	var userEvents sdk.Events
	for _, member := range msg.InitialMembers {
		// Add the user to the group
		k.AddUserToGroup(ctx, group.SubspaceID, group.ID, member)

		// Add the events to the list of to emit
		userEvents = append(userEvents, sdk.NewEvent(
			types.EventTypeAddUserToGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", group.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", group.ID)),
			sdk.NewAttribute(types.AttributeKeyUser, member),
		))
	}

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
	}.AppendEvents(userEvents))

	return &types.MsgCreateUserGroupResponse{GroupID: groupID}, nil
}

// EditUserGroup defines a rpc method for MsgEditUserGroup
func (k msgServer) EditUserGroup(goCtx context.Context, msg *types.MsgEditUserGroup) (*types.MsgEditUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permission to create a group
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionManageGroups) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in this subspace")
	}

	// Update the group and validate it
	updated := group.Update(types.NewGroupUpdate(msg.Name, msg.Description))
	err = updated.Validate()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
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

// MoveUserGroup defines a rpc method for MsgMoveUserGroup
func (k msgServer) MoveUserGroup(goCtx context.Context, msg *types.MsgMoveUserGroup) (*types.MsgMoveUserGroupResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.SubspaceID)
	}

	// Check if the destination section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.NewSectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.NewSectionID, msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions to manage the current section groups
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionManageGroups) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in this section")
	}

	// Check the permissions to manage the destination section groups
	if !k.HasPermission(ctx, msg.SubspaceID, msg.NewSectionID, signer.String(), types.PermissionManageGroups) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user groups in the destination section")
	}
	if !k.HasPermission(ctx, msg.SubspaceID, msg.NewSectionID, signer.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage permissions in the destination section")
	}

	// Update the group section
	group.SectionID = msg.NewSectionID

	// Save the group
	k.SaveUserGroup(ctx, group)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EvenTypeMoveUserGroup,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyUserGroupID, fmt.Sprintf("%d", msg.GroupID)),
		),
	})

	return &types.MsgMoveUserGroupResponse{}, nil
}

// SetUserGroupPermissions defines a rpc method for MsgSetUserGroupPermissions
func (k msgServer) SetUserGroupPermissions(goCtx context.Context, msg *types.MsgSetUserGroupPermissions) (*types.MsgSetUserGroupPermissionsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, found := k.GetSubspace(ctx, msg.SubspaceID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group with id %d not found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrapf(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the permission is valid
	if !types.ArePermissionsValid(msg.Permissions) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Make sure that the user is not part of the group they want to change the permissions for, unless they are the owner
	if subspace.Owner != msg.Signer && k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, signer.String()) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "cannot set the permissions for a group you are part of")
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
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check for permissions
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionManageGroups) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot delete user groups in this subspace")
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
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	user, err := sdk.AccAddressFromBech32(msg.User)
	if err != nil {
		return nil, err
	}

	// Check if the user is already part of the group
	if k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, user.String()) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "user is already part of group %d", msg.GroupID)
	}

	// Set the user group
	k.AddUserToGroup(ctx, msg.SubspaceID, msg.GroupID, user.String())

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
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the group exists
	group, found := k.GetUserGroup(ctx, msg.SubspaceID, msg.GroupID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", msg.GroupID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, group.SubspaceID, group.SectionID, signer.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrap(types.ErrPermissionDenied, "you cannot manage user group members in this subspace")
	}

	// Check if the user is already part of the group
	if !k.IsMemberOfGroup(ctx, msg.SubspaceID, msg.GroupID, msg.User) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "user is not part of group %d", msg.GroupID)
	}

	// Remove the user group
	k.RemoveUserFromGroup(ctx, msg.SubspaceID, msg.GroupID, msg.User)

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
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.SectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "section with id %d not found inside subspace %d", msg.SectionID, msg.SubspaceID)
	}

	signer, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid signer address: %s", msg.Signer)
	}

	// Check the permissions
	if !k.HasPermission(ctx, msg.SubspaceID, msg.SectionID, signer.String(), types.PermissionSetPermissions) {
		return nil, errors.Wrapf(types.ErrPermissionDenied, "you cannot manage permissions in this subspace")
	}

	// Make sure the permission is valid
	if !types.ArePermissionsValid(msg.Permissions) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "invalid permission value")
	}

	// Set the permissions
	if msg.Permissions == nil {
		// Remove the permission to clear the store if empty permissions are provided
		k.RemoveUserPermissions(ctx, msg.SubspaceID, msg.SectionID, msg.User)
	} else {
		k.Keeper.SetUserPermissions(ctx, msg.SubspaceID, msg.SectionID, msg.User, msg.Permissions)
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
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgSetUserPermissionsResponse{}, nil
}

// UpdateSubspaceFeeTokens defines a rpc method for MsgUpdateSubspaceFeeTokens
func (k msgServer) UpdateSubspaceFeeTokens(goCtx context.Context, msg *types.MsgUpdateSubspaceFeeTokens) (*types.MsgUpdateSubspaceFeeTokensResponse, error) {
	if msg.Authority != k.authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	subspace, exists := k.GetSubspace(ctx, msg.SubspaceID)
	if !exists {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Update the subspace and validate it
	updated := types.NewAdditionalFeeTokensUpdate(msg.AllowedFeeTokens...).Update(subspace)
	err := updated.Validate()
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	k.SaveSubspace(ctx, updated)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
		),
		sdk.NewEvent(
			types.EventTypeUpdateSubspaceFeeToken,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", updated.ID)),
			sdk.NewAttribute(types.AttributeKeyUser, msg.Authority),
		),
	})

	return &types.MsgUpdateSubspaceFeeTokensResponse{}, nil
}
