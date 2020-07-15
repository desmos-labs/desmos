package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/reports/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case types.MsgReportPost:
			return handleMsgReportPost(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgReportPost handles the reports of a post
func handleMsgReportPost(ctx sdk.Context, keeper Keeper, msg types.MsgReportPost) (*sdk.Result, error) {
	// check if the post to reports exists
	if exist := keeper.CheckPostExistence(ctx, msg.PostID); !exist {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with ID: %s doesn't exist", msg.PostID))
	}

	keeper.SaveReport(ctx, msg.PostID, msg.Report)

	createEvent := sdk.NewEvent(
		types.EventTypePostReported,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReportOwner, msg.Report.User.String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   []byte(fmt.Sprintf("post with ID: %s reported correctly", msg.PostID)),
		Events: ctx.EventManager().Events(),
	}
	return &result, nil
}
