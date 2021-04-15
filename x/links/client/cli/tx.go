package cli

import (
	"encoding/hex"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types/bech32"
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

	linksTxCmd.AddCommand(
		GetCmdCreateIBCAccountLink(),
	)

	return linksTxCmd
}

// GetCmdCreateIBCAccountConnection returns the command to create an account link on other chain with different private keys
func GetCmdCreateIBCAccountConnection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ibc-connection [src-port] [src-channel] [destination-chain] [destination-chain-path] [destination-key-name]",
		Short: "Create a new account link with different keys",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			srcPort := args[0]
			srcChannel := args[1]
			dstChain := args[2]
			dstKeyBasePath := args[3]
			dstKeyName := args[4]

			// get source key info from cli
			srcKeyName := clientCtx.GetFromName()
			srcKeybase := clientCtx.Keyring
			srcAddr := clientCtx.GetFromAddress().String()

			// get destination key info from path
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			dstKeyBase, err := keyring.New(dstChain, keyringBackend, dstKeyBasePath, clientCtx.Input)
			dstKey, err := dstKeyBase.Key(dstKeyName)
			if err != nil {
				return err
			}

			dstAddr, err := bech32.ConvertAndEncode(dstChain, dstKey.GetAddress().Bytes())
			if err != nil {
				return err
			}

			link := types.NewLink(srcAddr, dstAddr)
			linkBz, err := link.Marshal()
			if err != nil {
				return err
			}

			srcSig, srcPubKey, err := srcKeybase.Sign(srcKeyName, linkBz)
			if err != nil {
				return err
			}

			srcPubKeyHex := hex.EncodeToString(srcPubKey.Bytes())
			srcSigHex := hex.EncodeToString(srcSig)

			dstSig, _, err := dstKeyBase.Sign(dstKeyName, linkBz)
			if err != nil {
				return err
			}
			dstSigHex := hex.EncodeToString(dstSig)

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

			msg := types.NewMsgCreateIBCAccountConnection(
				srcPort,
				srcChannel,
				timeoutTimestamp,
				srcAddr,
				srcPubKeyHex,
				dstAddr,
				srcSigHex,
				dstSigHex,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

// GetCmdCreateIBCAccountLink return the command to create an account link on other chain
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

			// get source chain key info
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
