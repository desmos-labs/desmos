package cmd

import (
	"encoding/hex"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/desmos-labs/desmos/app"
	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	"github.com/desmos-labs/desmos/x/profiles/types"
)

// GetGenerateChainlinkJsonCmd returns the command allowing to generate the chain link json file for creating chain link
func GetGenerateChainlinkJsonCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate-chain-link-json",
		Short: "generate the chain link json for creating chain link with the key specified using the --from flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			err, chainLinkJson := GenerateChainLinkJson(
				clientCtx,
				app.Bech32MainPrefix,
			)

			cdc, _ := app.MakeCodecs()
			bz, err := cdc.MarshalJSON(&chainLinkJson)
			if err != nil {
				return err
			}

			filename, _ := cmd.Flags().GetString("filename")
			if strings.TrimSpace(filename) != "" {
				if err := ioutil.WriteFile("data.json", bz, 0644); err != nil {
					return err
				}
			}
			return clientCtx.PrintBytes(bz)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	cmd.Flags().String("filename", "data.json", "The name of output chain link json file. It does not generate the file if it is empty.")
	return cmd
}

func GenerateChainLinkJson(clientCtx client.Context, prefix string) (error, profilescliutils.ChainLinkJSON) {

	// generate signature
	addr, _ := sdk.Bech32ifyAddressBytes(app.Bech32MainPrefix, clientCtx.GetFromAddress())
	sig, pubkey, err := clientCtx.Keyring.Sign(clientCtx.GetFromName(), []byte(addr))
	if err != nil {
		return err, profilescliutils.ChainLinkJSON{}
	}

	// create chain link json
	cdc, _ := app.MakeCodecs()
	chainLinkJson := profilescliutils.NewChainLinkJSON(
		types.NewBech32Address(addr, app.Bech32MainPrefix),
		types.NewProof(pubkey, hex.EncodeToString(sig), addr),
		types.NewChainConfig(app.Bech32MainPrefix),
	)
	if err := chainLinkJson.UnpackInterfaces(cdc); err != nil {
		return err, profilescliutils.ChainLinkJSON{}
	}
	return nil, chainLinkJson
}
