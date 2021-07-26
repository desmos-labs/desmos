package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/subspaces/types"
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
	if k.DoesSubspaceExist(ctx, msg.SubspaceID) {
		return nil,
			sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "the subspaces with id %s already exists", msg.SubspaceID)
	}

	// Create and store the new subspaces
	creationTime := ctx.BlockTime()
	subspace := types.NewSubspace(
		msg.SubspaceID,
		msg.Name,
		msg.Description,
		msg.Logo,
		msg.Creator,
		msg.Creator,
		msg.SubspaceType,
		creationTime,
	)

	// Validate the subspace
	if err := subspace.Validate(); err != nil {
		return nil, err
	}

	err := k.SaveSubspace(ctx, subspace, subspace.Owner)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeCreateSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceName, msg.Name),
		sdk.NewAttribute(types.AttributeKeySubspaceCreator, msg.Creator),
		sdk.NewAttribute(types.AttributeKeyCreationTime, creationTime.Format(time.RFC3339)),
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
		WithName(msg.Name).
		WithDescription(msg.Description).
		WithLogo(msg.Logo).
		WithOwner(msg.Owner).
		WithSubspaceType(msg.SubspaceType)

	// Validate the subspace
	err := editedSubspace.Validate()
	if err != nil {
		return nil, err
	}

	err = k.SaveSubspace(ctx, editedSubspace, msg.Editor)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeEditSubspace,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.ID),
		sdk.NewAttribute(types.AttributeKeyNewOwner, editedSubspace.Owner),
		sdk.NewAttribute(types.AttributeKeySubspaceName, editedSubspace.Name),
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
		types.EventTypeRemoveAdmin,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeySubspaceRemovedAdmin, msg.Admin),
	))

	return &types.MsgRemoveAdminResponse{}, nil
}

func (k msgServer) RegisterUser(goCtx context.Context, msg *types.MsgRegisterUser) (*types.MsgRegisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RegisterUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRegisterUser,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyRegisteredUser, msg.User),
	))

	return &types.MsgRegisterUserResponse{}, nil
}

func (k msgServer) UnregisterUser(goCtx context.Context, msg *types.MsgUnregisterUser) (*types.MsgUnregisterUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.UnregisterUserFromSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnregisterUser,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyUnregisteredUser, msg.User),
	))

	return &types.MsgUnregisterUserResponse{}, nil
}

func (k msgServer) BanUser(goCtx context.Context, msg *types.MsgBanUser) (*types.MsgBanUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.BanUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBanUser,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyBanUser, msg.User),
	))

	return &types.MsgBanUserResponse{}, nil
}

func (k msgServer) UnbanUser(goCtx context.Context, msg *types.MsgUnbanUser) (*types.MsgUnbanUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.UnbanUserInSubspace(ctx, msg.SubspaceID, msg.User, msg.Admin)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnbanUser,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyUnbannedUser, msg.User),
	))

	return &types.MsgUnbanUserResponse{}, nil
}

func (k msgServer) SaveTokenomicsPair(goCtx context.Context, msg *types.MsgSaveTokenomicsPair) (*types.MsgSaveTokenomicsPairResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	tokenomicsPair := types.NewTokenomicsPair(msg.SubspaceID, msg.ContractAddress, msg.Admin)

	err := k.SaveSubspaceContractPair(ctx, tokenomicsPair)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeSaveTokenomicsPair,
		sdk.NewAttribute(types.AttributeKeySubspaceID, msg.SubspaceID),
		sdk.NewAttribute(types.AttributeKeyContractAddress, msg.ContractAddress),
		sdk.NewAttribute(types.AttributeKeySubspaceNewAdmin, msg.Admin),
	))

	return &types.MsgSaveTokenomicsPairResponse{}, nil
}
