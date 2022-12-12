package sign

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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
		Short: "Allows to sign the given value using the private key associated to the address or key specified using the --from flag",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			f := cmd.Flags()
			txFactory := tx.NewFactoryCLI(clientCtx, f)

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

			// Sign the data with the private key
			value := []byte(args[0])
			bz, pubKey, err := txFactory.Keybase().Sign(key.GetName(), value)
			if err != nil {
				return err
			}

			// Build the signature data output
			signatureData := SignatureData{
				Address:   strings.ToLower(pubKey.Address().String()),
				Signature: strings.ToLower(hex.EncodeToString(bz)),
				PubKey:    strings.ToLower(hex.EncodeToString(pubKey.Bytes())),
				Value:     hex.EncodeToString(value),
			}

			// Serialize the output as JSON and print it
			bz, err = json.Marshal(&signatureData)
			if err != nil {
				return err
			}

			return clientCtx.PrintBytes(bz)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
