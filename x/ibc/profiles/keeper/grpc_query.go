package keeper

import (
	"context"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/ibc/profiles/types"
)

var _ types.QueryServer = Keeper{}

// Link implements the Query/Link gRPC method
func (k Keeper) Link(ctx context.Context, req *types.QueryLinkRequest) (*types.QueryLinkResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	addr := req.SourceAddress
	if strings.TrimSpace(addr) == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "Address cannot be empty or blank")
	}

	sdkAddress, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, addr)
	}

	link, found := k.GetLink(sdkCtx, sdkAddress.String())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Link with sdkAddress %s doesn't exists", addr)
	}

	return &types.QueryLinkResponse{Link: link}, nil
}
