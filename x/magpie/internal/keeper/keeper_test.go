package keeper_test

import (
	"fmt"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"math"
)

func (suite *KeeperTestSuite) TestKeeper_SetDefaultSessionLength() {
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

		suite.Run(fmt.Sprintf("Default session length: %d", test.length), func() {
			err := suite.keeper.SetDefaultSessionLength(suite.ctx, test.length)

			if test.expErr == nil {
				suite.NoError(err)
				var stored int64
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.SessionLengthKey), &stored)
				suite.Equal(test.length, stored)
			}

			if test.expErr != nil {
				suite.Equal(test.expErr, err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDefaultSessionLength() {
	tests := []int64{0, 1, 2, math.MaxInt64}

	for _, length := range tests {
		length := length
		suite.Run(fmt.Sprintf("Get default session length: %d", length), func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)
			if length != 0 {
				store.Set(types.SessionLengthKey, suite.keeper.Cdc.MustMarshalBinaryBare(&length))
			}

			recovered := suite.keeper.GetDefaultSessionLength(suite.ctx)
			suite.Equal(length, recovered)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetLastSessionID() {
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
		suite.Run(test.name, func() {
			if test.existingID.Valid() {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.LastSessionIDStoreKey, suite.keeper.Cdc.MustMarshalBinaryBare(test.existingID))
			}

			suite.Equal(test.expID, suite.keeper.GetLastSessionID(suite.ctx))
		})
	}

	suite.Equal(types.SessionID(0), suite.keeper.GetLastSessionID(suite.ctx))
}

func (suite *KeeperTestSuite) TestKeeper_SetLastSessionID() {
	tests := []struct {
		name string
		id   types.SessionID
	}{
		{name: "set id session to 0", id: types.SessionID(0)},
		{name: "set id session to 3", id: types.SessionID(3)},
		{name: "set id session to num", id: types.SessionID(18446744073709551615)},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.StoreKey)

			suite.keeper.SetLastSessionID(suite.ctx, test.id)

			var stored types.SessionID
			suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &stored)
			suite.Equal(test.id, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveSession() {
	session := types.Session{Owner: testOwner, SessionID: types.SessionID(1)}

	suite.keeper.SaveSession(suite.ctx, session)

	var stored types.Session
	store := suite.ctx.KVStore(suite.keeper.StoreKey)
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(session.SessionID)), &stored)
	suite.Equal(session, stored)

	var storedLastID types.SessionID
	suite.keeper.Cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &storedLastID)
	suite.Equal(session.SessionID, storedLastID)
}

func (suite *KeeperTestSuite) TestKeeper_GetSession() {
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
		suite.Run(test.name, func() {
			if !(types.Session{}).Equals(test.storedSession) {
				store := suite.ctx.KVStore(suite.keeper.StoreKey)
				store.Set(types.SessionStoreKey(test.id), suite.keeper.Cdc.MustMarshalBinaryBare(&test.storedSession))
			}

			result, found := suite.keeper.GetSession(suite.ctx, types.SessionID(1))
			suite.Equal(test.expSession, result)
			suite.Equal(test.expFound, found)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetSessions() {
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
		suite.Run(test.name, func() {
			for _, s := range test.storedSessions {
				suite.keeper.SaveSession(suite.ctx, s)
			}

			sessions := suite.keeper.GetSessions(suite.ctx)
			suite.True(test.expSessions.Equals(sessions))
		})
	}
}
