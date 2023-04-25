package keeper

import (
	"context"
	"fmt"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v4/x/reactions/types"
	subspacestypes "github.com/desmos-labs/desmos/v4/x/subspaces/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the stored MsgServer interface for the provided keeper
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = &msgServer{}

// AddReaction defines the rpc method for Msg/AddReaction
func (k msgServer) AddReaction(goCtx context.Context, msg *types.MsgAddReaction) (*types.MsgAddReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the user has a profile
	if !k.HasProfile(ctx, msg.User) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you cannot add a reaction without a profile")
	}

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the post exists
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d not found inside subspace %d", msg.PostID, msg.SubspaceID)
	}

	// Make sure the post author has not blocked the user
	if k.HasUserBlocked(ctx, post.Owner, msg.User, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "the post author has blocked you on this subspace")
	}

	// Check the permission to react
	if !k.HasPermission(ctx, post.SubspaceID, post.SectionID, msg.User, types.PermissionsReact) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot add reactions inside this subspace")
	}

	value, ok := msg.Value.GetCachedValue().(types.ReactionValue)
	if !ok {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid reaction value: %s", msg.Value)
	}

	// Make sure the reaction does not exist already
	if k.HasReacted(ctx, msg.SubspaceID, msg.PostID, msg.User, value) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you have already reacted with the same value to this post")
	}

	// Get the next reaction id
	reactionID, err := k.GetNextReactionID(ctx, msg.SubspaceID, msg.PostID)
	if err != nil {
		return nil, err
	}

	// Create and validate the reaction
	reaction := types.NewReaction(
		msg.SubspaceID,
		msg.PostID,
		reactionID,
		value,
		msg.User,
	)
	err = k.ValidateReaction(ctx, reaction)
	if err != nil {
		return nil, err
	}

	// Store the reaction
	k.SaveReaction(ctx, reaction)

	// Update the id for the next reaction
	k.SetNextReactionID(ctx, msg.SubspaceID, msg.PostID, reaction.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.EventTypeAddReaction,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReactionID, fmt.Sprintf("%d", reaction.ID)),
			sdk.NewAttribute(types.AttributeKeyUser, msg.User),
		),
	})

	return &types.MsgAddReactionResponse{
		ReactionID: reaction.ID,
	}, nil
}

// RemoveReaction defines the rpc method for Msg/RemoveReaction
func (k msgServer) RemoveReaction(goCtx context.Context, msg *types.MsgRemoveReaction) (*types.MsgRemoveReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the post exists
	post, found := k.GetPost(ctx, msg.SubspaceID, msg.PostID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "post with id %d not found inside subspace %d", msg.PostID, msg.SubspaceID)
	}

	// Check if the reaction exists
	reaction, found := k.GetReaction(ctx, msg.SubspaceID, msg.PostID, msg.ReactionID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "reaction does not exist")
	}

	// Make sure the user matches the author
	if reaction.Author != msg.User {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "you are not the author of this reaction")
	}

	// Check the permission to remove reaction
	if !k.HasPermission(ctx, post.SubspaceID, post.SectionID, msg.User, types.PermissionsReact) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot remove reactions inside this subspace")
	}

	// Remove the reaction
	k.DeleteReaction(ctx, msg.SubspaceID, msg.PostID, msg.ReactionID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.EventTypeRemoveReaction,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyPostID, fmt.Sprintf("%d", msg.PostID)),
			sdk.NewAttribute(types.AttributeKeyReactionID, fmt.Sprintf("%d", msg.ReactionID)),
		),
	})

	return &types.MsgRemoveReactionResponse{}, nil
}

// AddRegisteredReaction defines the rpc method for Msg/AddRegisteredReaction
func (k msgServer) AddRegisteredReaction(goCtx context.Context, msg *types.MsgAddRegisteredReaction) (*types.MsgAddRegisteredReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the registered reactions
	if !k.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.User, types.PermissionManageRegisteredReactions) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the registered reactions inside this subspace")
	}

	// Get the next reaction id
	registeredReactionID, err := k.GetNextRegisteredReactionID(ctx, msg.SubspaceID)
	if err != nil {
		return nil, err
	}

	// Create the registered reaction and validate it
	reaction := types.NewRegisteredReaction(
		msg.SubspaceID,
		registeredReactionID,
		msg.ShorthandCode,
		msg.DisplayValue,
	)
	err = reaction.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the registered reaction
	k.SaveRegisteredReaction(ctx, reaction)

	// Update the id for the next reaction
	k.SetNextRegisteredReactionID(ctx, msg.SubspaceID, reaction.ID+1)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.ActionAddRegisteredReaction,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, fmt.Sprintf("%d", reaction.ID)),
		),
	})

	return &types.MsgAddRegisteredReactionResponse{
		RegisteredReactionID: reaction.ID,
	}, nil
}

// EditRegisteredReaction defines the rpc method for Msg/EditRegisteredReaction
func (k msgServer) EditRegisteredReaction(goCtx context.Context, msg *types.MsgEditRegisteredReaction) (*types.MsgEditRegisteredReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the registered reaction exists
	reaction, found := k.GetRegisteredReaction(ctx, msg.SubspaceID, msg.RegisteredReactionID)
	if !found {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "registered reaction with id %d not found", msg.RegisteredReactionID)
	}

	// Check the permission to manage the registered reactions
	if !k.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.User, types.PermissionManageRegisteredReactions) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the registered reactions inside this subspace")
	}

	// Update the reaction and validate it
	updated := reaction.Update(types.NewRegisteredReactionUpdate(msg.ShorthandCode, msg.DisplayValue))
	err := updated.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Save the reaction
	k.SaveRegisteredReaction(ctx, updated)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.ActionEditRegisteredReaction,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, fmt.Sprintf("%d", msg.RegisteredReactionID)),
		),
	})

	return &types.MsgEditRegisteredReactionResponse{}, nil
}

// RemoveRegisteredReaction defines the rpc method for Msg/RemoveRegisteredReaction
func (k msgServer) RemoveRegisteredReaction(goCtx context.Context, msg *types.MsgRemoveRegisteredReaction) (*types.MsgRemoveRegisteredReactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check if the registered reaction exists
	if !k.HasRegisteredReaction(ctx, msg.SubspaceID, msg.RegisteredReactionID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "registered reaction with id %d not found", msg.RegisteredReactionID)
	}

	// Check the permission to manage the registered reactions
	if !k.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.User, types.PermissionManageRegisteredReactions) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the registered reactions inside this subspace")
	}

	// Delete the registered reaction
	k.DeleteRegisteredReaction(ctx, msg.SubspaceID, msg.RegisteredReactionID)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.EventTypeRemoveRegisteredReaction,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
			sdk.NewAttribute(types.AttributeKeyRegisteredReactionID, fmt.Sprintf("%d", msg.RegisteredReactionID)),
		),
	})

	return &types.MsgRemoveRegisteredReactionResponse{}, nil
}

// SetReactionsParams defines the rpc method for Msg/SetReactionsParams
func (k msgServer) SetReactionsParams(goCtx context.Context, msg *types.MsgSetReactionsParams) (*types.MsgSetReactionsParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the subspace exists
	if !k.HasSubspace(ctx, msg.SubspaceID) {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", msg.SubspaceID)
	}

	// Check the permission to manage the reaction params
	if !k.HasPermission(ctx, msg.SubspaceID, subspacestypes.RootSectionID, msg.User, types.PermissionManageReactionParams) {
		return nil, errors.Wrap(subspacestypes.ErrPermissionDenied, "you cannot manage the reactions params inside this subspace")
	}

	// Create and validate the params
	params := types.NewSubspaceReactionsParams(msg.SubspaceID, msg.RegisteredReaction, msg.FreeText)
	err := params.Validate()
	if err != nil {
		return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest, err.Error())
	}

	// Store the params
	k.SaveSubspaceReactionsParams(ctx, params)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, sdk.MsgTypeURL(msg)),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.User),
		),
		sdk.NewEvent(
			types.EventTypeSetReactionsParams,
			sdk.NewAttribute(types.AttributeKeySubspaceID, fmt.Sprintf("%d", msg.SubspaceID)),
		),
	})

	return &types.MsgSetReactionsParamsResponse{}, nil
}
