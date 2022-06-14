package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	profileskeeper "github.com/desmos-labs/desmos/v3/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
)

var _ cosmwasm.Querier = ProfilesWasmQuerier{}

type ProfilesWasmQuerier struct {
	profilesKeeper profileskeeper.Keeper
	cdc            codec.Codec
}

func NewProfilesWasmQuerier(profilesKeeper profileskeeper.Keeper, cdc codec.Codec) ProfilesWasmQuerier {
	return ProfilesWasmQuerier{profilesKeeper: profilesKeeper, cdc: cdc}
}

func (ProfilesWasmQuerier) Query(_ sdk.Context, _ wasmvmtypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier ProfilesWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.ProfilesQuery
	err := json.Unmarshal(data, &query)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var response []byte
	switch {
	case query.Profile != nil:
		if response, err = querier.handleProfileRequest(ctx, *query.Profile); err != nil {
			return nil, err
		}
	case query.IncomingDtagTransferRequests != nil:
		if response, err = querier.handleIncomingDTagRequest(ctx, *query.IncomingDtagTransferRequests); err != nil {
			return nil, err
		}
	case query.ChainLinks != nil:
		if response, err = querier.handleChainLinksRequest(ctx, *query.ChainLinks); err != nil {
			return nil, err
		}
	case query.AppLinks != nil:
		if response, err = querier.handleAppLinksRequest(ctx, *query.AppLinks); err != nil {
			return nil, err
		}
	case query.ApplicationLinkByClientID != nil:
		if response, err = querier.handleApplicationLinkByClientIDRequest(ctx, *query.ApplicationLinkByClientID); err != nil {
			return nil, err
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}

	return response, nil
}

func (querier ProfilesWasmQuerier) handleProfileRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var profileReq types.QueryProfileRequest
	err = querier.cdc.UnmarshalJSON(request, &profileReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	profileResponse, err := querier.profilesKeeper.Profile(sdk.WrapSDKContext(ctx), &profileReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(profileResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleIncomingDTagRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var incomingDtagReq types.QueryIncomingDTagTransferRequestsRequest
	err = querier.cdc.UnmarshalJSON(request, &incomingDtagReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	incomingDtagTransferRequestsResponse, err := querier.profilesKeeper.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx), &incomingDtagReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(incomingDtagTransferRequestsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleChainLinksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var chainLinkReq types.QueryChainLinksRequest
	err = querier.cdc.UnmarshalJSON(request, &chainLinkReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	chainLinksResponse, err := querier.profilesKeeper.ChainLinks(sdk.WrapSDKContext(ctx), &chainLinkReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(chainLinksResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func (querier ProfilesWasmQuerier) handleAppLinksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var appLinksReq types.QueryApplicationLinksRequest
	err = querier.cdc.UnmarshalJSON(request, &appLinksReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	appLinksResponse, err := querier.profilesKeeper.ApplicationLinks(sdk.WrapSDKContext(ctx), &appLinksReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(appLinksResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleApplicationLinkByClientIDRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var applicationReq types.QueryApplicationLinkByClientIDRequest
	err = querier.cdc.UnmarshalJSON(request, &applicationReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	applicationLinkByChainIDResponse, err := querier.profilesKeeper.ApplicationLinkByClientID(
		sdk.WrapSDKContext(ctx),
		&applicationReq,
	)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(applicationLinkByChainIDResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
