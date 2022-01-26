package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v2/x/relationships/types"
)

var _ types.QueryServer = Keeper{}

// Relationships implements the Query/Relationships gRPC method
func (k Keeper) Relationships(ctx context.Context, request *types.QueryRelationshipsRequest) (*types.QueryRelationshipsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var relationships []types.Relationship

	// Get user relationships prefix store
	store := sdkCtx.KVStore(k.storeKey)

	storePrefix := types.UserRelationshipsPrefix(request.User)
	if request.User != "" {
		storePrefix = types.UserRelationshipsSubspacePrefix(request.User, request.SubspaceId)
	}
	relsStore := prefix.NewStore(store, storePrefix)

	// Get paginated user relationships
	pageRes, err := query.Paginate(relsStore, request.Pagination, func(key []byte, value []byte) error {
		var rel types.Relationship
		if err := k.cdc.Unmarshal(value, &rel); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		relationships = append(relationships, rel)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRelationshipsResponse{Relationships: relationships, Pagination: pageRes}, nil
}

// Blocks implements the Query/Blocks gRPC method
func (k Keeper) Blocks(ctx context.Context, request *types.QueryBlocksRequest) (*types.QueryBlocksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var userblocks []types.UserBlock

	// Get user blocks prefix store
	store := sdkCtx.KVStore(k.storeKey)

	storePrefix := types.BlockerPrefix(request.User)
	if request.User != "" {
		storePrefix = types.BlockerSubspacePrefix(request.User, request.SubspaceId)
	}
	userBlocksStore := prefix.NewStore(store, storePrefix)

	// Get paginated user blocks
	pageRes, err := query.Paginate(userBlocksStore, request.Pagination, func(key []byte, value []byte) error {
		var userBlock types.UserBlock
		if err := k.cdc.Unmarshal(value, &userBlock); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		userblocks = append(userblocks, userBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBlocksResponse{Blocks: userblocks, Pagination: pageRes}, nil
}
