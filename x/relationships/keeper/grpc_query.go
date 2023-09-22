package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

var _ types.QueryServer = Keeper{}

// Relationships implements the Query/Relationships gRPC method
func (k Keeper) Relationships(ctx context.Context, request *types.QueryRelationshipsRequest) (*types.QueryRelationshipsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var relationships []types.Relationship

	// Get user relationships prefix store
	store := sdkCtx.KVStore(k.storeKey)

	// If the user and counterparty are provided, get the specific relationship
	if request.User != "" && request.Counterparty != "" {
		relationship, found := k.GetRelationship(sdkCtx, request.User, request.Counterparty, request.SubspaceId)
		if found {
			relationships = append(relationships, relationship)
		}
		return &types.QueryRelationshipsResponse{Relationships: relationships, Pagination: nil}, nil
	}

	// Get the correct store prefix to be used
	storePrefix := types.SubspaceRelationshipsPrefix(request.SubspaceId)
	if request.User != "" {
		storePrefix = types.UserRelationshipsSubspacePrefix(request.SubspaceId, request.User)
	}

	// Get paginated user relationships
	relationshipsStore := prefix.NewStore(store, storePrefix)
	pageRes, err := query.Paginate(relationshipsStore, request.Pagination, func(key []byte, value []byte) error {
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
	var userBlocks []types.UserBlock

	// Get user blocks prefix store
	store := sdkCtx.KVStore(k.storeKey)

	// If the blocker and blocked are provided, get the specific block
	if request.Blocker != "" && request.Blocked != "" {
		userBlock, found := k.GetUserBlock(sdkCtx, request.Blocker, request.Blocked, request.SubspaceId)
		if found {
			userBlocks = append(userBlocks, userBlock)
		}
		return &types.QueryBlocksResponse{Blocks: userBlocks, Pagination: nil}, nil
	}

	// Get the correct store prefix to be used
	storePrefix := types.SubspaceBlocksPrefix(request.SubspaceId)
	if request.Blocker != "" {
		storePrefix = types.BlockerSubspacePrefix(request.SubspaceId, request.Blocker)
	}

	// Get paginated user blocks
	userBlocksStore := prefix.NewStore(store, storePrefix)
	pageRes, err := query.Paginate(userBlocksStore, request.Pagination, func(key []byte, value []byte) error {
		var userBlock types.UserBlock
		if err := k.cdc.Unmarshal(value, &userBlock); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		userBlocks = append(userBlocks, userBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBlocksResponse{Blocks: userBlocks, Pagination: pageRes}, nil
}
