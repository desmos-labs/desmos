package cosmwasm

import (
	"encoding/json"

	"github.com/CosmWasm/wasmd/x/wasm"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	poststypes "github.com/desmos-labs/desmos/v3/x/posts/types"
	profiletypes "github.com/desmos-labs/desmos/v3/x/profiles/types"
	reactionstypes "github.com/desmos-labs/desmos/v3/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v3/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v3/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v3/x/subspaces/types"
)

type Querier interface {
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
	QueryRouteProfiles      = profiletypes.ModuleName
	QueryRouteSubspaces     = subspacestypes.ModuleName
	QueryRouteRelationships = relationshipstypes.ModuleName
	QueryRoutePosts         = poststypes.ModuleName
	QueryRouteReports       = reportstypes.ModuleName
	QueryRouteReactions     = reactionstypes.ModuleName
)

func (q QuerierRouter) QueryCustom(ctx sdk.Context, data json.RawMessage) ([]byte, error) {
	var customQuery CustomQuery
	err := json.Unmarshal(data, &customQuery)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}

	if querier, ok := q.Queriers[customQuery.Route]; ok {
		return querier.QueryCustom(ctx, customQuery.QueryData)
	}

	return nil, sdkerrors.Wrap(wasm.ErrQueryFailed, customQuery.Route)
}
