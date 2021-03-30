package keeper

import (
	"context"

	"github.com/desmos-labs/desmos/x/links/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) Link(ctx context.Context, req *types.QueryLinkRequest) (*types.QueryLinkResponse, error) {
	response := &types.QueryLinkResponse{Link: nil}
	return response, nil
}
