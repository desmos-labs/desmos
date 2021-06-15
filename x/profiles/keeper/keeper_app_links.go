package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Connections are stored in two ways:
// 1. ApplicationLinkKey -> types.ApplicationLink
// 2. ClientID -> types.ClientRequest
//
// This allows to get connections by client id quickly

// SaveApplicationLink stores the given connection replacing any existing one for the same user and application
func (k Keeper) SaveApplicationLink(ctx sdk.Context, user string, link types.ApplicationLink) error {
	_, found, err := k.GetProfile(ctx, user)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrapf(types.ErrProfileNotFound, "a profile is required to link an application")
	}

	store := ctx.KVStore(k.storeKey)

	// Store the link
	linkBz, err := k.cdc.MarshalBinaryBare(&link)
	if err != nil {
		return err
	}
	store.Set(types.ApplicationLinkKey(user, link.Data.Application, link.Data.Username), linkBz)

	// Store the client request
	clientRequest := types.NewClientRequest(user, link.Data.Application, link.Data.Username)
	requestBz, err := k.cdc.MarshalBinaryBare(&clientRequest)
	if err != nil {
		return err
	}
	store.Set(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID), requestBz)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypesApplicationLinkSaved,
			sdk.NewAttribute(types.AttributeKeyUser, user),
			sdk.NewAttribute(types.AttributeKeyApplicationName, link.Data.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, link.Data.Username),
		),
	)

	return nil
}

// GetApplicationLink returns the link for the given user, application and username.
// If the link is not found returns an error instead
func (k Keeper) GetApplicationLink(ctx sdk.Context, user, application, username string) (types.ApplicationLink, error) {
	store := ctx.KVStore(k.storeKey)

	linkKey := types.ApplicationLinkKey(user, application, username)
	if !store.Has(linkKey) {
		return types.ApplicationLink{}, sdkerrors.Wrap(sdkerrors.ErrNotFound, "link not found")
	}

	var link types.ApplicationLink
	err := k.cdc.UnmarshalBinaryBare(store.Get(linkKey), &link)
	if err != nil {
		return types.ApplicationLink{}, err
	}

	return link, nil
}

// GetApplicationLinkByClientID returns the application link and user given a specific client id
func (k Keeper) GetApplicationLinkByClientID(ctx sdk.Context, clientID string) (string, types.ApplicationLink, error) {
	store := ctx.KVStore(k.storeKey)

	// Get the client request using the client id
	clientIDKey := types.ApplicationLinkClientIDKey(clientID)
	if !store.Has(clientIDKey) {
		return "", types.ApplicationLink{},
			sdkerrors.Wrapf(sdkerrors.ErrNotFound, "link for client id %s not found", clientID)
	}

	var clientRequest types.ClientRequest
	err := k.cdc.UnmarshalBinaryBare(store.Get(clientIDKey), &clientRequest)
	if err != nil {
		return "", types.ApplicationLink{}, err
	}

	link, err := k.GetApplicationLink(ctx, clientRequest.User, clientRequest.Application, clientRequest.Username)
	if err != nil {
		return "", types.ApplicationLink{}, err
	}

	return clientRequest.User, link, nil
}

// DeleteApplicationLink removes the application link associated to the given user,
// for the given application and username
func (k Keeper) DeleteApplicationLink(ctx sdk.Context, user string, application, username string) error {
	store := ctx.KVStore(k.storeKey)

	// Get the link to obtain the client id
	link, err := k.GetApplicationLink(ctx, user, application, username)
	if err != nil {
		return err
	}

	// Delete the client request
	store.Delete(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID))

	// Delete the link object
	store.Delete(types.ApplicationLinkKey(user, application, username))

	return nil
}
