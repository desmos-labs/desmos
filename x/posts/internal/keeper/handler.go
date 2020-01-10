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
		case types.MsgAddPostReaction:
			return handleMsgAddPostReaction(ctx, keeper, msg)
		case types.MsgRemovePostReaction:
			return handleMsgRemovePostReaction(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// handleMsgCreatePost handles the creation of a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) sdk.Result {

	var post types.Post

	if textMsg, ok := msg.(types.MsgCreateTextPost); ok {
		post = types.NewTextPost(
			keeper.GetLastPostID(ctx).Next(),
			textMsg.ParentID,
			textMsg.Message,
			textMsg.AllowsComments,
			textMsg.Subspace,
			textMsg.OptionalData,
			ctx.BlockHeight(),
			textMsg.Creator,
		)
	}

	if mediaMsg, ok := msg.(types.MsgCreateMediaPost); ok {
		textPost := types.NewTextPost(
			keeper.GetLastPostID(ctx).Next(),
			mediaMsg.MsgCreatePost.ParentID,
			mediaMsg.MsgCreatePost.Message,
			mediaMsg.MsgCreatePost.AllowsComments,
			mediaMsg.MsgCreatePost.Subspace,
			mediaMsg.MsgCreatePost.OptionalData,
			ctx.BlockHeight(),
			mediaMsg.MsgCreatePost.Creator,
		)

		post = types.NewMediaPost(textPost, mediaMsg.Medias)
	}

	// Check for double posting
	if _, found := keeper.GetPost(ctx, post.GetID()); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("TextPost with id %s already exists", post.GetID())).Result()
	}

	// If valid, check the parent post
	if post.GetParentID().Valid() {
		parentPost, found := keeper.GetPost(ctx, post.GetParentID())
		if !found {
			return sdk.ErrUnknownRequest(fmt.Sprintf("Parent post with id %s not found",
				post.GetParentID())).Result()
		}

		if !parentPost.CanComment() {
			return sdk.ErrUnknownRequest(fmt.Sprintf("TextPost with id %s does not allow comments",
				parentPost.GetID())).Result()
		}
	}

	keeper.SavePost(ctx, post)

	createEvent := sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.GetID().String()),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.GetParentID().String()),
		sdk.NewAttribute(types.AttributeKeyCreationTime, post.CreationTime().String()),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Owner().String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(post.GetID()),
		Events: sdk.Events{createEvent},
	}
}

// handleMsgEditPost handles MsgEditsPost messages
func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg types.MsgEditPost) sdk.Result {

	// Get the existing post
	existing, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("TextPost with id %s not found", msg.PostID)).Result()
	}

	// Checks if the the msg sender is the same as the current owner
	if !msg.Editor.Equals(existing.Owner()) {
		return sdk.ErrUnauthorized("Incorrect owner").Result()
	}

	// Check the validity of the current block height respect to the creation date of the post
	if existing.CreationTime().GT(sdk.NewInt(ctx.BlockHeight())) {
		return sdk.ErrUnknownRequest("Edit date cannot be before creation date").Result()
	}

	// Edit the post
	existing = existing.SetMessage(msg.Message)
	existing = existing.SetEditTime(sdk.NewInt(ctx.BlockHeight()))
	keeper.SavePost(ctx, existing)

	editEvent := sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.GetID().String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.GetEditTime().String()),
	)
	ctx.EventManager().EmitEvent(editEvent)

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.GetID()),
		Events: sdk.Events{editEvent},
	}
}

func handleMsgAddPostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgAddPostReaction) sdk.Result {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("TextPost with id %s not found", msg.PostID)).Result()
	}

	// Create and store the reaction
	reaction := types.NewReaction(msg.Value, ctx.BlockHeight(), msg.User)
	if err := keeper.SaveReaction(ctx, post.GetID(), reaction); err != nil {
		return err.Result()
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypeReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Value),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Data:   []byte("Reaction added properly"),
		Events: sdk.Events{event},
	}
}

func handleMsgRemovePostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgRemovePostReaction) sdk.Result {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("TextPost with id %s not found", msg.PostID)).Result()
	}

	// Remove the reaction
	if err := keeper.RemoveReaction(ctx, post.GetID(), msg.User, msg.Reaction); err != nil {
		return err.Result()
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Reaction),
	)
	ctx.EventManager().EmitEvent(event)

	return sdk.Result{
		Data:   []byte("Reaction removed properly"),
		Events: sdk.Events{event},
	}
}
