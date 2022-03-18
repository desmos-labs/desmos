package keeper

import (
	"bytes"
	"context"

	"github.com/desmos-labs/desmos/v3/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

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

// Subspace implements the Query/Subspace gRPC method
func (k Keeper) Subspace(ctx context.Context, request *types.QuerySubspaceRequest) (*types.QuerySubspaceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "subspace with id %d not found", request.SubspaceId)
	}

	return &types.QuerySubspaceResponse{Subspace: subspace}, nil
}

// UserGroups implements the Query/UserGroups gRPC method
func (k Keeper) UserGroups(ctx context.Context, request *types.QueryUserGroupsRequest) (*types.QueryUserGroupsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.GroupsStoreKey(request.SubspaceId)
	groupsStore := prefix.NewStore(store, storePrefix)

	var groups []types.UserGroup
	pageRes, err := query.Paginate(groupsStore, request.Pagination, func(key []byte, value []byte) error {
		var group types.UserGroup
		err := k.cdc.Unmarshal(value, &group)
		if err != nil {
			return err
		}

		groups = append(groups, group)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserGroupsResponse{Groups: groups, Pagination: pageRes}, nil
}

// UserGroup implements the Query/UserGroup gRPC method
func (k Keeper) UserGroup(ctx context.Context, request *types.QueryUserGroupRequest) (*types.QueryUserGroupResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	// Get the group
	group, found := k.GetUserGroup(sdkCtx, request.SubspaceId, request.GroupId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", request.GroupId)
	}

	return &types.QueryUserGroupResponse{Group: group}, nil
}

// UserGroupMembers implements the Query/UserGroupMembers gRPC method
func (k Keeper) UserGroupMembers(ctx context.Context, request *types.QueryUserGroupMembersRequest) (*types.QueryUserGroupMembersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	// Check if the group exists
	if !k.HasUserGroup(sdkCtx, request.SubspaceId, request.GroupId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "group %d could not be found", request.GroupId)
	}

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.GroupMembersStoreKey(request.SubspaceId, request.GroupId)
	membersStore := prefix.NewStore(store, storePrefix)

	var members []string
	pageRes, err := query.Paginate(membersStore, request.Pagination, func(key []byte, value []byte) error {
		member := types.GetAddressFromBytes(bytes.TrimPrefix(key, storePrefix))
		members = append(members, member.String())
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserGroupMembersResponse{Members: members, Pagination: pageRes}, nil
}

// UserPermissions implements the Query/UserPermissions gRPC method
func (k Keeper) UserPermissions(ctx context.Context, request *types.QueryUserPermissionsRequest) (*types.QueryUserPermissionsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	sdkAddr, err := sdk.AccAddressFromBech32(request.User)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid user address: %s", request.User)
	}

	// Get the user specific permissions
	userPermission := k.GetUserPermissions(sdkCtx, request.SubspaceId, sdkAddr)
	groupPermissions := k.GetGroupsInheritedPermissions(sdkCtx, request.SubspaceId, sdkAddr)
	permissionResult := types.CombinePermissions(userPermission, groupPermissions)

	// Get the details of all the permissions
	var details []types.PermissionDetail
	if userPermission != types.PermissionNothing {
		details = append(details, types.NewPermissionDetailUser(request.User, userPermission))
	}

	k.IterateSubspaceGroups(sdkCtx, request.SubspaceId, func(index int64, group types.UserGroup) (stop bool) {
		if k.IsMemberOfGroup(sdkCtx, request.SubspaceId, group.ID, sdkAddr) {
			details = append(details, types.NewPermissionDetailGroup(group.ID, group.Permissions))
		}
		return false
	})

	return &types.QueryUserPermissionsResponse{
		Permissions: permissionResult,
		Details:     details,
	}, nil
}
