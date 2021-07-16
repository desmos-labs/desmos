package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// Connections are stored using three keys:
// 1. UserApplicationLinkKey (user + application + username)  		-> types.ApplicationLink
// 2. ApplicationLinkClientIDKey (client_id)                  		-> UserApplicationLinkKey
// 3. ApplicationLinkExpirationKey (expiredBlockHeight + client_id) -> ApplicationLinkClientIDKey
//
// This allows to get connections by client id as well as by app + username quickly

// SaveApplicationLink stores the given connection replacing any existing one for the same user and application
func (k Keeper) SaveApplicationLink(ctx sdk.Context, link types.ApplicationLink) error {
	if !k.HasProfile(ctx, link.User) {
		return sdkerrors.Wrapf(types.ErrProfileNotFound, "a profile is required to link an application")
	}

	params := k.GetParams(ctx)

	// Get the keys
	userApplicationLinkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
	applicationLinkClientIDKey := types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID)
	applicationLinkExpirationKey := types.ExpiringApplicationLinkKey(ctx.BlockHeight()+params.ApplicationLink.ExpiryInterval, link.OracleRequest.ClientID)

	// Store the data
	store := ctx.KVStore(k.storeKey)
	store.Set(userApplicationLinkKey, types.MustMarshalApplicationLink(k.cdc, link))
	store.Set(applicationLinkClientIDKey, userApplicationLinkKey)
	store.Set(applicationLinkExpirationKey, applicationLinkClientIDKey)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypesApplicationLinkSaved,
			sdk.NewAttribute(types.AttributeKeyUser, link.User),
			sdk.NewAttribute(types.AttributeKeyApplicationName, link.Data.Application),
			sdk.NewAttribute(types.AttributeKeyApplicationUsername, link.Data.Username),
		),
	)

	return nil
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
	err := k.cdc.UnmarshalBinaryBare(store.Get(userApplicationLinkKey), &link)
	if err != nil {
		return types.ApplicationLink{}, false, err
	}

	return link, true, nil
}

// GetApplicationLinkByClientID returns the application link and user given a specific client id
func (k Keeper) GetApplicationLinkByClientID(ctx sdk.Context, clientID string) (types.ApplicationLink, error) {
	store := ctx.KVStore(k.storeKey)

	// Get the client request using the client id
	clientIDKey := types.ApplicationLinkClientIDKey(clientID)
	if !store.Has(clientIDKey) {
		return types.ApplicationLink{},
			sdkerrors.Wrapf(sdkerrors.ErrNotFound, "link for client id %s not found", clientID)
	}

	// Get the link key
	applicationLinkKey := store.Get(clientIDKey)

	// Read the link
	var link types.ApplicationLink
	err := k.cdc.UnmarshalBinaryBare(store.Get(applicationLinkKey), &link)
	if err != nil {
		return types.ApplicationLink{}, sdkerrors.Wrap(err, "error while reading application link")
	}

	return link, nil
}

// DeleteApplicationLink removes the application link associated to the given user,
// for the given application and username
func (k Keeper) DeleteApplicationLink(ctx sdk.Context, user string, application, username string) error {
	// Get the link to obtain the client id
	link, found, err := k.GetApplicationLink(ctx, user, application, username)
	if err != nil {
		return err
	}

	if !found {
		return sdkerrors.Wrap(sdkerrors.ErrNotFound, "application link not found")
	}

	if link.User != user {
		return sdkerrors.Wrap(sdkerrors.ErrorInvalidSigner, "cannot delete the application link of another user")
	}

	// Delete the data
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UserApplicationLinkKey(user, application, username))
	store.Delete(types.ApplicationLinkClientIDKey(link.OracleRequest.ClientID))
	store.Delete(types.ExpiringApplicationLinkKey(link.ExpirationBlockHeight, link.OracleRequest.ClientID))

	return nil
}

// DeleteAllUserApplicationLinks delete all the applications links associated with the given user
func (k Keeper) DeleteAllUserApplicationLinks(ctx sdk.Context, user string) {
	var links []types.ApplicationLink
	k.IterateUserApplicationLinks(ctx, user, func(index int64, link types.ApplicationLink) (stop bool) {
		links = append(links, link)
		return false
	})

	store := ctx.KVStore(k.storeKey)
	for _, link := range links {
		store.Delete(types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username))
	}
}

// UpdateExpiringApplicationLinks update the states of all the expiring application links to be expired
func (k Keeper) UpdateExpiringApplicationLinks(ctx sdk.Context) {
	k.IterateExpiringApplicationLinks(ctx, ctx.BlockHeight(), func(_ int64, link types.ApplicationLink) (stop bool) {
		store := ctx.KVStore(k.storeKey)
		link.State = types.AppLinkStateVerificationExpired
		userApplicationLinkKey := types.UserApplicationLinkKey(link.User, link.Data.Application, link.Data.Username)
		store.Set(userApplicationLinkKey, types.MustMarshalApplicationLink(k.cdc, link))

		store.Delete(types.ExpiringApplicationLinkKey(ctx.BlockHeight(), link.OracleRequest.ClientID))
		return false
	})
}
