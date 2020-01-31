package keeper

import (
	"fmt"
	"time"

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
		case types.MsgClosePollPost:
			return handleMsgClosePollPost(ctx, keeper, msg)
		case types.MsgAnswerPollPost:
			return handleMsgAnswerPollPost(ctx, keeper, msg)
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
		msg.PollData,
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

// handleMsgAnswerPollPost handles the answer to a poll post
func handleMsgAnswerPollPost(ctx sdk.Context, keeper Keeper, msg types.MsgAnswerPollPost) (*sdk.Result, error) {
	// checks if post exists
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Post with id %s doesn't exist", msg.PostID))
	}

	// checks if post has a poll
	if post.PollData == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No poll associated with ID: %s", msg.PostID))
	}

	// checks if the poll is already closed or not
	if time.Now().UTC().After(post.PollData.EndDate) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The poll associated with ID %s was closed at %s",
			post.PostID, post.PollData.EndDate))
	}

	// checks if the post's poll allows multiple answers
	if len(msg.UserAnswers) > 1 && !post.PollData.AllowsMultipleAnswers {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The poll associated with ID %s doesn't allow multiple answers",
			post.PostID))
	}

	// check if the user answers are more than the answers provided by the poll
	if len(msg.UserAnswers) > len(post.PollData.ProvidedAnswers) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("User's answers are more than the available ones in Poll"))
	}

	pollAnswers := keeper.GetPostPollAnswersByUser(ctx, post.PostID, msg.Answerer)

	// check if the poll allows to edit previous answers
	if len(pollAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Post with ID %s doesn't allow answers' edits", post.PostID))
	}

	userPollAnswers := types.NewAnswersDetails(msg.UserAnswers, msg.Answerer)

	keeper.SavePollUserAnswers(ctx, post.PostID, userPollAnswers)

	answerEvent := sdk.NewEvent(
		types.EventTypeAnsweredPoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPollAnswerer, msg.Answerer.String()),
	)

	ctx.EventManager().EmitEvent(answerEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed("Answered to poll correctly"),
		Events: sdk.Events{answerEvent},
	}
	return &result, nil
}

// handleMsgClosePollPost handles the closure of a poll post
func handleMsgClosePollPost(ctx sdk.Context, keeper Keeper, msg types.MsgClosePollPost) (*sdk.Result, error) {
	// check if post exists
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Post with id %s doesn't exists", msg.PostID))
	}

	// check if the creator of message is the owner of the post
	if !post.Creator.Equals(msg.Creator) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Only the poll creator can close it, %s", post.Creator))
	}

	// check if post has a poll
	if post.PollData == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("No poll associated with this post ID: %s", msg.PostID))
	}

	// check if the post has already been closed.
	if time.Now().UTC().After(post.PollData.EndDate) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("The poll associated with this post ID is already closed"))
	}

	keeper.ClosePollPost(ctx, msg.PostID)

	closeEvent := sdk.NewEvent(
		types.EventTypeClosePoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostOwner, msg.Creator.String()),
	)

	ctx.EventManager().EmitEvent(closeEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(fmt.Sprintf("Poll closed correctly, %s", msg.Message)),
		Events: sdk.Events{closeEvent},
	}

	return &result, nil
}
