package keeper_test

import (
	"fmt"
	"math"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

func (suite *KeeperTestSuite) TestKeeper_SetDefaultSessionLength() {
	tests := []struct {
		length           uint64
		expSessionLength uint64
		expErr           error
	}{
		{
			length:           0,
			expSessionLength: 0,
			expErr:           sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "cannot set 0 as default session length"),
		},
		{
			length:           1,
			expSessionLength: 1,
			expErr:           nil,
		},
		{
			length:           math.MaxInt64,
			expSessionLength: math.MaxInt64,
			expErr:           nil,
		},
	}

	for _, test := range tests {
		test := test

		suite.Run(fmt.Sprintf("Default session length: %d", test.length), func() {
			suite.SetupTest()

			err := suite.keeper.SetDefaultSessionLength(suite.ctx, test.length)
			suite.RequireErrorsEqual(test.expErr, err)

			stored := suite.keeper.GetDefaultSessionLength(suite.ctx)
			suite.Require().Equal(test.expSessionLength, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetDefaultSessionLength() {
	tests := []uint64{0, 1, 2, math.MaxUint64}

	for _, length := range tests {
		length := length
		suite.SetupTest()
		suite.Run(fmt.Sprintf("Get default session length: %d", length), func() {
			if length != 0 {
				err := suite.keeper.SetDefaultSessionLength(suite.ctx, length)
				suite.Require().NoError(err)
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
			expID: types.NewSessionID(0),
		},
		{
			name:       "Existing ID is returned properly",
			existingID: types.NewSessionID(18446744073709551615),
			expID:      types.NewSessionID(18446744073709551615),
		},
	}

	for _, test := range tests {
		test := test
		suite.SetupTest()
		suite.Run(test.name, func() {
			if test.existingID.Valid() {
				suite.keeper.SetLastSessionID(suite.ctx, test.existingID)
			}

			suite.Require().Equal(test.expID, suite.keeper.GetLastSessionID(suite.ctx))
		})
	}

	suite.SetupTest()
	suite.Require().Equal(types.NewSessionID(0), suite.keeper.GetLastSessionID(suite.ctx))
}

func (suite *KeeperTestSuite) TestKeeper_SetLastSessionID() {
	tests := []struct {
		name string
		id   types.SessionID
	}{
		{name: "set id session to 0", id: types.NewSessionID(0)},
		{name: "set id session to 3", id: types.NewSessionID(3)},
		{name: "set id session to num", id: types.NewSessionID(18446744073709551615)},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			store := suite.ctx.KVStore(suite.storeKey)

			suite.keeper.SetLastSessionID(suite.ctx, test.id)

			var stored types.SessionID
			suite.cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &stored)
			suite.Require().Equal(test.id, stored)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_SaveSession() {
	session := types.Session{Owner: suite.testData.owner, SessionId: types.NewSessionID(1)}

	suite.keeper.SaveSession(suite.ctx, session)

	var stored types.Session
	store := suite.ctx.KVStore(suite.storeKey)
	suite.cdc.MustUnmarshalBinaryBare(store.Get(types.SessionStoreKey(session.SessionId)), &stored)
	suite.Require().Equal(session, stored)

	var storedLastID types.SessionID
	suite.cdc.MustUnmarshalBinaryBare(store.Get(types.LastSessionIDStoreKey), &storedLastID)
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
			id:         types.NewSessionID(0),
			expFound:   false,
			expSession: types.Session{},
		},
		{
			name:          "Valid session is returned",
			storedSession: types.Session{Owner: suite.testData.owner, SessionId: types.NewSessionID(1)},
			id:            types.NewSessionID(1),
			expFound:      true,
			expSession:    types.Session{Owner: suite.testData.owner, SessionId: types.NewSessionID(1)},
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			empty := types.Session{}
			if !empty.Equal(test.storedSession) {
				store := suite.ctx.KVStore(suite.storeKey)
				store.Set(types.SessionStoreKey(test.id), suite.cdc.MustMarshalBinaryBare(&test.storedSession))
			}

			result, found := suite.keeper.GetSession(suite.ctx, types.NewSessionID(1))
			suite.Require().Equal(test.expSession, result)
			suite.Require().Equal(test.expFound, found)
		})
	}
}

func (suite *KeeperTestSuite) TestKeeper_GetSessions() {
	tests := []struct {
		name           string
		storedSessions []types.Session
		expSessions    []types.Session
	}{
		{
			name:           "Empty slice",
			storedSessions: nil,
			expSessions:    nil,
		},
		{
			name: "Non empty, non double items",
			storedSessions: []types.Session{
				{SessionId: types.NewSessionID(1)},
				{SessionId: types.NewSessionID(2)},
			},
			expSessions: []types.Session{
				{SessionId: types.NewSessionID(1)},
				{SessionId: types.NewSessionID(2)},
			},
		},
		{
			name: "Non empty, double items",
			storedSessions: []types.Session{
				{SessionId: types.NewSessionID(1)},
				{SessionId: types.NewSessionID(1)},
			},
			expSessions: []types.Session{
				{SessionId: types.NewSessionID(1)},
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
