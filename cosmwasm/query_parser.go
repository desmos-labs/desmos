package cosmwasm

import (
	"encoding/json"
	"fmt"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmvmtypes "github.com/CosmWasm/wasmvm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	profiletypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
	subspacestypes "github.com/desmos-labs/desmos/v2/x/subspaces/types"
)

type Querier interface {
	Query(ctx sdk.Context, request wasmvmtypes.QueryRequest) ([]byte, error)
	QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error)
}

type QuerierRouter struct {
	Queriers map[string]Querier
}

func NewQuerier(queriers map[string]Querier) QuerierRouter {
	return QuerierRouter{
		Queriers: queriers,
	}
}

type CustomQuery struct {
	Route     string          `json:"route"`
	QueryData json.RawMessage `json:"query_data"`
}

const (
	QueryRouteProfiles  = profiletypes.ModuleName
	QueryRouteSubspaces = subspacestypes.ModuleName
)

func (q QuerierRouter) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var customQuery CustomQuery
	err := json.Unmarshal(data, &customQuery)

	fmt.Println("[!] Cosmwasm contract query routed to module: ", customQuery.Route)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if querier, ok := q.Queriers[customQuery.Route]; ok {
		return querier.QueryCustom(ctx, customQuery.QueryData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrQueryFailed, customQuery.Route)
}
