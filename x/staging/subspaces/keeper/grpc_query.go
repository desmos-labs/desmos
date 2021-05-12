package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Subspace(ctx context.Context, request *types.QuerySubspaceRequest) (*types.QuerySubspaceResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspace, found := k.GetSubspace(sdkCtx, request.SubspaceId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "subspace with id %s not found", request.SubspaceId)
	}

	admins := k.GetAllSubspaceAdmins(sdkCtx, request.SubspaceId)

	return &types.QuerySubspaceResponse{Subspace: subspace, Admins: admins}, nil
}

func (k Keeper) SubspaceBlockedUsers(ctx context.Context, request *types.QuerySubspaceBlockedUsersRequest) (*types.QuerySubspaceBlockedUsersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockedUsers := k.GetSubspaceBlockedUsers(sdkCtx, request.SubspaceId)
	return &types.QuerySubspaceBlockedUsersResponse{Users: blockedUsers}, nil
}
