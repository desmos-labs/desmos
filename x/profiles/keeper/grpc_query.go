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

	"github.com/desmos-labs/desmos/v2/x/profiles/types"
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

func (k Keeper) IncomingDTagTransferRequests(ctx context.Context, request *types.QueryIncomingDTagTransferRequestsRequest) (*types.QueryIncomingDTagTransferRequestsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var requests []types.DTagTransferRequest

	// Get user requests prefix store
	store := sdkCtx.KVStore(k.storeKey)
	reqStore := prefix.NewStore(store, types.IncomingDTagTransferRequestsPrefix(request.Receiver))

	// Get paginated user requests
	pageRes, err := query.Paginate(reqStore, request.Pagination, func(key []byte, value []byte) error {
		var req types.DTagTransferRequest
		if err := k.cdc.Unmarshal(value, &req); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		requests = append(requests, req)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryIncomingDTagTransferRequestsResponse{Requests: requests, Pagination: pageRes}, nil
}

// Relationships implements the Query/Relationships gRPC method
func (k Keeper) Relationships(ctx context.Context, request *types.QueryRelationshipsRequest) (*types.QueryRelationshipsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var relationships []types.Relationship

	// Get user relationships prefix store
	store := sdkCtx.KVStore(k.storeKey)
	relsStore := prefix.NewStore(store, types.UserRelationshipsSubspacePrefix(request.User, request.SubspaceId))

	// Get paginated user relationships
	pageRes, err := query.Paginate(relsStore, request.Pagination, func(key []byte, value []byte) error {
		var rel types.Relationship
		if err := k.cdc.Unmarshal(value, &rel); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		relationships = append(relationships, rel)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryRelationshipsResponse{Relationships: relationships, Pagination: pageRes}, nil
}

// Blocks implements the Query/Blocks gRPC method
func (k Keeper) Blocks(ctx context.Context, request *types.QueryBlocksRequest) (*types.QueryBlocksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var userblocks []types.UserBlock

	// Get user blocks prefix store
	store := sdkCtx.KVStore(k.storeKey)
	userBlocksStore := prefix.NewStore(store, types.BlockerSubspacePrefix(request.User, request.SubspaceId))

	// Get paginated user blocks
	pageRes, err := query.Paginate(userBlocksStore, request.Pagination, func(key []byte, value []byte) error {
		var userBlock types.UserBlock
		if err := k.cdc.Unmarshal(value, &userBlock); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		userblocks = append(userblocks, userBlock)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryBlocksResponse{Blocks: userblocks, Pagination: pageRes}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// ChainLinks implements the Query/ChainLinks gRPC method
func (k Keeper) ChainLinks(ctx context.Context, request *types.QueryChainLinksRequest) (*types.QueryChainLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	var links []types.ChainLink

	// Get user chain links prefix store
	store := sdkCtx.KVStore(k.storeKey)
	linksStore := prefix.NewStore(store, types.UserChainLinksPrefix(request.User))

	// Get paginated user chain links
	pageRes, err := query.Paginate(linksStore, request.Pagination, func(key []byte, value []byte) error {
		var link types.ChainLink
		if err := k.cdc.Unmarshal(value, &link); err != nil {
			return status.Error(codes.Internal, err.Error())
		}
		links = append(links, link)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryChainLinksResponse{Links: links, Pagination: pageRes}, nil
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

// ApplicationLinks implements the Query/ApplicationLinks gRPC method
func (k Keeper) ApplicationLinks(ctx context.Context, request *types.QueryApplicationLinksRequest) (*types.QueryApplicationLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	var links []types.ApplicationLink

	// Get user links prefix store
	store := sdkCtx.KVStore(k.storeKey)
	linksStore := prefix.NewStore(store, types.UserApplicationLinksPrefix(request.User))

	// Get paginated user links
	pageRes, err := query.Paginate(linksStore, request.Pagination, func(key []byte, value []byte) error {
		var link types.ApplicationLink
		if err := k.cdc.Unmarshal(value, &link); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		links = append(links, link)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryApplicationLinksResponse{Links: links, Pagination: pageRes}, nil
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

// ApplicationLinkByClientID implements the Query/ApplicationLinkByClientID gRPC method
func (k Keeper) ApplicationLinkByClientID(ctx context.Context, request *types.QueryApplicationLinkByClientIDRequest) (*types.QueryApplicationLinkByClientIDResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	link, found, err := k.GetApplicationLinkByClientID(sdkCtx, request.ClientId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "link for client id %s not found", request.ClientId)
	}

	return &types.QueryApplicationLinkByClientIDResponse{Link: link}, nil
}
