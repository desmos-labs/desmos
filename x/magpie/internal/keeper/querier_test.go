package keeper_test

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/desmos-labs/desmos/x/magpie/internal/keeper"
	"github.com/desmos-labs/desmos/x/magpie/internal/types"
	"github.com/stretchr/testify/require"
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
			storedSession: testSession,
			query:         []string{keeper.QuerySessions, testSession.SessionID.String()},
			expRes:        testSession,
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
		t.Run(test.name, func(t *testing.T) {
			ctx, k := SetupTestInput()

			if !(types.Session{}).Equals(test.storedSession) {
				k.SaveSession(ctx, test.storedSession)
			}

			querier := keeper.NewQuerier(k)
			result, err := querier(ctx, test.query, request)

			if result != nil {
				require.Nil(t, err)

				expectedIndented, err := codec.MarshalJSONIndent(k.Cdc, &test.expRes)
				require.NoError(t, err)

				require.Equal(t, string(expectedIndented), string(result))
			}

			if result == nil {
				require.NotNil(t, err)
				require.Equal(t, test.expErr.Error(), err.Error())
				require.Nil(t, result)
			}
		})
	}
}
