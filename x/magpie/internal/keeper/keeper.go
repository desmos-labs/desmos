package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	StoreKey sdk.StoreKey // Unexposed key to access store from sdk.Context
	Cdc      *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		StoreKey: storeKey,
		Cdc:      cdc,
	}
}

// -------------
// --- Sessions
// -------------

func (k Keeper) getSessionStoreKey(id types.SessionID) []byte {
	return []byte(types.SessionStorePrefix + id.String())
}

// GetLastLikeId returns the last like id that has been used
func (k Keeper) GetLastSessionID(ctx sdk.Context) types.SessionID {
	store := ctx.KVStore(k.StoreKey)
	if !store.Has([]byte(types.LastSessionIDStoreKey)) {
		return types.SessionID(0)
	}

	var id types.SessionID
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIDStoreKey)), &id)
	return id
}

// SetLastSessionID allows to set the last used like id
func (k Keeper) SetLastSessionID(ctx sdk.Context, id types.SessionID) {
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.LastSessionIDStoreKey), k.Cdc.MustMarshalBinaryBare(&id))
}

// SaveSession allows to save a session inside the given context.
// It assumes the given session has already been validated.
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) sdk.Error {
	// Save the session
	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSessionStoreKey(session.SessionID), k.Cdc.MustMarshalBinaryBare(session))

	// Update the last used session id
	k.SetLastSessionID(ctx, session.SessionID)

	return nil
}

// GetSession returns the session having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id types.SessionID) (session types.Session, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getSessionStoreKey(id)
	if !store.Has(key) {
		return types.Session{}, false
	}

	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &session)
	return session, true
}

// GetSessions returns the list of all the sessions present inside the current context
func (k Keeper) GetSessions(ctx sdk.Context) types.Sessions {
	store := ctx.KVStore(k.StoreKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.SessionStorePrefix))

	var sessions []types.Session
	for ; iterator.Valid(); iterator.Next() {
		var session types.Session
		k.Cdc.MustUnmarshalBinaryBare(iterator.Value(), &session)
		sessions = append(sessions, session)
	}

	return sessions
}
