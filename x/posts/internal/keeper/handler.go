package keeper

import (
	"fmt"

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
		case types.MsgUnlikePost:
			return handleMsgUnLike(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgCreatePost handles the creation of a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) sdk.Result {
	post := types.NewPost(
		keeper.GetLastPostID(ctx).Next(),
		msg.ParentID,
		msg.Message,
		msg.AllowsComments,
		msg.Subspace,
		msg.OptionalData,
		ctx.BlockHeight(),
		msg.Creator,
	)

	// Check for double posting
	if _, found := keeper.GetPost(ctx, post.PostID); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s already exists", post.PostID)).Result()
	}

	// If valid, check the parent post
	if post.ParentID.Valid() {
		parentPost, found := keeper.GetPost(ctx, post.ParentID)
		if !found {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Parent post with id %s not found", post.ParentID)).Result()
		}

		if !parentPost.AllowsComments {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s does not allow comments", parentPost.PostID)).Result()
		}
	}

	keeper.SavePost(ctx, post)

	createEvent := sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentID.String()),
		sdk.NewAttribute(types.AttributeKeyCreationTime, post.Created.String()),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Owner.String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(post.PostID),
		Events: sdk.Events{createEvent},
	}
}

// handleMsgEditPost handles MsgEditsPost messages
func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg types.MsgEditPost) sdk.Result {

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
	if existing.Created.GT(sdk.NewInt(ctx.BlockHeight())) {
		return sdk.ErrUnknownRequest("Edit date cannot be before creation date").Result()
	}

	// Edit the post
	existing.Message = msg.Message
	existing.LastEdited = sdk.NewInt(ctx.BlockHeight())
	keeper.SavePost(ctx, existing)

	editEvent := sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.LastEdited.String()),
	)
	ctx.EventManager().EmitEvent(editEvent)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.PostID),
		Events: sdk.Events{editEvent},
	}
}

func handleMsgLike(ctx sdk.Context, keeper Keeper, msg types.MsgLikePost) sdk.Result {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Check the liker date to make sure it's before the post creation date.
	if post.Created.GT(sdk.NewInt(ctx.BlockHeight())) {
		return sdk.ErrUnknownRequest("Like cannot have a creation time before the post itself").Result()
	}

	// Create and store the liker
	like := types.NewLike(ctx.BlockHeight(), msg.Liker)
	if err := keeper.SaveLike(ctx, post.PostID, like); err != nil {
		return err.Result()
	}

	// Emit the event
	likeEvent := sdk.NewEvent(
		types.EventTypePostLiked,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyLikeOwner, msg.Liker.String()),
	)
	ctx.EventManager().EmitEvent(likeEvent)

	return sdk.Result{
		Data:   []byte("Like added properly"),
		Events: sdk.Events{likeEvent},
	}
}

func handleMsgUnLike(ctx sdk.Context, keeper Keeper, msg types.MsgUnlikePost) sdk.Result {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Check the liker date to make sure it's before the post creation date.
	if post.Created.GT(sdk.NewInt(ctx.BlockHeight())) {
		return sdk.ErrUnknownRequest("Cannot unlike a post before it's created").Result()
	}

	if err := keeper.RemoveLike(ctx, post.PostID, msg.Liker); err != nil {
		return err.Result()
	}

	// Emit the event
	unlikeEvent := sdk.NewEvent(
		types.EventTypePostUnliked,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyLikeOwner, msg.Liker.String()),
	)
	ctx.EventManager().EmitEvent(unlikeEvent)

	return sdk.Result{
		Data:   []byte("Like removed properly"),
		Events: sdk.Events{unlikeEvent},
	}
}
