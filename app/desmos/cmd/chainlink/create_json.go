package chainlink

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"

	chainlinktypes "github.com/desmos-labs/desmos/v2/app/desmos/cmd/chainlink/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
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
	value := []byte(addr)
	sig, pubkey, err := keyBase.Sign(keyName, value)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}
	sigData := &profilestypes.SingleSignatureData{
		Mode:      signing.SignMode_SIGN_MODE_DIRECT,
		Signature: sig,
	}

	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(pubkey, sigData, hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}

// getChainLinkJSONFromMultiSign generates the chain-link JSON from the multisign file and its raw transaction file
func getChainLinkJSONFromMultiSign(
	cmd *cobra.Command,
	txFile string,
	chain chainlinktypes.Chain,
) (profilescliutils.ChainLinkJSON, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	parsedTx, err := authclient.ReadTxFromFile(clientCtx, txFile)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txCfg := clientCtx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())
	if txFactory.SignMode() == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		txFactory = txFactory.WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	}

	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	if len(sigs) != 1 {
		return profilescliutils.ChainLinkJSON{}, fmt.Errorf("invalid number of signatures")
	}
	multisigSig := sigs[0]

	signingData := authsigning.SignerData{
		ChainID:       txFactory.ChainID(),
		AccountNumber: txFactory.AccountNumber(),
		Sequence:      txFactory.Sequence(),
	}
	// the bytes of plain text
	value, err := txCfg.SignModeHandler().GetSignBytes(txFactory.SignMode(), signingData, txBuilder.GetTx())
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	addr, _ := sdk.Bech32ifyAddressBytes(chain.Prefix, multisigSig.PubKey.Address().Bytes())
	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(multisigSig.PubKey, profilestypes.SignatureDataFromCosmosSignatureData(multisigSig.Data), hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), nil
}
