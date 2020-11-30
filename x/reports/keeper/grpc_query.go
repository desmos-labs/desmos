package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/reports/types"
)

var _ types.QueryServer = Keeper{}

// PostReports implements the Query/PostReports gRPC method
func (k Keeper) PostReports(
	ctx context.Context, request *types.QueryPostReportsRequest,
) (*types.QueryPostReportsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reports := k.GetPostReports(sdkCtx, request.PostId)
	return &types.QueryPostReportsResponse{Reports: reports}, nil
}
