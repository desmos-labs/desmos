package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v3/cosmwasm"
	reactionskeeper "github.com/desmos-labs/desmos/v3/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v3/x/reactions/types"
)

var _ cosmwasm.Querier = ReactionsWasmQuerier{}

type ReactionsWasmQuerier struct {
	reactionskeeper reactionskeeper.Keeper
	cdc             codec.Codec
}

func NewReactionsWasmQuerier(reactionskeeper reactionskeeper.Keeper, cdc codec.Codec) ReactionsWasmQuerier {
	return ReactionsWasmQuerier{reactionskeeper: reactionskeeper, cdc: cdc}
}

func (ReactionsWasmQuerier) Query(_ sdk.Context, _ wasmvmtypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier ReactionsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var req types.ReactionsQuery
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case req.Reactions != nil:
		return querier.handleReactionsRequest(ctx, *req.Reactions)
	case req.RegisteredReactions != nil:
		return querier.handleRegisteredReactionsRequest(ctx, *req.RegisteredReactions)
	case req.RegisteredReactions != nil:
		return querier.handleReactionsParamsRequest(ctx, *req.ReactionsParams)
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier ReactionsWasmQuerier) handleReactionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReactionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.Reactions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleRegisteredReactionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryRegisteredReactionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.RegisteredReactions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleReactionsParamsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReactionsParamsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.ReactionsParams(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}
