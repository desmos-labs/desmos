package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the magpie Querier
const (
	QuerySessions = "sessions"
)

// Params for queries:
// - 'custom/magpie/sessions'

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QuerySessions:
			return querySession(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown magpie query endpoint")
		}
	}
}

// querySession allows to return a Session object based on its id
// Query path: custom/magpie/sessions/{id}
func querySession(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, error) {
	id, err := types.ParseSessionID(path[0])
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("invalid session id: %s", path[0]))
	}

	session, found := keeper.GetSession(ctx, id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("session with id %s not found", id))
	}

	res, err := codec.MarshalJSONIndent(keeper.Cdc, &session)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
