package dwitter

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type GenesisState struct {
	Posts []Post `json:"posts"`
}

func NewGenesisState(posts []Post) GenesisState {
	return GenesisState{Posts: nil}
}

func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if record.Owner == nil {
			return fmt.Errorf("Invalid Posts: Value: %s. Error: Missing Owner", record.Owner)
		}
		if record.Message == "" {
			return fmt.Errorf("Invalid WhoisRecord: Owner: %s. Error: Missing Value", record.Message)
		}
		if record.Time.String() == "" {
			return fmt.Errorf("Invalid WhoisRecord: Value: %s. Error: Missing Price", record.Time)
		}
	}
	return nil
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Posts: []Post{},
	}
}

func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	for _, record := range data.Posts {
		keeper.SetPost(ctx, record)
	}
	return []abci.ValidatorUpdate{}
}

func ExportGenesis(ctx sdk.Context, k Keeper) GenesisState {
	var records []Post
	iterator := k.GetPostsIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		id := string(iterator.Key())
		var whois Post
		whois = k.GetPost(ctx, id)
		records = append(records, whois)
	}
	return GenesisState{Posts: records}
}
