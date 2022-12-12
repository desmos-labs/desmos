package keeper

import (
	"bytes"
	"context"

	"github.com/desmos-labs/desmos/v4/x/subspaces/types"

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

// Sections implements the Query/Sections gRPC method
func (k Keeper) Sections(ctx context.Context, request *types.QuerySectionsRequest) (*types.QuerySectionsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.SubspaceSectionsPrefix(request.SubspaceId)
	sectionsStore := prefix.NewStore(store, storePrefix)

	var sections []types.Section
	pageRes, err := query.Paginate(sectionsStore, request.Pagination, func(key []byte, value []byte) error {
		var section types.Section
		if err := k.cdc.Unmarshal(value, &section); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		sections = append(sections, section)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySectionsResponse{Sections: sections, Pagination: pageRes}, nil
}

// Section implements the Query/Section gRPC method
func (k Keeper) Section(ctx context.Context, request *types.QuerySectionRequest) (*types.QuerySectionResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	section, found := k.GetSection(sdkCtx, request.SubspaceId, request.SectionId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "section with id %d not found inside subspace %d", request.SectionId, request.SubspaceId)
	}

	return &types.QuerySectionResponse{Section: section}, nil
}

// UserGroups implements the Query/UserGroups gRPC method
func (k Keeper) UserGroups(ctx context.Context, request *types.QueryUserGroupsRequest) (*types.QueryUserGroupsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Check if the subspace exists
	if !k.HasSubspace(sdkCtx, request.SubspaceId) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %d not found", request.SubspaceId)
	}

	store := sdkCtx.KVStore(k.storeKey)
	storePrefix := types.SubspaceGroupsPrefix(request.SubspaceId)
	if request.SectionId != types.RootSectionID {
		storePrefix = types.SectionGroupsPrefix(request.SubspaceId, request.SectionId)
	}
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
	storePrefix := types.GroupMembersPrefix(request.SubspaceId, request.GroupId)
	membersStore := prefix.NewStore(store, storePrefix)

	var members []string
	pageRes, err := query.Paginate(membersStore, request.Pagination, func(key []byte, value []byte) error {
		member := types.GetAddressFromBytes(bytes.TrimPrefix(key, storePrefix))
		members = append(members, member)
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

	// Get the user specific permissions
	userPermissions := k.GetUserPermissions(sdkCtx, request.SubspaceId, request.SectionId, request.User)
	groupPermissions := k.GetGroupsInheritedPermissions(sdkCtx, request.SubspaceId, request.SectionId, request.User)
	permissionResult := types.CombinePermissions(append(userPermissions, groupPermissions...)...)

	// Get the details of all the permissions
	var details []types.PermissionDetail
	if userPermissions != nil {
		details = append(details, types.NewPermissionDetailUser(request.SubspaceId, request.SectionId, request.User, userPermissions))
	}

	k.IterateSubspaceUserGroups(sdkCtx, request.SubspaceId, func(group types.UserGroup) (stop bool) {
		if k.IsMemberOfGroup(sdkCtx, request.SubspaceId, group.ID, request.User) {
			details = append(details, types.NewPermissionDetailGroup(group.SubspaceID, group.SectionID, group.ID, group.Permissions))
		}
		return false
	})

	return &types.QueryUserPermissionsResponse{
		Permissions: permissionResult,
		Details:     details,
	}, nil
}

func (k Keeper) UserGrants(ctx context.Context, request *types.QueryUserGrantsRequest) (*types.QueryUserGrantsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// Get grants prefix store
	grantsPrefix := types.UserGrantPrefix
	switch {
	case request.SubspaceId != 0:
		grantsPrefix = types.SubspaceUserGrantPrefix(request.SubspaceId)
	case request.SubspaceId != 0 && request.Granter != "":
		grantsPrefix = types.GranterUserGrantPrefix(request.SubspaceId, request.Granter)
	case request.SubspaceId != 0 && request.Granter != "" && request.Grantee != "":
		grantsPrefix = types.UserGrantKey(request.SubspaceId, request.Granter, request.Grantee)
	}

	grantsStore := prefix.NewStore(store, grantsPrefix)
	var grants []types.UserGrant
	pageRes, err := query.FilteredPaginate(grantsStore, request.Pagination, func(key []byte, value []byte, acc bool) (bool, error) {
		var grant types.UserGrant
		if err := k.cdc.Unmarshal(value, &grant); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}
		if (request.SubspaceId != 0 && grant.Grantee == request.Grantee) || request.Grantee == "" {
			grants = append(grants, grant)
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserGrantsResponse{
		Grants:     grants,
		Pagination: pageRes,
	}, nil
}

func (k Keeper) GroupGrants(ctx context.Context, request *types.QueryGroupGrantsRequest) (*types.QueryGroupGrantsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// Get grants prefix store
	grantsPrefix := types.UserGrantPrefix
	switch {
	case request.SubspaceId != 0:
		grantsPrefix = types.SubspaceUserGrantPrefix(request.SubspaceId)
	case request.SubspaceId != 0 && request.Granter != "":
		grantsPrefix = types.GranterUserGrantPrefix(request.SubspaceId, request.Granter)
	case request.SubspaceId != 0 && request.Granter != "" && request.GroupId != 0:
		grantsPrefix = types.GroupGrantKey(request.SubspaceId, request.Granter, request.GroupId)
	}

	grantsStore := prefix.NewStore(store, grantsPrefix)
	var grants []types.GroupGrant
	pageRes, err := query.FilteredPaginate(grantsStore, request.Pagination, func(key []byte, value []byte, acc bool) (bool, error) {
		var grant types.GroupGrant
		if err := k.cdc.Unmarshal(value, &grant); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}
		if (request.SubspaceId != 0 && grant.GroupID == request.GroupId) || request.GroupId == 0 {
			grants = append(grants, grant)
			return true, nil
		}
		return false, nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryGroupGrantsResponse{
		Grants:     grants,
		Pagination: pageRes,
	}, nil
}
