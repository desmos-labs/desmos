package app

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/log"
	"cosmossdk.io/math"
	abci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/mock"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

func TestCheckProfilesKeeper(t *testing.T) {
	db := dbm.NewMemDB()
	app := NewDesmosApp(
		log.NewTestLogger(t), db, nil, true,
		simtestutil.NewAppOptionsWithFlagHome(t.TempDir()),
	)

	require.Equal(t, app.IBCKeeper.ChannelKeeper, app.ProfilesKeeper.ChannelKeeper)
	require.Equal(t, &app.IBCKeeper.PortKeeper, app.ProfilesKeeper.PortKeeper)
	require.Equal(t, app.ScopedProfilesKeeper, app.ProfilesKeeper.ScopedKeeper)
}

func TestSimAppExportAndBlockedAddrs(t *testing.T) {
	db := dbm.NewMemDB()
	logger := log.NewTestLogger(t)
	app := NewDesmosApp(
		logger, db, nil, true,
		simtestutil.NewAppOptionsWithFlagHome(t.TempDir()),
	)

	privVal := mock.NewPV()
	pubKey, err := privVal.GetPubKey()
	require.NoError(t, err)
	// create validator set with single validator
	validator := tmtypes.NewValidator(pubKey, 1)
	valSet := tmtypes.NewValidatorSet([]*tmtypes.Validator{validator})

	// generate genesis account
	senderPrivKey := secp256k1.GenPrivKey()
	acc := authtypes.NewBaseAccount(senderPrivKey.PubKey().Address().Bytes(), senderPrivKey.PubKey(), 0, 0)
	balance := banktypes.Balance{
		Address: acc.GetAddress().String(),
		Coins:   sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, math.NewInt(100000000000000))),
	}

	genesisState := NewDefaultGenesisState(app.AppCodec())
	genesisState, err = simtestutil.GenesisStateWithValSet(app.AppCodec(), genesisState, valSet, []authtypes.GenesisAccount{acc}, balance)
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	require.NoError(t, err)

	// Initialize the chain
	app.InitChain(
		&abci.RequestInitChain{
			Validators:      []abci.ValidatorUpdate{},
			ConsensusParams: simtestutil.DefaultConsensusParams,
			AppStateBytes:   stateBytes,
		},
	)

	// BlockedAddresses returns a map of addresses in app v1 and a map of modules name in app v2.
	for acc := range BlockedAddresses() {
		var addr sdk.AccAddress
		if modAddr, err := sdk.AccAddressFromBech32(acc); err == nil {
			addr = modAddr
		} else {
			addr = app.AccountKeeper.GetModuleAddress(acc)
		}

		require.True(
			t,
			app.BankKeeper.BlockedAddr(addr),
			fmt.Sprintf("ensure that blocked addresses are properly set in bank keeper: %s should be blocked", acc),
		)
	}

	// finalize block so we have CheckTx state set
	_, err = app.FinalizeBlock(&abci.RequestFinalizeBlock{
		Height: 1,
	})
	require.NoError(t, err)

	_, err = app.Commit()
	require.NoError(t, err)

	// Making a new app object with the db, so that initchain hasn't been called
	app2 := NewDesmosApp(
		logger.With("instance", "second"), db, nil, true, simtestutil.NewAppOptionsWithFlagHome(t.TempDir()),
	)
	_, err = app2.ExportAppStateAndValidators(false, []string{}, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
