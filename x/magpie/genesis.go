package magpie

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/internal/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Posts    []Post    `json:"posts"`
	Likes    []Like    `json:"likes"`
	Sessions []Session `json:"sessions"`
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ExportGenesis returns the GenesisState associated with the given context
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {
	return GenesisState{
		Posts:    k.GetPosts(ctx),
		Likes:    k.GetLikes(ctx),
		Sessions: k.GetSessions(ctx),
	}
}

// InitGenesis initializes the chain state based on the given GenesisState
// noinspection GoUnhandledErrorResult
func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, post := range data.Posts {
		keeper.SavePost(ctx, post)
	}

	for _, like := range data.Likes {
		keeper.SaveLike(ctx, like)
	}

	for _, session := range data.Sessions {
		keeper.SaveSession(ctx, session)
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

	for _, record := range data.Likes {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	return nil
}
