package wasm

import (
	"encoding/json"
	wasmTypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	postskeeper "github.com/desmos-labs/desmos/x/posts/keeper"
)

type Querier interface {
	Query(ctx sdk.Context, request wasmTypes.QueryRequest) ([]byte, error)
	QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error)
}

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

func (querier PostsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var desmosQuery PostsModuleQuery
	err := json.Unmarshal(data, &desmosQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var bz []byte

	if desmosQuery.Posts != nil {
		posts := querier.postsKeeper.GetPosts(ctx)
		convertedPosts := make([]Post, len(posts))
		for index, post := range posts {
			convertedPosts[index] = convertPost(post)
		}
		//fmt.Println(PostsResponse{Posts: posts})
		bz, err = json.Marshal(PostsResponse{Posts: convertedPosts})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	} else {
		return nil, sdkerrors.ErrInvalidRequest
	}

	return bz, nil
}
