package keeper

import (
	"context"

	"github.com/desmos-labs/desmos/x/posts/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Posts(ctx context.Context, request *types.QueryPostsRequest) (*types.QueryPostsResponse, error) {
	panic("implement me")
}

func (k Keeper) Post(ctx context.Context, request *types.QueryPostRequest) (*types.QueryPostResponse, error) {
	panic("implement me")
}

func (k Keeper) PollAnswers(ctx context.Context, request *types.QueryPollAnswersRequest) (*types.QueryPollAnswersResponse, error) {
	panic("implement me")
}

func (k Keeper) RegisteredReactions(ctx context.Context, request *types.QueryRegisteredReactionsRequest) (*types.QueryRegisteredReactionsResponse, error) {
	panic("implement me")
}

func (k Keeper) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	panic("implement me")
}
