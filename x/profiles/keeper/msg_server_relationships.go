package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
)

func (k msgServer) CreateRelationship(goCtx context.Context, msg *types.MsgCreateRelationship) (*types.MsgCreateRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the receiver has blocked the sender before
	if k.HasUserBlocked(ctx, msg.Receiver, msg.Sender, msg.SubspaceID) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "The user with address %s has blocked you", msg.Receiver)
	}

	// Save the relationship
	err := k.SaveRelationship(ctx, types.NewRelationship(msg.Sender, msg.Receiver, msg.SubspaceID))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, fmt.Sprintf("%d", msg.SubspaceID)),
	))

	return &types.MsgCreateRelationshipResponse{}, nil
}

func (k msgServer) DeleteRelationship(goCtx context.Context, msg *types.MsgDeleteRelationship) (*types.MsgDeleteRelationshipResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.RemoveRelationship(ctx, types.NewRelationship(msg.User, msg.Counterparty, msg.SubspaceID))
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.User),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, fmt.Sprintf("%d", msg.SubspaceID)),
	))

	return &types.MsgDeleteRelationshipResponse{}, nil
}
