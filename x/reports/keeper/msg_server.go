package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	posttypes "github.com/desmos-labs/desmos/x/posts/types"
	"github.com/desmos-labs/desmos/x/reports/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) ReportPost(goCtx context.Context, msg *types.MsgReportPost) (*types.MsgReportPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the post to stored exists
	postID, err := posttypes.ParsePostID(msg.PostId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	if exist := k.CheckPostExistence(ctx, postID); !exist {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with ID: %s doesn't exist", msg.PostId),
		)
	}

	report := types.NewReport(msg.PostId, msg.ReportType, msg.Message, msg.User)
	if err := k.SaveReport(ctx, report); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, err.Error())
	}

	createEvent := sdk.NewEvent(
		types.EventTypePostReported,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostId),
		sdk.NewAttribute(types.AttributeKeyReportOwner, msg.User),
	)
	ctx.EventManager().EmitEvent(createEvent)

	return &types.MsgReportPostResponse{}, nil
}
