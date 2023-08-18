package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	poststypes "github.com/desmos-labs/desmos/v6/x/posts/types"
	subspacestypes "github.com/desmos-labs/desmos/v6/x/subspaces/types"

	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.getSubspaceDataEntries(ctx),
		k.GetRegisteredReactions(ctx),
		k.getPostDataEntries(ctx),
		k.GetReactions(ctx),
		k.GetReactionsParams(ctx),
	)
}

// getSubspaceDataEntries returns the subspaces data entries stored in the given context
func (k Keeper) getSubspaceDataEntries(ctx sdk.Context) []types.SubspaceDataEntry {
	var entries []types.SubspaceDataEntry
	k.sk.IterateSubspaces(ctx, func(subspace subspacestypes.Subspace) (stop bool) {
		nextRegisteredReactionID, err := k.GetNextRegisteredReactionID(ctx, subspace.ID)
		if err != nil {
			nextRegisteredReactionID = 1
		}

		entries = append(entries, types.NewSubspaceDataEntry(subspace.ID, nextRegisteredReactionID))

		return false
	})
	return entries
}

// getPostDataEntries returns the post data entries stored in the given context
func (k Keeper) getPostDataEntries(ctx sdk.Context) []types.PostDataEntry {
	var entries []types.PostDataEntry
	k.pk.IteratePosts(ctx, func(post poststypes.Post) (stop bool) {
		nextReactionID, err := k.GetNextReactionID(ctx, post.SubspaceID, post.ID)
		if err != nil {
			nextReactionID = 1
		}

		entries = append(entries, types.NewPostDataEntry(post.SubspaceID, post.ID, nextReactionID))

		return false
	})
	return entries
}

// --------------------------------------------------------------------------------------------------------------------

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	// Initialize the subspaces data
	for _, entry := range data.SubspacesData {
		k.SetNextRegisteredReactionID(ctx, entry.SubspaceID, entry.RegisteredReactionID)
	}

	// Initialize all the registered reactions
	for _, reaction := range data.RegisteredReactions {
		k.SaveRegisteredReaction(ctx, reaction)
	}

	// Initialize the posts data
	for _, entry := range data.PostsData {
		k.SetNextReactionID(ctx, entry.SubspaceID, entry.PostID, entry.ReactionID)
	}

	// Initialize all the posts reactions
	for _, reaction := range data.Reactions {
		k.SaveReaction(ctx, reaction)
	}

	// Initialize the params
	for _, params := range data.SubspacesParams {
		k.SaveSubspaceReactionsParams(ctx, params)
	}
}
