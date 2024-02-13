package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"

	profileskeeper "github.com/desmos-labs/desmos/v7/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v7/x/profiles/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v7/cosmwasm"
)

var _ cosmwasm.Querier = ProfilesWasmQuerier{}

type ProfilesWasmQuerier struct {
	profilesKeeper *profileskeeper.Keeper
	cdc            codec.Codec
}

func NewProfilesWasmQuerier(profilesKeeper *profileskeeper.Keeper, cdc codec.Codec) ProfilesWasmQuerier {
	return ProfilesWasmQuerier{profilesKeeper: profilesKeeper, cdc: cdc}
}

func (querier ProfilesWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.ProfilesQuery
	err := json.Unmarshal(data, &query)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case query.Profile != nil:
		return querier.handleProfileRequest(ctx, *query.Profile)
	case query.IncomingDtagTransferRequests != nil:
		return querier.handleIncomingDTagRequest(ctx, *query.IncomingDtagTransferRequests)
	case query.ChainLinks != nil:
		return querier.handleChainLinksRequest(ctx, *query.ChainLinks)
	case query.ChainLinkOwners != nil:
		return querier.handleChainLinkOwnersRequest(ctx, *query.ChainLinkOwners)
	case query.DefaultExternalAddresses != nil:
		return querier.handleDefaultExternalAddressesRequest(ctx, *query.DefaultExternalAddresses)
	case query.ApplicationLinks != nil:
		return querier.handleApplicationLinksRequest(ctx, *query.ApplicationLinks)
	case query.ApplicationLinkByClientID != nil:
		return querier.handleApplicationLinkByClientIDRequest(ctx, *query.ApplicationLinkByClientID)
	case query.ApplicationLinkOwners != nil:
		return querier.handleApplicationLinkOwnersRequest(ctx, *query.ApplicationLinkOwners)
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier ProfilesWasmQuerier) handleProfileRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryProfileRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.Profile(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleIncomingDTagRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryIncomingDTagTransferRequestsRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleChainLinksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryChainLinksRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.ChainLinks(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func (querier ProfilesWasmQuerier) handleChainLinkOwnersRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryChainLinkOwnersRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.ChainLinkOwners(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func (querier ProfilesWasmQuerier) handleDefaultExternalAddressesRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryDefaultExternalAddressesRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.DefaultExternalAddresses(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}

func (querier ProfilesWasmQuerier) handleApplicationLinksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryApplicationLinksRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.ApplicationLinks(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleApplicationLinkByClientIDRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryApplicationLinkByClientIDRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.ApplicationLinkByClientID(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleApplicationLinkOwnersRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryApplicationLinkOwnersRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.profilesKeeper.ApplicationLinkOwners(
		sdk.WrapSDKContext(ctx),
		&req,
	)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
