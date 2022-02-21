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
	subspaceQuery := desmosQuery.Subspaces
	var bz []byte
	switch {
	case subspaceQuery.Subspaces != nil:
		subspacesResponse, err := querier.subspacesKeeper.Subspaces(sdk.WrapSDKContext(ctx), subspaceQuery.Subspaces)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(subspacesResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	case subspaceQuery.Subspace != nil:
		subspaceResponse, err := querier.subspacesKeeper.Subspace(sdk.WrapSDKContext(ctx), subspaceQuery.Subspace)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
		}
		bz, err = json.Marshal(subspaceResponse)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}

	default:
		return nil, sdkerrors.ErrInvalidRequest
	}
	return bz, nil
}
