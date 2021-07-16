package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

func (k msgServer) CreateRelationship(goCtx context.Context, msg *types.MsgCreateRelationship) (*types.MsgCreateRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.CheckUserPermissionsInSubspace(ctx, msg.Subspace, msg.Sender); err != nil {
		return nil, err
	}

	if err := k.CheckUserPermissionsInSubspace(ctx, msg.Subspace, msg.Receiver); err != nil {
		return nil, err
	}

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

	return &types.MsgCreateRelationshipResponse{}, nil
}

func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.MsgDeleteRelationshipResponse, error) {
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

	return &types.MsgDeleteRelationshipResponse{}, nil
}
