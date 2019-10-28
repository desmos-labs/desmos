package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kwunyeung/desmos/x/magpie/internal/types"
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
func queryPost(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParsePostId(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid post id: %s", path[0]))
	}

	post, found := keeper.GetPost(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest("could not get post")
	}

	bz, err2 := codec.MarshalJSONIndent(keeper.cdc, &post)
	if err2 != nil {
		panic("could not marshal result to JSON")
	}

	return bz, nil
}

// nolint: unparam
func queryLike(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParseLikeId(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid like id: %s", path[0]))
	}

	like, found := keeper.GetLike(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest("could not get like")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, &like)
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func querySession(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParseSessionId(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid session id: %s", path[0]))
	}

	session, found := keeper.GetSession(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest("could not get session")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, &session)

	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
