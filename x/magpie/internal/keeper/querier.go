package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the magpie Querier
const (
	QuerySessions = "sessions"
)

// Params for queries:
// - 'custom/magpie/sessions'

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QuerySessions:
			return querySession(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown magpie query endpoint")
		}
	}
}

// querySession allows to return a Session object based on its id
// Query path: custom/magpie/sessions/{id}
func querySession(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	id, err := types.ParseSessionID(path[0])
	if err != nil {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Invalid session id: %s", path[0]))
	}

	session, found := keeper.GetSession(ctx, id)
	if !found {
		return nil, sdk.ErrUnknownRequest(fmt.Sprintf("Session with id %s not found", id))
	}

	res, err := codec.MarshalJSONIndent(keeper.Cdc, &session)
	if err != nil {
		return nil, sdk.ErrInternal("Could not marshal result to JSON")
	}

	return res, nil
}
