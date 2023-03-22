package app

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/CosmWasm/wasmd/x/wasm"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"

	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/log"
	"github.com/stretchr/testify/require"
)

func TestSimAppExport(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewDesmosApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		DefaultNodeHome, 0, MakeTestEncodingConfig(), simtestutil.NewAppOptionsWithFlagHome(t.TempDir()), wasm.EnableAllProposals,
	)

	genesisState := NewDefaultGenesisState()
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	app.Commit()

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewDesmosApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, map[int64]bool{},
		DefaultNodeHome, 0, MakeTestEncodingConfig(), simtestutil.NewAppOptionsWithFlagHome(t.TempDir()), wasm.EnableAllProposals,
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}

func TestGetMaccPerms(t *testing.T) {
	dup := GetMaccPerms()
	require.Equal(t, maccPerms, dup, "duplicated module account permissions differed from actual module account permissions")
}
