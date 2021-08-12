package wasm

import (
	"encoding/json"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmTypes "github.com/CosmWasm/wasmvm/types"
	postsTypes "github.com/desmos-labs/desmos/x/posts/types"
)

type Querier interface {
	Query(ctx sdk.Context, request wasmTypes.QueryRequest) ([]byte, error)
	QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error)
}

type QueriersMap struct {
	Queriers map[string]Querier
}

func NewQuerier(queriers map[string]Querier) QueriersMap {
	return QueriersMap{
		Queriers: queriers,
	}
}

type CustomQuery struct {
	Route     string          `json:"route"`
	QueryData json.RawMessage `json:"query_data"`
}

const (
	QueryRoutePosts = postsTypes.ModuleName
)

func (q QueriersMap) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var customQuery CustomQuery
	err := json.Unmarshal(data, &customQuery)
	fmt.Println("[!] Wasm query routed to module: ", customQuery.Route)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if querier, ok := q.Queriers[customQuery.Route]; ok {
		return querier.QueryCustom(ctx, customQuery.QueryData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrQueryFailed, customQuery.Route)
}
