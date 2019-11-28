package posts

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/keeper"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Posts Posts            `json:"posts"`
	Likes map[string]Likes `json:"likes"`
}

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
		postID, err := types.ParsePostID(key)
		if err != nil {
			panic(err)
		}
		likesMap[postID] = value
	}
	return likesMap
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {
	return GenesisState{
		Posts: k.GetPosts(ctx),
		Likes: convertLikesMap(k.GetLikes(ctx)),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) []abci.ValidatorUpdate {
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
			postID, err := types.ParsePostID(postID)
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

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, likes := range data.Likes {
		for _, record := range likes {
			if err := record.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
