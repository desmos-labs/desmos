package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/spf13/cobra"
)

func GetCmdSavePermissionedContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save [contract_address]",
		Args:  cobra.ExactArgs(1),
		Short: "Save a permissioned contract with you as an admin",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSavePermissionedContractReference(args[0], clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
