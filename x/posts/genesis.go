package posts

import (
	"fmt"
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts:               k.GetPosts(ctx),
		UsersPollAnswers:    k.GetPollAnswersMap(ctx),
		PostReactions:       k.GetReactions(ctx),
		RegisteredReactions: k.ListReactions(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {

	// Sort the posts so that they are inserted based on their IDs
	sort.Sort(data.Posts)
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	for postID, usersAnswersDetails := range data.UsersPollAnswers {
		for _, userAnswersDetails := range usersAnswersDetails {
			postID := types.PostID(postID)
			if !postID.Valid() {
				panic(fmt.Sprintf("invalid parent ID: %s", postID))
			}
			keeper.SavePollAnswers(ctx, postID, userAnswersDetails)
		}
	}

	for _, reaction := range data.RegisteredReactions {
		if _, found := keeper.DoesReactionForShortCodeExist(ctx, reaction.ShortCode, reaction.Subspace); !found {
			keeper.RegisterReaction(ctx, reaction)
		}
	}

	for postID, postReactions := range data.PostReactions {
		for _, postReaction := range postReactions {
			postID := types.PostID(postID)
			if !postID.Valid() {
				panic(fmt.Sprintf("invalid parent ID: %s", postID))
			}
			if err := keeper.SavePostReaction(ctx, postID, postReaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
