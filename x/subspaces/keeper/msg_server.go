package keeper

import (
	"context"
	"time"

	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface
// for the provided keeper
func NewMsgServerImpl(keeper Keeper) types2.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types2.MsgServer = msgServer{}

func (k msgServer) CreateSubspace(goCtx context.Context, msg *types2.MsgCreateSubspace) (*types2.MsgCreateSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the if the subspace already exists
	if k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil,
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s already exists", msg.SubspaceID)
	}

	// Create and store the new subspaces
	creationTime := ctx.BlockTime()
	subspace := types2.NewSubspace(msg.SubspaceID, msg.Name, msg.Creator, msg.Creator, msg.SubspaceType, creationTime)

	// Validate the subspace
	if err := subspace.Validate(); err != nil {
		return nil, err
	}

	_ = k.SaveSubspace(ctx, subspace, subspace.Owner)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeCreateSubspace,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeySubspaceName, msg.Name),
		sdk.NewAttribute(types2.AttributeKeySubspaceCreator, msg.Creator),
		sdk.NewAttribute(types2.AttributeKeyCreationTime, creationTime.Format(time.RFC3339)),
	))

	return &types2.MsgCreateSubspaceResponse{}, nil
}

func (k msgServer) EditSubspace(goCtx context.Context, msg *types2.MsgEditSubspace) (*types2.MsgEditSubspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check the if the subspace exists
	subspace, exist := k.GetSubspace(ctx, msg.ID)
	if !exist {
		return nil,
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s doesn't exists", msg.ID)
	}

	editedSubspace := subspace.
		WithName(msg.Name).
		WithOwner(msg.Owner).
		WithSubspaceType(msg.SubspaceType)

	// Validate the subspace
	if err := editedSubspace.Validate(); err != nil {
		return nil, err
	}

	if err := k.SaveSubspace(ctx, editedSubspace, msg.Editor); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeEditSubspace,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.ID),
		sdk.NewAttribute(types2.AttributeKeyNewOwner, editedSubspace.Owner),
		sdk.NewAttribute(types2.AttributeKeySubspaceName, editedSubspace.Name),
	))

	return &types2.MsgEditSubspaceResponse{}, nil
}

func (k msgServer) AddAdmin(goCtx context.Context, msg *types2.MsgAddAdmin) (*types2.MsgAddAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddAdminToSubspace(ctx, msg.SubspaceID, msg.Admin, msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeAddAdmin,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeySubspaceNewAdmin, msg.Admin),
	))

	return &types2.MsgAddAdminResponse{}, nil
}

func (k msgServer) RemoveAdmin(goCtx context.Context, msg *types2.MsgRemoveAdmin) (*types2.MsgRemoveAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveAdminFromSubspace(ctx, msg.SubspaceID, msg.Admin, msg.Owner)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeRemoveAdmin,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeySubspaceRemovedAdmin, msg.Admin),
	))

	return &types2.MsgRemoveAdminResponse{}, nil
}

func (k msgServer) RegisterUser(goCtx context.Context, msg *types2.MsgRegisterUser) (*types2.MsgRegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.RegisterUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeRegisterUser,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeyRegisteredUser, msg.User),
	))

	return &types2.MsgRegisterUserResponse{}, nil
}

func (k msgServer) UnregisterUser(goCtx context.Context, msg *types2.MsgUnregisterUser) (*types2.MsgUnregisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.UnregisterUserFromSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeUnregisterUser,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeyUnregisteredUser, msg.User),
	))

	return &types2.MsgUnregisterUserResponse{}, nil
}

func (k msgServer) BanUser(goCtx context.Context, msg *types2.MsgBanUser) (*types2.MsgBanUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.BanUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeBanUser,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeyBanUser, msg.User),
	))

	return &types2.MsgBanUserResponse{}, nil
}

func (k msgServer) UnbanUser(goCtx context.Context, msg *types2.MsgUnbanUser) (*types2.MsgUnbanUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.UnbanUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin); err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types2.EventTypeUnbanUser,
		sdk.NewAttribute(types2.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types2.AttributeKeyUnbannedUser, msg.User),
	))

	return &types2.MsgUnbanUserResponse{}, nil
}
