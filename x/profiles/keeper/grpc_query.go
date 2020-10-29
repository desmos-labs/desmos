package keeper

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/profiles/types"
	"strings"
)

var _ types.QueryServer = Keeper{}

// Profiles implements the Query/Profiles gRPC method
func (k Keeper) Profiles(ctx context.Context, request *types.QueryProfilesRequest) (*types.QueryProfilesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	profiles := k.GetProfiles(sdkCtx)
	return &types.QueryProfilesResponse{Profiles: profiles}, nil
}

// Profiles implements the Query/Profiles gRPC method
func (k Keeper) Profile(ctx context.Context, request *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	dTagOrAddress := request.User
	if strings.TrimSpace(dTagOrAddress) == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or sdkAddress cannot be empty or blank")
	}

	sdkAddress, err := sdk.AccAddressFromBech32(dTagOrAddress)
	if err != nil {
		addr := k.GetDtagRelatedAddress(sdkCtx, dTagOrAddress)
		if addr == "" {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"No address related to this DTag: %s", dTagOrAddress)
		}

		sdkAddress, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, addr)
		}
	}

	account, found := k.GetProfile(sdkCtx, sdkAddress.String())
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Profile with sdkAddress %s doesn't exists", dTagOrAddress)
	}

	return &types.QueryProfileResponse{Profile: account}, nil
}

// DTagTransfers implements the Query/DTagTransfers gRPC method
func (k Keeper) DTagTransfers(ctx context.Context, request *types.QueryDTagTransfersRequest) (*types.QueryDTagTransfersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	user, err := sdk.AccAddressFromBech32(request.User)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, request.User)
	}

	requests := k.GetUserDTagTransferRequests(sdkCtx, user.String())
	return &types.QueryDTagTransfersResponse{Requests: requests}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}
