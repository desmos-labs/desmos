package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// Connections are stored using three keys:
// 1. UserApplicationLinkKey (user + application + username)  -> types.ApplicationLink
// 2. ApplicationLinkClientIDKey (client_id)                  -> UserApplicationLinkKey
// 3. ApplicationLinkOwnerKey (application + username + user) -> 0x01
//
// This allows to get connections by client id as well as by app + username quickly

// SaveApplicationLink stores the given connection replacing any existing one for the same user and application
func (k Keeper) SaveApplicationLink(ctx sdk.Context, link types.ApplicationLink) error {
	if !k.HasProfile(ctx, link.User) {
		return sdkerrors.Wrapf(types.ErrProfileNotFound, "a profile is required to link an application")
	}

	// Store the data
	store := ctx.KVStore(k.storeKey)
	userApplicationLinkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
	store.Set(userApplicationLinkKey, types.MustMarshalApplicationLink(k.cdc, link))
	store.Set(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID), userApplicationLinkKey)
	store.Set(types.ApplicationLinkOwnerKey(link.Data.Application, link.Data.Username, link.User), []byte{0x01})
	applicationLinkExpiringTimeKey := types.ApplicationLinkExpiringTimeKey(link.ExpirationTime, link.OracleRequest.ClientID)
	store.Set(applicationLinkExpiringTimeKey, []byte(link.OracleRequest.ClientID))

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypesApplicationLinkSaved,
			sdk.NewAttribute(types.AttributeKeyUser, link.User),
			sdk.NewAttribute(types.AttributeKeyApplicationName, link.Data.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, link.Data.Username),
		),
	)

	k.AfterApplicationLinkSaved(ctx, link)

	return nil
}

// HasApplicationLink tells whether the given application link exists
func (k Keeper) HasApplicationLink(ctx sdk.Context, user, application, username string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.UserApplicationLinkKey(user, application, username))
}

// GetApplicationLink returns the link for the given application and username.
// If the link is not found returns an error instead.
func (k Keeper) GetApplicationLink(ctx sdk.Context, user, application, username string) (types.ApplicationLink, bool, error) {
	store := ctx.KVStore(k.storeKey)

	// Check to see if the key exists
	userApplicationLinkKey := types.UserApplicationLinkKey(user, application, username)
	if !store.Has(userApplicationLinkKey) {
		return types.ApplicationLink{}, false, nil
	}

	var link types.ApplicationLink
	err := k.cdc.Unmarshal(store.Get(userApplicationLinkKey), &link)
	if err != nil {
		return types.ApplicationLink{}, false, err
	}

	return link, true, nil
}

// GetApplicationLinkByClientID returns the application link and user given a specific client id.
// If the link is not found, returns false instead.
func (k Keeper) GetApplicationLinkByClientID(ctx sdk.Context, clientID string) (types.ApplicationLink, bool, error) {
	store := ctx.KVStore(k.storeKey)

	// Get the client request using the client id
	clientIDKey := types.ApplicationLinkClientIDKey(clientID)
	if !store.Has(clientIDKey) {
		return types.ApplicationLink{}, false, nil
	}

	// Get the link key
	applicationLinkKey := store.Get(clientIDKey)

	// Read the link
	var link types.ApplicationLink
	err := k.cdc.Unmarshal(store.Get(applicationLinkKey), &link)
	if err != nil {
		return types.ApplicationLink{}, true, sdkerrors.Wrap(err, "error while reading application link")
	}

	return link, true, nil
}

// DeleteApplicationLink removes the application link associated to the given user,
// for the given application and username
func (k Keeper) DeleteApplicationLink(ctx sdk.Context, appLink types.ApplicationLink) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserApplicationLinkKey(appLink.User, appLink.Data.Application, appLink.Data.Username))
	store.Delete(types.ApplicationLinkClientIDKey(appLink.OracleRequest.ClientID))
	store.Delete(types.ApplicationLinkOwnerKey(appLink.Data.Application, appLink.Data.Username, appLink.User))
	store.Delete(types.ApplicationLinkExpiringTimeKey(appLink.ExpirationTime, appLink.OracleRequest.ClientID))

	k.AfterApplicationLinkDeleted(ctx, appLink)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeApplicationLinkDeleted,
			sdk.NewAttribute(types.AttributeKeyUser, appLink.User),
			sdk.NewAttribute(types.AttributeKeyApplicationName, appLink.Data.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, appLink.Data.Username),
			sdk.NewAttribute(types.AttributeKeyApplicationLinkExpirationTime, appLink.ExpirationTime.Format(time.RFC3339)),
		),
	)
}

// DeleteAllUserApplicationLinks delete all the applications links associated with the given user
func (k Keeper) DeleteAllUserApplicationLinks(ctx sdk.Context, user string) {
	var links []types.ApplicationLink
	k.IterateUserApplicationLinks(ctx, user, func(link types.ApplicationLink) (stop bool) {
		links = append(links, link)
		return false
	})

	for _, link := range links {
		k.DeleteApplicationLink(ctx, link)
	}
}
