package main

import (
	"encoding/json"
	"io"

	app2 "github.com/desmos-labs/desmos/app"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccscli "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/spf13/cobra"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"
)

func main() {
	cobra.EnableCommandSorting = false

	cdc := app2.MakeCodec()

	config := sdk.GetConfig()
	app2.SetBech32AddressPrefixes(config)

	// 852 is the international dialing code of Hong Kong
	// Following the coin type registered at https://github.com/satoshilabs/slips/blob/master/slip-0044.md
	config.SetCoinType(852)
	config.Seal()

	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "desmosd",
		Short:             "Desmos App Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}
	// CLI commands to initialize the chain
	rootCmd.AddCommand(genutilcli.InitCmd(ctx, cdc, app2.ModuleBasics, app2.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app2.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
	rootCmd.AddCommand(
		genutilcli.GenTxCmd(
			ctx, cdc, app2.ModuleBasics, staking.AppModuleBasic{},
			genaccounts.AppModuleBasic{}, app2.DefaultNodeHome, app2.DefaultCLIHome,
		),
	)
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, app2.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx, cdc, app2.DefaultNodeHome, app2.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx, cdc, app2.ModuleBasics, genaccounts.AppModuleBasic{}))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "DM", app2.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, _ io.Writer) abci.Application {
	return app2.NewDesmosApp(logger, db)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, _ io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		desmosApp := app2.NewDesmosApp(logger, db)
		err := desmosApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return desmosApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	desmosApp := app2.NewDesmosApp(logger, db)

	return desmosApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
