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

func queryPost(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParsePostID(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid post id: %s", path[0]))
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest("could not get post")
	}

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
	postResponse := NewPostResponse(post, postLikes, childrenIDs)

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

	bz, err := codec.MarshalJSONIndent(keeper.Cdc, keeper.GetPostsFiltered(ctx, params))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("failed to JSON marshal result: %s", err.Error()))
	}

	return bz, nil
}
