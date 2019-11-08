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
	ctx, k := SetupTestInput()

	path := []string{keeper.QuerySessions, types.SessionID(1).String()}

	querier := keeper.NewQuerier(k)
	_, err := querier(ctx, path, request)

	assert.Error(t, err)
	assert.Contains(t, err.Result().Log, "Session with id 1 not found")
}

func Test_querySession_ValidIdReturnsAssociatedSession(t *testing.T) {
	ctx, k := SetupTestInput()

	// Store a test session
	store := ctx.KVStore(k.StoreKey)
	store.Set([]byte(types.SessionStorePrefix+testSession.SessionID.String()), k.Cdc.MustMarshalBinaryBare(&testSession))

	path := []string{keeper.QuerySessions, testSession.SessionID.String()}

	querier := keeper.NewQuerier(k)
	actualBz, err := querier(ctx, path, request)

	assert.NoError(t, err)

	var actual types.Session
	k.Cdc.MustUnmarshalJSON(actualBz, &actual)
	assert.Equal(t, testSession, actual)
}
