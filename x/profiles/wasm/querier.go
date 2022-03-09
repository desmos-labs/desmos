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
	var route types.DesmosQueryRoute
	err := json.Unmarshal(data, &route)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	query := route.Profiles
	var response []byte
	switch {
	case query.Profile != nil:
		if response, err = querier.handleProfileRequest(ctx, query.Profile); err != nil {
			return nil, err
		}
	case query.Relationships != nil:
		if response, err = querier.handleRelationshipsRequest(ctx, query.Relationships); err != nil {
			return nil, err
		}
	case query.IncomingDtagTransferRequests != nil:
		if response, err = querier.handleIncomingDTagRequest(ctx, query.IncomingDtagTransferRequests); err != nil {
			return nil, err
		}
	case query.Blocks != nil:
		if response, err = querier.handleBlocksRequest(ctx, query.Blocks); err != nil {
			return nil, err
		}
	case query.ChainLinks != nil:
		if response, err = querier.handleChainLinksRequest(ctx, query.ChainLinks); err != nil {
			return nil, err
		}
	case query.UserChainLink != nil:
		if response, err = querier.handleUserChainLinkRequest(ctx, query.UserChainLink); err != nil {
			return nil, err
		}
	case query.AppLinks != nil:
		if response, err = querier.handleAppLinksRequest(ctx, query.AppLinks); err != nil {
			return nil, err
		}
	case query.UserAppLinks != nil:
		if response, err = querier.handleUserAppLinkRequest(ctx, query.UserAppLinks); err != nil {
			return nil, err
		}
	case query.ApplicationLinkByClientID != nil:
		if response, err = querier.handleApplicationLinkByClientIDRequest(ctx, query.ApplicationLinkByClientID); err != nil {
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

func (querier ProfilesWasmQuerier) handleUserChainLinkRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var chainLink types.QueryUserChainLinkRequest
	err = querier.cdc.UnmarshalJSON(request, &chainLink)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	userChainLinkResponse, err := querier.profilesKeeper.UserChainLink(sdk.WrapSDKContext(ctx), &chainLink)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(userChainLinkResponse)
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

func (querier ProfilesWasmQuerier) handleUserAppLinkRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var userAppLink types.QueryUserApplicationLinkRequest
	err = querier.cdc.UnmarshalJSON(request, &userAppLink)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	userAppLinksResponse, err := querier.profilesKeeper.UserApplicationLink(sdk.WrapSDKContext(ctx), &userAppLink)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(userAppLinksResponse)
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

func (querier ProfilesWasmQuerier) handleRelationshipsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var relationshipsReq types.QueryRelationshipsRequest
	err = querier.cdc.UnmarshalJSON(request, &relationshipsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	relationshipsResponse, err := querier.profilesKeeper.Relationships(sdk.WrapSDKContext(ctx), &relationshipsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(relationshipsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier ProfilesWasmQuerier) handleBlocksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var blockReq types.QueryBlocksRequest
	err = querier.cdc.UnmarshalJSON(request, &blockReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	blocksResponse, err := querier.profilesKeeper.Blocks(sdk.WrapSDKContext(ctx), &blockReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(blocksResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
