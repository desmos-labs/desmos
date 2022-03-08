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
	var route types.SubspacesQueryRoute
	err := json.Unmarshal(data, &route)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	req := route.Query
	switch {
	case req.Subspaces != nil:
		return querier.handleSubspacesRequest(ctx, req.Subspaces)

	case req.Subspace != nil:
		return querier.handleSubspaceRequest(ctx, req.Subspace)

	case req.UserGroups != nil:
		return querier.handleUserGroupsRequest(ctx, req.UserGroups)

	case req.UserGroup != nil:
		return querier.handleUserGroupRequest(ctx, req.UserGroup)

	case req.UserGroupMembers != nil:
		return querier.handleUserGroupMembersRequest(ctx, req.UserGroupMembers)

	case req.UserPermissions != nil:
		return querier.handleUserPermissionsRequest(ctx, req.UserPermissions)

	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier SubspacesWasmQuerier) handleSubspacesRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QuerySubspacesRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.Subspaces(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}

func (querier SubspacesWasmQuerier) handleSubspaceRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QuerySubspaceRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.Subspace(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}

func (querier SubspacesWasmQuerier) handleUserGroupsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.UserGroups(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}

func (querier SubspacesWasmQuerier) handleUserGroupRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.UserGroup(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}

func (querier SubspacesWasmQuerier) handleUserGroupMembersRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupMembersRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.UserGroupMembers(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}

func (querier SubspacesWasmQuerier) handleUserPermissionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserPermissionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	subspacesResponse, err := querier.subspacesKeeper.UserPermissions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(subspacesResponse)
}
