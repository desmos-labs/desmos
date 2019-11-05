package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

// --------------
// --- Sessions
// --------------

func defaultSessionID() types.SessionID {
	return types.SessionID(1)
}

func TestKeeper_GetLastSessionId_FirstId(t *testing.T) {
	ctx, k := SetupTestInput()
	assert.Equal(t, types.SessionID(0), k.GetLastSessionID(ctx))
}

func TestKeeper_GetLastSessionId_Existing(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.SessionID{types.SessionID(0), types.SessionID(3), types.SessionID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		store.Set([]byte(types.LastSessionIDStoreKey), k.Cdc.MustMarshalBinaryBare(id))
		assert.Equal(t, id, k.GetLastSessionID(ctx))
	}
}

func TestKeeper_SetLastSessionId(t *testing.T) {
	ctx, k := SetupTestInput()

	ids := []types.SessionID{types.SessionID(0), types.SessionID(3), types.SessionID(18446744073709551615)}

	store := ctx.KVStore(k.StoreKey)
	for _, id := range ids {
		k.SetLastSessionID(ctx, id)
		var stored types.SessionID
		k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIDStoreKey)), &stored)
		assert.Equal(t, id, stored)
	}
}

func TestKeeper_CreateSession_ExistingId(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionID()}

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.SessionStorePrefix+session.SessionID.String()), k.Cdc.MustMarshalBinaryBare(&session))

	err := k.CreateSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Session with id 1 already exists")
}

func TestKeeper_CreatePost_ValidSession(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionID(), Owner: testOwner}
	err := k.CreateSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.SessionStorePrefix+session.SessionID.String())), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_SaveSession_EmptyOwner(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{SessionID: defaultSessionID()}
	err := k.SaveSession(ctx, session)
	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Owner address cannot be empty")
}

func TestKeeper_SaveSession_ValidSession(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{Owner: testOwner, SessionID: defaultSessionID()}

	err := k.SaveSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.SessionStorePrefix+session.SessionID.String())), &stored)
	assert.Equal(t, session, stored)
}

func TestKeeper_GetSession_NonExistent(t *testing.T) {
	ctx, k := SetupTestInput()

	_, found := k.GetSession(ctx, defaultSessionID())
	assert.False(t, found)
}

func TestKeeper_GetSession_Existent(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{Owner: testOwner, SessionID: defaultSessionID()}

	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.SessionStorePrefix+session.SessionID.String()), k.Cdc.MustMarshalBinaryBare(&session))

	stored, found := k.GetSession(ctx, session.SessionID)
	assert.True(t, found)
	assert.Equal(t, session, stored)
}
