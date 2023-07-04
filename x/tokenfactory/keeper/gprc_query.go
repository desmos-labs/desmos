package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	subspacestypes "github.com/desmos-labs/desmos/v5/x/subspaces/types"
	"github.com/desmos-labs/desmos/v5/x/tokenfactory/types"
)

var _ types.QueryServer = Keeper{}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(ctx context.Context, request *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)

	return &types.QueryParamsResponse{Params: params}, nil
}

// SubspaceDenoms implements the Query/SubspaceDenoms gRPC method
func (k Keeper) SubspaceDenoms(ctx context.Context, request *types.QuerySubspaceDenomsRequest) (*types.QuerySubspaceDenomsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	treasury := subspacestypes.GetTreasuryAddress(request.SubspaceId)
	denoms := k.GetDenomsFromCreator(sdkCtx, treasury.String())

	return &types.QuerySubspaceDenomsResponse{Denoms: denoms}, nil
}
