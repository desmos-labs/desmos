package wasm

import (
	"encoding/json"

	wasmTypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
)

var _ Querier = PostsWasmQuerier{}

type PostsWasmQuerier struct {
	postsKeeper postskeeper.Keeper
}

func NewPostsWasmQuerier(postsKeeper postskeeper.Keeper) PostsWasmQuerier {
	return PostsWasmQuerier{postsKeeper: postsKeeper}
}

func (PostsWasmQuerier) Query(_ sdk.Context, _ wasmTypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

// nolint
func (querier PostsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var desmosQuery PostsModuleQuery
	err := json.Unmarshal(data, &desmosQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var bz []byte

	if desmosQuery.Posts != nil {
		posts := querier.postsKeeper.GetPosts(ctx)
		bz, err = json.Marshal(PostsResponse{Posts: convertPosts(posts)})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	} else if desmosQuery.Reports != nil {
		reports := querier.postsKeeper.GetPostReports(ctx, desmosQuery.Reports.PostID)
		bz, err = json.Marshal(ReportsResponse{Reports: convertReports(reports)})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	} else { // Possible future queries before this
		return nil, sdkerrors.ErrInvalidRequest
	}

	return bz, nil
}
