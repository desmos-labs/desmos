package keeper

import (
	"context"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

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
	var relationships []types.Relationship

	// Get user relationships prefix store
	store := sdkCtx.KVStore(k.storeKey)
	relsStore := prefix.NewStore(store, types.UserRelationshipsSubspacePrefix(request.User, request.Subspace))

	// Get paginated user relationships
	pageRes, err := query.FilteredPaginate(relsStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var rel types.Relationship
		if err := k.cdc.UnmarshalBinaryBare(value, &rel); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		if accumulate {
			relationships = append(relationships, rel)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserRelationshipsResponse{User: request.User, Relationships: relationships, Pagination: pageRes}, nil
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

// UserChainLinks implements the Query/UserChainLinks gRPC method
func (k Keeper) UserChainLinks(ctx context.Context, request *types.QueryUserChainLinksRequest) (*types.QueryUserChainLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var links []types.ChainLink

	// Get user chain links prefix store
	store := sdkCtx.KVStore(k.storeKey)
	linksStore := prefix.NewStore(store, types.UserChainLinksPrefix(request.User))

	// Get paginated user chain links
	pageRes, err := query.FilteredPaginate(linksStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		link := types.MustUnmarshalChainLink(k.cdc, value)
		if accumulate {
			links = append(links, link)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserChainLinksResponse{Links: links, Pagination: pageRes}, nil
}

// UserChainLink implements the Query/UserChainLink gRPC method
func (k Keeper) UserChainLink(ctx context.Context, request *types.QueryUserChainLinkRequest) (*types.QueryUserChainLinkResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	link, found := k.GetChainLink(sdkCtx, request.User, request.ChainName, request.Target)
	if !found {
		return nil, status.Error(codes.NotFound, "link not found")
	}

	return &types.QueryUserChainLinkResponse{Link: link}, nil
}

// UserApplicationLinks implements the Query/UserApplicationLinks gRPC method
func (k Keeper) UserApplicationLinks(ctx context.Context, request *types.QueryUserApplicationLinksRequest) (*types.QueryUserApplicationLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var links []types.ApplicationLink

	// Get user links prefix store
	store := sdkCtx.KVStore(k.storeKey)
	linksStore := prefix.NewStore(store, types.UserApplicationLinksPrefix(request.User))

	// Get paginated user links
	pageRes, err := query.FilteredPaginate(linksStore, request.Pagination, func(key []byte, value []byte, accumulate bool) (bool, error) {
		var link types.ApplicationLink
		if err := k.cdc.UnmarshalBinaryBare(value, &link); err != nil {
			return false, status.Error(codes.Internal, err.Error())
		}

		if accumulate {
			links = append(links, link)
		}
		return true, nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryUserApplicationLinksResponse{Links: links, Pagination: pageRes}, nil
}

// UserApplicationLink implements the Query/UserApplicationLink gRPC method
func (k Keeper) UserApplicationLink(ctx context.Context, request *types.QueryUserApplicationLinkRequest) (*types.QueryUserApplicationLinkResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	link, found, err := k.GetApplicationLink(sdkCtx, request.User, request.Application, request.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "link not found")
	}

	return &types.QueryUserApplicationLinkResponse{Link: link}, nil
}
