package wasm

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/cosmwasm"
	reportskeeper "github.com/desmos-labs/desmos/v3/x/reports/keeper"
	"github.com/desmos-labs/desmos/v3/x/reports/types"
)

var _ cosmwasm.Querier = ReportsWasmQuerier{}

type ReportsWasmQuerier struct {
	reportskeeper reportskeeper.Keeper
	cdc           codec.Codec
}

func NewReportsWasmQuerier(reportskeeper reportskeeper.Keeper, cdc codec.Codec) ReportsWasmQuerier {
	return ReportsWasmQuerier{reportskeeper: reportskeeper, cdc: cdc}
}

func (querier ReportsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.ReportsQuery
	err := json.Unmarshal(data, &query)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case query.Reports != nil:
		return querier.handleReportsRequest(ctx, *query.Reports)
	case query.Reasons != nil:
		return querier.handleReasonsRequest(ctx, *query.Reasons)
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier ReportsWasmQuerier) handleReportsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReportsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reportskeeper.Reports(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReportsWasmQuerier) handleReasonsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReasonsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reportskeeper.Reasons(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}
