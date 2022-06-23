package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/cosmwasm"
	subspaceskeeper "github.com/desmos-labs/desmos/v3/x/subspaces/keeper"
	"github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

var _ cosmwasm.Querier = SubspacesWasmQuerier{}

type SubspacesWasmQuerier struct {
	subspacesKeeper subspaceskeeper.Keeper
	cdc             codec.Codec
}

func NewSubspacesWasmQuerier(subspacesKeeper subspaceskeeper.Keeper, cdc codec.Codec) SubspacesWasmQuerier {
	return SubspacesWasmQuerier{subspacesKeeper: subspacesKeeper, cdc: cdc}
}

func (querier SubspacesWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.SubspacesQuery
	err := json.Unmarshal(data, &query)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case query.Subspaces != nil:
		return querier.handleSubspacesRequest(ctx, *query.Subspaces)

	case query.Subspace != nil:
		return querier.handleSubspaceRequest(ctx, *query.Subspace)

	case query.UserGroups != nil:
		return querier.handleUserGroupsRequest(ctx, *query.UserGroups)

	case query.UserGroup != nil:
		return querier.handleUserGroupRequest(ctx, *query.UserGroup)

	case query.UserGroupMembers != nil:
		return querier.handleUserGroupMembersRequest(ctx, *query.UserGroupMembers)

	case query.UserPermissions != nil:
		return querier.handleUserPermissionsRequest(ctx, *query.UserPermissions)

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
	res, err := querier.subspacesKeeper.Subspaces(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier SubspacesWasmQuerier) handleSubspaceRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QuerySubspaceRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.subspacesKeeper.Subspace(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier SubspacesWasmQuerier) handleUserGroupsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.subspacesKeeper.UserGroups(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier SubspacesWasmQuerier) handleUserGroupRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.subspacesKeeper.UserGroup(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier SubspacesWasmQuerier) handleUserGroupMembersRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserGroupMembersRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.subspacesKeeper.UserGroupMembers(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier SubspacesWasmQuerier) handleUserPermissionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryUserPermissionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.subspacesKeeper.UserPermissions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}
