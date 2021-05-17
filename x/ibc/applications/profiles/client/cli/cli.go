package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetQueryCmd returns the query commands for IBC profile connections
func GetQueryCmd() *cobra.Command {
	queryCmd := &cobra.Command{
		Use:                        "ibc-profiles",
		Short:                      "IBC profiles query subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
	}

	queryCmd.AddCommand()

	return queryCmd
}

// NewTxCmd returns the transaction commands for IBC fungible token transfer
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "ibc-profiles",
		Short:                      "IBC profiles transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewConnectProfileTxCmd(),
	)

	return txCmd
}
