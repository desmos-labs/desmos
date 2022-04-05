package keeper

import (
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v3/x/supply/types"
)

var _ types.QueryServer = Keeper{}

// TotalSupply implements the Query/TotalSupply gRPC method
func (k Keeper) TotalSupply(ctx context.Context, request *types.QueryTotalSupplyRequest) (*wrapperspb.StringValue, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	totalSupply := k.GetConvertedTotalSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return wrapperspb.String(totalSupply.String()), nil
}

// CirculatingSupply implements the Query/CirculatingSupply gRPC method
func (k Keeper) CirculatingSupply(ctx context.Context, request *types.QueryCirculatingSupplyRequest) (*wrapperspb.StringValue, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	circulatingSupply := k.CalculateCirculatingSupply(sdkCtx, request.Denom, types.NewDividerPoweredByExponent(request.DividerExponent))
	return wrapperspb.String(circulatingSupply.String()), nil
}
