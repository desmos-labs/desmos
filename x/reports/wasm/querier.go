package wasm

import (
	"encoding/json"

	wasmTypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	reportsKeeper "github.com/desmos-labs/desmos/x/reports/keeper"
)

type Querier interface {
	Query(ctx sdk.Context, request wasmTypes.QueryRequest) ([]byte, error)
	QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error)
}

var _ Querier = ReportsWasmQuerier{}

type ReportsWasmQuerier struct {
	reportsKeeper reportsKeeper.Keeper
}

func NewReportsWasmQuerier(reportsKeeper reportsKeeper.Keeper) ReportsWasmQuerier {
	return ReportsWasmQuerier{reportsKeeper: reportsKeeper}
}

func (ReportsWasmQuerier) Query(_ sdk.Context, _ wasmTypes.QueryRequest) ([]byte, error) {
	return nil, nil
}

func (querier ReportsWasmQuerier) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var reportsQuery ReportsModuleQuery
	err := json.Unmarshal(data, &reportsQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	var bz []byte

	if reportsQuery.Reports != nil {
		reports := querier.reportsKeeper.GetPostReports(ctx, reportsQuery.Reports.PostID)
		bz, err = json.Marshal(ReportsResponse{Reports: reports})
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
		}
	} else {
		return nil, sdkerrors.ErrInvalidRequest
	}

	return bz, nil
}
