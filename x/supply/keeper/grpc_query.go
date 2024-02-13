package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v7/x/supply/types"
)

var _ types.QueryServer = Keeper{}

// Total implements the Query/Total gRPC method
func (k Keeper) Total(ctx context.Context, request *types.QueryTotalRequest) (*types.QueryTotalResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	totalSupply := k.GetTotalSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return &types.QueryTotalResponse{TotalSupply: totalSupply}, nil
}

// Circulating implements the Query/Circulating gRPC method
func (k Keeper) Circulating(ctx context.Context, request *types.QueryCirculatingRequest) (*types.QueryCirculatingResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	circulatingSupply := k.GetCirculatingSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return &types.QueryCirculatingResponse{CirculatingSupply: circulatingSupply}, nil
}
