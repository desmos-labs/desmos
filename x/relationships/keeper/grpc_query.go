package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/relationships/types"
)

var _ types.QueryServer = Keeper{}

// Relationships implements the Query/Session gRPC method
func (k Keeper) Relationships(ctx context.Context, _ *types.QueryRelationshipsRequest) (*types.QueryRelationshipsResult, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	relationships, err := k.GetAllRelationships(sdkCtx)
	if err != nil {
		return nil, err
	}

	return &types.QueryRelationshipsResult{Relationships: relationships}, nil
}

// UserRelationships implements the Query/Session gRPC method
func (k Keeper) UserRelationships(ctx context.Context, request *types.QueryUserRelationshipsRequest) (*types.QueryRelationshipsResult, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	relationships, err := k.GetUserRelationships(sdkCtx, request.User)
	if err != nil {
		return nil, err
	}

	return &types.QueryRelationshipsResult{Relationships: relationships}, nil
}

// UserBlocks implements the Query/Session gRPC method
func (k Keeper) UserBlocks(ctx context.Context, request *types.QueryUserBlocksRequest) (*types.QueryBlocksResult, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blocks, err := k.GetUserBlocks(sdkCtx, request.User)
	if err != nil {
		return nil, err
	}

	return &types.QueryBlocksResult{Blocks: blocks}, nil
}
