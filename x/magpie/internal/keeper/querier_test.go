package keeper_test

import (
	"testing"

	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/assert"
	abci "github.com/tendermint/tendermint/abci/types"
)

var request abci.RequestQuery

// ----------------------------------
// --- Sessions
// ----------------------------------

func Test_querySession_InvalidIdReturnsError(t *testing.T) {
	tests := []struct {
		name          string
		storedSession types.Session
		query         []string
		expErr        string
		expRes        types.Session
	}{
		{
			name:   "Not found session returns error",
			query:  []string{keeper.QuerySessions, types.SessionID(50).String()},
			expErr: "Session with id 50 not found",
		},
		{
			name:          "Existing session is returned",
			storedSession: testSession,
			query:         []string{keeper.QuerySessions, testSession.SessionID.String()},
			expRes:        testSession,
		},
		{
			name:   "Invalid id",
			query:  []string{keeper.QuerySessions, "invalid-id"},
			expErr: "Invalid session id: invalid-id",
		},
		{
			name:   "Unknown endpoint",
			query:  []string{"endpoint"},
			expErr: "Unknown magpie query endpoint",
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if !(types.Session{}).Equals(test.storedSession) {
				_ = k.SaveSession(ctx, test.storedSession)
			}

			querier := keeper.NewQuerier(k)
			res, err := querier(ctx, test.query, request)

			if len(test.expErr) != 0 {
				assert.Error(t, err)
				assert.Contains(t, err.Result().Log, test.expErr)

				assert.Nil(t, res)
			}

			if !(types.Session{}).Equals(test.expRes) {
				assert.NoError(t, err)

				var returned types.Session
				k.Cdc.MustUnmarshalJSON(res, &returned)
				assert.Equal(t, test.expRes, returned)
			}
		})
	}
}
