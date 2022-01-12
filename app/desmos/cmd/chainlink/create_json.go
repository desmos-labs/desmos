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

			isSingleSignatureAccount, err := getter.IsSingleSignatureAccount()
			if err != nil {
				return err
			}

			chain, err := getter.GetChain()
			if err != nil {
				return err
			}

			var chainLinkJSON profilescliutils.ChainLinkJSON
			if isSingleSignatureAccount {
				// Get the data
				mnemonic, err := getter.GetMnemonic()
				if err != nil {
					return err
				}

				// Build the chain link JSON
				chainLinkJSON, err = generateChainLinkJSON(mnemonic, chain)
				if err != nil {
					return err
				}
			} else {
				multiSignedTxFile, err := getter.GetMultiSignedTxFile()
				if err != nil {
					return err
				}
				signedChainID, err := getter.GetSignedChainID()
				if err != nil {
					return err
				}
				txCfg, txBuilder, txFactory, err := getMultisignedTxReference(cmd, multiSignedTxFile, signedChainID)
				if err != nil {
					return err
				}

				chainLinkJSON, err = getChainLinkJSONFromMultiSign(txCfg, txBuilder, txFactory, chain)
				if err != nil {
					return err
				}
			}

			cdc, _ := app.MakeCodecs()
			bz, err := cdc.MarshalJSON(&chainLinkJSON)
			if err != nil {
				return err
			}

			filename, err := getter.GetFilename()
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

// generateChainLinkJSON build a new ChainLinkJSON instance using the provided mnemonic and chain configuration
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

// getMultisignedTxReference returns the multisigned tx reference to build the chain link json
func getMultisignedTxReference(
	cmd *cobra.Command,
	txFile string,
	signedChainID string,
) (client.TxConfig, client.TxBuilder, tx.Factory, error) {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return nil, nil, tx.Factory{}, err
	}
	parsedTx, err := authclient.ReadTxFromFile(clientCtx, txFile)
	if err != nil {
		return nil, nil, tx.Factory{}, err
	}
	txCfg := clientCtx.TxConfig
	txBuilder, err := txCfg.WrapTxBuilder(parsedTx)
	if err != nil {
		return nil, txBuilder, tx.Factory{}, err
	}
	txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithChainID(signedChainID)
	if txFactory.SignMode() == signing.SignMode_SIGN_MODE_UNSPECIFIED {
		txFactory = txFactory.WithSignMode(signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON)
	}
	return txCfg, txBuilder, txFactory, err
}

// getChainLinkJSONFromMultiSign generates the chain-link JSON from the multisign tx reference
func getChainLinkJSONFromMultiSign(
	txCfg client.TxConfig,
	txBuilder client.TxBuilder,
	txFactory tx.Factory,
	chain chainlinktypes.Chain,
) (profilescliutils.ChainLinkJSON, error) {
	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}
	// make sure there is only one signature for the multisig account
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

	addr, err := sdk.Bech32ifyAddressBytes(chain.Prefix, multisigSig.PubKey.Address().Bytes())
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	sigData, err := profilestypes.SignatureDataFromCosmosSignatureData(multisigSig.Data)
	return profilescliutils.NewChainLinkJSON(
		profilestypes.NewBech32Address(addr, chain.Prefix),
		profilestypes.NewProof(multisigSig.PubKey, sigData, hex.EncodeToString(value)),
		profilestypes.NewChainConfig(chain.Name),
	), err
}
