package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/v3/x/supply/types"
)

var _ types.QueryServer = Keeper{}

// Total implements the Query/Total gRPC method
func (k Keeper) Total(ctx context.Context, request *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	totalSupply := k.GetTotalSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return &types.QueryTotalSupplyResponse{TotalSupply: totalSupply}, nil
}

// Circulating implements the Query/Circulating gRPC method
func (k Keeper) Circulating(ctx context.Context, request *types.QueryCirculatingSupplyRequest) (*types.QueryCirculatingSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	circulatingSupply := k.GetCirculatingSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return &types.QueryCirculatingSupplyResponse{CirculatingSupply: circulatingSupply}, nil
}
