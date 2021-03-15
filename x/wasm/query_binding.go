package wasm

import (
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmTypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
	reportsTypes "github.com/desmos-labs/desmos/x/reports/types"
)

type WasmQuerier interface {
	Query(ctx sdk.Context, request wasmTypes.QueryRequest) ([]byte, error)
	QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error)
}

type QueriersMap struct {
	Queriers map[string]WasmQuerier
}

func NewQuerier() QueriersMap {
	return QueriersMap{
		Queriers: make(map[string]WasmQuerier),
	}
}

type WasmCustomQuery struct {
	Route     string          `json:"route"`
	QueryData json.RawMessage `json:"query_data"`
}

const (
	WasmQueryRoutePosts   = postsTypes.ModuleName
	WasmQueryRouteReports = reportsTypes.ModuleName
)

func (q QueriersMap) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var customQuery WasmCustomQuery
	err := json.Unmarshal(data, &customQuery)
	fmt.Println("[!] Wasm query routed to module: ", customQuery.Route)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if querier, ok := q.Queriers[customQuery.Route]; ok {
		return querier.QueryCustom(ctx, customQuery.QueryData)
	} else {
		return nil, sdkerrors.Wrap(wasm.ErrQueryFailed, customQuery.Route)
	}
}
