package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the relationships MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (k msgServer) CreateRelationship(goCtx context.Context, msg *types.MsgCreateRelationship) (*types.CreateRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the receiver has blocked the sender before
	if k.HasUserBlocked(ctx, msg.Receiver, msg.Sender, msg.Subspace) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "The user with address %s has blocked you", msg.Receiver)
	}

	// Save the relationship
	err := k.SaveRelationship(ctx, types.NewRelationship(msg.Sender, msg.Receiver, msg.Subspace))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	return &types.CreateRelationshipResponse{}, nil
}

func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.DeleteRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveRelationship(ctx, types.NewRelationship(msg.User, msg.Counterparty, msg.Subspace))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.User),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	return &types.DeleteRelationshipResponse{}, nil
}

func (k msgServer) BlockUser(goCtx context.Context, msg *types.MsgBlockUser) (*types.BlockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	userBlock := types.NewUserBlock(msg.Blocker, msg.Blocked, msg.Reason, msg.Subspace)
	err := k.SaveUserBlock(ctx, userBlock)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
		sdk.NewAttribute(types.AttributeUserBlockReason, msg.Reason),
	))

	return &types.BlockUserResponse{}, nil
}

func (k msgServer) UnblockUser(goCtx context.Context, msg *types.MsgUnblockUser) (*types.UnblockUserResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.DeleteUserBlock(ctx, msg.Blocker, msg.Blocked, msg.Subspace)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
	))

	return &types.UnblockUserResponse{}, nil
}
