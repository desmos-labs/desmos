package wasm

import (
	"encoding/json"
	cosmwasm "github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
	reportsKeeper "github.com/desmos-labs/desmos/x/reports/keeper"
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"
)

func DesmosQuerier(postsKeeper postskeeper.Keeper, reportsKeeper reportsKeeper.Keeper) cosmwasm.CustomQuerier {
	return func(ctx sdk.Context, request json.RawMessage) ([]byte, error) {
		var desmosQuery = DesmosQuery{
			Posts:   &PostsQuery{},
			Reports: &ReportsQuery{},
		}

		_ = json.Unmarshal(request, desmosQuery.Posts)
		_ = json.Unmarshal(request, desmosQuery.Reports)

		if desmosQuery.Posts != nil {
			posts := postsKeeper.GetPosts(ctx)
			convertedPosts := make([]Post, len(posts))
			for index, post := range posts {
				convertedPosts[index] = convertPost(post)
			}
			postsResponse := PostsResponse{Posts: convertedPosts}
			return json.Marshal(postsResponse)
		}

		if desmosQuery.Reports != nil {
			reports := reportsKeeper.GetPostReports(ctx, desmosQuery.Reports.PostID)
			reportsResponse := ReportsResponse{Reports: reports}
			return json.Marshal(reportsResponse)
		}
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unknown wasm request")
	}

}

type DesmosQuery struct {
	Posts   *PostsQuery   `json:"posts"`
	Reports *ReportsQuery `json:"reports"`
}

type ReportsQuery struct {
	PostID string `json:"post_id"`
}

type PostsQuery struct{}

type PostsResponse struct {
	Posts []Post `json:"posts"`
}

type ReportsResponse struct {
	Reports []reportsTypes.Report `json:"reports"`
}
