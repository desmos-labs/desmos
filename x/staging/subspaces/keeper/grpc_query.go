package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Subspace(ctx context.Context, request *types.QuerySubspaceRequest) (*types.QuerySubspaceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspaces with id %s not found", request.SubspaceId)
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
