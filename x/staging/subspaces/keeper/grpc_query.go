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

	pageRes, err := query.FilteredPaginate(subspacesStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var subspace types.Subspace
		if err := k.cdc.UnmarshalBinaryBare(value, &subspace); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		if accumulate {
			subspaces = append(subspaces, subspace)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspacesResponse{Subspaces: subspaces, Pagination: pageRes}, nil
}

func (k Keeper) Admins(goCtx context.Context, request *types.QuerySubspaceAdminsRequest) (*types.QuerySubspaceAdminsResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var admins []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceAdminsPrefix(request.SubspaceId))

	pageRes, err := query.FilteredPaginate(subspacesStore, request.Pagination, func(_ []byte, value []byte, accumulate bool) (bool, error) {
		admin := string(value)
		if accumulate {
			admins = append(admins, admin)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspaceAdminsResponse{Admins: admins, Pagination: pageRes}, nil
}

func (k Keeper) RegisteredUsers(goCtx context.Context, request *types.QuerySubspaceRegisteredUsersRequest) (*types.QuerySubspaceRegisteredUsersResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceRegisteredUsersPrefix(request.SubspaceId))

	pageRes, err := query.FilteredPaginate(subspacesStore, request.Pagination, func(_ []byte, value []byte, accumulate bool) (bool, error) {
		user := string(value)
		if accumulate {
			users = append(users, user)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspaceRegisteredUsersResponse{Users: users, Pagination: pageRes}, nil
}

func (k Keeper) BannedUsers(goCtx context.Context, request *types.QuerySubspaceBannedUsersRequest) (*types.QuerySubspaceBannedUsersResponse, error) {
	if !types.IsValidSubspace(request.SubspaceId) {
		return nil, sdkerrors.Wrap(types.ErrInvalidSubspaceID, request.SubspaceId)
	}

	var users []string
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	subspacesStore := prefix.NewStore(store, types.SubspaceBannedUsersPrefix(request.SubspaceId))

	pageRes, err := query.FilteredPaginate(subspacesStore, request.Pagination, func(_ []byte, value []byte, accumulate bool) (bool, error) {
		user := string(value)
		if accumulate {
			users = append(users, user)
		}

		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QuerySubspaceBannedUsersResponse{Users: users, Pagination: pageRes}, nil
}
