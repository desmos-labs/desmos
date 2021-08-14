package types

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cosmos/go-bip39"

	"github.com/manifoldco/promptui"
)

// ChainLinkReferenceGetter allows to get all the data needed to generate a ChainLinkJSON instance
type ChainLinkReferenceGetter interface {
	// GetMnemonic returns the mnemonic
	GetMnemonic() (string, error)

	// GetChain returns Chain instance
	GetChain() (Chain, error)

	// GetFilename returns filename to save
	GetFilename() (string, error)
}

// ChainLinkReferencePrompt is a ChainLinkReferenceGetter implemented with an interactive prompt
type ChainLinkReferencePrompt struct {
	ChainLinkReferenceGetter
	cfg Config
}

// NewChainLinkReferencePrompt returns an instance implementing ChainLinkReferencePrompt
func NewChainLinkReferencePrompt() *ChainLinkReferencePrompt {
	return &ChainLinkReferencePrompt{
		cfg: DefaultConfig(),
	}
}

// GetMnemonic implements ChainLinkReferenceGetter
func (cp ChainLinkReferencePrompt) GetMnemonic() (string, error) {
	mnemonic, err := cp.getMnemonic()
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// GetChain implements ChainLinkReferenceGetter
func (cp ChainLinkReferencePrompt) GetChain() (Chain, error) {
	chain, err := cp.selectChain()
	if err != nil {
		return Chain{}, err
	}

	if chain.ID == "Other" {
		newChain, err := cp.getCustomChain(chain)
		if err != nil {
			return Chain{}, err
		}
		chain = newChain
	}

	return chain, nil
}

// GetFilename implements ChainLinkReferenceGetter
func (cp ChainLinkReferencePrompt) GetFilename() (string, error) {
	filename, err := cp.getFilename()
	if err != nil {
		return "", err
	}
	return filename, nil
}

// getMnemonic asks the user the mnemonic and then returns it
func (cp ChainLinkReferencePrompt) getMnemonic() (string, error) {
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

// selectChain asks the user to select a predefined Chain instance, and returns it
func (cp ChainLinkReferencePrompt) selectChain() (Chain, error) {
	cfg := cp.cfg
	prompt := promptui.Select{
		Label: "Select a target chain",
		Items: cfg.Chains,
		Templates: &promptui.SelectTemplates{
			Active:   "\U00002713 {{ .ID | cyan }}",
			Inactive: "  {{ .ID | cyan }}",
			Selected: "Module: \U00002713 {{ .ID | cyan }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return Chain{}, err
	}

	return cfg.Chains[index], nil
}

// getCustomChain asks the user to input the data to build a custom Chain instance, and then returns it
func (cp ChainLinkReferencePrompt) getCustomChain(chain Chain) (Chain, error) {
	chainName, err := cp.getChainName()
	if err != nil {
		return Chain{}, err
	}

	prefix, err := cp.getBech32Prefix()
	if err != nil {
		return Chain{}, err
	}

	derivationPath, err := cp.getDerivationPath()
	if err != nil {
		return Chain{}, err
	}

	chain.Name = chainName
	chain.Prefix = prefix
	chain.DerivationPath = derivationPath

	return chain, nil
}

// getChainName asks the user to input a chain name, and returns it
func (cp ChainLinkReferencePrompt) getChainName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please input the name of the chain you want to link with",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("chain name cannot be empty or blank")
			}
			if strings.ToLower(s) != s {
				return fmt.Errorf("chain name should be lowercase")
			}
			return nil
		},
	}
	return prompt.Run()
}

// getBech32Prefix asks the user to input the Bech32 prefix of the address, and then returns it
func (cp ChainLinkReferencePrompt) getBech32Prefix() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please input the Bech32 account address prefix used inside the the chain",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("bech32 prefix cannot be empty or blank")
			}
			return nil
		},
	}
	return prompt.Run()
}

// getDerivationPath asks the user to input the derivation path of the account, and then returns it
func (cp ChainLinkReferencePrompt) getDerivationPath() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Please input the derivation path used by the chain to generate the accounts",
		Default: "m/44'/118'/0'/0/0",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("derivation path cannot be empty or blank")
			}
			return nil
		},
	}
	return prompt.Run()
}

// getFilename asks the user to input the filename where to store the chain link, and then returns it
func (cp ChainLinkReferencePrompt) getFilename() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	prompt := promptui.Prompt{
		Label:   "Please insert where the chain link JSON object should be stored (fully qualified path)",
		Default: path.Join(wd, "data.json"),
	}
	return prompt.Run()
}
