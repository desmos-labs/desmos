package cmd

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/desmos-labs/desmos/app"
	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetGenerateChainlinkJSONCmd returns the command allowing to generate the chain link json file for creating chain link
func GetGenerateChainlinkJSONCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-chain-link-json",
		Short: "Generate the chain link json for creating chain link with the key specified using the --from flag",
		Long: strings.TrimSpace(fmt.Sprintf(`Generate the chain link json for creating chain link.
If you want to create chain link with the key from other chain, you have to use --prefix and --target-chain flags.
E.g.
$ %s generate-chain-link-json --prefix cosmos --target-chain cosmos --from <your-key> --home ~/.gaia`, version.AppName)),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			prefix, _ := cmd.Flags().GetString("prefix")
			targetChain, _ := cmd.Flags().GetString("target-chain")
			chainLinkJSON, err := GenerateChainLinkJSON(
				clientCtx,
				prefix,
				targetChain,
			)
			if err != nil {
				return err
			}

			cdc, _ := app.MakeCodecs()
			bz, err := cdc.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			filename, _ := cmd.Flags().GetString("filename")
			if strings.TrimSpace(filename) != "" {
				if err := ioutil.WriteFile("data.json", bz, 0600); err != nil {
					return err
				}
			}
			return clientCtx.PrintBytes(bz)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String("filename", "data.json", "The name of output chain link json file. It does not generate the file if it is empty.")
	cmd.Flags().String("prefix", app.Bech32MainPrefix, "The bech32 prefix of the target chain.")
	cmd.Flags().String("target-chain", app.Bech32MainPrefix, "The name of the target chain.")
	return cmd
}

// GenerateChainLinkJSON returns ChainLinkJSON instance for creating chain link
func GenerateChainLinkJSON(clientCtx client.Context, prefix, chainName string) (profilescliutils.ChainLinkJSON, error) {

	// generate signature
	addr, _ := sdk.Bech32ifyAddressBytes(prefix, clientCtx.GetFromAddress())
	sig, pubkey, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte(addr))
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	// create chain link json
	cdc, _ := app.MakeCodecs()
	chainLinkJSON := profilescliutils.NewChainLinkJSON(
		types.NewBech32Address(addr, prefix),
		types.NewProof(pubkey, hex.EncodeToString(sig), addr),
		types.NewChainConfig(chainName),
	)
	if err := chainLinkJSON.UnpackInterfaces(cdc); err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}
	return chainLinkJSON, nil
}
