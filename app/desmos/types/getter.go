package types

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
)

// ChainLinkReferenceGetter is an interface to get reference for creating ChainLinkJSON
type ChainLinkReferenceGetter interface {
	// GetMnemonic returns the mnemonic
	GetMnemonic() (string, error)

	// GetChain returns Chain instance
	GetChain() (Chain, error)

	// GetFilename returns filename to save
	GetFilename() (string, error)
}

// ChainLinkReferencePrompt is a ChainGetter implemented by promptui
type ChainLinkReferencePrompt struct {
	ChainLinkReferenceGetter
	cfg Config
}

// NewChainLinkReferencePrompt returns an instance implementing ChainLinkReferencePrompt
func NewChainLinkReferencePrompt(cfg Config) *ChainLinkReferencePrompt {
	return &ChainLinkReferencePrompt{cfg: cfg}
}

func (cp ChainLinkReferencePrompt) GetMnemonic() (string, error) {
	mnemonic, err := cp.getMnemonic()
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}

// GetChain returns Chain instance from the prompt
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

// GetFilename returns filename to save from the prompt
func (cp ChainLinkReferencePrompt) GetFilename() (string, error) {
	filename, err := cp.getFilename()
	if err != nil {
		return "", err
	}
	return filename, nil
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

// selectChain returns Chain instance from the prompt
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

// getCustomChain returns Chain instance not in the default config from the prompt
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

// getFilename returns filename to save from the prompt
func (cp ChainLinkReferencePrompt) getFilename() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	prompt := promptui.Prompt{
		Label:   "Please input the output filename if provided",
		Default: path.Join(wd, "data.json"),
	}
	return prompt.Run()
}
