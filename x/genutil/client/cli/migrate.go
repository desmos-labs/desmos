package cli

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	extypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
	tm "github.com/tendermint/tendermint/types"

	"github.com/desmos-labs/desmos/x/genutil/types"
)

// migrationMap contains the list of migrations that should be performed when migrating
// a version of the chain to the next one. It contains an array as we need to support Cosmos SDK migrations
// too if needed.
var migrationMap = map[string]types.MigrationCallback{}

const (
	flagGenesisTime   = "genesis-time"
	flagChainID       = "chain-id"
	flagBlockInterval = "block-interval"
)

func MigrationsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations-list",
		Short: "Lists all the available migrations",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			var migrations []string
			for migration := range migrationMap {
				migrations = append(migrations, migration)
			}

			sort.Strings(migrations)
			for _, m := range migrations {
				fmt.Println(m)
			}

			return nil
		},
	}
}

func MigrateGenesisCmd(cdc codec.Marshaler) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate [target-version] [genesis-file]",
		Short: "Migrate genesis to a specified target version",
		Long: fmt.Sprintf(`Migrate the source genesis into the target version and print to STDOUT.

Please note that migrations should be only be done sequentially. As a reference, suppose we have the following versions:
- v0.1.0
- v0.2.0
- v0.3.0

If you want to migrate from version v0.1.0 to v0.3.0, you need to execute two migrations:
1. From v0.1.0 to v0.2.0
   $ %s migrate v0.2.0 ...
2. From v0.2.0 to v0.3.0
   $ %s migrate v0.3.0 ...

To see get a full list of available migrations, use the migrations-list command.

Example:
$ %s migrate v0.2.0 /path/to/genesis.json --chain-id=morpheus-XXXXX --genesis-time=2019-11-31T18:00:00Z
`, version.AppName, version.AppName, version.AppName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			target := args[0]
			importGenesis := args[1]

			genDoc, err := tm.GenesisDocFromFile(importGenesis)
			if err != nil {
				return err
			}

			var initialState extypes.AppMap
			if err := json.Unmarshal(genDoc.AppState, &initialState); err != nil {
				return errors.Wrap(err, "failed to JSON unmarshal initial genesis state")
			}

			migration := migrationMap[target]
			if migration == nil {
				return fmt.Errorf("unknown migration function version: %s", target)
			}

			// Get the genesis time
			// We get this year since that the migrations might want to get it
			var genesisTime time.Time
			genesisTimeStr := cmd.Flag(flagGenesisTime).Value.String()
			if genesisTimeStr != "" {
				err := genesisTime.UnmarshalText([]byte(genesisTimeStr))
				if err != nil {
					return err
				}
			}

			// Perform the migration
			newGenState := migration(initialState, genesisTime)

			genDoc.AppState, err = json.Marshal(newGenState)
			if err != nil {
				return errors.Wrap(err, "failed to JSON marshal migrated genesis state")
			}

			// Set genesis time
			if !genesisTime.IsZero() {
				genDoc.GenesisTime = genesisTime
			}

			// Set chain id
			chainID := cmd.Flag(flagChainID).Value.String()
			if chainID != "" {
				genDoc.ChainID = chainID
			}

			out, err := json.MarshalIndent(genDoc, "", "  ")
			if err != nil {
				return err
			}

			fmt.Println(string(sdk.MustSortJSON(out)))
			return nil
		},
	}

	cmd.Flags().String(flagGenesisTime, "", "Override genesis_time with this flag")
	cmd.Flags().String(flagChainID, "", "Override chain_id with this flag")
	cmd.Flags().Int(flagBlockInterval, 0, "Block interval of seconds to consider while computing timestamps dates")

	return cmd
}
