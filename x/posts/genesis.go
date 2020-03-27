package posts

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func convertReactionsMap(reactions map[PostID]PostReactions) map[string]PostReactions {
	reactionsMap := make(map[string]PostReactions, len(reactions))
	for key, value := range reactions {
		reactionsMap[key.String()] = value
	}
	return reactionsMap
}

func convertGenesisReactions(reactions map[string]PostReactions) map[PostID]PostReactions {
	reactionsMap := make(map[PostID]PostReactions, len(reactions))
	for key, value := range reactions {
		postID, err := ParsePostID(key)
		if err != nil {
			panic(err)
		}
		reactionsMap[postID] = value
	}
	return reactionsMap
}

func convertPostPollAnswersMap(answers map[PostID]UserAnswers) map[string]UserAnswers {
	answersMap := make(map[string]UserAnswers, len(answers))
	for key, value := range answers {
		answersMap[key.String()] = value
	}
	return answersMap
}

func convertGenesisPostPollAnswers(pollAnswers map[string]UserAnswers) map[PostID]UserAnswers {
	answersMap := make(map[PostID]UserAnswers, len(pollAnswers))
	for key, value := range pollAnswers {
		postID, err := ParsePostID(key)
		if err != nil {
			panic(err)
		}
		answersMap[postID] = value
	}
	return answersMap
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts:               k.GetPosts(ctx),
		PollAnswers:         convertPostPollAnswersMap(k.GetPollAnswersMap(ctx)),
		PostReactions:       convertReactionsMap(k.GetReactions(ctx)),
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

	pollAnswersMap := convertGenesisPostPollAnswers(data.PollAnswers)
	for postID, usersAnswersDetails := range pollAnswersMap {
		for _, userAnswersDetails := range usersAnswersDetails {
			keeper.SavePollAnswers(ctx, postID, userAnswersDetails)
		}
	}

	for _, reaction := range data.RegisteredReactions {
		if _, found := keeper.DoesReactionForShortCodeExist(ctx, reaction.ShortCode, reaction.Subspace); !found {
			keeper.RegisterReaction(ctx, reaction)
		}
	}

	postReactionsMap := convertGenesisReactions(data.PostReactions)
	for postID, postReactions := range postReactionsMap {
		for _, postReaction := range postReactions {
			if err := keeper.SavePostReaction(ctx, postID, postReaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
