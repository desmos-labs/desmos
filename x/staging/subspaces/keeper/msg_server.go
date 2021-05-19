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

	// Check the if the subspace already exists
	if k.DoesSubspaceExists(ctx, msg.SubspaceID) {
		return nil,
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s already exists", msg.SubspaceID)
	}

	// Create and store the new subspaces
	subspace := types.NewSubspace(msg.SubspaceID, msg.Name, msg.Creator, msg.Creator, msg.Open, ctx.BlockTime())

	k.SaveSubspace(ctx, subspace)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceName, msg.Name),
		sdk.NewAttribute(types.AttributeKeySubspaceCreator, msg.Creator),
	))

	return &types.MsgCreateSubspaceResponse{}, nil
}

func (k msgServer) EditSubspace(goCtx context.Context, msg *types.MsgEditSubspace) (*types.MsgEditSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the if the subspace exists
	subspace, exist := k.GetSubspace(ctx, msg.ID)
	if !exist {
		return nil,
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s doesn't exists", msg.ID)
	}

	editedSubspace := subspace.
		WithName(msg.NewName).
		WithOwner(msg.NewOwner)

	k.SaveSubspace(ctx, editedSubspace)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEditSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.ID),
		sdk.NewAttribute(types.AttributeKeyNewOwner, subspace.Owner),
		sdk.NewAttribute(types.AttributeKeySubspaceName, subspace.Name),
	))

	return &types.MsgEditSubspaceResponse{}, nil
}

func (k msgServer) AddAdmin(goCtx context.Context, msg *types.MsgAddAdmin) (*types.MsgAddAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddAdminToSubspace(ctx, msg.SubspaceID, msg.Admin, msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAddAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.Admin),
	))

	return &types.MsgAddAdminResponse{}, nil
}

func (k msgServer) RemoveAdmin(goCtx context.Context, msg *types.MsgRemoveAdmin) (*types.MsgRemoveAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveAdminFromSubspace(ctx, msg.SubspaceID, msg.Admin, msg.Owner)
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

func (k msgServer) RegisterUser(goCtx context.Context, msg *types.MsgRegisterUser) (*types.MsgRegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.RegisterUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRegisterUser,
		sdk.NewAttribute(types.AttributeKeyRegisteredUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgRegisterUserResponse{}, nil
}

func (k msgServer) UnregisterUser(goCtx context.Context, msg *types.MsgUnregisterUser) (*types.MsgUnregisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.UnregisterUserFromSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnregisterUser,
		sdk.NewAttribute(types.AttributeKeyUnregisteredUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgUnregisterUserResponse{}, nil
}

func (k msgServer) BlockUser(goCtx context.Context, msg *types.MsgBlockUser) (*types.MsgBlockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.BlockUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUser,
		sdk.NewAttribute(types.AttributeKeyBlockedUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgBlockUserResponse{}, nil
}

func (k msgServer) UnblockUser(goCtx context.Context, msg *types.MsgUnblockUser) (*types.MsgUnblockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.BlockUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeKeyUnblockedUser, msg.User),
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
	))

	return &types.MsgUnblockUserResponse{}, nil
}
