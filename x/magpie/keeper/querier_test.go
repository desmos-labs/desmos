package keeper_test

import (
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/desmos-labs/desmos/x/magpie/keeper"
	"github.com/desmos-labs/desmos/x/magpie/types"
)

var request abci.RequestQuery

// ----------------------------------
// --- Sessions
// ----------------------------------

func (suite *KeeperTestSuite) TestQuerier_QuerySessions() {
	tests := []struct {
		name              string
		storedSessions    []types.Session
		query             []string
		expErr            error
		expStoredSessions []types.Session
	}{
		{
			name:   "Not found session returns expError",
			query:  []string{keeper.QuerySessions, "50"},
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "session with id 50 not found"),
		},
		{
			name: "Existing session is returned",
			storedSessions: []types.Session{
				suite.testData.session,
			},
			query: []string{keeper.QuerySessions, fmt.Sprint(suite.testData.session.SessionId.Value)},
			expStoredSessions: []types.Session{
				suite.testData.session,
			},
		},
		{
			name:   "Invalid id",
			query:  []string{keeper.QuerySessions, "invalid-id"},
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid session id: invalid-id"),
		},
		{
			name:   "Unknown endpoint",
			query:  []string{"endpoint"},
			expErr: sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown magpie query endpoint"),
		},
	}

	for _, test := range tests {
		test := test
		suite.Run(test.name, func() {
			suite.SetupTest()

			for _, session := range test.storedSessions {
				suite.keeper.SaveSession(suite.ctx, session)
			}

			querier := keeper.NewQuerier(suite.keeper, suite.legacyAmino)
			_, err := querier(suite.ctx, test.query, request)
			suite.RequireErrorsEqual(test.expErr, err)

			stored := suite.keeper.GetSessions(suite.ctx)
			suite.Require().Equal(test.expStoredSessions, stored)
		})
	}
}
