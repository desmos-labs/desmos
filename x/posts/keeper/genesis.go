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
		k.GetAllUserAnswers(ctx),
		k.GetAllPostReactions(ctx),
		k.GetRegisteredReactions(ctx),
		k.GetAllReports(ctx),
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

	// Save user answers
	for _, answer := range data.UsersPollAnswers {
		if !types.IsValidPostID(answer.PostID) {
			panic(fmt.Errorf("invalid poll answer post id: %s", answer.PostID))
		}
		k.SaveUserAnswer(ctx, answer)
	}

	// Save post reactions
	for _, reaction := range data.PostsReactions {
		if !types.IsValidPostID(reaction.PostID) {
			fmt.Println("here")
			panic(fmt.Errorf("invalid post id: %s", reaction.PostID))
		}
		err := k.SavePostReaction(ctx, reaction)
		if err != nil {
			panic(err)
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

	// Save posts reports
	for _, report := range data.Reports {
		err := k.SaveReport(ctx, report)
		if err != nil {
			panic(err)
		}
	}
}
