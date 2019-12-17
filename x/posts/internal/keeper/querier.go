package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {

		case types.QueryPost:
			return queryPost(ctx, path[1:], req, keeper)

		case types.QueryPosts:
			return queryPosts(ctx, req, keeper)

		default:
			return nil, sdk.ErrUnknownRequest("unknown magpie query endpoint")
		}
	}
}

func getPostResponse(ctx sdk.Context, keeper Keeper, post types.Post) PostQueryResponse {
	// Get the likes
	postLikes := keeper.GetPostLikes(ctx, post.PostID)
	if postLikes == nil {
		postLikes = types.Likes{}
	}

	// Get the children
	childrenIDs := keeper.GetPostChildrenIDs(ctx, post.PostID)
	if childrenIDs == nil {
		childrenIDs = types.PostIDs{}
	}

	// Crete the response object
	return NewPostResponse(post, postLikes, childrenIDs)
}

func queryPost(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParsePostID(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid post id: %s", path[0]))
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Post with id %s not found", id))
	}

	postResponse := getPostResponse(ctx, keeper, post)
	bz, err2 := codec.MarshalJSONIndent(keeper.Cdc, &postResponse)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

func queryPosts(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	var params types.QueryPostsParams

	err := keeper.Cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("failed to parse params", err.Error()))
	}

	posts := keeper.GetPostsFiltered(ctx, params)

	postResponses := make([]PostQueryResponse, len(posts))
	for index, post := range posts {
		postResponses[index] = getPostResponse(ctx, keeper, post)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, &postResponses)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}

	return bz, nil
}
