package keeper

import (
	"context"
	"fmt"
	"sort"
	"time"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/desmos-labs/desmos/v6/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"
)

type msgServer struct {
	*Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface
// for the provided k
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = &msgServer{}

// CreatePost defines the rpc method for Msg/CreatePost
func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the author has a profile
	if !k.HasProfile(ctx, msg.Author) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot create a post without having a profile")
	}

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the section exists
	if !k.HasSection(ctx, msg.SubspaceID, msg.SectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace section with id %d not found", msg.SectionID)
	}

	// Check the permission to create this post
	canWrite := k.HasPermission(ctx, msg.SubspaceID, msg.SectionID, msg.Author, types.PermissionWrite)
	canComment := msg.ConversationID != 0 && k.HasPermission(ctx, msg.SubspaceID, msg.SectionID, msg.Author, types.PermissionComment)
	if !canWrite && !canComment {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot create posts nor comment inside this section")
	}

	// Get the next post id
	postID, err := k.GetNextPostID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create and validate the post
	post := types.NewPost(
		msg.SubspaceID,
		msg.SectionID,
		postID,
		msg.ExternalID,
		msg.Text,
		msg.Author,
		msg.ConversationID,
		msg.Entities,
		msg.Tags,
		msg.ReferencedPosts,
		msg.ReplySettings,
		ctx.BlockTime(),
		nil,
		msg.Author,
	)
	err = k.ValidatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	// Store the post
	k.SavePost(ctx, post)

	// Update the id for the next post
	k.SetNextPostID(ctx, msg.SubspaceID, post.ID+1)

	// Unpack the attachments
	attachments, err := types.UnpackAttachments(k.cdc, msg.Attachments)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachments: %s", err)
	}

	// Store the attachments
	for _, content := range attachments {
		_, err = k.storePostAttachment(ctx, post.SubspaceID, post.ID, content)
		if err != nil {
			return nil, err
		}
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Author),
		),
		sdk.NewEvent(
			types.EventTypeCreatePost,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeySectionID, fmt.Sprintf("%d", msg.SectionID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", post.ID)),
			sdk.NewAttribute(types.AttributeKeyAuthor, msg.Author),
			sdk.NewAttribute(types.AttributeKeyCreationTime, post.CreationDate.Format(time.RFC3339)),
		),
	})

	return &types.MsgCreatePostResponse{
		PostID:       post.ID,
		CreationDate: post.CreationDate,
	}, nil
}

// EditPost defines the rpc method for Msg/EditPost
func (k msgServer) EditPost(goCtx context.Context, msg *types.MsgEditPost) (*types.MsgEditPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d not found", msg.PostID)
	}

	// Make sure the editor matches the owner
	if post.Owner != msg.Editor {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you are not the owner of this post")
	}

	// Check the permission to create content
	if !k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Editor, types.PermissionEditOwnContent) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot edit content inside this subspace")
	}

	// Update the post and validate it
	updateTime := ctx.BlockTime()
	update := types.NewPostUpdate(msg.Text, msg.Entities, msg.Tags, updateTime)
	updatedPost := post.Update(update)
	err := k.ValidatePost(ctx, updatedPost)
	if err != nil {
		return nil, err
	}

	// Store the update
	k.SavePost(ctx, updatedPost)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Editor),
		),
		sdk.NewEvent(
			types.EventTypeEditPost,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyLastEditTime, updateTime.Format(time.RFC3339)),
		),
	})

	return &types.MsgEditPostResponse{
		EditDate: updateTime,
	}, nil
}

// DeletePost defines the rpc method for Msg/DeletePost
func (k msgServer) DeletePost(goCtx context.Context, msg *types.MsgDeletePost) (*types.MsgDeletePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Check the permission to remove the post
	isModerator := k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Signer, types.PermissionModerateContent)
	canEdit := post.Owner == msg.Signer && k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Signer, types.PermissionEditOwnContent)
	if !isModerator && !canEdit {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot edit content inside this subspace")
	}

	// Delete the post
	k.Keeper.DeletePost(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeDeletePost,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
		),
	})

	return &types.MsgDeletePostResponse{}, nil
}

// storePostAttachment allows to easily store a post attachment, returning the attachment id used and any error
func (k msgServer) storePostAttachment(ctx sdk.Context, subspaceID uint64, postID uint64, content types.AttachmentContent) (attachmentID uint32, err error) {
	// Perform poll checks
	if poll, ok := content.(*types.Poll); ok {
		// Make sure no tally results are provided
		if poll.FinalTallyResults != nil {
			return 0, errors.Wrapf(sdkerrors.ErrInvalidRequest, "poll tally results must be nil")
		}

		// Make sure the end date is in the future
		if poll.EndDate.Before(ctx.BlockTime()) {
			return 0, errors.Wrapf(sdkerrors.ErrInvalidRequest, "poll end date must be in the future")
		}
	}

	// Get the next attachment id
	attachmentID, err = k.GetNextAttachmentID(ctx, subspaceID, postID)
	if err != nil {
		return
	}

	// Create the attachment and validate it
	attachment := types.NewAttachment(subspaceID, postID, attachmentID, content)
	err = attachment.Validate()
	if err != nil {
		return 0, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment content: %s", err)
	}

	// Save the attachment
	k.SaveAttachment(ctx, attachment)

	// Store the poll inside the active queue
	if types.IsPoll(attachment) {
		k.InsertActivePollQueue(ctx, attachment)
	}

	// Update the id for the next attachment
	k.SetNextAttachmentID(ctx, subspaceID, postID, attachment.ID+1)

	return attachmentID, nil
}

// AddPostAttachment defines the rpc method for Msg/AddPostAttachment
func (k msgServer) AddPostAttachment(goCtx context.Context, msg *types.MsgAddPostAttachment) (*types.MsgAddPostAttachmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Make sure the editor matches the owner
	if post.Owner != msg.Editor {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you are not the author of this post")
	}

	// Check the permission to edit content
	if !k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Editor, types.PermissionEditOwnContent) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot edit content inside this subspace")
	}

	// Unpack the content
	var content types.AttachmentContent
	err := k.cdc.UnpackAny(msg.Content, &content)
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid attachment content: %s", err)
	}

	// Save the attachment
	attachmentID, err := k.storePostAttachment(ctx, msg.SubspaceID, msg.PostID, msg.Content.GetCachedValue().(types.AttachmentContent))
	if err != nil {
		return nil, err
	}

	// Update the post edit time
	updateTime := ctx.BlockTime()
	post.LastEditedDate = &updateTime
	err = k.ValidatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	k.SavePost(ctx, post)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Editor),
		),
		sdk.NewEvent(
			types.EventTypeAddPostAttachment,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyAttachmentID, fmt.Sprintf("%d", attachmentID)),
			sdk.NewAttribute(types.AttributeKeyLastEditTime, post.LastEditedDate.Format(time.RFC3339)),
		),
	})

	return &types.MsgAddPostAttachmentResponse{
		AttachmentID: attachmentID,
		EditDate:     updateTime,
	}, nil
}

// RemovePostAttachment defines the rpc method for Msg/RemovePostAttachment
func (k msgServer) RemovePostAttachment(goCtx context.Context, msg *types.MsgRemovePostAttachment) (*types.MsgRemovePostAttachmentResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Check the permission to remove the attachment
	isModerator := k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Editor, types.PermissionModerateContent)
	canEdit := post.Owner == msg.Editor && k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Editor, types.PermissionEditOwnContent)
	if !isModerator && !canEdit {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot edit content inside this subspace")
	}

	// Check if the attachment exists
	if !k.HasAttachment(ctx, msg.SubspaceID, msg.PostID, msg.AttachmentID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "attachment with id %d not found", msg.AttachmentID)
	}

	// Remove the post attachment
	k.DeleteAttachment(ctx, msg.SubspaceID, msg.PostID, msg.AttachmentID)

	// Update the post edit time and validate it
	updateTime := ctx.BlockTime()
	post.LastEditedDate = &updateTime
	err := k.ValidatePost(ctx, post)
	if err != nil {
		return nil, err
	}

	// Save the post
	k.SavePost(ctx, post)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Editor),
		),
		sdk.NewEvent(
			types.EventTypeRemovePostAttachment,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyAttachmentID, fmt.Sprintf("%d", msg.AttachmentID)),
			sdk.NewAttribute(types.AttributeKeyLastEditTime, post.LastEditedDate.Format(time.RFC3339)),
		),
	})

	return &types.MsgRemovePostAttachmentResponse{
		EditDate: updateTime,
	}, nil
}

// AnswerPoll defines the rpc method for Msg/AnswerPoll
func (k msgServer) AnswerPoll(goCtx context.Context, msg *types.MsgAnswerPoll) (*types.MsgAnswerPollResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the author has a profile
	if !k.HasProfile(ctx, msg.Signer) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot answer a poll without having a profile")
	}

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Make sure the post exists
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Check the permission to interact with content
	if !k.HasPermission(ctx, msg.SubspaceID, post.SectionID, msg.Signer, types.PermissionInteractWithContent) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot interact with content inside this subspace")
	}

	// Get the poll
	poll, found := k.GetPoll(ctx, msg.SubspaceID, msg.PostID, msg.PollID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "poll with id %d does not exist", msg.PollID)
	}

	// Make sure the poll is still active
	if ctx.BlockTime().After(poll.EndDate) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "the poll voting period has already ended")
	}

	alreadyAnswered := k.HasUserAnswer(ctx, msg.SubspaceID, msg.PostID, msg.PollID, msg.Signer)

	// Make sure the user is not trying to edit the answer when the poll does not allow it
	if alreadyAnswered && !poll.AllowsAnswerEdits {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot edit this poll's answer")
	}

	// Make sure the user not answering with multiple options when the poll does not allow it
	if len(msg.AnswersIndexes) > 1 && !poll.AllowsMultipleAnswers {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "only one answer is allowed on this post")
	}

	// Sort the answers indexes
	sort.Slice(msg.AnswersIndexes, func(i, j int) bool {
		return msg.AnswersIndexes[i] < msg.AnswersIndexes[j]
	})

	// Make sure the answer indexes exist
	maxProvidedIndex := uint32(len(poll.ProvidedAnswers) - 1)
	maxAnswerIndex := msg.AnswersIndexes[len(msg.AnswersIndexes)-1]
	if maxAnswerIndex > maxProvidedIndex {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid answer index: %d", maxAnswerIndex)
	}

	// Store the user answer
	answer := types.NewUserAnswer(msg.SubspaceID, msg.PostID, msg.PollID, msg.AnswersIndexes, msg.Signer)
	k.SaveUserAnswer(ctx, answer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Signer),
		),
		sdk.NewEvent(
			types.EventTypeAnswerPoll,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyPollID, fmt.Sprintf("%d", msg.PollID)),
		),
	})

	return &types.MsgAnswerPollResponse{}, nil
}

// UpdateParams updates the module parameters
func (m msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	authority := m.authority
	if authority != msg.Authority {
		return nil, errors.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", authority, msg.Authority)
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	m.SetParams(ctx, msg.Params)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Authority),
		),
	})

	return &types.MsgUpdateParamsResponse{}, nil
}

// MovePost defines the rpc method for Msg/MovePost
func (k msgServer) MovePost(goCtx context.Context, msg *types.MsgMovePost) (*types.MsgMovePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Make sure the post exists
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Check if the target subspace exists
	if !k.HasSubspace(ctx, msg.TargetSubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.TargetSubspaceID)
	}

	// Check if the target section exists
	if !k.HasSection(ctx, msg.TargetSubspaceID, msg.TargetSectionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace section with id %d not found", msg.TargetSectionID)
	}

	// Check the sender is the owner of the post
	if post.Owner != msg.Owner {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you are not the owner of the post")
	}

	// Check the permission to move the post
	if !k.HasPermission(ctx, msg.TargetSubspaceID, msg.TargetSectionID, msg.Owner, types.PermissionWrite) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you don't have write permission on the target section")
	}

	// Get the post id within the new subspace
	newPostID, err := k.GetNextPostID(ctx, msg.TargetSubspaceID)
	if err != nil {
		return nil, err
	}

	// Update the post and validate it
	updateTime := ctx.BlockTime()
	postMove := types.NewPostMove(msg.TargetSubspaceID, msg.TargetSectionID, newPostID, updateTime)
	updatedPost := postMove.Update(post)
	err = k.ValidatePost(ctx, updatedPost)
	if err != nil {
		return nil, err
	}

	k.Keeper.SavePost(ctx, updatedPost)

	// Update the id for the next post
	k.SetNextPostID(ctx, msg.TargetSubspaceID, newPostID+1)

	// Move all post attachments
	maxAttachmentID := uint32(0)
	k.IteratePostAttachments(ctx, msg.SubspaceID, msg.PostID, func(attachment types.Attachment) (stop bool) {
		attachmentMove := types.NewAttachmentMove(updatedPost.SubspaceID, updatedPost.ID)
		updatedAttachment := attachmentMove.Update(attachment)
		k.SaveAttachment(ctx, updatedAttachment)

		if updatedAttachment.ID > maxAttachmentID {
			maxAttachmentID = updatedAttachment.ID
		}

		if k.isActivePoll(ctx, attachment) {
			k.InsertActivePollQueue(ctx, updatedAttachment)
		}
		return false
	})

	// Update the id for the next attachment
	k.SetNextAttachmentID(ctx, msg.TargetSubspaceID, newPostID, maxAttachmentID+1)

	// Delete the post
	k.Keeper.DeletePost(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner),
		),
		sdk.NewEvent(
			types.EventTypeMovePost,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyNewSubspaceID, fmt.Sprintf("%d", msg.TargetSubspaceID)),
			sdk.NewAttribute(types.AttributeKeyNewPostID, fmt.Sprintf("%d", updatedPost.ID)),
		),
	})

	return &types.MsgMovePostResponse{}, nil
}

// RequestPostOwnerTransfer defines the rpc method for Msg/RequestPostOwnerTransfer
func (k msgServer) RequestPostOwnerTransfer(goCtx context.Context, msg *types.MsgRequestPostOwnerTransfer) (*types.MsgRequestPostOwnerTransferResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the sender has profile
	if !k.HasProfile(ctx, msg.Sender) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post without having a profile")
	}

	// Check if the receiver has profile
	if !k.HasProfile(ctx, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post to the user having no profile")
	}

	// Make sure the receiver has not blocked the sender
	if k.HasUserBlocked(ctx, msg.Receiver, msg.Sender, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "receiver has blocked you")
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Make sure the sender matches the owner
	if post.Owner != msg.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot transfer a post that you do not own")
	}

	// Check if the post owner transfer request exists
	if k.HasPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID) {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request already exists")
	}

	// Save the post owner transfer request
	k.SavePostOwnerTransferRequest(ctx, types.NewPostOwnerTransferRequest(
		msg.SubspaceID,
		msg.PostID,
		msg.Receiver,
		msg.Sender,
	))

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeRequestPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgRequestPostOwnerTransferResponse{}, nil
}

// CancelPostOwnerTransferRequest defines the rpc method for Msg/CancelPostOwnerTransfer
func (k msgServer) CancelPostOwnerTransferRequest(goCtx context.Context, msg *types.MsgCancelPostOwnerTransferRequest) (*types.MsgCancelPostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request sender matches the sender
	if request.Sender != msg.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot cancel a post owner transfer request that you are not the sender")
	}

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.EventTypeCancelPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeySender, msg.Sender),
		),
	})

	return &types.MsgCancelPostOwnerTransferRequestResponse{}, nil
}

// AcceptPostOwnerTransferRequest defines the rpc method for Msg/AcceptPostOwnerTransfer
func (k msgServer) AcceptPostOwnerTransferRequest(goCtx context.Context, msg *types.MsgAcceptPostOwnerTransferRequest) (*types.MsgAcceptPostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the receiver has profile
	if !k.HasProfile(ctx, msg.Receiver) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot accept a post owner transfer request without having a profile")
	}

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request receiver matches the receiver
	if request.Receiver != msg.Receiver {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot accept a post owner transfer request that you are not the receiver")
	}

	// Get the post
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d does not exist", msg.PostID)
	}

	// Make sure the post owner matches the sender
	if post.Owner != request.Sender {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "the sender of the post owner transfer request does not own the post")
	}

	// Update the post and validate it
	newPost := types.NewOwnerTransfer(msg.Receiver, ctx.BlockTime()).Update(post)
	err := newPost.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the update
	k.SavePost(ctx, newPost)

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeAcceptPostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	})

	return &types.MsgAcceptPostOwnerTransferRequestResponse{}, nil
}

// RefusePostOwnerTransferRequest defines the rpc method for Msg/RefusePostOwnerTransfer
func (k msgServer) RefusePostOwnerTransferRequest(goCtx context.Context, msg *types.MsgRefusePostOwnerTransferRequest) (*types.MsgRefusePostOwnerTransferRequestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Get the post owner transfer request
	request, found := k.GetPostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "post owner transfer request does not exists")
	}

	// Make sure the request receiver matches the receiver
	if request.Receiver != msg.Receiver {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot refuse a post owner transfer request that you are not the receiver")
	}

	// Delete the post owner transfer request
	k.DeletePostOwnerTransferRequest(ctx, msg.SubspaceID, msg.PostID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Receiver),
		),
		sdk.NewEvent(
			types.EventTypeRefusePostOwnerTransfer,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
		),
	})

	return &types.MsgRefusePostOwnerTransferRequestResponse{}, nil
}
