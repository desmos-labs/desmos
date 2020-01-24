package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/store"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/desmos-labs/desmos/app"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/server"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	desmosgenutil "github.com/desmos-labs/desmos/x/genutil/client/cli"
)

const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	// Initialize the app overriding the various methods we want to customize
	app.Init()

	config := sdk.GetConfig()
	app.SetupConfig(config)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "desmosd",
		Short:             "Desmos App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, app.DefaultNodeHome))
	rootCmd.AddCommand(
		genutilcli.GenTxCmd(
			ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
			auth.GenesisAccountIterator{}, app.DefaultNodeHome, app.DefaultCLIHome,
		),
	)
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(flags.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx, cdc, app.ModuleBasics, auth.GenesisAccountIterator{}))
	rootCmd.AddCommand(replayCmd())
	rootCmd.AddCommand(debug.Cmd(cdc))

	// Custom commands
	rootCmd.AddCommand(desmosgenutil.MigrationsListCmd())
	rootCmd.AddCommand(desmosgenutil.MigrateGenesisCmd(cdc))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "GA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, _ io.Writer) abci.Application {
	var cache sdk.MultiStorePersistentCache

	if viper.GetBool(server.FlagInterBlockCache) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(h)] = true
	}

	return app.NewDesmosApp(
		logger, db, skipUpgradeHeights,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(viper.GetUint64(server.FlagHaltHeight)),
		baseapp.SetHaltTime(viper.GetUint64(server.FlagHaltTime)),
		baseapp.SetInterBlockCache(cache),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, _ io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		gapp := app.NewDesmosApp(logger, db, map[int64]bool{})
		err := gapp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return gapp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	gapp := app.NewDesmosApp(logger, db, map[int64]bool{})
	return gapp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
