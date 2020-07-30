package keeper

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/posts/types"
)

// NewHandler returns a handler for "magpie" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

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
		case types.MsgRegisterReaction:
			return handleMsgRegisterReaction(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized Posts message type: %v", msg.Type())
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

// ComputeID returns a sha256 hash of the msg's json representation
// nolint: interfacer
func ComputeID(parentID types.PostID, message, subspace string, allowsComment bool,
	creationTime time.Time, creator sdk.AccAddress) types.PostID {
	bz := []byte(parentID.String() + message + subspace + strconv.FormatBool(allowsComment) + creationTime.String() +
		creator.String())
	hash := sha256.Sum256(bz)
	return types.PostID(hex.EncodeToString(hash[:]))
}

// handleMsgCreatePost handles the creation of a new post
func handleMsgCreatePost(ctx sdk.Context, keeper Keeper, msg types.MsgCreatePost) (*sdk.Result, error) {
	post := types.NewPost(
		ComputeID(msg.ParentID, msg.Message, msg.Subspace, msg.AllowsComments, ctx.BlockTime(), msg.Creator),
		msg.ParentID,
		msg.Message,
		msg.AllowsComments,
		msg.Subspace,
		msg.OptionalData,
		ctx.BlockTime(),
		msg.Creator,
	).WithAttachments(msg.Attachments)

	if msg.PollData != nil {
		post = post.WithPollData(*msg.PollData)
	}

	// Check for double posting
	if existing, found := keeper.GetPost(ctx, post.PostID); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("the provided post conflicts with the one having id %s", existing.PostID))
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

	if err := ValidatePost(ctx, keeper, post); err != nil {
		return nil, err
	}

	keeper.SavePost(ctx, post)

	createEvent := sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentID.String()),
		sdk.NewAttribute(types.AttributeKeyPostCreationTime, post.Created.Format(time.RFC3339)),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator.String()),
	)
	ctx.EventManager().EmitEvent(createEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(post.PostID),
		Events: ctx.EventManager().Events(),
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
	if existing.Created.After(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date")
	}

	// Edit the post
	existing.Message = msg.Message
	existing.LastEdited = ctx.BlockTime()

	if err := ValidatePost(ctx, keeper, existing); err != nil {
		return nil, err
	}
	keeper.SavePost(ctx, existing)

	editEvent := sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.LastEdited.Format(time.RFC3339)),
	)
	ctx.EventManager().EmitEvent(editEvent)

	result := sdk.Result{
		Data:   keeper.Cdc.MustMarshalBinaryLengthPrefixed(existing.PostID),
		Events: ctx.EventManager().Events(),
	}
	return &result, nil
}

// extractReactionValueAndShortcode parse the given reaction returning its correct value and shortcode
func extractReactionValueAndShortcode(keeper Keeper, ctx sdk.Context, reaction string, subspace string) (string, string, error) {
	var reactionShortcode, reactionValue string

	// Parse reaction adding the variation selector-16 to let the emoji being readable
	parsedReaction := strings.ReplaceAll(reaction, "️", "")

	if emojiReact, found := types.GetEmojiByShortCodeOrValue(reaction); found {
		reactionShortcode = emojiReact.Shortcodes[0]
		reactionValue = emojiReact.Value
	} else {
		// The reaction is a shortcode that should be registered
		regReaction, registered := keeper.GetRegisteredReaction(ctx, reaction, subspace)
		if !registered {
			return "", "", sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
				fmt.Sprintf("short code %s must be registered before using it", parsedReaction))
		}

		reactionShortcode = regReaction.ShortCode
		reactionValue = regReaction.Value
	}

	return reactionShortcode, reactionValue, nil
}

// handleMsgAddPostReaction handles the adding of a reaction to a post
func handleMsgAddPostReaction(ctx sdk.Context, keeper Keeper, msg types.MsgAddPostReaction) (*sdk.Result, error) {
	// Get the post
	post, found := keeper.GetPost(ctx, msg.PostID)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("post with id %s not found", msg.PostID))
	}

	reactionShortcode, reactionValue, err := extractReactionValueAndShortcode(keeper, ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	postReaction := types.NewPostReaction(reactionShortcode, reactionValue, msg.User)
	if err := keeper.SavePostReaction(ctx, post.PostID, postReaction); err != nil {
		return nil, err
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypePostReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("postReaction added properly"),
		Events: ctx.EventManager().Events(),
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

	reactionShortcode, reactionValue, err := extractReactionValueAndShortcode(keeper, ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	// Remove the reaction
	reaction := types.NewPostReaction(reactionShortcode, reactionValue, msg.User)
	if err := keeper.RemovePostReaction(ctx, post.PostID, reaction); err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Emit the event
	event := sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostID.String()),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User.String()),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction removed properly"),
		Events: ctx.EventManager().Events(),
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
			"user's answers are more than the available ones in Poll",
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
		Events: ctx.EventManager().Events(),
	}
	return &result, nil
}

// handleMsgRegisterReaction handles the reaction registration
func handleMsgRegisterReaction(ctx sdk.Context, keeper Keeper, msg types.MsgRegisterReaction) (*sdk.Result, error) {
	// Check if the shortcode is associated with an emoji
	if _, found := types.GetEmojiByShortCodeOrValue(msg.ShortCode); found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf(
			"shortcode %s represents an emoji and thus can't be used to register a new reaction", msg.ShortCode))
	}

	// Make sure the reaction is already registered
	if _, isAlreadyRegistered := keeper.GetRegisteredReaction(ctx, msg.ShortCode, msg.Subspace); isAlreadyRegistered {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf(
			"reaction with shortcode %s and subspace %s has already been registered", msg.ShortCode, msg.Subspace))
	}

	reaction := types.NewReaction(msg.Creator, msg.ShortCode, msg.Value, msg.Subspace)
	keeper.RegisterReaction(ctx, reaction)

	event := sdk.NewEvent(
		types.EventTypeRegisterReaction,
		sdk.NewAttribute(types.AttributeKeyReactionCreator, reaction.Creator.String()),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reaction.ShortCode),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reaction.Value),
		sdk.NewAttribute(types.AttributeKeyReactionSubSpace, reaction.Subspace),
	)
	ctx.EventManager().EmitEvent(event)

	result := sdk.Result{
		Data:   []byte("reaction registered properly"),
		Events: ctx.EventManager().Events(),
	}

	return &result, nil
}
