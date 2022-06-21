package wasm

import (
	"encoding/json"

	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/v3/cosmwasm"
	relationshipskeeper "github.com/desmos-labs/desmos/v3/x/relationships/keeper"
	"github.com/desmos-labs/desmos/v3/x/relationships/types"
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

func (RelationshipsWasmQuerier) Query(_ sdk.Context, _ wasmvmtypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier RelationshipsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var query types.RelationshipsQuery
	err := json.Unmarshal(data, &query)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var response []byte
	switch {
	case query.Relationships != nil:
		if response, err = querier.handleRelationshipsRequest(ctx, *query.Relationships); err != nil {
			return nil, err
		}
	case query.Blocks != nil:
		if response, err = querier.handleBlocksRequest(ctx, *query.Blocks); err != nil {
			return nil, err
		}
	default:
		return nil, sdkerrors.ErrInvalidRequest
	}

	return response, nil
}

func (querier RelationshipsWasmQuerier) handleRelationshipsRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var relationshipsReq types.QueryRelationshipsRequest
	err = querier.cdc.UnmarshalJSON(request, &relationshipsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	relationshipsResponse, err := querier.relationshipsKeeper.Relationships(sdk.WrapSDKContext(ctx), &relationshipsReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(relationshipsResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}

func (querier RelationshipsWasmQuerier) handleBlocksRequest(ctx sdk.Context, request json.RawMessage) (bz []byte, err error) {
	var blockReq types.QueryBlocksRequest
	err = querier.cdc.UnmarshalJSON(request, &blockReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	blocksResponse, err := querier.relationshipsKeeper.Blocks(sdk.WrapSDKContext(ctx), &blockReq)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}
	bz, err = querier.cdc.MarshalJSON(blocksResponse)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return bz, nil
}
