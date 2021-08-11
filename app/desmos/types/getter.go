package types

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

// ChainLinkReferenceGetter is an interface to get mnemonic and chain type
type ChainLinkReferenceGetter interface {
	// GetReference returns mnemonic and ChainType instance for creating ChainLinkJSON
	GetReference() (string, ChainType, error)
}

// ChainLinkReferencePrompt is a ChainTypeGetter implemented by promptui
type ChainLinkReferencePrompt struct {
	ChainLinkReferenceGetter
	cfg Config
}

// NewChainTypePrompt returns a ChainTypePrompt instance
func NewChainTypePrompt(cfg Config) *ChainLinkReferencePrompt {
	return &ChainLinkReferencePrompt{cfg: cfg}
}

// GetMnemonicAndChainType returns mnemonic and ChainType instance from the prompt
func (cp ChainLinkReferencePrompt) GetReference() (string, ChainType, error) {
	mnemonic, _ := cp.getMnemonic()
	chain, _ := cp.selectChain()

	if chain.ID == "Other" {
		newChain, err := cp.getCustomChain(chain)
		if err != nil {
			return "", ChainType{}, err
		}
		chain = newChain
	}

	return mnemonic, chain, nil
}

// getMnemonic returns mnemonic from the prompt
func (cp ChainLinkReferencePrompt) getMnemonic() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please enter your mnemonic",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("mnemonic cannot be empty or blank")
			}
			return nil
		},
	}
	return prompt.Run()
}

// selectChain returns ChainType instance from the prompt
func (cp ChainLinkReferencePrompt) selectChain() (ChainType, error) {
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
		return ChainType{}, err
	}

	return cfg.Chains[index], nil
}

// getCustomChain returns ChainType instance not in the default config from the prompt
func (cp ChainLinkReferencePrompt) getCustomChain(chain ChainType) (ChainType, error) {
	chainName, err := cp.getChainName()
	if err != nil {
		return ChainType{}, err
	}

	prefix, err := cp.getBech32Prefix()
	if err != nil {
		return ChainType{}, err
	}

	derivationPath, err := cp.getDerivationPath()
	if err != nil {
		return ChainType{}, err
	}

	chain.Name = chainName
	chain.Prefix = prefix
	chain.DerivationPath = derivationPath

	return chain, nil
}

// getChainName returns chain name from the prompt
func (cp ChainLinkReferencePrompt) getChainName() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please input the name of the chain",
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

// getBech32Prefix returns bech32 prefix from the prompt
func (cp ChainLinkReferencePrompt) getBech32Prefix() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please input the bech32 prefix of the chain",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("bech32 prefix cannot be empty or blank")
			}
			return nil
		},
	}
	return prompt.Run()
}

// getDerivationPath returns derivation path from the prompt
func (cp ChainLinkReferencePrompt) getDerivationPath() (string, error) {
	prompt := promptui.Prompt{
		Label:   "Please input the derivation path of the chain",
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
