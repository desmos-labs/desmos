package magpie

import (
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the magpie Querier
const (
	QueryPost = "post"
	QueryLike = "like"
)

// Params for queries:
// - 'custom/magpie/post'
// - 'custom/magpie/like'

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryPost:
			return queryPost(ctx, path[1:], req, keeper)
		case QueryLike:
			return queryLike(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown magpie query endpoint")
		}
	}
}

// nolint: unparam
func queryPost(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id := path[0]

	post := keeper.GetPost(ctx, id)

	if post.String() == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not get post")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, QueryResPost{post.ID, post.Message, post.Owner, post.Time, post.Likes})
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// nolint: unparam
func queryLike(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	like := keeper.GetLike(ctx, path[0])

	if like.String() == "" {
		return []byte{}, sdk.ErrUnknownRequest("could not get like")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, QueryResLike{like.ID, like.PostID, like.Owner, like.Time})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

// func queryNames(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
// 	var namesList QueryResNames

// 	iterator := keeper.GetNamesIterator(ctx)

// 	for ; iterator.Valid(); iterator.Next() {
// 		namesList = append(namesList, string(iterator.Key()))
// 	}

// 	res, err := codec.MarshalJSONIndent(keeper.cdc, namesList)
// 	if err != nil {
// 		panic("could not marshal result to JSON")
// 	}

// 	return res, nil
// }
