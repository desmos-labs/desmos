package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case types.MsgCreatePost:
			return handleMsgCreatePost(ctx, keeper, msg)
		case types.MsgEditPost:
			return handleMsgEditPost(ctx, keeper, msg)
		case types.MsgLikePost:
			return handleMsgLike(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgCreatePost handles the creation of a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) sdk.Result {

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionCreatePost),
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
		),
	)

	post := types.Post{
		PostID:        keeper.GetLastPostID(ctx).Next(),
		ParentID:      msg.ParentID,
		Message:       msg.Message,
		Created:       ctx.BlockHeight(),
		Owner:         msg.Owner,
		Namespace:     msg.Namespace,
		ExternalOwner: msg.ExternalOwner,
	}

	if err := keeper.SavePost(ctx, post); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCreatePost,
			sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
			sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentID.String()),
			sdk.NewAttribute(types.AttributeKeyCreationTime, strconv.FormatInt(post.Created, 10)),
			sdk.NewAttribute(types.AttributeKeyPostOwner, post.Owner.String()),
			sdk.NewAttribute(types.AttributeKeyNamespace, post.Namespace),
			sdk.NewAttribute(types.AttributeKeyExternalOwner, post.ExternalOwner),
		),
	)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(post.PostID),
		Events: ctx.EventManager().Events(),
	}
}

// handleMsgEditPost handles MsgEditsPost messages
func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg types.MsgEditPost) sdk.Result {

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionEditPost),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Editor.String()),
		),
	)

	// Get the existing post
	existing, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Checks if the the msg sender is the same as the current owner
	if !msg.Editor.Equals(existing.Owner) {
		return sdk.ErrUnauthorized("Incorrect owner").Result()
	}

	// Check the validity of the current block height respect to the creation date of the post
	if ctx.BlockHeight() < existing.Created {
		return sdk.ErrUnknownRequest("Edit date cannot be before creation date").Result()
	}

	// Edit the post
	existing.Message = msg.Message
	existing.LastEdited = ctx.BlockHeight()
	if err := keeper.SavePost(ctx, existing); err != nil {
		return err.Result()
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeEditPost,
			sdk.NewAttribute(types.AttributeKeyPostID, existing.PostID.String()),
			sdk.NewAttribute(types.AttributeKeyPostEditTime, strconv.FormatInt(existing.LastEdited, 10)),
		),
	)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.PostID),
		Events: ctx.EventManager().Events(),
	}
}

func handleMsgLike(ctx sdk.Context, keeper Keeper, msg types.MsgLikePost) sdk.Result {

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.ActionLikePost),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Liker.String()),
		),
	)

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Check the like date to make sure it's before the post creation date.
	if ctx.BlockHeight() < post.Created {
		return sdk.ErrUnknownRequest("Like cannot have a creation time before the post itself").Result()
	}

	// Create and store the like
	like := types.Like{
		Created: ctx.BlockHeight(),
		Owner:   msg.Liker,
	}

	if err := keeper.SaveLike(ctx, post.PostID, like); err != nil {
		return err.Result()
	}

	// Emit the event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeLikePost,
			sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
			sdk.NewAttribute(types.AttributeKeyLikeOwner, msg.Liker.String()),
		),
	)

	return sdk.Result{
		Data:   []byte("Like added properly"),
		Events: ctx.EventManager().Events(),
	}
}
