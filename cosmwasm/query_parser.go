package cosmwasm

import (
	"encoding/json"

	"cosmossdk.io/errors"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	poststypes "github.com/desmos-labs/desmos/v7/x/posts/types"
	profiletypes "github.com/desmos-labs/desmos/v7/x/profiles/types"
	reactionstypes "github.com/desmos-labs/desmos/v7/x/reactions/types"
	relationshipstypes "github.com/desmos-labs/desmos/v7/x/relationships/types"
	reportstypes "github.com/desmos-labs/desmos/v7/x/reports/types"
	subspacestypes "github.com/desmos-labs/desmos/v7/x/subspaces/types"
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
	Profiles      *json.RawMessage `json:"profiles"`
	Subspaces     *json.RawMessage `json:"subspaces"`
	Relationships *json.RawMessage `json:"relationships"`
	Posts         *json.RawMessage `json:"posts"`
	Reports       *json.RawMessage `json:"reports"`
	Reactions     *json.RawMessage `json:"reactions"`
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
		return nil, errors.Wrap(sdkerrors.ErrJSONUnmarshal, err.Error())
	}
	// get route and query from data
	var route string
	var query json.RawMessage
	switch {
	case customQuery.Profiles != nil:
		route = QueryRouteProfiles
		query = *customQuery.Profiles
	case customQuery.Subspaces != nil:
		route = QueryRouteSubspaces
		query = *customQuery.Subspaces
	case customQuery.Relationships != nil:
		route = QueryRouteRelationships
		query = *customQuery.Relationships
	case customQuery.Posts != nil:
		route = QueryRoutePosts
		query = *customQuery.Posts
	case customQuery.Reports != nil:
		route = QueryRouteReports
		query = *customQuery.Reports
	case customQuery.Reactions != nil:
		route = QueryRouteReactions
		query = *customQuery.Reactions
	}

	if querier, ok := q.Queriers[route]; ok {
		return querier.QueryCustom(ctx, query)
	}
	return nil, errors.Wrap(wasmtypes.ErrQueryFailed, "unimplemented route")
}
