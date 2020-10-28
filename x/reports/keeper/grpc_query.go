package keeper

import (
	"context"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/reports/types"
)

var _ types.QueryServer = Keeper{}

// PostReports implements the Query/Session gRPC method
func (k Keeper) PostReports(ctx context.Context, request *types.QueryPostReportsRequest) (*types.QueryPostReportsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	reports, err := k.GetPostReports(sdkCtx, request.PostId)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic,
			fmt.Sprintf("no stored found for post with id %s", request.PostId))
	}

	return &types.QueryPostReportsResponse{Reports: reports}, nil
}
