package cmd

import (
	"encoding/hex"
	"io/ioutil"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/desmos/app"
	"github.com/desmos-labs/desmos/app/desmos/types"

	profilescliutils "github.com/desmos-labs/desmos/x/profiles/client/utils"
	profilestypes "github.com/desmos-labs/desmos/x/profiles/types"
)

// GetCreateChainlinkJSON returns the command allowing to generate the chain link json file for creating chain link
func GetCreateChainlinkJSON(generator types.ChainLinkReferenceGetter) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain-link-json",
		Short: "Generate the chain link json for creating chain link with the key",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			mnemonic, chain, err := generator.GetReference()
			if err != nil {
				return err
			}

			chainLinkJSON, err := generateChainLink(mnemonic, chain)
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
	return cmd
}

// generateChainLink returns ChainLinkJSON for creating chain link
func generateChainLink(mnemonic string, chain types.ChainType) (profilescliutils.ChainLinkJSON, error) {
	// generate keybase for signing
	keyBase := keyring.NewInMemory()
	keyName := "chainlink"
	_, err := keyBase.NewAccount("chainlink", mnemonic, "", chain.DerivationPath, hd.Secp256k1)
	if err != nil {
		return profilescliutils.ChainLinkJSON{}, err
	}

	// create the proof with the key
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
