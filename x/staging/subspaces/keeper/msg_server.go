package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
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

	// Create and store the new subspaces
	subspace := types.NewSubspace(ctx.BlockTime(), msg.SubspaceID, msg.Name, msg.Owner)

	// Return error if it has already been created
	err := k.SaveSubspace(ctx, subspace)
	if err != nil {
		return nil, err
	}

	// this error should never happen, adding the creator to the admins list to better handle admins checks
	if err = k.AddAdminToSubspace(ctx, subspace.ID, subspace.Owner); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceCreator, msg.Owner),
	))

	return &types.MsgCreateSubspaceResponse{}, nil
}

func (k msgServer) AddAdmin(goCtx context.Context, msg *types.MsgAddAdmin) (*types.MsgAddAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckSubspaceExistenceAndOwnerValidity(ctx, msg.SubspaceID, msg.Owner)
	if err != nil {
		return nil, err
	}

	err = k.AddAdminToSubspace(ctx, msg.SubspaceID, msg.NewAdmin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.NewAdmin),
	))

	return &types.MsgAddAdminResponse{}, nil
}

func (k msgServer) RemoveAdmin(goCtx context.Context, msg *types.MsgRemoveAdmin) (*types.MsgRemoveAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckSubspaceExistenceAndOwnerValidity(ctx, msg.SubspaceID, msg.Owner)
	if err != nil {
		return nil, err
	}

	err = k.RemoveAdminFromSubspace(ctx, msg.SubspaceID, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceRemovedAdmin, msg.Admin),
	))

	return &types.MsgRemoveAdminResponse{}, nil
}

func (k msgServer) EnableUserPosts(goCtx context.Context, msg *types.MsgEnableUserPosts) (*types.MsgEnableUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckSubspaceExistenceAndAdminValidity(ctx, msg.Admin, msg.SubspaceID); err != nil {
		return nil, err
	}

	if err := k.UnblockPostsForUser(ctx, msg.User, msg.SubspaceID); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEnableUserPosts,
		sdk.NewAttribute(types.AttributeKeyEnabledToPostUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgEnableUserPostsResponse{}, nil
}

func (k msgServer) DisableUserPosts(goCtx context.Context, msg *types.MsgDisableUserPosts) (*types.MsgDisableUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckSubspaceExistenceAndAdminValidity(ctx, msg.Admin, msg.SubspaceID); err != nil {
		return nil, err
	}

	if err := k.BlockPostsForUser(ctx, msg.User, msg.SubspaceID); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUserPosts,
		sdk.NewAttribute(types.AttributeKeyDisabledToPostUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgDisableUserPostsResponse{}, nil
}

func (k msgServer) TransferOwnership(goCtx context.Context, msg *types.MsgTransferOwnership) (*types.MsgTransferOwnershipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckSubspaceExistenceAndOwnerValidity(ctx, msg.Owner, msg.SubspaceID); err != nil {
		return nil, err
	}

	k.TransferSubspaceOwnership(ctx, msg.SubspaceID, msg.NewOwner)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeTransferOwnership,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyNewOwner, msg.NewOwner),
	))

	return &types.MsgTransferOwnershipResponse{}, nil
}
