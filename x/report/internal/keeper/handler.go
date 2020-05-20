package keeper

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/report/internal/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgReportPost:
			return handleMsgReportPost(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgReportPost handles the report of a post
// TODO check if the posts exists before saving the reports
func handleMsgReportPost(ctx sdk.Context, keeper Keeper, msg types.MsgReportPost) (*sdk.Result, error) {
	// Save the report if it's not been made before
	if saved := keeper.SaveReport(ctx, msg.PostID, msg.Report); !saved {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("report to the post with id %s has already been made by user: %s",
			msg.PostID, msg.Report.User))
	}

	createEvent := sdk.NewEvent(
		types.EventTypePostReported,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReportOwner, msg.Report.User.String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   []byte(fmt.Sprintf("post with ID: %s reported correctly", msg.PostID)),
		Events: sdk.Events{createEvent},
	}
	return &result, nil
}
