package cmd

// DONTCOVER

import (
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	srvconfig "github.com/cosmos/cosmos-sdk/server/config"
	"github.com/cosmos/cosmos-sdk/types/module"
	cosmosgenutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
)

// initCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func initCmd(mbm module.BasicManager, defaultNodeHome string) *cobra.Command {
	cmd := cosmosgenutilcli.InitCmd(mbm, defaultNodeHome)
	defaultRun := cmd.RunE
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		clientCtx := client.GetClientContextFromCmd(cmd)
		err := defaultRun(cmd, args)

		// Set app.toml file
		appConfig := srvconfig.DefaultConfig()
		appConfig.StateSync.SnapshotInterval = 500
		srvconfig.WriteConfigFile(filepath.Join(clientCtx.HomeDir, "config", "app.toml"), appConfig)
		return err
	}
	return cmd
}
