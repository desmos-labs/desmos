package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	keeper2 "github.com/desmos-labs/desmos/x/posts/keeper"
	types2 "github.com/desmos-labs/desmos/x/posts/types"
)

// NewHandler returns a handler for "posts" type messages.
func NewHandler(k keeper2.Keeper) sdk.Handler {
	msgServer := keeper2.NewMsgServerImpl(k)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		case *types2.MsgCreatePost:
			res, err := msgServer.CreatePost(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types2.MsgEditPost:
			res, err := msgServer.EditPost(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types2.MsgAddPostReaction:
			res, err := msgServer.AddPostReaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types2.MsgRemovePostReaction:
			res, err := msgServer.RemovePostReaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types2.MsgAnswerPoll:
			res, err := msgServer.AnswerPoll(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *types2.MsgRegisterReaction:
			res, err := msgServer.RegisterReaction(sdk.WrapSDKContext(ctx), msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
				"unrecognized %s message type: %v", types2.ModuleName, msg.Type())
		}
	}
}
