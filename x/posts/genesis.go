package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts:               k.GetPosts(ctx),
		PollAnswers:         k.GetPollAnswersMap(ctx),
		PostReactions:       k.GetReactions(ctx),
		RegisteredReactions: k.ListReactions(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	// Sort the posts so that they are inserted based on their IDs
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	for postID, usersAnswersDetails := range data.PollAnswers {
		for _, userAnswersDetails := range usersAnswersDetails {
			parsedID, err := ParsePostID(postID)
			if err != nil {
				panic(err)
			}
			keeper.SavePollAnswers(ctx, parsedID, userAnswersDetails)
		}
	}

	for _, reaction := range data.RegisteredReactions {
		if _, found := keeper.DoesReactionForShortCodeExist(ctx, reaction.ShortCode, reaction.Subspace); !found {
			keeper.RegisterReaction(ctx, reaction)
		}
	}

	for postID, postReactions := range data.PostReactions {
		for _, postReaction := range postReactions {
			parsedID, err := ParsePostID(postID)
			if err != nil {
				panic(err)
			}
			if err := keeper.SavePostReaction(ctx, parsedID, postReaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
