package wasm

import (
	"encoding/json"

	wasmTypes "github.com/CosmWasm/wasmvm/types"
	"github.com/desmos-labs/desmos/wasm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
)

var _ wasm.Querier = PostsWasmQuerier{}

type PostsWasmQuerier struct {
	postsKeeper postskeeper.Keeper
}

func NewPostsWasmQuerier(postsKeeper postskeeper.Keeper) PostsWasmQuerier {
	return PostsWasmQuerier{postsKeeper: postsKeeper}
}

func (PostsWasmQuerier) Query(_ sdk.Context, _ wasmTypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier PostsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var desmosQuery PostsModuleQueryRoutes
	err := json.Unmarshal(data, &desmosQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var bz []byte

	switch {
	case desmosQuery.Posts != nil:
		posts := querier.postsKeeper.GetPosts(ctx)
		bz, err = json.Marshal(PostsResponse{Posts: convertPosts(posts)})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.Reports != nil:
		reports := querier.postsKeeper.GetPostReports(ctx, desmosQuery.Reports.PostID)
		bz, err = json.Marshal(ReportsResponse{Reports: convertReports(reports)})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.Reactions != nil:
		reactions := querier.postsKeeper.GetPostReactions(ctx, desmosQuery.Reactions.PostID)
		bz, err = json.Marshal(ReactionsResponse{Reactions: reactions})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}

	return bz, nil
}
