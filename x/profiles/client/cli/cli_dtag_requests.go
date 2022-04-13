package cli

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

// GetCmdRequestDTagTransfer returns the command to create a DTag transfer request
func GetCmdRequestDTagTransfer() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "request-dtag-transfer [address]",
		Short:   "Make a request to get the DTag of the user having the given address",
		Example: fmt.Sprintf(`%s tx profiles request-dtag-transfer desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(1),
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
		Use:     "cancel-dtag-transfer-request [recipient]",
		Short:   "Cancel a DTag transfer made to the given recipient address",
		Example: fmt.Sprintf(`%s tx profiles cancel-dtag-transfer-request desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(1),
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
		Use: "accept-dtag-transfer-request [DTag] [address]",
		Short: `Accept a DTag transfer request made by the user with the given address.
When accepting the request, you can specify the request recipient DTag as your new DTag. 
If this happens, your DTag and the other user's one will be effectively swapped.`,
		Example: fmt.Sprintf(`%s tx profiles accept-dtag-transfer-request "leoDiCaprio" desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			receivingUser, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgAcceptDTagTransferRequest(args[0], receivingUser.String(), clientCtx.FromAddress.String())
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
		Use:     "refuse-dtag-transfer-request [sender]",
		Short:   "Refuse a DTag transfer made by the given sender address",
		Example: fmt.Sprintf(`%s tx profiles refuse-dtag-transfer-request desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud`, version.AppName),
		Args:    cobra.ExactArgs(1),
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
		Use:   "incoming-dtag-transfer-requests [[receiver]]",
		Short: "Retrieve the DTag transfer requests with optional address and pagination",
		Example: fmt.Sprintf(`%s tx profiles incoming-dtag-transfer-requests
%s tx profiles incoming-dtag-transfer-requests --page=2 --limit=100
%s tx profiles incoming-dtag-transfer-requests desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
`, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var receiver string
			if len(args) == 1 {
				receiver = args[0]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.IncomingDTagTransferRequests(
				context.Background(),
				types.NewQueryIncomingDTagTransferRequestsRequest(receiver, pageReq),
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
