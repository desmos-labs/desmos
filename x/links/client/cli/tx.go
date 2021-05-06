package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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
		GetCmdCreateIBCAccountConnection(),
	)

	return linksTxCmd
}

// GetCmdCreateIBCAccountConnection returns the command to create an account link on other chain with different private keys
func GetCmdCreateIBCAccountConnection() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ibc-connection [src-port] [src-channel] [dest-chain-prefix] [dest-keybase-path] [destination-key-name]",
		Short: "Create a new account link with different keys",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			srcPort := args[0]
			srcChannel := args[1]
			destChainPrefix := args[2]
			destKeyBasePath := args[3]
			destKeyName := args[4]

			// get source key info from cli
			srcKeybase, srcKeyName, srcAddr := getSourceKeyInfo(clientCtx)

			// get destination key info from path
			keyringBackend, _ := cmd.Flags().GetString(flags.FlagKeyringBackend)
			destKeybase, err := keyring.New(destChainPrefix, keyringBackend, destKeyBasePath, clientCtx.Input)
			if err != nil {
				return err
			}
			destKey, err := destKeybase.Key(destKeyName)
			if err != nil {
				return err
			}

			// Get bech32 encoded address on destination chain
			destAddr, err := bech32.ConvertAndEncode(destChainPrefix, destKey.GetAddress().Bytes())
			if err != nil {
				return err
			}

			link := types.NewLink(srcAddr, destAddr)
			linkBz, _ := link.Marshal()

			// Create signature by keys
			srcPubKey, srcSig, destSig, err := accountConnectionSign(
				linkBz,
				srcKeybase,
				srcKeyName,
				destKeybase,
				destKeyName,
			)
			if err != nil {
				return err
			}

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
				types.NewIBCAccountConnectionPacketData(
					sdk.GetConfig().GetBech32AccountAddrPrefix(),
					srcAddr,
					hex.EncodeToString(srcPubKey.Bytes()),
					destAddr,
					hex.EncodeToString(srcSig),
					hex.EncodeToString(destSig),
				),
				timeoutTimestamp,
			)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(
		FlagPacketTimeoutTimestamp,
		DefaultRelativePacketTimeoutTimestamp,
		"Packet timeout timestamp in nanoseconds. Default is 10 minutes. The timeout is disabled when set to 0.",
	)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func accountConnectionSign(
	msg []byte,
	srcKeybase keyring.Keyring,
	srcKeyName string,
	destKeyBase keyring.Keyring,
	destKeyName string,
) (cryptotypes.PubKey, []byte, []byte, error) {
	// Create signature by src key
	srcSig, srcPubKey, err := srcKeybase.Sign(srcKeyName, msg)
	if err != nil {
		return nil, nil, nil, err
	}

	// Create signature by dest key
	destSig, _, err := destKeyBase.Sign(destKeyName, srcSig)
	if err != nil {
		return nil, nil, nil, err
	}
	return srcPubKey, srcSig, destSig, nil
}

// GetCmdCreateIBCAccountLink return the command to create an account link on other chain
func GetCmdCreateIBCAccountLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-ibc-link [src-port] [src-channel] [dest-chain-prefix]",
		Short: "Create a new ibc account link",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			srcPort := args[0]
			srcChannel := args[1]
			destChainPrefix := args[2]

			destAddr, err := bech32.ConvertAndEncode(destChainPrefix, clientCtx.GetFromAddress().Bytes())
			if err != nil {
				return err
			}

			// get source chain key info
			keybase, keyName, srcAddr := getSourceKeyInfo(clientCtx)

			link := types.NewLink(srcAddr, destAddr)
			linkBz, _ := link.Marshal()

			sig, pubKey, err := keybase.Sign(keyName, linkBz)
			if err != nil {
				return err
			}

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
				types.NewIBCAccountLinkPacketData(
					sdk.GetConfig().GetBech32AccountAddrPrefix(),
					srcAddr,
					hex.EncodeToString(pubKey.Bytes()),
					hex.EncodeToString(sig),
				),
				timeoutTimestamp,
			)

			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().Uint64(FlagPacketTimeoutTimestamp, DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes. The timeout is disabled when set to 0.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getSourceKeyInfo(clientCtx client.Context) (keyring.Keyring, string, string) {
	keybase := clientCtx.Keyring
	keyName := clientCtx.GetFromName()
	addr := clientCtx.GetFromAddress().String()
	return keybase, keyName, addr
}
