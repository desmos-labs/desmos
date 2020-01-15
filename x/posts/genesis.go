package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func convertReactionsMap(reactions map[PostID]Reactions) map[string]Reactions {
	reactionsMap := make(map[string]Reactions, len(reactions))
	for key, value := range reactions {
		reactionsMap[key.String()] = value
	}
	return reactionsMap
}

func convertGenesisReactions(reactions map[string]Reactions) map[PostID]Reactions {
	reactionsMap := make(map[PostID]Reactions, len(reactions))
	for key, value := range reactions {
		postID, err := ParsePostID(key)
		if err != nil {
			panic(err)
		}
		reactionsMap[postID] = value
	}
	return reactionsMap
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts:     k.GetPosts(ctx),
		Reactions: convertReactionsMap(k.GetReactions(ctx)),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	reactionsMap := convertGenesisReactions(data.Reactions)
	for postID, reactions := range reactionsMap {
		for _, reaction := range reactions {
			if err := keeper.SaveReaction(ctx, postID, reaction); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
