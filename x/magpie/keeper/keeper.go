package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// Keeper maintains the link to data storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	storeKey sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc      codec.BinaryMarshaler // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the magpie Keeper
func NewKeeper(cdc codec.BinaryMarshaler, storeKey sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: storeKey,
		cdc:      cdc,
	}
}

// -------------
// --- Params
// -------------

// SetDefaultSessionLength allows to set a default session length for new magpie sessions.
// The specified length is intended to be in number of blocks.
func (k Keeper) SetDefaultSessionLength(ctx sdk.Context, length uint64) error {
	if length < 1 {
		return fmt.Errorf("cannot set %d as default session length", length)
	}

	store := ctx.KVStore(k.storeKey)

	var bz []byte
	binary.LittleEndian.PutUint64(bz, length)
	store.Set(types.SessionLengthKey, bz)
	return nil
}

// GetDefaultSessionLength returns the default session length in number of blocks.
func (k Keeper) GetDefaultSessionLength(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)

	length := uint64(0)
	if store.Has(types.SessionLengthKey) {
		length = binary.LittleEndian.Uint64(store.Get(types.SessionLengthKey))
	}

	return length
}

// -------------
// --- Sessions
// -------------

// GetLastLikeId returns the last like id that has been used
func (k Keeper) GetLastSessionID(ctx sdk.Context) types.SessionID {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.LastSessionIDStoreKey) {
		return types.SessionID{Value: 0}
	}

	var id types.SessionID
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &id)
	return id
}

// SetLastSessionID allows to set the last used like id
func (k Keeper) SetLastSessionID(ctx sdk.Context, id types.SessionID) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LastSessionIDStoreKey, k.cdc.MustMarshalBinaryBare(&id))
}

// SaveSession allows to save a session inside the given context.
// It assumes the given session has already been validated.
func (k Keeper) SaveSession(ctx sdk.Context, session types.Session) {
	// Save the session
	store := ctx.KVStore(k.storeKey)
	store.Set(types.SessionStoreKey(session.SessionId), k.cdc.MustMarshalBinaryBare(&session))

	// Update the last used session id
	k.SetLastSessionID(ctx, session.SessionId)
}

// GetSession returns the session having the specified id
func (k Keeper) GetSession(ctx sdk.Context, id types.SessionID) (session types.Session, found bool) {
	store := ctx.KVStore(k.storeKey)

	key := types.SessionStoreKey(id)
	if !store.Has(key) {
		return types.Session{}, false
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &session)
	return session, true
}

// GetSessions returns the list of all the sessions present inside the current context
func (k Keeper) GetSessions(ctx sdk.Context) types.Sessions {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.SessionStorePrefix)
	defer iterator.Close()

	sessions := make(types.Sessions, 0)
	for ; iterator.Valid(); iterator.Next() {
		var session types.Session
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &session)
		sessions = append(sessions, session)
	}

	return sessions
}
