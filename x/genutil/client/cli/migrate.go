package cli

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	extypes "github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/desmos-labs/desmos/x/genutil/internal/types"
	v020 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.2.0"
	v030 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.3.0"
	v040 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.4.0"
	v050 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.5.0"
	v060 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.6.0"
	v080 "github.com/desmos-labs/desmos/x/genutil/legacy/v0.8.0"
	"github.com/spf13/cobra"
	tm "github.com/tendermint/tendermint/types"
)

// migrationMap contains the list of migrations that should be performed when migrating
// a version of the chain to the next one. It contains an array as we need to support Cosmos SDK migrations
// too if needed.
var migrationMap = map[string]types.MigrationCallback{
	"v0.2.0": v020.Migrate,
	"v0.3.0": v030.Migrate,
	"v0.4.0": v040.Migrate,
	"v0.5.0": v050.Migrate,
	"v0.6.0": v060.Migrate,
	"v0.8.0": v080.Migrate,
}

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

func MigrateGenesisCmd(cdc *codec.Codec) *cobra.Command {
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
`, version.ServerName, version.ServerName, version.ServerName),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			target := args[0]
			importGenesis := args[1]

			genDoc, err := tm.GenesisDocFromFile(importGenesis)
			if err != nil {
				return err
			}

			var initialState extypes.AppMap
			cdc.MustUnmarshalJSON(genDoc.AppState, &initialState)

			migration := migrationMap[target]
			if migration == nil {
				return fmt.Errorf("unknown migration function version: %s", target)
			}

			newGenState := initialState

			// v0.2.0 migration needs to have the previous version's genesis time and the
			// block interval to convert the block height dates into timestamps
			if target == "v0.2.0" {
				blockInterval, err := strconv.Atoi(cmd.Flag(flagBlockInterval).Value.String())
				if err != nil {
					panic(err)
				}

				newGenState = migration(newGenState, genDoc.GenesisTime, blockInterval)
			} else {
				newGenState = migration(newGenState)
			}

			genDoc.AppState = cdc.MustMarshalJSON(newGenState)

			genesisTime := cmd.Flag(flagGenesisTime).Value.String()
			if genesisTime != "" {
				var t time.Time

				err := t.UnmarshalText([]byte(genesisTime))
				if err != nil {
					return err
				}

				genDoc.GenesisTime = t
			}

			chainID := cmd.Flag(flagChainID).Value.String()
			if chainID != "" {
				genDoc.ChainID = chainID
			}

			out, err := cdc.MarshalJSONIndent(genDoc, "", "  ")
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
