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

// SaveApplicationLink stores the given connection replacing any existing one for the same user and application
func (k Keeper) SaveApplicationLink(ctx sdk.Context, link *types.ApplicationLink) error {
	store := ctx.KVStore(k.storeKey)

	bz, err := k.cdc.MarshalBinaryBare(link)
	if err != nil {
		return err
	}

	connectionKey := types.ConnectionKey(link)
	store.Set(connectionKey, bz)
	store.Set(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientId), connectionKey)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeConnectionSaved,
			sdk.NewAttribute(types.AttributeKeyApplicationName, link.Application.Name),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, link.Application.Username),
		),
	)

	return nil
}

func (k Keeper) GetApplicationLinkByClientID(ctx sdk.Context, clientID string) (*types.ApplicationLink, error) {
	store := ctx.KVStore(k.storeKey)

	clientIDKey := types.ApplicationLinkClientIDKey(clientID)
	if !store.Has(clientIDKey) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "link for client id %s not found", clientID)
	}

	connectionKey := store.Get(clientIDKey)
	if !store.Has(connectionKey) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "link key by client id found but no link found")
	}

	var link types.ApplicationLink
	err := k.cdc.UnmarshalBinaryBare(store.Get(connectionKey), &link)
	if err != nil {
		return nil, err
	}

	return &link, nil
}

// GetUserApplicationsLinks returns all the connections of a given user
func (k Keeper) GetUserApplicationsLinks(ctx sdk.Context, user string) ([]*types.ApplicationLink, error) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.UserApplicationsLinksPrefix(user))
	defer iterator.Close()

	var links []*types.ApplicationLink
	for ; iterator.Valid(); iterator.Next() {
		var link types.ApplicationLink
		err := k.cdc.UnmarshalBinaryBare(iterator.Value(), &link)
		if err != nil {
			return nil, err
		}

		links = append(links, &link)
	}

	return links, nil
}
