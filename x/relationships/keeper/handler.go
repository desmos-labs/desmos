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
		case types.MsgCreateMonoDirectionalRelationship:
			return handleMsgCreateRelationship(ctx, keeper, msg)
		case types.MsgDeleteRelationship:
			return handleMsgDeleteRelationship(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Relationships message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateRelationship handles the creation of a mono directional relationship
func handleMsgCreateRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgCreateMonoDirectionalRelationship) (*sdk.Result, error) {
	// Save the relationship
	err := keeper.StoreRelationship(ctx, msg.Sender, msg.Receiver)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeMonoDirectionalRelationshipCreated,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Receiver.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Receiver),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgDeleteRelationship handles the relationship's deletion
func handleMsgDeleteRelationship(ctx sdk.Context, keeper Keeper, msg types.MsgDeleteRelationship) (*sdk.Result, error) {
	keeper.DeleteRelationship(ctx, msg.Sender, msg.Counterparty)

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRelationshipsDeleted,
		sdk.NewAttribute(types.AttributeRelationshipSender, msg.Sender.String()),
		sdk.NewAttribute(types.AttributeRelationshipReceiver, msg.Counterparty.String()),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Counterparty),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
