package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/reports/types"
)

// NewTxCmd returns a new command allowing to perform reports transactions
func NewTxCmd() *cobra.Command {
	postsTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Reports transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	postsTxCmd.AddCommand(
		GetCmdReportPost(),
	)
	return postsTxCmd
}

// GetCmdReportPost returns the command allowing to report a post
func GetCmdReportPost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [post-id] [reports-type] [reports-message]",
		Short: "reports a post",
		Long: fmt.Sprintf(`
Report an existent post specifying its ID, the reports's type and message.

E.g.
%s tx reports create a4469741bb0c0622627810082a5f2e4e54fbbb888f25a4771a5eebc697d30cfc scam "this post is a scam" 
`, version.AppName),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgReportPost(args[0], args[1], args[2], clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
