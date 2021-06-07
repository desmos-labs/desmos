package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetQueryCmd returns the command allowing to perform queries
func GetQueryCmd() *cobra.Command {
	profileQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the profiles module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	profileQueryCmd.AddCommand(
		GetCmdQueryProfile(),
		GetCmdQueryDTagRequests(),
		GetCmdQueryUserRelationships(),
		GetCmdQueryUserBlocks(),
		GetCmdQueryParams(),
		GetCmdQueryProfileByChainLink(),
	)
	return profileQueryCmd
}
