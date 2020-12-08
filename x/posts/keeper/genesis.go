package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetPosts(ctx),
		k.GetUserAnswersEntries(ctx),
		k.GetPostReactionsEntries(ctx),
		k.GetRegisteredReactions(ctx),
		k.GetParams(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) {
	k.SetParams(ctx, data.Params)

	// Save posts
	for _, post := range data.Posts {
		err := k.ValidatePost(ctx, post)
		if err != nil {
			panic(err)
		}

		k.SavePost(ctx, post)
	}

	// Save poll answers
	for _, entry := range data.UsersPollAnswers {
		if !types.IsValidPostID(entry.PostId) {
			panic(fmt.Errorf("invalid postID: %s", entry.PostId))
		}

		for _, answer := range entry.UserAnswers {
			k.SavePollAnswers(ctx, entry.PostId, answer)
		}
	}

	// Save post reactions
	for _, entry := range data.PostsReactions {
		if !types.IsValidPostID(entry.PostId) {
			panic(fmt.Errorf("invalid post id: %s", entry.PostId))
		}

		for _, reaction := range entry.Reactions {
			err := k.SavePostReaction(ctx, entry.PostId, reaction)
			if err != nil {
				panic(err)
			}
		}
	}

	// Save registered reactions
	for _, reaction := range data.RegisteredReactions {
		_, found := k.GetRegisteredReaction(ctx, reaction.ShortCode, reaction.Subspace)
		if found {
			panic(fmt.Errorf("registeredReactions with shortcode %s already existing", reaction.ShortCode))
		}

		k.SaveRegisteredReaction(ctx, reaction)
	}
}
