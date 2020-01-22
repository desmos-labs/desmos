package keeper_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestKeeper_SetDefaultSessionLength(t *testing.T) {
	tests := []int64{-1, 0, 1, math.MaxInt32}

	for _, length := range tests {
		length := length

		t.Run(fmt.Sprintf("Default session length: %d", length), func(t *testing.T) {
			ctx, k := SetupTestInput()
			k.SetDefaultSessionLength(ctx, length)

			var stored int64
			store := ctx.KVStore(k.StoreKey)
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.SessionLengthKey)), &stored)
			assert.Equal(t, length, stored)
		})
	}
}

func TestKeeper_GetDefaultSessionLength(t *testing.T) {
	tests := []int64{-1, -2, 0, 1, 2}

	for _, length := range tests {
		length := length
		t.Run(fmt.Sprintf("Get default session length: %d", length), func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if length != 0 {
				store.Set([]byte(types.SessionLengthKey), k.Cdc.MustMarshalBinaryBare(&length))
			}

			recovered := k.GetDefaultSessionLength(ctx)
			assert.Equal(t, length, recovered)
		})
	}
}

func TestKeeper_GetLastSessionID(t *testing.T) {
	tests := []struct {
		name       string
		existingID types.SessionID
		expID      types.SessionID
	}{
		{
			name:  "First ID is returned properly",
			expID: types.SessionID(0),
		},
		{
			name:       "Existing ID is returned properly",
			existingID: types.SessionID(18446744073709551615),
			expID:      types.SessionID(18446744073709551615),
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if test.existingID.Valid() {
				store := ctx.KVStore(k.StoreKey)
				store.Set([]byte(types.LastSessionIDStoreKey), k.Cdc.MustMarshalBinaryBare(test.existingID))
			}

			assert.Equal(t, test.expID, k.GetLastSessionID(ctx))
		})
	}

	ctx, k := SetupTestInput()
	assert.Equal(t, types.SessionID(0), k.GetLastSessionID(ctx))
}

func TestKeeper_SetLastSessionID(t *testing.T) {
	tests := []struct {
		id types.SessionID
	}{
		{id: types.SessionID(0)},
		{id: types.SessionID(3)},
		{id: types.SessionID(18446744073709551615)},
	}

	for _, test := range tests {
		test := test
		t.Run(t.Name(), func(t *testing.T) {
			ctx, k := SetupTestInput()
			store := ctx.KVStore(k.StoreKey)

			k.SetLastSessionID(ctx, test.id)

			var stored types.SessionID
			k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIDStoreKey)), &stored)
			assert.Equal(t, test.id, stored)
		})
	}
}

func TestKeeper_SaveSession(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{Owner: testOwner, SessionID: types.SessionID(1)}

	err := k.SaveSession(ctx, session)
	assert.NoError(t, err)

	var stored types.Session
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.SessionStorePrefix+session.SessionID.String())), &stored)
	assert.Equal(t, session, stored)

	var storedLastID types.SessionID
	k.Cdc.MustUnmarshalBinaryBare(store.Get([]byte(types.LastSessionIDStoreKey)), &storedLastID)
	assert.Equal(t, session.SessionID, storedLastID)
}

func TestKeeper_GetSession(t *testing.T) {
	tests := []struct {
		name          string
		storedSession types.Session
		id            types.SessionID
		expFound      bool
		expSession    types.Session
	}{
		{
			name:       "Non existent session",
			id:         types.SessionID(0),
			expFound:   false,
			expSession: types.Session{},
		},
		{
			name:          "Valid session is returned",
			storedSession: types.Session{Owner: testOwner, SessionID: types.SessionID(1)},
			id:            types.SessionID(1),
			expFound:      true,
			expSession:    types.Session{Owner: testOwner, SessionID: types.SessionID(1)},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if !(types.Session{}).Equals(test.storedSession) {
				store := ctx.KVStore(k.StoreKey)
				store.Set([]byte(types.SessionStorePrefix+test.id.String()), k.Cdc.MustMarshalBinaryBare(&test.storedSession))
			}

			result, found := k.GetSession(ctx, types.SessionID(1))
			assert.Equal(t, test.expSession, result)
			assert.Equal(t, test.expFound, found)
		})
	}
}

func TestKeeper_GetSessions(t *testing.T) {
	tests := []struct {
		name           string
		storedSessions types.Sessions
		expSessions    types.Sessions
	}{
		{
			name:           "Empty slice",
			storedSessions: types.Sessions{},
			expSessions:    types.Sessions{},
		},
		{
			name: "Non empty, non double items",
			storedSessions: types.Sessions{
				types.Session{SessionID: types.SessionID(1)},
				types.Session{SessionID: types.SessionID(2)},
			},
			expSessions: types.Sessions{
				types.Session{SessionID: types.SessionID(1)},
				types.Session{SessionID: types.SessionID(2)},
			},
		},
		{
			name: "Non empty, double items",
			storedSessions: types.Sessions{
				types.Session{SessionID: types.SessionID(1)},
				types.Session{SessionID: types.SessionID(1)},
			},
			expSessions: types.Sessions{
				types.Session{SessionID: types.SessionID(1)},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			for _, s := range test.storedSessions {
				_ = k.SaveSession(ctx, s)
			}

			sessions := k.GetSessions(ctx)
			assert.True(t, test.expSessions.Equals(sessions))
		})
	}
}
