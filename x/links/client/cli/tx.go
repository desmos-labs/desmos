package cli

import (
	"encoding/hex"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	channelutils "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/client/utils"
	"github.com/desmos-labs/desmos/x/links/types"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	linksTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Links transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	linksTxCmd.AddCommand()

	return linksTxCmd
}

// GetCmdCreateIBCAccountLink return the command to create a account link on other chain
func GetCmdCreateIBCAccountLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ibc-link [src-port] [src-channel] [destination-address]",
		Short: "Create a new ibc account link",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			srcPort := args[0]
			srcChannel := args[1]
			dstAddr := args[2]

			keyName := clientCtx.GetFromName()
			keybase := clientCtx.Keyring
			srcAddr := clientCtx.GetFromAddress().String()

			link := types.NewLink(srcAddr, dstAddr)
			linkBz, err := link.Marshal()
			if err != nil {
				return nil
			}
			sig, srcPubKey, err := keybase.Sign(keyName, linkBz)

			srcPubKeyHex := hex.EncodeToString(srcPubKey.Bytes())
			sigHex := hex.EncodeToString(sig)

			// Get the relative timeout timestamp
			timeoutTimestamp, err := cmd.Flags().GetUint64(FlagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}
			consensusState, _, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
			if err != nil {
				return err
			}
			if timeoutTimestamp != 0 {
				timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
			}

			msg := types.NewMsgCreateIBCAccountLink(
				srcPort,
				srcChannel,
				timeoutTimestamp,
				srcAddr,
				srcPubKeyHex,
				sigHex,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
