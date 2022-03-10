package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/v2/x/coingecko/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) CirculatingSupply(ctx context.Context, request *types.QueryCirculatingSupplyRequest) (*types.QueryCirculatingSupplyResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	circulatingSupply := k.CalculateCirculatingSupply(sdkCtx, request.Denom)
	return &types.QueryCirculatingSupplyResponse{CirculatingSupply: circulatingSupply}, nil
}
