package keeper

import (
	"fmt"
	"time"

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
		case types.MsgClosePollPost:
			return handleMsgClosePollPost(ctx, keeper, msg)
		case types.MsgAnswerPollPost:
			return handleMsgAnswerPollPost(ctx, keeper, msg)
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
		msg.CreationDate,
		msg.Creator,
		msg.Medias,
		msg.PollData,
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
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator.String()),
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
	if !msg.Editor.Equals(existing.Creator) {
		return sdk.ErrUnauthorized("Incorrect owner").Result()
	}

	// Check the validity of the current block height respect to the creation date of the post
	if existing.Created.After(msg.EditDate) {
		return sdk.ErrUnknownRequest("Edit date cannot be before creation date").Result()
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

	return sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.PostID),
		Events: sdk.Events{editEvent},
	}
}

func handleMsgAddPostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgAddPostReaction) sdk.Result {

	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Create and store the reaction
	reaction := types.NewReaction(msg.Value, msg.User)
	if err := keeper.SaveReaction(ctx, post.PostID, reaction); err != nil {
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
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", msg.PostID)).Result()
	}

	// Remove the reaction
	if err := keeper.RemoveReaction(ctx, post.PostID, msg.User, msg.Reaction); err != nil {
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

// handleMsgAnswerPollPost handles the answer to a poll post
func handleMsgAnswerPollPost(ctx sdk.Context, keeper Keeper, msg types.MsgAnswerPollPost) sdk.Result {
	// checks if post exists
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s doesn't exists", msg.PostID)).Result()
	}

	// checks if post has a poll
	if post.PollData == nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("No poll associated with ID: %s", msg.PostID)).Result()
	}

	// checks if the poll is already closed or not
	if time.Now().UTC().After(post.PollData.EndDate) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("The poll associated with ID %s was closed at %s",
			post.PostID, post.PollData.EndDate)).Result()
	}

	// checks if the post's poll allows multiple answers
	if len(msg.UserAnswers) > 1 && !post.PollData.AllowsMultipleAnswers {
		return sdk.ErrUnknownRequest(fmt.Sprintf("The poll associated with ID %s doesn't support multiple answers",
			post.PostID)).Result()
	}

	// check if the user answers are more than the answers provided by the poll
	if len(msg.UserAnswers) > len(post.PollData.ProvidedAnswers) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("User's answers are more than the available ones in Poll")).Result()
	}

	userAnswers := keeper.GetPollPostUserAnswers(ctx, post.PostID, msg.Answerer)

	// check if the poll allows to edit previous answers
	if len(userAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with ID %s doesn't allow answers' edits", post.PostID)).Result()
	}

	keeper.SavePollPostAnswers(ctx, post.PostID, msg.UserAnswers, msg.Answerer)

	event := sdk.NewEvent(
		types.EventTypeAnsweredPoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPollAnswerer, msg.Answerer.String()),
	)

	return sdk.Result{
		Data:   []byte("Answered to poll correctly"),
		Events: sdk.Events{event},
	}
}

// handleMsgClosePollPost handles the closure of a poll post
func handleMsgClosePollPost(ctx sdk.Context, keeper Keeper, msg types.MsgClosePollPost) sdk.Result {
	// check if post exists
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s doesn't exists", msg.PostID)).Result()
	}

	// check if the creator of message is the owner of the post
	if post.Creator.Equals(msg.Creator) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Only the poll creator can close it, %s", post.Creator)).Result()
	}

	// check if post has a poll
	if post.PollData == nil {
		return sdk.ErrUnknownRequest(fmt.Sprintf("No poll associated with this post ID: %s", msg.PostID)).Result()
	}

	// check if the post has already been closed.
	if time.Now().UTC().After(post.PollData.EndDate) {
		return sdk.ErrUnknownRequest(fmt.Sprintf("The poll associated with this post ID is already closed")).Result()
	}

	keeper.ClosePollPost(ctx, msg.PostID)

	event := sdk.NewEvent(
		types.EventTypeClosePoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostOwner, msg.Creator.String()),
	)

	return sdk.Result{
		Data:   []byte(fmt.Sprintf("Poll closed correctly, %s", msg.Message)),
		Events: sdk.Events{event},
	}
}
