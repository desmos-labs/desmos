package keeper

import (
	"fmt"

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

// CreateSession allows to create a new testSession checking that no other testSession
// with the same id already exist
func (k Keeper) CreateSession(ctx sdk.Context, session types.Session) sdk.Error {
	// Check for any previously existing testSession
	if _, found := k.GetSession(ctx, session.SessionID); found {
		return sdk.ErrUnknownRequest(fmt.Sprintf("Session with id %s already exists", session.SessionID))
	}

	return k.SaveSession(ctx, session)
}

// SaveSession allows to save a testSession inside the given context
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) sdk.Error {
	if session.Owner.Empty() {
		return sdk.ErrInvalidAddress("Owner address cannot be empty")
	}

	// Save the testSession
	store := ctx.KVStore(k.StoreKey)
	store.Set(k.getSessionStoreKey(session.SessionID), k.Cdc.MustMarshalBinaryBare(session))

	// Update the last used testSession id
	k.SetLastSessionID(ctx, session.SessionID)

	return nil
}

// GetSession returns the testSession having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id types.SessionID) (session types.Session, found bool) {
	store := ctx.KVStore(k.StoreKey)

	key := k.getSessionStoreKey(id)
	if !store.Has(key) {
		return types.NewSession(), false
	}

	bz := store.Get(key)
	k.Cdc.MustUnmarshalBinaryBare(bz, &session)
	return session, true
}

// GetSessions returns the list of all the sessions present inside the current context
func (k Keeper) GetSessions(ctx sdk.Context) []types.Session {
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
