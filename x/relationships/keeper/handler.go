package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// NewHandler returns a handler for "profile" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreateRelationship:
			return handleMsgCreateRelationship(ctx, keeper, msg)
		case types.MsgDeleteRelationship:
			return handleMsgDeleteRelationship(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Relationships message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateRelationship handles the creation of a relationship
func handleMsgCreateRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgCreateRelationship) (*sdk.Result, error) {
	// Save the relationship
	err := keeper.StoreRelationship(ctx, msg.Sender, types.NewRelationship(msg.Receiver, msg.Subspace))
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver.String()),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Receiver),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgDeleteRelationship handles the relationship's deletion
func handleMsgDeleteRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteRelationship) (*sdk.Result, error) {
	keeper.DeleteRelationship(ctx, msg.Sender, types.NewRelationship(msg.Counterparty, msg.Subspace))

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty.String()),
		sdk.NewAttribute(types.AttributeRelationshipSubspace, msg.Subspace),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Counterparty),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
