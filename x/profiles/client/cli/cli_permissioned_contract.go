package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/desmos-labs/desmos/v2/x/profiles/types"
	"github.com/spf13/cobra"
	"io/ioutil"
	"strings"
)

func GetCmdSavePermissionedContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save [contract_address] [message_path]",
		Args:  cobra.ExactArgs(2),
		Short: "Save a permissioned contract and its executable message",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			messagePath := args[1]
			if strings.TrimSpace(messagePath) == "" {
				return fmt.Errorf("invalid message file path")
			}

			contractMessage, err := ioutil.ReadFile(messagePath)
			if err != nil {
				return err
			}

			msg := types.NewMsgSavePermissionedContractReference(
				args[0],
				clientCtx.FromAddress.String(),
				contractMessage,
			)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
