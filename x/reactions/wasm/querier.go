package wasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v6/cosmwasm"
	reactionskeeper "github.com/desmos-labs/desmos/v6/x/reactions/keeper"
	"github.com/desmos-labs/desmos/v6/x/reactions/types"
)

var _ cosmwasm.Querier = ReactionsWasmQuerier{}

type ReactionsWasmQuerier struct {
	reactionskeeper reactionskeeper.Keeper
	cdc             codec.Codec
}

func NewReactionsWasmQuerier(reactionskeeper reactionskeeper.Keeper, cdc codec.Codec) ReactionsWasmQuerier {
	return ReactionsWasmQuerier{reactionskeeper: reactionskeeper, cdc: cdc}
}

func (querier ReactionsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.ReactionsQuery
	err := json.Unmarshal(data, &query)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case query.Reactions != nil:
		return querier.handleReactionsRequest(ctx, *query.Reactions)
	case query.Reaction != nil:
		return querier.handleReactionRequest(ctx, *query.Reaction)
	case query.RegisteredReactions != nil:
		return querier.handleRegisteredReactionsRequest(ctx, *query.RegisteredReactions)
	case query.RegisteredReaction != nil:
		return querier.handleRegisteredReactionRequest(ctx, *query.RegisteredReaction)
	case query.ReactionsParams != nil:
		return querier.handleReactionsParamsRequest(ctx, *query.ReactionsParams)
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier ReactionsWasmQuerier) handleReactionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReactionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.Reactions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleReactionRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReactionRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.Reaction(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleRegisteredReactionsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryRegisteredReactionsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.RegisteredReactions(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleRegisteredReactionRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryRegisteredReactionRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.RegisteredReaction(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}

func (querier ReactionsWasmQuerier) handleReactionsParamsRequest(ctx sdk.Context, data json.RawMessage) (json.RawMessage, error) {
	var req types.QueryReactionsParamsRequest
	err := querier.cdc.UnmarshalJSON(data, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.reactionskeeper.ReactionsParams(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	return querier.cdc.MarshalJSON(res)
}
