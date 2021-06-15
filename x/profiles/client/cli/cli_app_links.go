package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channelutils "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/client/utils"

	"github.com/desmos-labs/desmos/x/profiles/types"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
)

// GetCmdLinkApplication returns the command to create a NewMsgLinkApplication transaction
func GetCmdLinkApplication() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-app [src-port] [src-channel] [application] [username] [verification-call-data]",
		Short: "Link a centralized application account to your Desmos profile",
		Long: strings.TrimSpace(`Connect a Desmos profile to a centralized social network account through IBC. 
Timeouts can be specified as absolute or relative using the "absolute-timeouts" flag. 
Timeout height can be set by passing in the height string in the form {revision}-{height} using the "packet-timeout-height" flag. 
Relative timeouts are added to the block height and block timestamp queried from the latest consensus state corresponding 
to the counterparty channel. Any timeout set to 0 is disabled.`),
		Example: fmt.Sprintf(
			"%s tx profiles link-app [src-port] [src-channel] [application] [username] [verification-call-data]",
			version.AppName),
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()
			srcPort := args[0]
			srcChannel := args[1]
			linkData := types.NewData(args[2], args[3])
			oracleRequestCallData := types.NewOracleRequestCallData(args[2], args[4])

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			absoluteTimeouts, err := cmd.Flags().GetBool(flagAbsoluteTimeouts)
			if err != nil {
				return err
			}

			// if the timeouts are not absolute, retrieve latest block height and block timestamp
			// for the consensus state connected to the destination port/channel
			if !absoluteTimeouts {
				consensusState, height, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
				if err != nil {
					return err
				}

				if !timeoutHeight.IsZero() {
					absoluteHeight := height
					absoluteHeight.RevisionNumber += timeoutHeight.RevisionNumber
					absoluteHeight.RevisionHeight += timeoutHeight.RevisionHeight
					timeoutHeight = absoluteHeight
				}

				if timeoutTimestamp != 0 {
					timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
				}
			}

			msg := types.NewMsgLinkApplication(
				linkData, oracleRequestCallData, sender,
				srcPort, srcChannel, timeoutHeight, timeoutTimestamp,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagPacketTimeoutHeight, types.DefaultRelativePacketTimeoutHeight, "Packet timeout block height. The timeout is disabled when set to 0-0.")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, types.DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes. The timeout is disabled when set to 0.")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "Timeout flags are used as absolute timeouts.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdUnlinkApplication returns the command allowing to unlink a centralized application from a Desmos profile
func GetCmdUnlinkApplication() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unlink-app [application] [username]",
		Short:   "Unlink a centralized application account from your Desmos profile",
		Example: fmt.Sprintf("%s tx profiles unlink-app [application] [username]", version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			sender := clientCtx.GetFromAddress()

			msg := types.NewMsgUnlinkApplication(args[0], args[1], sender)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// -------------------------------------------------------------------------------------------------------------------

// GetCmdQueryAppLinks returns the command allowing to query the application links associated with a profile
func GetCmdQueryAppLinks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app-links [user]",
		Short: "Get all the application links associated to the given username with optional pagination",
		Args:  cobra.ExactArgs(2),
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

			res, err := queryClient.UserApplicationLinks(
				context.Background(),
				&types.QueryUserApplicationLinksRequest{User: args[0], Pagination: pageReq},
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
