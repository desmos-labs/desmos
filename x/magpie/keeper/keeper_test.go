package keeper_test

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetDefaultSessionLength() {
	tests := []struct {
		length uint64
		expErr error
	}{
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
			suite.SetupTest() // reset
			err := suite.keeper.SetDefaultSessionLength(suite.ctx, test.length)

			if test.expErr == nil {
				suite.Require().NoError(err)

				store := suite.ctx.KVStore(suite.keeper.storeKey)
				stored := binary.LittleEndian.Uint64(store.Get(types.SessionLengthKey))
				suite.Require().Equal(test.length, stored)
			}

			if test.expErr != nil {
				suite.Require().Equal(test.expErr, err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDefaultSessionLength() {
	tests := []uint64{0, 1, 2, math.MaxUint64}

	for _, length := range tests {
		length := length
		suite.Run(fmt.Sprintf("Get default session length: %d", length), func() {
			suite.SetupTest() // reset
			store := suite.ctx.KVStore(suite.keeper.storeKey)
			if length != 0 {
				var bz []byte
				binary.LittleEndian.PutUint64(bz, length)
				store.Set(types.SessionLengthKey, bz)
			}

			recovered := suite.keeper.GetDefaultSessionLength(suite.ctx)
			suite.Require().Equal(length, recovered)
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
			expID: types.SessionID(uint64(0)),
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
			suite.SetupTest() // reset
			if test.existingID.Valid() {
				store := suite.ctx.KVStore(suite.keeper.storeKey)

				var bz []byte
				binary.LittleEndian.PutUint64(bz, test.existingID.Value)
				store.Set(types.LastSessionIDStoreKey, bz)
			}

			suite.Require().Equal(test.expID, suite.keeper.GetLastSessionID(suite.ctx))
		})
	}

	suite.SetupTest() // reset
	suite.Require().Equal(types.SessionID(uint64(0)), suite.keeper.GetLastSessionID(suite.ctx))
}

func (suite *KeeperTestSuite) TestKeeper_SetLastSessionID() {
	tests := []struct {
		name string
		id   types.SessionID
	}{
		{name: "set id session to 0", id: types.SessionID(uint64(0))},
		{name: "set id session to 3", id: types.SessionID(uint64(3))},
		{name: "set id session to num", id: types.SessionID(18446744073709551615)},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.keeper.storeKey)

			suite.keeper.SetLastSessionID(suite.ctx, test.id)

			var stored types.SessionID
			suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &stored)
			suite.Require().Equal(test.id, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveSession() {
	session := types.Session{Owner: suite.testData.owner, SessionId: types.SessionID(uint64(1))}

	suite.keeper.SaveSession(suite.ctx, session)

	var stored types.Session
	store := suite.ctx.KVStore(suite.keeper.storeKey)
	suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(session.SessionId)), &stored)
	suite.Require().Equal(session, stored)

	var storedLastID types.SessionID
	suite.keeper.cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &storedLastID)
	suite.Require().Equal(session.SessionId, storedLastID)
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
			id:         types.SessionID(uint64(0)),
			expFound:   false,
			expSession: types.Session{},
		},
		{
			name:          "Valid session is returned",
			storedSession: types.Session{Owner: suite.testData.owner, SessionId: types.SessionID(uint64(1))},
			id:            types.SessionID(uint64(1)),
			expFound:      true,
			expSession:    types.Session{Owner: suite.testData.owner, SessionId: types.SessionID(uint64(1))},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			empty := types.Session{}
			if !empty.Equal(test.storedSession) {
				store := suite.ctx.KVStore(suite.keeper.storeKey)
				store.Set(types.SessionStoreKey(test.id), suite.keeper.cdc.MustMarshalBinaryBare(&test.storedSession))
			}

			result, found := suite.keeper.GetSession(suite.ctx, types.SessionID(uint64(1)))
			suite.Require().Equal(test.expSession, result)
			suite.Require().Equal(test.expFound, found)
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
				types.Session{SessionId: types.SessionID(uint64(1))},
				types.Session{SessionId: types.SessionID(uint64(2))},
			},
			expSessions: types.Sessions{
				types.Session{SessionId: types.SessionID(uint64(1))},
				types.Session{SessionId: types.SessionID(uint64(2))},
			},
		},
		{
			name: "Non empty, double items",
			storedSessions: types.Sessions{
				types.Session{SessionId: types.SessionID(uint64(1))},
				types.Session{SessionId: types.SessionID(uint64(1))},
			},
			expSessions: types.Sessions{
				types.Session{SessionId: types.SessionID(uint64(1))},
			},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest() // reset
			for _, s := range test.storedSessions {
				suite.keeper.SaveSession(suite.ctx, s)
			}

			sessions := suite.keeper.GetSessions(suite.ctx)
			suite.Require().Equal(test.expSessions, sessions)

		})
	}
}
