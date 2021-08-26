package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Subspace(ctx context.Context, request *types.QuerySubspaceRequest) (*types.QuerySubspaceResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "subspace with id %s not found", request.SubspaceId)
	}

	return &types.QuerySubspaceResponse{Subspace: subspace}, nil
}

func (k Keeper) Subspaces(goCtx context.Context, request *types.QuerySubspacesRequest) (*types.QuerySubspacesResponse, error) {
	var subspaces []types.Subspace
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceStorePrefix)

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

func (k Keeper) Admins(goCtx context.Context, request *types.QueryAdminsRequest) (*types.QueryAdminsResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var admins []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceAdminsPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		admin := string(value)
		admins = append(admins, admin)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAdminsResponse{Admins: admins, Pagination: pageRes}, nil
}

func (k Keeper) RegisteredUsers(goCtx context.Context, request *types.QueryRegisteredUsersRequest) (*types.QueryRegisteredUsersResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceRegisteredUsersPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		user := string(value)
		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRegisteredUsersResponse{Users: users, Pagination: pageRes}, nil
}

func (k Keeper) BannedUsers(goCtx context.Context, request *types.QueryBannedUsersRequest) (*types.QueryBannedUsersResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceBannedUsersPrefix(request.SubspaceId))

	pageRes, err := query.Paginate(subspacesStore, request.Pagination, func(_ []byte, value []byte) error {
		user := string(value)
		users = append(users, user)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBannedUsersResponse{Users: users, Pagination: pageRes}, nil
}
