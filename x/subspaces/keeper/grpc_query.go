package keeper

import (
	"context"

	types2 "github.com/desmos-labs/desmos/v2/x/subspaces/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types2.QueryServer = Keeper{}

func (k Keeper) Subspace(ctx context.Context, request *types2.QuerySubspaceRequest) (*types2.QuerySubspaceResponse, error) {
	if !types2.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidSubspaceID, request.SubspaceId)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "subspace with id %s not found", request.SubspaceId)
	}

	return &types2.QuerySubspaceResponse{Subspace: subspace}, nil
}

func (k Keeper) Subspaces(goCtx context.Context, request *types2.QuerySubspacesRequest) (*types2.QuerySubspacesResponse, error) {
	var subspaces []types2.Subspace
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types2.SubspaceStorePrefix)

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(key []byte, value []byte) error {
		var subspace types2.Subspace
		if err := k.cdc.Unmarshal(value, &subspace); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		subspaces = append(subspaces, subspace)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types2.QuerySubspacesResponse{Subspaces: subspaces, Pagination: pageRes}, nil
}

func (k Keeper) Admins(goCtx context.Context, request *types2.QueryAdminsRequest) (*types2.QueryAdminsResponse, error) {
	if !types2.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var admins []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types2.SubspaceAdminsPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		admin := string(value)
		admins = append(admins, admin)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types2.QueryAdminsResponse{Admins: admins, Pagination: pageRes}, nil
}

func (k Keeper) RegisteredUsers(goCtx context.Context, request *types2.QueryRegisteredUsersRequest) (*types2.QueryRegisteredUsersResponse, error) {
	if !types2.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types2.SubspaceRegisteredUsersPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		user := string(value)
		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types2.QueryRegisteredUsersResponse{Users: users, Pagination: pageRes}, nil
}

func (k Keeper) BannedUsers(goCtx context.Context, request *types2.QueryBannedUsersRequest) (*types2.QueryBannedUsersResponse, error) {
	if !types2.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types2.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types2.SubspaceBannedUsersPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		user := string(value)
		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types2.QueryBannedUsersResponse{Users: users, Pagination: pageRes}, nil
}
