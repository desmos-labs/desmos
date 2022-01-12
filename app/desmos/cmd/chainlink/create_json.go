package chainlink

import (
	"fmt"
	"io/ioutil"

	"github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/builder"
	chainlinktypes "github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/getter"

	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v2/app"
)

// GetCreateChainLinkJSON returns the command allowing to generate the chain link JSON
// file that is required by the link-chain command
func GetCreateChainLinkJSON(
	getter chainlinktypes.ChainLinkReferenceGetter,
	provider builder.ChainLinkJSONBuilderProvider,
) *cobra.Command {
	return &cobra.Command{
		Use:   "create-chain-link-json",
		Short: "Start an interactive prompt to create a new chain link JSON object",
		Long: `Start an interactive prompt to create a new chain link JSON object that can be used to later link your Desmos profile to another chain.
Once you have built the JSON object using this command, you can then run the following command to complete the linkage:

desmos tx profiles link-chain [/path/to/json/file.json]

Note that this command will ask you the mnemonic that should be used to generate the private key of the address you want to link.
The mnemonic is only used temporarily and never stored anywhere.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			isSingleSignatureAccount, err := getter.IsSingleSignatureAccount()
			if err != nil {
				return err
			}

			chain, err := getter.GetChain()
			if err != nil {
				return err
			}

			chainLinkJSON, err := provider(isSingleSignatureAccount).BuildChainLinkJSON(chain)
			if err != nil {
				return err
			}

			// Marshal the chain link JSON
			bz, err := app.MakeTestEncodingConfig().Marshaler.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			// Write the chain link JSON to a file
			filename, err := getter.GetFilename()
			if err != nil {
				return err
			}

			err = ioutil.WriteFile(filename, bz, 0600)
			if err != nil {
				return err
			}

			cmd.Println(fmt.Sprintf("Chain link JSON file stored at %s", filename))
			return nil
		},
	}
}
