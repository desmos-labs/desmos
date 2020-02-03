package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
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
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// handleMsgCreatePost handles the creation of a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) (*sdk.Result, error) {
	post := types.NewPost(
		keeper.GetLastPostID(ctx).Next(),
		msg.ParentID,
		msg.Message,
		msg.AllowsComments,
		msg.Subspace,
		msg.OptionalData,
		msg.CreationDate,
		msg.Creator,
		msg.Medias,
	)

	// Check for double posting
	if _, found := keeper.GetPost(ctx, post.PostID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s already exists", post.PostID))
	}

	// If valid, check the parent post
	if post.ParentID.Valid() {
		parentPost, found := keeper.GetPost(ctx, post.ParentID)
		if !found {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("parent post with id %s not found", post.ParentID))
		}

		if !parentPost.AllowsComments {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s does not allow comments", parentPost.PostID))
		}
	}

	// Check the post does not exist yet
	if existing, exists := keeper.DoesPostExit(ctx, post); exists {
		errMsg := fmt.Sprintf("post with same data already exists and has ID %s", existing.PostID)
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, errMsg)
	}

	keeper.SavePost(ctx, post)

	createEvent := sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentID.String()),
		sdk.NewAttribute(types.AttributeKeyCreationTime, post.Created.String()),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator.String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(post.PostID),
		Events: sdk.Events{createEvent},
	}
	return &result, nil
}

// handleMsgEditPost handles MsgEditsPost messages
func handleMsgEditPost(ctx sdk.Context, keeper Keeper, msg types.MsgEditPost) (*sdk.Result, error) {

	// Get the existing post
	existing, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	// Checks if the the msg sender is the same as the current owner
	if !msg.Editor.Equals(existing.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check the validity of the current block height respect to the creation date of the post
	if existing.Created.After(msg.EditDate) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date")
	}

	// Edit the post
	existing.Message = msg.Message
	existing.LastEdited = msg.EditDate
	keeper.SavePost(ctx, existing)

	editEvent := sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.LastEdited.String()),
	)
	ctx.EventManager().EmitEvent(editEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.PostID),
		Events: sdk.Events{editEvent},
	}
	return &result, nil
}

func handleMsgAddPostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgAddPostReaction) (*sdk.Result, error) {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	// Create and store the reaction
	reaction := types.NewReaction(msg.Value, msg.User)
	if err := keeper.SaveReaction(ctx, post.PostID, reaction); err != nil {
		return nil, err
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypeReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Value),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction added properly"),
		Events: sdk.Events{event},
	}
	return &result, nil
}

func handleMsgRemovePostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgRemovePostReaction) (*sdk.Result, error) {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	// Remove the reaction
	if err := keeper.RemoveReaction(ctx, post.PostID, msg.User, msg.Reaction); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyReactionValue, msg.Reaction),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction removed properly"),
		Events: sdk.Events{event},
	}
	return &result, nil
}
