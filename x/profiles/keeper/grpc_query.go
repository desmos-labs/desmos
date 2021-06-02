package keeper

import (
	"context"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

var _ types.QueryServer = Keeper{}

// Profile implements the Query/Profile gRPC method
func (k Keeper) Profile(ctx context.Context, request *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	dTagOrAddress := request.User
	if strings.TrimSpace(dTagOrAddress) == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank")
	}

	sdkAddress, err := sdk.AccAddressFromBech32(dTagOrAddress)
	if err != nil {
		addr := k.GetAddressFromDTag(sdkCtx, dTagOrAddress)
		if addr == "" {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
				"No address related to this DTag: %s", dTagOrAddress)
		}

		sdkAddress, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, addr)
		}
	}

	account, found, err := k.GetProfile(sdkCtx, sdkAddress.String())
	if err != nil {
		return nil, err
	}

	if !found {
		return &types.QueryProfileResponse{Profile: nil}, nil
	}

	accountAny, err := codectypes.NewAnyWithValue(account)
	if err != nil {
		return nil, err
	}

	return &types.QueryProfileResponse{Profile: accountAny}, nil
}

// DTagTransfers implements the Query/DTagTransfers gRPC method
func (k Keeper) DTagTransfers(ctx context.Context, request *types.QueryDTagTransfersRequest) (*types.QueryDTagTransfersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	user, err := sdk.AccAddressFromBech32(request.User)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, request.User)
	}

	requests := k.GetUserIncomingDTagTransferRequests(sdkCtx, user.String())
	return &types.QueryDTagTransfersResponse{Requests: requests}, nil
}

// UserRelationships implements the Query/UserRelationships gRPC method
func (k Keeper) UserRelationships(ctx context.Context, request *types.QueryUserRelationshipsRequest) (*types.QueryUserRelationshipsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	relationships := k.GetUserRelationships(sdkCtx, request.User)
	return &types.QueryUserRelationshipsResponse{User: request.User, Relationships: relationships}, nil
}

// UserBlocks implements the Query/UserBlocks gRPC method
func (k Keeper) UserBlocks(ctx context.Context, request *types.QueryUserBlocksRequest) (*types.QueryUserBlocksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	blocks := k.GetUserBlocks(sdkCtx, request.User)
	return &types.QueryUserBlocksResponse{Blocks: blocks}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) ProfileByChainLink(ctx context.Context, request *types.QueryProfileByChainLinkRequest) (*types.QueryProfileByChainLinkResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	account, found := k.GetAccountByChainLink(sdkCtx, request.ChainName, request.Target)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"No link related to this address: %s", request.Target)
	}

	profile, found, err := k.GetProfile(sdkCtx, account.String())
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest,
			"Profile with address %s doesn't exists", account.String())
	}

	profileAny, err := codectypes.NewAnyWithValue(profile)
	if err != nil {
		return nil, err
	}

	return &types.QueryProfileByChainLinkResponse{Profile: profileAny}, nil
}
