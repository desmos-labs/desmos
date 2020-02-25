package keeper

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/desmos-labs/desmos/x/posts/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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
		case types.MsgAnswerPoll:
			return handleMsgAnswerPollPost(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func GetPostHashtags(message string) []string {
	re := regexp.MustCompile(`(?:|^)#[A-Za-z0-9]+(?:|$)`)
	hashtags := re.FindAllStringSubmatch(message, -1)

	hts := []string{}
	for _, hashtagSlice := range hashtags {
		hts = append(hts, hashtagSlice[0])
	}

	return hts
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
	).WithMedias(msg.Medias)

	if msg.PollData != nil {
		post = post.WithPollData(*msg.PollData)
	}

	// Check for double posting
	if existing, found := keeper.IsPostConflicting(ctx, post); found {
		msg := `the provided post conflicts with the one having id %s. Please check that either their creation date, subspace or creator are different`
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf(msg, existing.PostID))
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

	// Handle post hashtags
	hashtags := GetPostHashtags(msg.Message)
	for _, hashtag := range hashtags {
		keeper.SavePostHashtag(ctx, hashtag, post.PostID)
	}

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

// handleMsgEditPost handles the edit of posts
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
	// Edit hashtags
	if existingHashtags := GetPostHashtags(existing.Message); len(existingHashtags) != 0 {
		keeper.RemovePostHashtags(ctx, existing.PostID, existingHashtags)
	}

	// Edit the post
	existing.Message = msg.Message
	existing.LastEdited = msg.EditDate
	keeper.SavePost(ctx, existing)

	newHashtags := GetPostHashtags(msg.Message)
	for _, newHashtag := range newHashtags {
		keeper.SavePostHashtag(ctx, newHashtag, existing.PostID)
	}

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

// handleMsgAddPostReaction handles the adding of a reaction to a post
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

// handleMsgRemovePostReaction handles the removal of a reaction from a post
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

// checkPostPollValid performs all the checks to ensure the post with the given id exists, contains a poll and such poll has not been closed
func checkPostPollValid(ctx sdk.Context, id types.PostID, keeper Keeper) (*types.Post, error) {
	// checks if post exists
	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s doesn't exist", id))
	}

	// checks if post has a poll
	if post.PollData == nil {
		return &post, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("no poll associated with ID: %s", id))
	}

	// checks if the poll is already closed or not
	if !post.PollData.Open {
		return &post, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("the poll associated with ID %s was closed at %s", post.PostID, post.PollData.EndDate),
		)
	}

	return &post, nil
}

// answerExist checks if the answer is contained in providedAnswers slice
func answerExist(providedAnswers []types.AnswerID, answer types.AnswerID) bool {
	for _, ans := range providedAnswers {
		if ans == answer {
			return true
		}
	}
	return false
}

// handleMsgAnswerPollPost handles the answer to a poll post
func handleMsgAnswerPollPost(ctx sdk.Context, keeper Keeper, msg types.MsgAnswerPoll) (*sdk.Result, error) {

	post, err := checkPostPollValid(ctx, msg.PostID, keeper)
	if err != nil {
		return nil, err
	}

	// checks if the post's poll allows multiple answers
	if len(msg.UserAnswers) > 1 && !post.PollData.AllowsMultipleAnswers {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("the poll associated with ID %s doesn't allow multiple answers",
				post.PostID),
		)
	}

	// check if the user answers are more than the answers provided by the poll
	if len(msg.UserAnswers) > len(post.PollData.ProvidedAnswers) {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("user's answers are more than the available ones in Poll"),
		)
	}

	for _, answer := range msg.UserAnswers {
		if found := answerExist(post.PollData.ProvidedAnswers.ExtractAnswersIDs(), answer); !found {
			return nil, sdkerrors.Wrap(
				sdkerrors.ErrInvalidRequest,
				fmt.Sprintf(
					"answer with ID %s isn't one of the poll's provided answers",
					strconv.FormatUint(uint64(answer), 10)),
			)
		}
	}

	pollAnswers := keeper.GetPollAnswersByUser(ctx, post.PostID, msg.Answerer)

	// check if the poll allows to edit previous answers
	if len(pollAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return nil, sdkerrors.Wrap(
			sdkerrors.ErrInvalidRequest,
			fmt.Sprintf("post with ID %s doesn't allow answers' edits", post.PostID),
		)
	}

	userPollAnswers := types.NewUserAnswer(msg.UserAnswers, msg.Answerer)

	keeper.SavePollAnswers(ctx, post.PostID, userPollAnswers)

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
