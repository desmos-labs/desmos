package magpie

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/types"
	"github.com/rs/xid"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgCreatePost:
			return handleMsgCreatePost(ctx, keeper, msg)
		case MsgEditPost:
			return handleMsgEditPost(ctx, keeper, msg)
		case MsgLike:
			return handleMsgLike(ctx, keeper, msg)
		// case MsgUnlike:
		default:
			errMsg := fmt.Sprintf("Unrecognized Magpie Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// Handle creating a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg MsgCreatePost) sdk.Result {
	// if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
	// 	return sdk.ErrUnauthorized("Incorrect Owner").Result()
	// }
	post := Post{
		ID:      xid.New().String(),
		Message: msg.Message,
		Time:    msg.Time,
		Likes:   0,
		Owner:   msg.Owner,
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
			sdk.NewAttribute(types.AttributeKeyPostID, post.ID),
		),
	)

	keeper.SetPost(ctx, post)
	return sdk.Result{
		Data:   keeper.cdc.MustMarshalBinaryLengthPrefixed(post.ID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg MsgEditPost) sdk.Result {
	if !msg.Owner.Equals(keeper.GetPostOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}

	keeper.EditPost(ctx, msg.ID, msg.Message)
	return sdk.Result{}
}

func handleMsgLike(ctx sdk.Context, keeper Keeper, msg MsgLike) sdk.Result {

	post := keeper.GetPost(ctx, msg.PostID)

	if msg.PostID != post.ID {
		return sdk.ErrUnknownRequest("Post doesn't exist").Result()
	}

	like := Like{
		ID:     xid.New().String(),
		Time:   msg.Time,
		PostID: msg.PostID,
		Owner:  msg.Liker,
	}

	keeper.SetLike(ctx, like.ID, like)
	return sdk.Result{}
}

//
