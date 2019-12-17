package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func convertLikesMap(likes map[PostID]Likes) map[string]Likes {
	likesMap := make(map[string]Likes, len(likes))
	for key, value := range likes {
		likesMap[key.String()] = value
	}
	return likesMap
}

func convertGenesisLikes(likes map[string]Likes) map[PostID]Likes {
	likesMap := make(map[PostID]Likes, len(likes))
	for key, value := range likes {
		postID, err := ParsePostID(key)
		if err != nil {
			panic(err)
		}
		likesMap[postID] = value
	}
	return likesMap
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	return GenesisState{
		Posts: k.GetPosts(ctx),
		Likes: convertLikesMap(k.GetLikes(ctx)),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	likesMap := convertGenesisLikes(data.Likes)
	for postID, likes := range likesMap {
		for _, like := range likes {
			if err := keeper.SaveLike(ctx, postID, like); err != nil {
				panic(err)
			}
		}
	}

	for postID, likes := range data.Likes {
		for _, like := range likes {
			postID, err := ParsePostID(postID)
			if err != nil {
				panic(err)
			}
			if err := keeper.SaveLike(ctx, postID, like); err != nil {
				panic(err)
			}
		}
	}

	return []abci.ValidatorUpdate{}
}
