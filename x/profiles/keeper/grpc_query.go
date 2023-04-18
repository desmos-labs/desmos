package keeper

import (
	"context"
	"strings"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

var _ types.QueryServer = Keeper{}

// Profile implements the Query/Profile gRPC method
func (k Keeper) Profile(ctx context.Context, request *types.QueryProfileRequest) (*types.QueryProfileResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	dTagOrAddress := request.User
	if strings.TrimSpace(dTagOrAddress) == "" {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, "DTag or address cannot be empty or blank")
	}

	sdkAddress, err := sdk.AccAddressFromBech32(dTagOrAddress)
	if err != nil {
		addr := k.GetAddressFromDTag(sdkCtx, dTagOrAddress)
		if addr == "" {
			return nil, errors.Wrapf(sdkerrors.ErrInvalidRequest,
				"No address related to this DTag: %s", dTagOrAddress)
		}

		sdkAddress, err = sdk.AccAddressFromBech32(addr)
		if err != nil {
			return nil, errors.Wrapf(sdkerrors.ErrInvalidAddress, addr)
		}
	}

	account, found, err := k.GetProfile(sdkCtx, sdkAddress.String())
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, status.Errorf(codes.NotFound, "profile for dtag/address %s not found", dTagOrAddress)
	}

	accountAny, err := codectypes.NewAnyWithValue(account)
	if err != nil {
		return nil, err
	}

	return &types.QueryProfileResponse{Profile: accountAny}, nil
}

func (k Keeper) IncomingDTagTransferRequests(ctx context.Context, request *types.QueryIncomingDTagTransferRequestsRequest) (*types.QueryIncomingDTagTransferRequestsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get user requests prefix store
	store := sdkCtx.KVStore(k.storeKey)
	reqStore := prefix.NewStore(store, types.DTagTransferRequestPrefix)
	if request.Receiver != "" {
		reqStore = prefix.NewStore(store, types.IncomingDTagTransferRequestsPrefix(request.Receiver))
	}

	// Get paginated user requests
	var requests []types.DTagTransferRequest
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

// ChainLinks implements the Query/ChainLinks gRPC method
func (k Keeper) ChainLinks(ctx context.Context, request *types.QueryChainLinksRequest) (*types.QueryChainLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// Get user chain links prefix store
	linksPrefix := types.ChainLinksPrefix
	switch {
	case request.User != "" && request.ChainName != "" && request.Target != "":
		linksPrefix = types.ChainLinksStoreKey(request.User, request.ChainName, request.Target)
	case request.User != "" && request.ChainName != "":
		linksPrefix = types.UserChainLinksChainPrefix(request.User, request.ChainName)
	case request.User != "":
		linksPrefix = types.UserChainLinksPrefix(request.User)
	}

	// Get paginated user chain links
	var links []types.ChainLink
	linksStore := prefix.NewStore(store, linksPrefix)
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

// ChainLinkOwners implements the Query/ChainLinkOwners gRPC method
func (k Keeper) ChainLinkOwners(ctx context.Context, request *types.QueryChainLinkOwnersRequest) (*types.QueryChainLinkOwnersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	ownersPrefix := types.ChainLinkChainPrefix
	switch {
	case request.ChainName != "" && request.Target != "":
		ownersPrefix = types.ChainLinkChainAddressKey(request.ChainName, request.Target)
	case request.ChainName != "":
		ownersPrefix = types.ChainLinkChainKey(request.ChainName)
	}

	var owners []types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails
	ownersStore := prefix.NewStore(store, ownersPrefix)
	pageRes, err := query.Paginate(ownersStore, request.Pagination, func(key []byte, value []byte) error {
		// Re-add the prefix because the prefix store trims it out, and we need it to get the data
		keyWithPrefix := append([]byte(nil), ownersPrefix...)
		keyWithPrefix = append(keyWithPrefix, key...)
		chainName, target, user := types.GetChainLinkOwnerData(keyWithPrefix)

		owners = append(owners, types.QueryChainLinkOwnersResponse_ChainLinkOwnerDetails{
			User:      user,
			ChainName: chainName,
			Target:    target,
		})

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryChainLinkOwnersResponse{Owners: owners, Pagination: pageRes}, nil
}

// DefaultExternalAddresses implements the Query/DefaultExternalAddresses gRPC method
func (k Keeper) DefaultExternalAddresses(ctx context.Context, request *types.QueryDefaultExternalAddressesRequest) (*types.QueryDefaultExternalAddressesResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	defaultPrefix := types.DefaultExternalAddressPrefix
	switch {
	case request.Owner != "" && request.ChainName != "":
		defaultPrefix = types.DefaultExternalAddressKey(request.Owner, request.ChainName)
	case request.Owner != "":
		defaultPrefix = types.OwnerDefaultExternalAddressPrefix(request.Owner)
	}

	var links []types.ChainLink
	defaultStore := prefix.NewStore(store, defaultPrefix)
	pageRes, err := query.Paginate(defaultStore, request.Pagination, func(key []byte, value []byte) error {
		// Re-add the prefix because the prefix store trims it out, and we need it to get the data
		keyWithPrefix := append([]byte(nil), defaultPrefix...)
		keyWithPrefix = append(keyWithPrefix, key...)
		owner, chainName := types.GetDefaultExternalAddressData(keyWithPrefix)
		link, found := k.GetChainLink(sdkCtx, owner, chainName, string(value))
		if !found {
			return errors.Wrap(sdkerrors.ErrNotFound, "chain link not found")
		}
		links = append(links, link)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryDefaultExternalAddressesResponse{Links: links, Pagination: pageRes}, nil
}

// ApplicationLinks implements the Query/ApplicationLinks gRPC method
func (k Keeper) ApplicationLinks(ctx context.Context, request *types.QueryApplicationLinksRequest) (*types.QueryApplicationLinksResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	// Get user links prefix store
	linksPrefix := types.ApplicationLinkPrefix
	switch {
	case request.User != "" && request.Application != "" && request.Username != "":
		linksPrefix = types.UserApplicationLinkKey(request.User, request.Application, request.Username)
	case request.User != "" && request.Application != "":
		linksPrefix = types.UserApplicationLinksApplicationPrefix(request.User, request.Application)
	case request.User != "":
		linksPrefix = types.UserApplicationLinksPrefix(request.User)
	}

	// Get paginated user links
	var links []types.ApplicationLink
	linksStore := prefix.NewStore(store, linksPrefix)
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

// ApplicationLinkOwners implements the Query/ApplicationLinkOwners gRPC method
func (k Keeper) ApplicationLinkOwners(ctx context.Context, request *types.QueryApplicationLinkOwnersRequest) (*types.QueryApplicationLinkOwnersResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	store := sdkCtx.KVStore(k.storeKey)

	ownersPrefix := types.ApplicationLinkAppPrefix
	switch {
	case request.Application != "" && request.Username != "":
		ownersPrefix = types.ApplicationLinkAppUsernameKey(request.Application, request.Username)
	case request.Application != "":
		ownersPrefix = types.ApplicationLinkAppKey(request.Application)
	}

	var owners []types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails
	ownersStore := prefix.NewStore(store, ownersPrefix)
	pageRes, err := query.Paginate(ownersStore, request.Pagination, func(key []byte, value []byte) error {
		// Re-add the prefix because the prefix store trims it out, and we need it to get the data
		keyWithPrefix := append([]byte(nil), ownersPrefix...)
		keyWithPrefix = append(keyWithPrefix, key...)
		application, username, user := types.GetApplicationLinkOwnerData(keyWithPrefix)

		owners = append(owners, types.QueryApplicationLinkOwnersResponse_ApplicationLinkOwnerDetails{
			User:        user,
			Application: application,
			Username:    username,
		})

		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryApplicationLinkOwnersResponse{Owners: owners, Pagination: pageRes}, nil
}

// Params implements the Query/Params gRPC method
func (k Keeper) Params(ctx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	params := k.GetParams(sdkCtx)
	return &types.QueryParamsResponse{Params: params}, nil
}
