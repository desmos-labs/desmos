package keeper

import (
	"bytes"
	"context"

	"github.com/desmos-labs/desmos/v2/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

// Subspace implements the Query/Subspace gRPC method
func (k Keeper) Subspace(ctx context.Context, request *types.QuerySubspaceRequest) (*types.QuerySubspaceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "subspace with id %d not found", request.SubspaceId)
	}

	return &types.QuerySubspaceResponse{Subspace: subspace}, nil
}

// Subspaces implements the Query/Subspaces gRPC method
func (k Keeper) Subspaces(ctx context.Context, request *types.QuerySubspacesRequest) (*types.QuerySubspacesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspacePrefix)

	var subspaces []types.Subspace
	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(key []byte, value []byte) error {
		var subspace types.Subspace
		if err := k.cdc.Unmarshal(value, &subspace); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		subspaces = append(subspaces, subspace)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspacesResponse{Subspaces: subspaces, Pagination: pageRes}, nil
}

// UserGroups implements the Query/UserGroups gRPC method
func (k Keeper) UserGroups(ctx context.Context, request *types.QueryUserGroupsRequest) (*types.QueryUserGroupsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.GroupsStoreKey(request.SubspaceId)
	groupsStore := prefix.NewStore(store, storePrefix)

	var groups []string
	pageRes, err := query.Paginate(groupsStore, request.Pagination, func(key []byte, value []byte) error {
		groupName := types.GetGroupNameFromBytes(bytes.TrimPrefix(key, storePrefix))
		groups = append(groups, groupName)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserGroupsResponse{Groups: groups, Pagination: pageRes}, nil
}

// UserGroupMembers implements the Query/UserGroupMembers gRPC method
func (k Keeper) UserGroupMembers(ctx context.Context, request *types.QueryUserGroupMembersRequest) (*types.QueryUserGroupMembersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.GroupMembersStoreKey(request.SubspaceId, request.GroupName)
	membersStore := prefix.NewStore(store, storePrefix)

	var members []string
	pageRes, err := query.Paginate(membersStore, request.Pagination, func(key []byte, value []byte) error {
		member := types.GetGroupMemberFromBytes(bytes.TrimPrefix(key, storePrefix))
		members = append(members, member)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserGroupMembersResponse{Members: members, Pagination: pageRes}, nil
}

// Permissions implements the Query/Permissions gRPC method
func (k Keeper) Permissions(ctx context.Context, request *types.QueryPermissionsRequest) (*types.QueryPermissionsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	permission := k.GetPermissions(sdkCtx, request.SubspaceId, request.Target)
	return &types.QueryPermissionsResponse{Permissions: permission}, nil
}
