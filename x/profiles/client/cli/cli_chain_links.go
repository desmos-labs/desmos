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

	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetCmdLinkChainAccount returns the command allowing to link an external chain account
func GetCmdLinkChainAccount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "link [data-file]",
		Args:  cobra.ExactArgs(1),
		Short: "Link a Desmos profile to an external chain account using the data written inside the given file",
		Long: strings.TrimSpace(fmt.Sprintf(`Link an external account to a Desmos profile.
The link data must be supplied via a JSON file.

Example:
$ %s tx profiles link <path/to/data.json> --from=<key_or_address>

Where data.json contains:

{
   "address":{
      "@type":"/desmos.profiles.v1beta1.Bech32Address",
      "value":"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm",
      "prefix":"cosmos"
   },
   "proof":{
      "pub_key":{
         "@type":"/cosmos.crypto.secp256k1.PubKey",
         "key":"A58DXR/lXKVkIjLofXgST/OHi+pkOQbVIiOjnTy7Zoqo"
      },
      "signature":"ecc6175e730917fb289d3a9f4e49a5630a44b42d972f481342f540e09def2ec5169780d85c4e060d52cc3ffb3d677745a4d56cd385760735bc6db0f1816713be",
      "plain_text":"cosmos15uc89vnzufu5kuhhsxdkltt38zfx8vcyggzwfm"
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

			data, err := profilescliutils.ParseChainLinkJSON(clientCtx.JSONMarshaler, args[0])
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
		Use:   "unlink [chain-name] [address]",
		Short: "Unlink the external account having the given chain name and address",
		Args:  cobra.ExactArgs(2),
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

// --------------------------------------------------------------------------------------------------------------------

// GetCmdQueryProfileByChainLink returns the command allowing to query a profile
// given an external chain name and address
func GetCmdQueryProfileByChainLink() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "chain-link [chain-name] [address]",
		Short: "Get the profile linked with the external account having the specified chain name and address",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.ProfileByChainLink(
				context.Background(),
				&types.QueryProfileByChainLinkRequest{ChainName: args[0], TargetAddress: args[1]},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
