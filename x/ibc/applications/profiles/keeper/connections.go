package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/ibc/applications/profiles/types"
)

// Connections are stored in two ways:
// 1. ConnectionKey -> *types.Connection
// 2. ClientID -> ConnectionKey
//
// This allows to get connections by client id quickly

// StoreConnection stores the given connection replacing any existing one for the same user and application
func (k Keeper) StoreConnection(ctx sdk.Context, connection *types.Connection) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.MarshalBinaryBare(connection)
	if err != nil {
		return err
	}

	connectionKey := types.ConnectionKey(connection)
	store.Set(connectionKey, bz)
	store.Set(types.ConnectionClientIDKey(connection.OracleRequest.ClientId), connectionKey)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeConnectionSaved,
			sdk.NewAttribute(types.AttributeKeyApplicationName, connection.Application.Name),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, connection.Application.Username),
		),
	)

	return nil
}

func (k Keeper) GetConnectionByClientID(ctx sdk.Context, clientID string) (*types.Connection, error) {
	store := ctx.KVStore(k.storeKey)

	clientIDKey := types.ConnectionClientIDKey(clientID)
	if !store.Has(clientIDKey) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "connection for client id %s not found", clientID)
	}

	connectionKey := store.Get(clientIDKey)
	if !store.Has(connectionKey) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "connection key by client id found but no connection found")
	}

	var connection types.Connection
	err := k.cdc.UnmarshalBinaryBare(store.Get(connectionKey), &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// GetUserConnections returns all the connections of a given user
func (k Keeper) GetUserConnections(ctx sdk.Context, user string) ([]*types.Connection, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UserConnectionsPrefix(user))
	defer iterator.Close()

	var connections []*types.Connection
	for ; iterator.Valid(); iterator.Next() {
		var connection types.Connection
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &connection)
		if err != nil {
			return nil, err
		}

		connections = append(connections, &connection)
	}

	return connections, nil
}
