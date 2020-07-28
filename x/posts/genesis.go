package posts

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/keeper"
	"github.com/desmos-labs/desmos/x/posts/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) types.GenesisState {
	posts := k.GetPosts(ctx)
	sort.Sort(posts)
	return types.GenesisState{
		Posts:               posts,
		UsersPollAnswers:    k.GetPollAnswersMap(ctx),
		PostReactions:       k.GetReactions(ctx),
		RegisteredReactions: k.GetRegisteredReactions(ctx),
		Params:              k.GetParams(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, k keeper.Keeper, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	for _, post := range data.Posts {
		if err := keeper.ValidatePost(ctx, k, post); err != nil {
			panic(err)
		}
		k.SavePost(ctx, post)
	}

	for postID, usersAnswersDetails := range data.UsersPollAnswers {
		for _, userAnswersDetails := range usersAnswersDetails {
			postID := types.PostID(postID)
			if !postID.Valid() {
				panic(fmt.Errorf("invalid postID: %s", postID))
			}
			k.SavePollAnswers(ctx, postID, userAnswersDetails)
		}
	}

	for _, reaction := range data.RegisteredReactions {
		if _, found := k.GetRegisteredReaction(ctx, reaction.ShortCode, reaction.Subspace); !found {
			k.RegisterReaction(ctx, reaction)
		}
	}

	for postID, postReactions := range data.PostReactions {
		for _, postReaction := range postReactions {
			postID := types.PostID(postID)
			if !postID.Valid() {
				panic(fmt.Errorf("invalid postID: %s", postID))
			}
			if err := k.SavePostReaction(ctx, postID, postReaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
