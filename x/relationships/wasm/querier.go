package wasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v6/cosmwasm"
	relationshipskeeper "github.com/desmos-labs/desmos/v6/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v6/x/relationships/types"
)

var _ cosmwasm.Querier = RelationshipsWasmQuerier{}

type RelationshipsWasmQuerier struct {
	relationshipsKeeper relationshipskeeper.Keeper
	cdc                 codec.Codec
}

func NewRelationshipsWasmQuerier(relationshipsKeeper relationshipskeeper.Keeper, cdc codec.Codec) RelationshipsWasmQuerier {
	return RelationshipsWasmQuerier{
		relationshipsKeeper: relationshipsKeeper,
		cdc:                 cdc,
	}
}

func (querier RelationshipsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.RelationshipsQuery
	err := json.Unmarshal(data, &query)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	switch {
	case query.Relationships != nil:
		return querier.handleRelationshipsRequest(ctx, *query.Relationships)
	case query.Blocks != nil:
		return querier.handleBlocksRequest(ctx, *query.Blocks)
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
}

func (querier RelationshipsWasmQuerier) handleRelationshipsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryRelationshipsRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.relationshipsKeeper.Relationships(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier RelationshipsWasmQuerier) handleBlocksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var req types.QueryBlocksRequest
	err = querier.cdc.UnmarshalJSON(request, &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	res, err := querier.relationshipsKeeper.Blocks(sdk.WrapSDKContext(ctx), &req)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(res)
	if err != nil {
		return nil, errors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
