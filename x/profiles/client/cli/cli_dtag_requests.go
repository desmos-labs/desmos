package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetCmdRequestDTagTransfer returns the command to create a DTag transfer request
func GetCmdRequestDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer-dtag [address]",
		Short: "Make a request to get the DTag of the user having the given address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			requestRecipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgRequestDTagTransfer(clientCtx.FromAddress.String(), requestRecipient.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdCancelDTagTransfer returns the command to cancel an outgoing DTag transfer request
func GetCmdCancelDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel-dtag-transfer [recipient]",
		Short: "Cancel a DTag transfer made to the given recipient address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			msg := types.NewMsgCancelDTagTransferRequest(clientCtx.FromAddress.String(), owner.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdAcceptDTagTransfer returns the command to accept a DTag transfer request
func GetCmdAcceptDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept-dtag-transfer [newDTag] [address]",
		Short: "Accept a DTag transfer request made by the user with the given address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receivingUser, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptDTagTransfer(args[0], receivingUser.String(), clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdRefuseDTagTransfer returns the command to refuse an incoming DTag transfer request
func GetCmdRefuseDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "refuse-dtag-transfer [sender]",
		Short: "Refuse a DTag transfer made by the given sender address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgRefuseDTagTransferRequest(args[0], clientCtx.FromAddress.String())
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// --------------------------------------------------------------------------------------------------------------------

// GetCmdQueryDTagRequests returns the command allowing to query all the DTag transfer requests made towards a user
func GetCmdQueryDTagRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dtag-requests [address]",
		Short: "Retrieve the requests made to the given address to transfer its profile's DTag with optional pagination",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.IncomingDTagTransferRequests(
				context.Background(),
				&types.QueryIncomingDTagTransferRequestsRequest{Receiver: args[0], Pagination: pageReq},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "DTag transfer requests")

	return cmd
}
