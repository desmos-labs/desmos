package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
)

var _ types.QueryServer = Keeper{}

// TotalSupply implements the Query/TotalSupply gRPC method
func (k Keeper) TotalSupply(ctx context.Context, request *types.QueryTotalSupplyRequest) (*types.QueryTotalSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	totalSupply := k.GetConvertedTotalSupply(sdkCtx, request.Denom, sdk.NewInt(request.Divider))
	return &types.QueryTotalSupplyResponse{TotalSupply: totalSupply}, nil
}

// CirculatingSupply implements the Query/CirculatingSupply gRPC method
func (k Keeper) CirculatingSupply(ctx context.Context, request *types.QueryCirculatingSupplyRequest) (*types.QueryCirculatingSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	circulatingSupply := k.CalculateCirculatingSupply(sdkCtx, request.Denom, sdk.NewInt(request.Divider))
	return &types.QueryCirculatingSupplyResponse{CirculatingSupply: circulatingSupply}, nil
}
