package simulation

// DONTCOVER

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	postssim "github.com/desmos-labs/desmos/v7/x/posts/simulation"
	poststypes "github.com/desmos-labs/desmos/v7/x/posts/types"
	"github.com/desmos-labs/desmos/v7/x/reactions/types"
	subspacessim "github.com/desmos-labs/desmos/v7/x/subspaces/simulation"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
)

// RandomizedGenState generates a random GenesisState
func RandomizedGenState(simState *module.SimulationState) {
	// Read the subspaces data
	subspacesGenesisBz := simState.GenState[subspacestypes.ModuleName]
	var subspacesGenesis subspacestypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(subspacesGenesisBz, &subspacesGenesis)

	// Read the posts data
	postsGenesisBz := simState.GenState[poststypes.ModuleName]
	var postsGenesis poststypes.GenesisState
	simState.Cdc.MustUnmarshalJSON(postsGenesisBz, &postsGenesis)

	registeredReactions := randomRegisteredReactions(simState.Rand, subspacesGenesis.Subspaces)
	subspacesDataEntries := getSubspacesData(subspacesGenesis.Subspaces, registeredReactions)
	reactions := randomReactions(simState.Rand, simState.Accounts, subspacesGenesis.Subspaces, postsGenesis.Posts, registeredReactions)
	postsDataEntries := getPostsData(postsGenesis.Posts, reactions)
	params := randomReactionsParams(simState.Rand, subspacesGenesis.Subspaces)

	// Create the genesis and sanitize it
	reactionsGenesis := types.NewGenesisState(
		subspacesDataEntries,
		registeredReactions,
		postsDataEntries,
		reactions,
		params,
	)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(reactionsGenesis)
}

// randomRegisteredReactions returns a slice of randomly generated registered reactions
func randomRegisteredReactions(r *rand.Rand, subspaces []subspacestypes.Subspace) []types.RegisteredReaction {
	if len(subspaces) == 0 {
		return nil
	}

	reactionsNumber := r.Intn(50)
	reactions := make([]types.RegisteredReaction, reactionsNumber)
	for i := 0; i < reactionsNumber; i++ {
		subspace := subspacessim.RandomSubspace(r, subspaces)
		reactions[i] = types.NewRegisteredReaction(
			subspace.ID,
			uint32(i+1),
			GenerateRandomShorthandCode(r),
			GenerateRandomDisplayValue(r),
		)
	}
	return reactions
}

// getSubspacesData returns the SubspacesDataEntry slice based on the given data
func getSubspacesData(subspaces []subspacestypes.Subspace, reactions []types.RegisteredReaction) []types.SubspaceDataEntry {
	if len(subspaces) == 0 {
		return nil
	}

	entries := make([]types.SubspaceDataEntry, len(subspaces))
	for i, subspace := range subspaces {
		// Get the max reaction id
		maxReactionID := uint32(0)
		for _, reaction := range reactions {
			if reaction.SubspaceID == subspace.ID && reaction.ID > maxReactionID {
				maxReactionID = reaction.ID
			}
		}

		// Generate the entry
		entries[i] = types.NewSubspaceDataEntry(
			subspace.ID,
			maxReactionID+1,
		)
	}
	return entries
}

// randomReactions returns a slice of randomly generated reactions
func randomReactions(r *rand.Rand, accounts []simtypes.Account, subspaces []subspacestypes.Subspace, posts []poststypes.Post, registeredReactions []types.RegisteredReaction) (reactions []types.Reaction) {
	if len(subspaces) == 0 || len(posts) == 0 {
		return nil
	}

	reactionsNumber := r.Intn(100)
	for i := 0; i < reactionsNumber; i++ {
		subspace := subspacessim.RandomSubspace(r, subspaces)
		subspacePosts := getSubspacePosts(subspace.ID, posts)
		if len(subspacePosts) == 0 {
			// No posts inside the subspace
			continue
		}
		post := postssim.RandomPost(r, subspacePosts)

		var value types.ReactionValue
		if r.Intn(101) < 50 {
			// 50% chance of a registered reaction value
			subspaceReactions := getSubspaceRegisteredReactions(subspace.ID, registeredReactions)
			if len(subspaceReactions) == 0 {
				// No registered reactions inside the subspace
				continue
			}
			registeredReaction := RandomRegisteredReaction(r, subspaceReactions)
			value = types.NewRegisteredReactionValue(registeredReaction.ID)
		} else {
			// 50% chance of a free text value
			value = types.NewFreeTextValue(GetRandomFreeTextValue(r, 2))
		}

		author, _ := simtypes.RandomAcc(r, accounts)

		reactions = append(reactions, types.NewReaction(
			subspace.ID,
			post.ID,
			uint32(i+1),
			value,
			author.Address.String(),
		))
	}
	return reactions
}

// getSubspacePosts returns all the posts for the given subspace
func getSubspacePosts(subspaceID uint64, posts []poststypes.Post) (subspacePosts []poststypes.Post) {
	for _, post := range posts {
		if post.SubspaceID == subspaceID {
			subspacePosts = append(subspacePosts, post)
		}
	}
	return subspacePosts
}

// getSubspaceRegisteredReactions returns all the registered reactions for the given subspace
func getSubspaceRegisteredReactions(subspaceID uint64, reactions []types.RegisteredReaction) (subspaceReactions []types.RegisteredReaction) {
	for _, reaction := range reactions {
		if reaction.SubspaceID == subspaceID {
			subspaceReactions = append(subspaceReactions, reaction)
		}
	}
	return subspaceReactions
}

// getPostsData uses the given posts and reactions to return a PostDataEntry slice
func getPostsData(posts []poststypes.Post, reactions []types.Reaction) (entries []types.PostDataEntry) {
	if len(posts) == 0 {
		return nil
	}

	type postReference struct {
		SubspaceID uint64
		PostID     uint64
	}

	// Get the max attachment id for each post that has an attachment
	maxReactionIDs := map[postReference]uint32{}
	for _, reaction := range reactions {
		key := postReference{SubspaceID: reaction.SubspaceID, PostID: reaction.PostID}
		maxReactionID, ok := maxReactionIDs[key]
		if !ok || maxReactionID < reaction.ID {
			maxReactionIDs[key] = reaction.ID
		}
	}

	entries = make([]types.PostDataEntry, len(posts))
	for i, post := range posts {
		key := postReference{SubspaceID: post.SubspaceID, PostID: post.ID}
		maxReactionID, ok := maxReactionIDs[key]
		if !ok {
			maxReactionID = 0
		}
		entries[i] = types.NewPostDataEntry(post.SubspaceID, post.ID, maxReactionID+1)
	}
	return entries
}

// randomReactionsParams generates a slice of random params for the given subspaces
func randomReactionsParams(r *rand.Rand, subspaces []subspacestypes.Subspace) (params []types.SubspaceReactionsParams) {
	params = make([]types.SubspaceReactionsParams, len(subspaces))
	for i, subspace := range subspaces {
		params[i] = GenerateRandomSubspaceReactionsParams(r, subspace.ID)
	}
	return params
}
