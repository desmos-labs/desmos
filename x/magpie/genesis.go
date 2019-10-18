package magpie

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/internal/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Posts []Post `json:"posts"`
	Likes []Like `json:"likes"`
}

func NewGenesisState(posts []Post) GenesisState {
	return GenesisState{Posts: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if record.ID == "" {
			return fmt.Errorf("Invalid Post: ID: %s. Error: Missing ID", record.ID)
		}
		if record.Owner == nil {
			return fmt.Errorf("Invalid Post: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.Message == "" {
			return fmt.Errorf("Invalid Post: Message: %s. Error: Missing Message", record.Message)
		}
		if record.Created.String() == "" {
			return fmt.Errorf("Invalid Post: Created: %s. Error: Missing Created Time", record.Created)
		}

		if record.Modified.String() == "" {
			return fmt.Errorf("Invalid Post: Modified: %s. Error: Missing Modified Time", record.Modified)
		}

		if record.Namespace == "" {
			return fmt.Errorf("Invalid Post: Namespace: %s. Error: Missing Namespace", record.Namespace)
		}

		if record.ExternalOwner == "" {
			return fmt.Errorf("Invalid Post: ExternalOwner: %s. Error: Missing ExternalOwner", record.ExternalOwner)
		}
	}

	for _, record := range data.Likes {
		if record.Owner == nil {
			return fmt.Errorf("Invalid Like: Owner: %s. Error: Missing Owner", record.Owner)
		}
		if record.ID == "" {
			return fmt.Errorf("Invalid Like: ID: %s. Error: Missing ID", record.ID)
		}
		if record.Created.String() == "" {
			return fmt.Errorf("Invalid Like: Created: %s. Error: Missing Created Time", record.Created)
		}
		if record.PostID == "" {
			return fmt.Errorf("Invalid Like: PostID: %s. Error: Missing Post ID", record.PostID)
		}

	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Posts: []Post{},
	}
}

func InitGenesis(ctx sdk.Context, keeper keeper.Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Posts {
		keeper.SetPost(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k keeper.Keeper) GenesisState {
	var posts []Post
	iterator := k.GetPostsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var post Post
		post = k.GetPost(ctx, id)
		posts = append(posts, post)
	}

	// var likes []Like
	// iterator := k.Get(ctx)
	// for ; iterator.Valid(); iterator.Next() {
	// 	id := string(iterator.Key())
	// 	var post Post
	// 	post = k.GetPost(ctx, id)
	// 	posts = append(posts, post)
	// }
	return GenesisState{Posts: posts}
}
