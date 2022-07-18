package cli

import (
	"context"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/version"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	profilescliutils "github.com/desmos-labs/desmos/v4/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/v4/x/profiles/types"
)

// GetCmdLinkChainAccount returns the command allowing to link an external chain account
func GetCmdLinkChainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link-chain [data-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Link a Desmos profile to an external chain account using the data written inside the given file",
		Long: strings.TrimSpace(fmt.Sprintf(`Link an external account to a Desmos profile.
The link data must be supplied via a JSON file.

Example:
$ %s tx profiles link-chain <path/to/data.json> --from=<key_or_address>

Where data.json contains:

{
   "address":{
      "@type":"/desmos.profiles.v3.Bech32Address",
      "value":"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
      "prefix":"cosmos"
   },
   "proof":{
      "pub_key":{
         "@type":"/cosmos.crypto.secp256k1.PubKey",
         "key":"A58DXR/lXKVkIjLofXgST/OHi+pkOQbVIiOjnTy7Zoqo"
      },
      "signature":"ecc6175e730917fb289d3a9f4e49a5630a44b42d972f481342f540e09def2ec5169780d85c4e060d52cc3ffb3d677745a4d56cd385760735bc6db0f1816713be",
      "plain_text":"636f736d6f73313575633839766e7a756675356b7568687378646b6c747433387a66783876637967677a77666d"
   },
   "chain_config":{
      "name":"cosmos"
   }
}
`, version.AppName)),

		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			data, err := profilescliutils.ParseChainLinkJSON(clientCtx.Codec, args[0])
			if err != nil {
				return err
			}

			addr := data.Address.GetCachedValue().(types.AddressData)
			msg := types.NewMsgLinkChainAccount(
				addr,
				data.Proof,
				data.ChainConfig,
				clientCtx.GetFromAddress().String(),
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

// GetCmdUnlinkChainAccount returns the command allowing to unlink an external chain account from a profile
func GetCmdUnlinkChainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "unlink-chain [chain-name] [address]",
		Short:   "Unlink the external account having the given chain name and address",
		Example: fmt.Sprintf(`%s tx profiles unlink-chain "cosmos" cosmos18xnmlzqrqr6zt526pnczxe65zk3f4xgmndpxn2`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUnlinkChainAccount(clientCtx.FromAddress.String(), args[0], args[1])
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// GetCmdSetDefaultExternalAddress returns the command allowing to set a default external address for a user
func GetCmdSetDefaultExternalAddress() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-default-external-address [chain-name] [target]",
		Short:   "Set the default external account having the given chain name and external address",
		Example: fmt.Sprintf(`%s tx profiles set-default-external-address "cosmos" cosmos18xnmlzqrqr6zt526pnczxe65zk3f4xgmndpxn2`, version.AppName),
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSetDefaultExternalAddress(args[0], args[1], clientCtx.FromAddress.String())
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

// GetCmdQueryChainLinks returns the command allowing to query the chain links, optionally associated with a user
func GetCmdQueryChainLinks() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-links [[user]] [[chain_name]] [[target]]",
		Short: "Retrieve all chain links with optional user address, chain name, target and pagination",
		Example: fmt.Sprintf(`%s query profiles chain-links
%s query profiles chain-links --page=2 --limit=100
%s query profiles chain-links desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud
%s query profiles chain-links desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud "cosmos"
%s query profiles chain-links desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud "cosmos" cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw
`, version.AppName, version.AppName, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(0, 3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var user string
			if len(args) > 0 {
				user = args[0]
			}

			var chainName string
			if len(args) > 1 {
				chainName = args[1]
			}

			var target string
			if len(args) > 2 {
				target = args[2]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.ChainLinks(
				context.Background(),
				types.NewQueryChainLinksRequest(user, chainName, target, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "chain links")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryChainLinkOwners returns the command allowing to query the chain link owners, optionally associated with a target
func GetCmdQueryChainLinkOwners() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-link-owners [[chain_name]] [[target]]",
		Short: "Retrieve all chain link owners with optional chain name, target and pagination",
		Example: fmt.Sprintf(`%s query profiles chain-link-owners
%s query profiles chain-link-owners --page=2 --limit=100
%s query profiles chain-link-owners "cosmos"
%s query profiles chain-link-owners "cosmos" cosmos19s242dxhxgzlsdmfjjg38jgfwhxca7569g84sw
`, version.AppName, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var chainName string
			if len(args) > 0 {
				chainName = args[0]
			}

			var target string
			if len(args) > 1 {
				target = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.ChainLinkOwners(
				context.Background(),
				types.NewQueryChainLinkOwnersRequest(chainName, target, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "chain link owners")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDefaultExternalAddresses returns the command allowing to query the default chain links, optionally associated with a owner and chain name
func GetCmdQueryDefaultExternalAddresses() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "default-external-addresses [[owner]] [[chain_name]]",
		Short: "Retrieve all default addresses in chain link with optional owner, chain name and pagination",
		Example: fmt.Sprintf(`%s query profiles chain-link-owners
%s query profiles default-external-address --page=2 --limit=100
%s query profiles default-external-address "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud"
%s query profiles default-external-address "desmos13p5pamrljhza3fp4es5m3llgmnde5fzcpq6nud" "cosmos"
`, version.AppName, version.AppName, version.AppName, version.AppName),
		Args: cobra.RangeArgs(0, 2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			var owner string
			if len(args) > 0 {
				owner = args[0]
			}

			var chainName string
			if len(args) > 1 {
				chainName = args[1]
			}

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			res, err := queryClient.DefaultExternalAddresses(
				context.Background(),
				types.NewQueryDefaultExternalAddressesRequest(owner, chainName, pageReq),
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, "default external address")
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
