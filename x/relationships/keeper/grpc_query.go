package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/relationships/types"
)

var _ types.QueryServer = Keeper{}

// UserRelationships implements the Query/UserRelationships gRPC method
func (k Keeper) UserRelationships(
	ctx context.Context, request *types.QueryUserRelationshipsRequest,
) (*types.QueryUserRelationshipsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	relationships := k.GetUserRelationships(sdkCtx, request.User)
	return &types.QueryUserRelationshipsResponse{User: request.User, Relationships: relationships}, nil
}

// UserBlocks implements the Query/UserBlocks gRPC method
func (k Keeper) UserBlocks(
	ctx context.Context, request *types.QueryUserBlocksRequest,
) (*types.QueryUserBlocksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blocks := k.GetUserBlocks(sdkCtx, request.User)
	return &types.QueryUserBlocksResponse{Blocks: blocks}, nil
}
