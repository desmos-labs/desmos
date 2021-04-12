package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/staging/posts/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the posts MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper}
}

func computePostID(ctx sdk.Context, msg *types.MsgCreatePost) string {
	post := types.Post{
		ParentId:       msg.ParentId,
		Message:        msg.Message,
		Created:        ctx.BlockTime(),
		AllowsComments: msg.AllowsComments,
		Subspace:       msg.Subspace,
		OptionalData:   msg.OptionalData,
		Creator:        msg.Creator,
		Attachments:    msg.Attachments,
		PollData:       msg.PollData,
	}

	bytes, err := post.Marshal()
	if err != nil {
		panic(err)
	}
	hash := sha256.Sum256(bytes)
	return hex.EncodeToString(hash[:])
}

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	post := types.NewPost(
		computePostID(ctx, msg),
		msg.ParentId,
		msg.Message,
		msg.AllowsComments,
		msg.Subspace,
		msg.OptionalData,
		msg.Attachments,
		msg.PollData,
		time.Time{},
		ctx.BlockTime(),
		msg.Creator,
	)

	// Validate the post
	if err := k.ValidatePost(ctx, post); err != nil {
		return nil, err
	}

	// Check for double posting
	if k.DoesPostExist(ctx, post.PostId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the provided post conflicts with the one having id %s", post.PostId)
	}

	// Check if any of the tags have blocked the post creator
	if err := k.IsCreatorBlockedBySomeTags(ctx, post.Attachments, post.Creator, post.Subspace); err != nil {
		return nil, err
	}

	// If valid, check the parent post
	if types.IsValidPostID(post.ParentId) {
		parentPost, found := k.GetPost(ctx, post.ParentId)
		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"parent post with id %s not found", post.ParentId)
		}

		if !parentPost.AllowsComments {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"post with id %s does not allow comments", parentPost.PostId)
		}
	}

	// Save the post
	k.SavePost(ctx, post)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostCreated,
		sdk.NewAttribute(types.AttributeKeyPostID, post.PostId),
		sdk.NewAttribute(types.AttributeKeyPostParentID, post.ParentId),
		sdk.NewAttribute(types.AttributeKeyPostCreationTime, post.Created.Format(time.RFC3339)),
		sdk.NewAttribute(types.AttributeKeyPostOwner, post.Creator),
	))

	return &types.MsgCreatePostResponse{}, nil
}

func (k msgServer) EditPost(goCtx context.Context, msg *types.MsgEditPost) (*types.MsgEditPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the existing post
	existing, found := k.GetPost(ctx, msg.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostId)
	}

	// Checks if the the msg sender is the same as the current owner
	if msg.Editor != existing.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	// Check the validity of the current block height respect to the creation date of the post
	if existing.Created.After(ctx.BlockTime()) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "edit date cannot be before creation date")
	}

	// Edit the post
	existing.Message = msg.Message

	if msg.Attachments != nil {
		// Check if any of the tags have blocked the post creator
		if err := k.IsCreatorBlockedBySomeTags(ctx, msg.Attachments, existing.Creator, existing.Subspace); err != nil {
			return nil, err
		}
		existing.Attachments = msg.Attachments
	}

	if msg.PollData != nil {
		existing.PollData = msg.PollData
	}

	existing.LastEdited = ctx.BlockTime()

	if err := k.ValidatePost(ctx, existing); err != nil {
		return nil, err
	}
	k.SavePost(ctx, existing)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostEdited,
		sdk.NewAttribute(types.AttributeKeyPostID, existing.PostId),
		sdk.NewAttribute(types.AttributeKeyPostEditTime, existing.LastEdited.Format(time.RFC3339)),
	))

	return &types.MsgEditPostResponse{}, nil
}

func (k msgServer) AddPostReaction(goCtx context.Context, msg *types.MsgAddPostReaction) (*types.MsgAddPostReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post
	post, found := k.GetPost(ctx, msg.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostId)
	}

	reactionShortcode, reactionValue, err := k.ExtractReactionValueAndShortcode(ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	postReaction := types.NewPostReaction(reactionShortcode, reactionValue, msg.User)
	if err := k.SavePostReaction(ctx, post.PostId, postReaction); err != nil {
		return nil, err
	}

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostReactionAdded,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostId),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	))

	return &types.MsgAddPostReactionResponse{}, nil
}

func (k msgServer) RemovePostReaction(goCtx context.Context, msg *types.MsgRemovePostReaction) (*types.MsgRemovePostReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post
	post, found := k.GetPost(ctx, msg.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %s not found", msg.PostId)
	}

	reactionShortcode, reactionValue, err := k.ExtractReactionValueAndShortcode(ctx, msg.Reaction, post.Subspace)
	if err != nil {
		return nil, err
	}

	// Remove the registeredReactions
	reaction := types.NewPostReaction(reactionShortcode, reactionValue, msg.User)
	if err := k.DeletePostReaction(ctx, post.PostId, reaction); err != nil {
		return nil, err
	}

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypePostReactionRemoved,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostId),
		sdk.NewAttribute(types.AttributeKeyPostReactionOwner, msg.User),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reactionValue),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reactionShortcode),
	))

	return &types.MsgRemovePostReactionResponse{}, nil
}

func (k msgServer) RegisterReaction(goCtx context.Context, msg *types.MsgRegisterReaction) (*types.MsgRegisterReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the shortcode is associated with an emoji
	if _, found := types.GetEmojiByShortCodeOrValue(msg.ShortCode); found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"shortcode %s represents an emoji and thus can't be used to register a new registeredReactions", msg.ShortCode)
	}

	// Make sure the given reaction isn't already registered
	if _, isAlreadyRegistered := k.GetRegisteredReaction(ctx, msg.ShortCode, msg.Subspace); isAlreadyRegistered {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"registeredReactions with shortcode %s and subspace %s has already been registered", msg.ShortCode, msg.Subspace)
	}

	reaction := types.NewRegisteredReaction(msg.Creator, msg.ShortCode, msg.Value, msg.Subspace)
	k.SaveRegisteredReaction(ctx, reaction)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeRegisterReaction,
		sdk.NewAttribute(types.AttributeKeyReactionCreator, reaction.Creator),
		sdk.NewAttribute(types.AttributeKeyReactionShortCode, reaction.ShortCode),
		sdk.NewAttribute(types.AttributeKeyPostReactionValue, reaction.Value),
		sdk.NewAttribute(types.AttributeKeyReactionSubSpace, reaction.Subspace),
	))

	return &types.MsgRegisterReactionResponse{}, nil
}

func (k msgServer) AnswerPoll(goCtx context.Context, msg *types.MsgAnswerPoll) (*types.MsgAnswerPollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks if the post exists
	post, found := k.GetPost(ctx, msg.PostId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"post with id %s doesn't exist", msg.PostId)
	}

	// Make sure the post has a poll
	if post.PollData == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"no poll associated with ID: %s", msg.PostId)
	}

	// Make sure the poll is not closed
	if post.PollData.EndDate.Before(ctx.BlockTime()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the poll associated with ID %s was closed at %s", post.PostId, post.PollData.EndDate)
	}

	// Check if the poll allows multiple answers
	if len(msg.UserAnswers) > 1 && !post.PollData.AllowsMultipleAnswers {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"the poll associated with ID %s doesn't allow multiple answers", post.PostId)
	}

	// Check if the user answers are more than the answers provided by the poll
	if len(msg.UserAnswers) > len(post.PollData.ProvidedAnswers) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest,
			"user's answers are more than the available ones inside the poll")
	}

	// Make sure that each answer provided by the user matches with one of the provided ones by the poll creator
	for _, answer := range msg.UserAnswers {
		var found = false
		for _, providedAnswer := range post.PollData.ProvidedAnswers {
			if answer == providedAnswer.ID {
				found = true
				break
			}
		}

		if !found {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"answer with ID %s isn't one of the poll's provided answers", answer)
		}
	}

	pollAnswers := k.GetPollAnswersByUser(ctx, post.PostId, msg.Answerer)

	// Check if the poll allows to edit previous answers
	if len(pollAnswers) > 0 && !post.PollData.AllowsAnswerEdits {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"post with ID %s doesn't allow answers' edits", post.PostId)
	}

	userPollAnswers := types.NewUserAnswer(msg.UserAnswers, msg.Answerer)

	k.SavePollAnswers(ctx, post.PostId, userPollAnswers)

	// Emit the event
	ctx.EventManager().EmitEvent(sdk.NewEvent(
		types.EventTypeAnsweredPoll,
		sdk.NewAttribute(types.AttributeKeyPostID, msg.PostId),
		sdk.NewAttribute(types.AttributeKeyPollAnswerer, msg.Answerer),
	))

	return &types.MsgAnswerPollResponse{}, nil
}
