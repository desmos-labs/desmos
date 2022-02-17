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

	var bz []byte

	switch {
	case desmosQuery.Profile != nil:
		profileResponse, err := querier.profilesKeeper.Profile(sdk.WrapSDKContext(ctx), desmosQuery.Profile)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(profileResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.Relationships != nil:
		relationshipsResponse, err := querier.profilesKeeper.Relationships(sdk.WrapSDKContext(ctx), desmosQuery.Relationships)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(relationshipsResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.IncomingDtagTransferRequests != nil:
		incomingDtagTransferRequestsResponse, err := querier.profilesKeeper.IncomingDTagTransferRequests(sdk.WrapSDKContext(ctx),
			desmosQuery.IncomingDtagTransferRequests)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(incomingDtagTransferRequestsResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.Blocks != nil:
		blocksResponse, err := querier.profilesKeeper.Blocks(sdk.WrapSDKContext(ctx), desmosQuery.Blocks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(blocksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.ChainLinks != nil:
		chainLinksResponse, err := querier.profilesKeeper.ChainLinks(sdk.WrapSDKContext(ctx), desmosQuery.ChainLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(chainLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.UserChainLink != nil:
		userChainLinkResponse, err := querier.profilesKeeper.UserChainLink(sdk.WrapSDKContext(ctx), desmosQuery.UserChainLink)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(userChainLinkResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.AppLinks != nil:
		appLinksResponse, err := querier.profilesKeeper.ApplicationLinks(sdk.WrapSDKContext(ctx), desmosQuery.AppLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(appLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.UserAppLinks != nil:
		userAppLinksResponse, err := querier.profilesKeeper.UserApplicationLink(sdk.WrapSDKContext(ctx), desmosQuery.UserAppLinks)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(userAppLinksResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	case desmosQuery.ApplicationLinkByClientID != nil:
		applicationLinkByChainIDResponse, err := querier.profilesKeeper.ApplicationLinkByClientID(
			sdk.WrapSDKContext(ctx),
			desmosQuery.ApplicationLinkByClientID,
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
