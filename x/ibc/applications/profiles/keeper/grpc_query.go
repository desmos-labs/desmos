package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"
)

var _ types.QueryServer = Keeper{}

// Connections implements the Query/Connections gRPC method
func (k Keeper) Connections(
	ctx context.Context, request *types.QueryUserApplicationsLinksRequest,
) (*types.QueryUserApplicationsLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	links, err := k.GetUserApplicationsLinks(sdkCtx, request.User)
	if err != nil {
		return nil, err
	}

	return &types.QueryUserApplicationsLinksResponse{
		ApplicationLinks: links,
	}, nil
}
