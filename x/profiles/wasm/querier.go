package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	profileskeeper "github.com/desmos-labs/desmos/v2/x/profiles/keeper"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v2/cosmwasm"
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
	var desmosQuery types.ProfilesQueryRoutes
	err := json.Unmarshal(data, &desmosQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	routes := desmosQuery.Profiles
	var bz []byte

	switch {
	case routes.Profile != nil:
		profileResponse, err := querier.profilesKeeper.Profile(sdk.WrapSDKContext(ctx), routes.Profile)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(profileResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.Relationships != nil:
		relationshipsResponse, err := querier.profilesKeeper.Relationships(sdk.WrapSDKContext(ctx), routes.Relationships)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(relationshipsResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.IncomingDtagTransferRequests != nil:
		incomingDtagTransferRequestsResponse, err := querier.profilesKeeper.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx),
			routes.IncomingDtagTransferRequests)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(incomingDtagTransferRequestsResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.Blocks != nil:
		blocksResponse, err := querier.profilesKeeper.Blocks(sdk.WrapSDKContext(ctx), routes.Blocks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(blocksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.ChainLinks != nil:
		chainLinksResponse, err := querier.profilesKeeper.ChainLinks(sdk.WrapSDKContext(ctx), routes.ChainLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(chainLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.UserChainLink != nil:
		userChainLinkResponse, err := querier.profilesKeeper.UserChainLink(sdk.WrapSDKContext(ctx), routes.UserChainLink)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(userChainLinkResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.AppLinks != nil:
		appLinksResponse, err := querier.profilesKeeper.ApplicationLinks(sdk.WrapSDKContext(ctx), routes.AppLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(appLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.UserAppLinks != nil:
		userAppLinksResponse, err := querier.profilesKeeper.UserApplicationLink(sdk.WrapSDKContext(ctx), routes.UserAppLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(userAppLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case routes.ApplicationLinkByClientID != nil:
		applicationLinkByChainIDResponse, err := querier.profilesKeeper.ApplicationLinkByClientID(
			sdk.WrapSDKContext(ctx),
			routes.ApplicationLinkByClientID,
		)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(applicationLinkByChainIDResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}

	return bz, nil
}
