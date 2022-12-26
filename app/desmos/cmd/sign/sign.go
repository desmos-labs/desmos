package sign

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/spf13/cobra"
)

type SignatureData struct {
	Address   string `json:"address"`
	PubKey    string `json:"pub_key"`
	Signature string `json:"signature"`
	Value     string `json:"value"`
}

// GetSignCmd returns the command allowing to sign an arbitrary for later verification
func GetSignCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign [value]",
		Short: "Sign the given value using the private key associated to either the address or the private key provided with the --from flag",
		Long: `
Sign the given value using the private key associated to either the address or the private key provided with the --from flag.

If the provided address/key name is associated to a key that leverages a Ledger device, the signed value will be placed inside the memo field of a transaction before being signed.
Otherwise, the provided value will be converted to raw bytes and then signed without any further transformation.

In both cases, after the signature the following data will be printed inside a JSON object:
- the hex-encoded address associated to the key used to sign the value
- the hex-encoded public key associated to the private key used to sign the value
- the hex-encoded signed value 
- the hex-encoded signature value

The printed JSON object can be safely used as the verification proof when connecting a Desmos profile to a centralized application.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Build a tx factory
			txFactory := tx.NewFactoryCLI(clientCtx, cmd.Flags())

			// Get the value of the "from" flag
			from, _ := cmd.Flags().GetString(flags.FlagFrom)
			_, fromName, _, err := client.GetFromFields(clientCtx, txFactory.Keybase(), from)
			if err != nil {
				return fmt.Errorf("error getting account from keybase: %w", err)
			}

			// Get the key from the keybase
			key, err := txFactory.Keybase().Key(fromName)
			if err != nil {
				return err
			}

			// Sign the value based on the signing mode
			var valueBz, sigBz []byte
			if txFactory.SignMode() == signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON {
				valueBz, sigBz, err = signAmino(clientCtx, txFactory, key, args[0])
			} else {
				valueBz, sigBz, err = signRaw(txFactory, key, args[0])
			}

			if err != nil {
				return err
			}

			// Build the signature data output
			pubKey := key.GetPubKey()
			signatureData := SignatureData{
				Address:   strings.ToLower(pubKey.Address().String()),
				Signature: strings.ToLower(hex.EncodeToString(sigBz)),
				PubKey:    strings.ToLower(hex.EncodeToString(pubKey.Bytes())),
				Value:     hex.EncodeToString(valueBz),
			}

			// Serialize the output as JSON and print it
			bz, err := json.Marshal(&signatureData)
			if err != nil {
				return err
			}
			return clientCtx.PrintBytes(bz)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// signRaw signs the given value directly by converting it into raw bytes
func signRaw(txFactory tx.Factory, key keyring.Info, value string) (valueBz []byte, sigBz []byte, err error) {
	valueBz = []byte(value)
	sigBz, _, err = txFactory.Keybase().Sign(key.GetName(), valueBz)
	return valueBz, sigBz, err
}

// signAmino puts the given value into a transaction memo field, and signs the transaction using the Amino encoding
func signAmino(clientCtx client.Context, txFactory tx.Factory, key keyring.Info, value string) (valueBz []byte, sigBz []byte, err error) {
	// Set a fake chain id
	txFactory = txFactory.WithChainID("desmos")

	// Set the memo to be the value to be signed
	txFactory = txFactory.WithMemo(value)

	// Build the fake transaction
	txBuilder, err := txFactory.BuildUnsignedTx()
	if err != nil {
		return
	}

	// Sign the data with the private key
	err = tx.Sign(txFactory, key.GetName(), txBuilder, true)
	if err != nil {
		return
	}

	// Encode the transaction
	signMode := signing.SignMode_SIGN_MODE_LEGACY_AMINO_JSON
	signerData := authsigning.SignerData{
		ChainID:       txFactory.ChainID(),
		AccountNumber: txFactory.AccountNumber(),
		Sequence:      txFactory.Sequence(),
	}
	valueBz, err = clientCtx.TxConfig.SignModeHandler().GetSignBytes(signMode, signerData, txBuilder.GetTx())
	if err != nil {
		return
	}

	// Get the signature bytes
	sigs, err := txBuilder.GetTx().GetSignaturesV2()
	if err != nil {
		return
	}
	sigBz = sigs[0].Data.(*signing.SingleSignatureData).Signature
	return valueBz, sigBz, nil
}
