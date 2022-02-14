package cli

import (
	"context"
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
		Use:   "save-contract [contract_address] [message_path]",
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

func GetCmdQueryPermissionedContract() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pc [admin] [address]",
		Short: "Get the permissioned contract associated with admin and address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.PermissionedContract(
				context.Background(),
				&types.QueryPermissionedContractRequest{Admin: args[0], ContractAddress: args[1]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "app links")

	return cmd
}
