package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/staging/subspaces/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Subspaces(ctx context.Context, _ *types.QuerySubspacesRequest) (*types.QuerySubspacesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspaces := k.GetAllSubspaces(sdkCtx)
	return &types.QuerySubspacesResponse{Subspaces: subspaces}, nil
}

func (k Keeper) SubspaceAdmins(ctx context.Context, request *types.QuerySubspaceAdminsRequest) (*types.QuerySubspaceAdminsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	subspaceAdmins := k.GetAllSubspaceAdmins(sdkCtx, request.SubspaceId)
	return &types.QuerySubspaceAdminsResponse{Admins: subspaceAdmins}, nil
}

func (k Keeper) SubspaceBlockedUsers(ctx context.Context, request *types.QuerySubspaceBlockedUsersRequest) (*types.QuerySubspaceBlockedUsersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blockedUsers := k.GetSubspaceBlockedUsers(sdkCtx, request.SubspaceId)
	return &types.QuerySubspaceBlockedUsersResponse{Users: blockedUsers}, nil
}
