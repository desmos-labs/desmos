package keeper_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/types"
	"github.com/stretchr/testify/require"
)

func TestKeeper_SetDefaultSessionLength(t *testing.T) {
	tests := []struct {
		length int64
		expErr error
	}{
		{
			length: -1,
			expErr: fmt.Errorf("cannot set -1 as default session length"),
		},
		{
			length: 0,
			expErr: fmt.Errorf("cannot set 0 as default session length"),
		},
		{
			length: 1,
			expErr: nil,
		},
		{
			length: math.MaxInt64,
			expErr: nil,
		},
	}

	for _, test := range tests {
		test := test

		t.Run(fmt.Sprintf("Default session length: %d", test.length), func(t *testing.T) {
			ctx, k := SetupTestInput()
			err := k.SetDefaultSessionLength(ctx, test.length)

			if test.expErr == nil {
				require.NoError(t, err)
				var stored int64
				store := ctx.KVStore(k.StoreKey)
				k.Cdc.MustUnmarshalBinaryBare(store.Get(types.SessionLengthKey), &stored)
				require.Equal(t, test.length, stored)
			}

			if test.expErr != nil {
				require.Equal(t, test.expErr, err)
			}
		})
	}
}

func TestKeeper_GetDefaultSessionLength(t *testing.T) {
	tests := []int64{0, 1, 2, math.MaxInt64}

	for _, length := range tests {
		length := length
		t.Run(fmt.Sprintf("Get default session length: %d", length), func(t *testing.T) {
			ctx, k := SetupTestInput()

			store := ctx.KVStore(k.StoreKey)
			if length != 0 {
				store.Set(types.SessionLengthKey, k.Cdc.MustMarshalBinaryBare(&length))
			}

			recovered := k.GetDefaultSessionLength(ctx)
			require.Equal(t, length, recovered)
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
				store.Set(types.LastSessionIDStoreKey, k.Cdc.MustMarshalBinaryBare(test.existingID))
			}

			require.Equal(t, test.expID, k.GetLastSessionID(ctx))
		})
	}

	ctx, k := SetupTestInput()
	require.Equal(t, types.SessionID(0), k.GetLastSessionID(ctx))
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
			k.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &stored)
			require.Equal(t, test.id, stored)
		})
	}
}

func TestKeeper_SaveSession(t *testing.T) {
	ctx, k := SetupTestInput()

	session := types.Session{Owner: testOwner, SessionID: types.SessionID(1)}

	k.SaveSession(ctx, session)

	var stored types.Session
	store := ctx.KVStore(k.StoreKey)
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(session.SessionID)), &stored)
	require.Equal(t, session, stored)

	var storedLastID types.SessionID
	k.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &storedLastID)
	require.Equal(t, session.SessionID, storedLastID)
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
				store.Set(types.SessionStoreKey(test.id), k.Cdc.MustMarshalBinaryBare(&test.storedSession))
			}

			result, found := k.GetSession(ctx, types.SessionID(1))
			require.Equal(t, test.expSession, result)
			require.Equal(t, test.expFound, found)
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
				k.SaveSession(ctx, s)
			}

			sessions := k.GetSessions(ctx)
			require.True(t, test.expSessions.Equals(sessions))
		})
	}
}
