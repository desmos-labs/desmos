package cli

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	tmjson "github.com/cometbft/cometbft/libs/json"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/version"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/spf13/cobra"
)

const chainUpgradeGuide = "https://docs.cosmos.network/master/migrations/chain-upgrade-guide-040.html"

// migrationMap contains the list of migrations that should be performed when migrating
// a version of the chain to the next one. It contains an array as we need to support Cosmos SDK migrations
// too if needed.
var migrationMap = map[string]genutiltypes.MigrationCallback{}

const (
	flagGenesisTime = "genesis-time"
	flagChainID     = "chain-id"
)

// GetMigrationCallback returns a MigrationCallback for a given version.
func GetMigrationCallback(version string) genutiltypes.MigrationCallback {
	return migrationMap[version]
}

// GetMigrationVersions get all migration version in a sorted slice.
func GetMigrationVersions() []string {
	versions := make([]string, len(migrationMap))

	var i int

	for version := range migrationMap {
		versions[i] = version
		i++
	}

	sort.Strings(versions)

	return versions
}

func MigrationsListCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrations-list",
		Short: "Lists all the available migrations",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {

			migrations := GetMigrationVersions()
			for _, m := range migrations {
				fmt.Println(m)
			}

			return nil
		},
	}
}

func MigrateGenesisCmd() *cobra.Command {
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
			clientCtx := client.GetClientContextFromCmd(cmd)

			var err error

			target := args[0]
			importGenesis := args[1]

			genDoc, err := validateGenDoc(importGenesis)
			if err != nil {
				return err
			}

			// Since some default values are valid values, we just print to
			// make sure the user didn't forget to update these values.
			if genDoc.ConsensusParams.Evidence.MaxBytes == 0 {
				fmt.Printf("Warning: consensus_params.evidence.max_bytes is set to 0. If this is"+
					" deliberate, feel free to ignore this warning. If not, please have a look at the chain"+
					" upgrade guide at %s.\n", chainUpgradeGuide)
			}

			var initialState genutiltypes.AppMap
			if err := json.Unmarshal(genDoc.AppState, &initialState); err != nil {
				return errors.Wrap(err, "failed to JSON unmarshal initial genesis state")
			}

			migrationFunc := GetMigrationCallback(target)
			if migrationFunc == nil {
				return fmt.Errorf("unknown migration function for version: %s", target)
			}

			// TODO: handler error from migrationFunc call
			newGenState := migrationFunc(initialState, clientCtx)

			genDoc.AppState, err = json.Marshal(newGenState)
			if err != nil {
				return errors.Wrap(err, "failed to JSON marshal migrated genesis state")
			}

			genesisTime, _ := cmd.Flags().GetString(flagGenesisTime)
			if genesisTime != "" {
				var t time.Time

				err := t.UnmarshalText([]byte(genesisTime))
				if err != nil {
					return errors.Wrap(err, "failed to unmarshal genesis time")
				}

				genDoc.GenesisTime = t
			}

			chainID, _ := cmd.Flags().GetString(flags.FlagChainID)
			if chainID != "" {
				genDoc.ChainID = chainID
			}

			bz, err := tmjson.Marshal(genDoc)
			if err != nil {
				return errors.Wrap(err, "failed to marshal genesis doc")
			}

			sortedBz, err := sdk.SortJSON(bz)
			if err != nil {
				return errors.Wrap(err, "failed to sort JSON genesis doc")
			}

			fmt.Println(string(sortedBz))
			return nil
		},
	}

	cmd.Flags().String(flagGenesisTime, "", "Override genesis_time with this flag")
	cmd.Flags().String(flagChainID, "", "Override chain_id with this flag")

	return cmd
}

// validateGenDoc reads a genesis file and validates that it is a correct
// Tendermint GenesisDoc. This function does not do any cosmos-related
// validation.
func validateGenDoc(importGenesisFile string) (*tmtypes.GenesisDoc, error) {
	genDoc, err := tmtypes.GenesisDocFromFile(importGenesisFile)
	if err != nil {
		return nil, fmt.Errorf("%s. Make sure that"+
			" you have correctly migrated all Tendermint consensus params, please see the"+
			" chain migration guide at %s for more info",
			err.Error(), chainUpgradeGuide,
		)
	}

	return genDoc, nil
}
