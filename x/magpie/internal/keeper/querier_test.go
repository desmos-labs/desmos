package keeper_test

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

var request abci.RequestQuery

// ----------------------------------
// --- Sessions
// ----------------------------------

func (suite *KeeperTestSuite) Test_querySession_InvalidIdReturnsError() {
	tests := []struct {
		name          string
		storedSession types.Session
		query         []string
		expErr        error
		expRes        types.Session
	}{
		{
			name:   "Not found session returns error",
			query:  []string{keeper.QuerySessions, types.SessionID(50).String()},
			expErr: sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "session with id 50 not found"),
		},
		{
			name:          "Existing session is returned",
			storedSession: suite.testData.session,
			query:         []string{keeper.QuerySessions, suite.testData.session.SessionID.String()},
			expRes:        suite.testData.session,
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
			suite.SetupTest() // reset
			if !(types.Session{}).Equals(test.storedSession) {
				suite.keeper.SaveSession(suite.ctx, test.storedSession)
			}

			querier := keeper.NewQuerier(suite.keeper)
			result, err := querier(suite.ctx, test.query, request)

			if result != nil {
				suite.Nil(err)

				expectedIndented, err := codec.MarshalJSONIndent(suite.keeper.Cdc, &test.expRes)
				suite.NoError(err)

				suite.Equal(string(expectedIndented), string(result))
			}

			if result == nil {
				suite.NotNil(err)
				suite.Equal(test.expErr.Error(), err.Error())
				suite.Nil(result)
			}
		})
	}
}
