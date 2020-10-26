package relationships

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/relationships/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

// NewHandler returns a handler for "profile" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgCreateRelationship:
			return handleMsgCreateRelationship(ctx, keeper, msg)
		case types.MsgDeleteRelationship:
			return handleMsgDeleteRelationship(ctx, keeper, msg)
		case types.MsgBlockUser:
			return handleMsgBlockUser(ctx, keeper, msg)
		case types.MsgUnblockUser:
			return handleMsgUnblockUser(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Relationships message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreateRelationship handles the creation of a relationship
func handleMsgCreateRelationship(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgCreateRelationship) (*sdk.Result, error) {
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
func handleMsgDeleteRelationship(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgDeleteRelationship) (*sdk.Result, error) {
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

// handleMsgBlockUser handles the process to block a user
func handleMsgBlockUser(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgBlockUser) (*sdk.Result, error) {
	userBlock := types.NewUserBlock(msg.Blocker, msg.Blocked, msg.Reason, msg.Subspace)

	if err := keeper.SaveUserBlock(ctx, userBlock); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeBlockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker.String()),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked.String()),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
		sdk.NewAttribute(types.AttributeUserBlockReason, msg.Reason),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Blocker),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}

// handleMsgUnblockUser handles the process to unblock a user
func handleMsgUnblockUser(ctx sdk.Context, keeper keeper.Keeper, msg types.MsgUnblockUser) (*sdk.Result, error) {
	if err := keeper.UnblockUser(ctx, msg.Blocker, msg.Blocked, msg.Subspace); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeUnblockUser,
		sdk.NewAttribute(types.AttributeUserBlockBlocker, msg.Blocker.String()),
		sdk.NewAttribute(types.AttributeUserBlockBlocked, msg.Blocked.String()),
		sdk.NewAttribute(types.AttributeSubspace, msg.Subspace),
	))

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(msg.Blocker),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
