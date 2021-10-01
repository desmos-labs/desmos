package chainlink

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	chainlinktypes "github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/v2/app"
	profilescliutils "github.com/desmos-labs/desmos/v2/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
)

// GetCreateChainLinkJSON returns the command allowing to generate the chain link JSON
// file that is required by the link-chain command
func GetCreateChainLinkJSON(getter chainlinktypes.ChainLinkReferenceGetter) *cobra.Command {
	return &cobra.Command{
		Use:   "create-chain-link-json",
		Short: "Start an interactive prompt to create a new chain link JSON object",
		Long: `Start an interactive prompt to create a new chain link JSON object that can be used to later link your Desmos profile to another chain.
Once you have built the JSON object using this command, you can then run the following command to complete the linkage:

desmos tx profiles link-chain [/path/to/json/file.json]

Note that this command will ask you the mnemonic that should be used to generate the private key of the address you want to link.
The mnemonic is only used temporarily and never stored anywhere.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the data
			mnemonic, err := getter.GetMnemonic()
			if err != nil {
				return err
			}

			chain, err := getter.GetChain()
			if err != nil {
				return err
			}

			filename, err := getter.GetFilename()
			if err != nil {
				return err
			}

			// Build che chain link JSON
			chainLinkJSON, err := generateChainLinkJSON(mnemonic, chain)
			if err != nil {
				return err
			}

			cdc, _ := app.MakeCodecs()
			bz, err := cdc.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			if filename != "" {
				err = ioutil.WriteFile(filename, bz, 0600)
				if err != nil {
					return err
				}
			}

			cmd.Println(fmt.Sprintf("Chain link JSON file stored at %s", filename))

			return nil
		},
	}
}

// generateChainLinkJSON returns build a new ChainLinkJSON intance using the provided mnemonic and chain configuration
func generateChainLinkJSON(mnemonic string, chain chainlinktypes.Chain) (profilescliutils.ChainLinkJSON, error) {
	// Create an in-memory keybase for signing
	keyBase := keyring.NewInMemory()
	keyName := "chainlink"
	_, err := keyBase.NewAccount(keyName, mnemonic, "", chain.DerivationPath, hd.Secp256k1)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	// Generate the proof signing it with the key
	key, _ := keyBase.Key(keyName)
	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, key.GetAddress())
	sig, pubkey, err := keyBase.Sign(keyName, []byte(addr))
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, hex.EncodeToString(sig), addr),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
