package cli

import (
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// NewTxCmd set the tx commands
func NewTxCmd() *cobra.Command {
	magpieTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Magpie transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	magpieTxCmd.AddCommand(GetCmdCreateSession())

	return magpieTxCmd
}

// GetCmdCreateSession is the CLI command for creating a session for create post
func GetCmdCreateSession() *cobra.Command {
	return &cobra.Command{
		Use:   "create-session [namespace] [external address] [pubkey] [external signer signature]",
		Short: "Creates a session for an external service to post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadTxCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSession(clientCtx.FromAddress.String(), args[0], args[1], args[2], args[3])
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
}
