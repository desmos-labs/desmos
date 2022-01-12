package single

import (
	"fmt"
	"strings"

	"github.com/cosmos/go-bip39"
	"github.com/manifoldco/promptui"
)

// ChainLinkJSONInfoGetter implements SingleSignatureAccountReferenceGetter
type ChainLinkJSONInfoGetter struct {
}

// NewChainLinkJSONInfoGetter returns a new ChainLinkJSONInfoGetter instance
func NewChainLinkJSONInfoGetter() *ChainLinkJSONInfoGetter {
	return &ChainLinkJSONInfoGetter{}
}

// getMnemonic asks the user the mnemonic and then returns it
func (cp *ChainLinkJSONInfoGetter) getMnemonic() (string, error) {
	prompt := promptui.Prompt{
		Label:       "Please enter the mnemonic that should be used to generate the address you want to link",
		HideEntered: true,
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("mnemonic cannot be empty or blank")
			} else if _, err := bip39.MnemonicToByteArray(s); err != nil {
				return fmt.Errorf("invalid mnemonic")
			}
			return nil
		},
	}
	return prompt.Run()
}

// GetMnemonic implements SingleSignatureAccountReferenceGetter
func (cp *ChainLinkJSONInfoGetter) GetMnemonic() (string, error) {
	mnemonic, err := cp.getMnemonic()
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}
