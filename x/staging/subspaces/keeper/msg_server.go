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

	// Create and store the new subspace
	subspace := types.NewSubspace(ctx.BlockTime(), msg.SubspaceID, msg.Name, msg.Creator)

	// Return error if it has already been created
	err := k.SaveSubspace(ctx, subspace)
	if err != nil {
		return nil, err
	}

	// this error should never happen, adding the creator to the admins list to better handle admins checks
	if err = k.AddAdminToSubspace(ctx, subspace.ID, subspace.Creator); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceCreator, msg.Creator),
	))

	return &types.MsgCreateSubspaceResponse{}, nil
}

func (k msgServer) AddSubspaceAdmin(goCtx context.Context, msg *types.MsgAddAdmin) (*types.MsgAddAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckSubspaceExistenceAndCreatorValidity(ctx, msg.SubspaceID, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.AddAdminToSubspace(ctx, msg.SubspaceID, msg.NewAdmin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.NewAdmin),
	))

	return &types.MsgAddAdminResponse{}, nil
}

func (k msgServer) RemoveSubspaceAdmin(goCtx context.Context, msg *types.MsgRemoveAdmin) (*types.MsgRemoveAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.CheckSubspaceExistenceAndCreatorValidity(ctx, msg.SubspaceID, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.RemoveAdminFromSubspace(ctx, msg.SubspaceId, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceId),
		sdk.NewAttribute(types.AttributeKeySubspaceRemovedAdmin, msg.Admin),
	))

	return &types.MsgRemoveAdminResponse{}, nil
}

func (k msgServer) EnableUserPosts(goCtx context.Context, msg *types.MsgEnableUserPosts) (*types.MsgEnableUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckSubspaceExistenceAndAdminValidity(ctx, msg.Admin, msg.SubspaceId); err != nil {
		return nil, err
	}

	if err := k.UnblockPostsForUser(ctx, msg.User, msg.SubspaceId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEnableUserPosts,
		sdk.NewAttribute(types.AttributeKeyEnabledToPostUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceId),
	))

	return &types.MsgEnableUserPostsResponse{}, nil
}

func (k msgServer) DisableUserPosts(goCtx context.Context, msg *types.MsgDisableUserPosts) (*types.MsgDisableUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckSubspaceExistenceAndAdminValidity(ctx, msg.Admin, msg.SubspaceId); err != nil {
		return nil, err
	}

	if err := k.BlockPostsForUser(ctx, msg.User, msg.SubspaceId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUserPosts,
		sdk.NewAttribute(types.AttributeKeyDisabledToPostUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceId),
	))

	return &types.MsgDisableUserPostsResponse{}, nil
}
