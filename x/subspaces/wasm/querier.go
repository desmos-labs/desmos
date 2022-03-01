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
	var req types.SubspacesQueryRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	var bz []byte
	switch {
	case req.Subspaces != nil:
		subspacesResponse, err := querier.subspacesKeeper.Subspaces(sdk.WrapSDKContext(ctx), req.Subspaces)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(subspacesResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case req.Subspace != nil:
		response, err := querier.subspacesKeeper.Subspace(sdk.WrapSDKContext(ctx), req.Subspace)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case req.UserGroups != nil:
		response, err := querier.subspacesKeeper.UserGroups(sdk.WrapSDKContext(ctx), req.UserGroups)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case req.UserGroup != nil:
		response, err := querier.subspacesKeeper.UserGroup(sdk.WrapSDKContext(ctx), req.UserGroup)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case req.UserGroupMembers != nil:
		response, err := querier.subspacesKeeper.UserGroupMembers(sdk.WrapSDKContext(ctx), req.UserGroupMembers)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case req.UserPermissions != nil:
		response, err := querier.subspacesKeeper.UserPermissions(sdk.WrapSDKContext(ctx), req.UserPermissions)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = querier.cdc.MarshalJSON(response)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
	return bz, nil
}
