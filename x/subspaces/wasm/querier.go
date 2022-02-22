package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v2/cosmwasm"
	subspaceskeeper "github.com/desmos-labs/desmos/v2/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

var _ cosmwasm.Querier = SubspacesWasmQuerier{}

type SubspacesWasmQuerier struct {
	subspacesKeeper subspaceskeeper.Keeper
	cdc             codec.Codec
}

func NewSubspacesWasmQuerier(subspacesKeeper subspaceskeeper.Keeper, cdc codec.Codec) SubspacesWasmQuerier {
	return SubspacesWasmQuerier{subspacesKeeper: subspacesKeeper, cdc: cdc}
}

func (SubspacesWasmQuerier) Query(_ sdk.Context, _ wasmvmtypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier SubspacesWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var desmosQuery types.SubspacesQueryRoutes
	err := json.Unmarshal(data, &desmosQuery)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	routes := desmosQuery.Subspaces
	var bz []byte
	switch {
	case routes.Subspaces != nil:
		subspacesResponse, err := querier.subspacesKeeper.Subspaces(sdk.WrapSDKContext(ctx), routes.Subspaces)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(subspacesResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case routes.Subspace != nil:
		response, err := querier.subspacesKeeper.Subspace(sdk.WrapSDKContext(ctx), routes.Subspace)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case routes.UserGroups != nil:
		response, err := querier.subspacesKeeper.UserGroups(sdk.WrapSDKContext(ctx), routes.UserGroups)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case routes.UserGroup != nil:
		response, err := querier.subspacesKeeper.UserGroup(sdk.WrapSDKContext(ctx), routes.UserGroup)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case routes.UserGroupMembers != nil:
		response, err := querier.subspacesKeeper.UserGroupMembers(sdk.WrapSDKContext(ctx), routes.UserGroupMembers)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case routes.UserPermissions != nil:
		response, err := querier.subspacesKeeper.UserPermissions(sdk.WrapSDKContext(ctx), routes.UserPermissions)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
	return bz, nil
}
