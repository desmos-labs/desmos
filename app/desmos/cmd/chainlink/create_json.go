package chainlink

import (
	"fmt"
	"io/ioutil"

	"github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/builder"
	chainlinktypes "github.com/desmos-labs/desmos/v3/app/desmos/cmd/chainlink/getter"

	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v3/app"
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

--- Single signature accounts ---
Note that this command will ask you the mnemonic that should be used to generate the private key of the address you want to link.
The mnemonic is only used temporarily and never stored anywhere.

--- Multi signature accounts ---
If you have are using a multi-signature account, you will be required to provide the path to a signed transaction file. 
That transaction must be signed as normal, except for the specified "account-number" and "sequence" values which should be both set to 0. 
Providing an invalid transaction (either with an account-number or sequence not set to 0, or not signed correctly) will result in a failing linkage later on.
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			isSingleSignatureAccount, err := getter.IsSingleSignatureAccount()
			if err != nil {
				return err
			}

			chain, err := getter.GetChain()
			if err != nil {
				return err
			}

			owner, err := getter.GetOwner()
			if err != nil {
				return err
			}

			chainLinkJSON, err := provider(owner, isSingleSignatureAccount).BuildChainLinkJSON(chain)
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
