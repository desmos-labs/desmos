package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/posts/types"
)

// ExportGenesis returns the GenesisState associated with the given context
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	return types.NewGenesisState(
		k.GetPosts(ctx),
		k.GetPollAnswersMap(ctx),
		k.GetReactions(ctx),
		k.GetRegisteredReactions(ctx),
		k.GetParams(ctx),
	)
}

// InitGenesis initializes the chain state based on the given GenesisState
func (k Keeper) InitGenesis(ctx sdk.Context, data types.GenesisState) []abci.ValidatorUpdate {
	k.SetParams(ctx, data.Params)

	for _, post := range data.Posts {
		if err := k.ValidatePost(ctx, k, post); err != nil {
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
			k.SaveRegisteredReaction(ctx, reaction)
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
