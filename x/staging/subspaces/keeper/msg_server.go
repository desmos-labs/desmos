package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
	subspace := types.NewSubspace(ctx.BlockTime(), msg.Id, msg.Creator)
	err := k.SaveSubspace(ctx, subspace)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.Id),
		sdk.NewAttribute(types.AttributeKeySubspaceCreator, msg.Creator),
	))

	return &types.MsgCreateSubspaceResponse{}, nil
}

// checkSubspaceExistenceAndCreatorValidity check if the subspace with the given id exists and if the creator is valid.
func (k msgServer) checkSubspaceExistenceAndCreatorValidity(ctx sdk.Context, subspaceId, creator string) error {
	// Check if the subspace exists
	subspace, exist := k.GetSubspace(ctx, subspaceId)

	// If it doesn't, return error
	if !exist {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceId,
		)
	}

	// Check if the creator is the actual one
	if subspace.Creator != creator {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"you are not the subspace creator and you can't add any admin to it",
		)
	}

	return nil
}

func (k msgServer) AddSubspaceAdmin(goCtx context.Context, msg *types.MsgAddAdmin) (*types.MsgAddAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.checkSubspaceExistenceAndCreatorValidity(ctx, msg.SubspaceId, msg.Creator)
	if err != nil {
		return nil, err
	}

	err = k.AddAdminToSubspace(ctx, msg.SubspaceId, msg.NewAdmin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.NewAdmin),
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.NewAdmin),
	))

	return &types.MsgAddAdminResponse{}, nil
}

func (k msgServer) RemoveSubspaceAdmin(goCtx context.Context, msg *types.MsgRemoveAdmin) (*types.MsgRemoveAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.checkSubspaceExistenceAndCreatorValidity(ctx, msg.SubspaceId, msg.Creator)
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
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.Admin),
	))

	return &types.MsgRemoveAdminResponse{}, nil
}

func (k msgServer) checkSubspaceAndAdmin(ctx sdk.Context, admin, subspaceId string) error {
	if !k.DoesSubspaceExists(ctx, subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the subspace with id %s doesn't exist", subspaceId,
		)
	}

	if !k.IsAdmin(ctx, admin, subspaceId) {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the user: %s is not an admin and can't perform this operation on the subspace: %s",
			admin, subspaceId)
	}

	return nil
}

func (k msgServer) AllowUserPosts(goCtx context.Context, msg *types.MsgAllowUserPosts) (*types.MsgAllowUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkSubspaceAndAdmin(ctx, msg.Admin, msg.SubspaceId); err != nil {
		return nil, err
	}

	if err := k.EnableUserPosts(ctx, msg.User, msg.SubspaceId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAllowUserPosts,
		sdk.NewAttribute(types.AttributeKeyAllowedUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceId),
	))

	return &types.MsgAllowUserPostsResponse{}, nil
}

func (k msgServer) BlockUserPosts(goCtx context.Context, msg *types.MsgBlockUserPosts) (*types.MsgBlockUserPostsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.checkSubspaceAndAdmin(ctx, msg.Admin, msg.SubspaceId); err != nil {
		return nil, err
	}

	if err := k.DisableUserPosts(ctx, msg.User, msg.SubspaceId); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUserPosts,
		sdk.NewAttribute(types.AttributeKeyBlockedUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceId, msg.SubspaceId),
	))

	return &types.MsgBlockUserPostsResponse{}, nil
}
