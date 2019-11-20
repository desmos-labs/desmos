package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/posts/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the magpie Querier
const (
	QueryPost = "post"
)

// Params for queries:
// - 'custom/magpie/post'

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryPost:
			return queryPost(ctx, path[1:], req, keeper)
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

	// Get the children
	childrenIDs := keeper.GetPostChildrenIDs(ctx, post.PostID)
	postChildren := make(types.Posts, len(childrenIDs))
	for index, childrenID := range childrenIDs {
		children, _ := keeper.GetPost(ctx, childrenID)
		postChildren[index] = children
	}

	// Crete the response object
	postResponse := NewPostResponse(post, postLikes, postChildren)

	bz, err2 := codec.MarshalJSONIndent(keeper.Cdc, &postResponse)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}
