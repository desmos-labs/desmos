package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/desmos-labs/desmos/x/magpie/types"
)

// NewTxCmd returns a new command allowing to perform magpie transactions
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

// GetCmdCreateSession returns the command allowing to create a session
func GetCmdCreateSession() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-session [namespace] [external address] [pubkey] [external signer signature]",
		Short: "Creates a session for an external service to post",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateSession(clientCtx.FromAddress.String(), args[0], args[1], args[2], args[3])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
