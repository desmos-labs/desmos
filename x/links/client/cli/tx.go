package cli

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
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

			srcKeybase, srcKey, err := GetSourceKeyInfo(clientCtx)
			if err != nil {
				return err
			}

			packet, err := GetIBCAccountConnectionPacket(
				srcKeybase,
				srcKey,
				destKeybase,
				destKey,
				destChainPrefix,
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
				packet,
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

func GetIBCAccountConnectionPacket(
	srcKeybase keyring.Keyring,
	srcKey keyring.Info,
	destKeybase keyring.Keyring,
	destKey keyring.Info,
	destChainPrefix string,
) (types.IBCAccountConnectionPacketData, error) {
	// Get bech32 encoded address on destination chain
	destAddr, err := bech32.ConvertAndEncode(destChainPrefix, destKey.GetAddress().Bytes())
	if err != nil {
		return types.IBCAccountConnectionPacketData{}, err
	}

	srcAddr := srcKey.GetAddress().String()
	link := types.NewLink(srcAddr, destAddr)
	linkBz, _ := link.Marshal()

	// Create signature by src keys
	srcSig, srcPubKey, err := srcKeybase.Sign(srcKey.GetName(), linkBz)
	if err != nil {
		return types.IBCAccountConnectionPacketData{}, err
	}

	// Create signature by dest key
	destSig, _, err := destKeybase.Sign(destKey.GetName(), srcSig)
	if err != nil {
		return types.IBCAccountConnectionPacketData{}, err
	}

	packet := types.NewIBCAccountConnectionPacketData(
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		srcAddr,
		hex.EncodeToString(srcPubKey.Bytes()),
		destAddr,
		hex.EncodeToString(srcSig),
		hex.EncodeToString(destSig),
	)

	return packet, nil
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

			srcKeybase, srcKey, err := GetSourceKeyInfo(clientCtx)
			if err != nil {
				return err
			}

			packet, err := GetIBCAccountLinkPacket(
				srcKeybase,
				srcKey,
				destChainPrefix,
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

			msg := types.NewMsgCreateIBCAccountLink(
				srcPort,
				srcChannel,
				packet,
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

func GetIBCAccountLinkPacket(
	srcKeybase keyring.Keyring,
	srcKey keyring.Info,
	destChainPrefix string,
) (types.IBCAccountLinkPacketData, error) {
	destAddr, err := bech32.ConvertAndEncode(destChainPrefix, srcKey.GetAddress().Bytes())
	if err != nil {
		return types.IBCAccountLinkPacketData{}, err
	}

	srcAddr := srcKey.GetAddress().String()
	link := types.NewLink(srcAddr, destAddr)
	linkBz, _ := link.Marshal()

	sig, pubKey, err := srcKeybase.Sign(srcKey.GetName(), linkBz)
	if err != nil {
		return types.IBCAccountLinkPacketData{}, err
	}

	packet := types.NewIBCAccountLinkPacketData(
		sdk.GetConfig().GetBech32AccountAddrPrefix(),
		srcAddr,
		hex.EncodeToString(pubKey.Bytes()),
		hex.EncodeToString(sig),
	)

	return packet, nil
}

func GetSourceKeyInfo(clientCtx client.Context) (keyring.Keyring, keyring.Info, error) {
	keybase := clientCtx.Keyring
	keyName := clientCtx.GetFromName()
	key, err := keybase.Key(keyName)
	if err != nil {
		return nil, nil, err
	}
	return keybase, key, nil
}
