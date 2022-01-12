package multi

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/manifoldco/promptui"
)

// ChainLinkJSONInfoGetter implements MultiSignatureAccountReferenceGetter
type ChainLinkJSONInfoGetter struct {
}

// NewChainLinkJSONInfoGetter returns a new ChainLinkJSONInfoGetter instance
func NewChainLinkJSONInfoGetter() *ChainLinkJSONInfoGetter {
	return &ChainLinkJSONInfoGetter{}
}

// getMultiSignedTxFile asks the user the path of the multisigned transaction file, and then returns it
func (cp *ChainLinkJSONInfoGetter) getMultiSignedTxFile() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	prompt := promptui.Prompt{
		Label:   "Please insert the path of multisigned tx file (fully qualified path)",
		Default: path.Join(wd, "tx.json"),
	}
	return prompt.Run()
}

// GetMultiSignedTxFilePath implements MultiSignatureAccountReferenceGetter
func (cp *ChainLinkJSONInfoGetter) GetMultiSignedTxFilePath() (string, error) {
	return cp.getMultiSignedTxFile()
}

// getSignedChainID asks the user the chain id that is used to sign the transaction file, and then returns it
func (cp *ChainLinkJSONInfoGetter) getSignedChainID() (string, error) {
	prompt := promptui.Prompt{
		Label:       "Please enter the chain id that is used to sign the multisigned transaction file",
		HideEntered: true,
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("signed chain id cannot be empty or blank")
			}
			return nil
		},
	}
	return prompt.Run()
}

// GetSignedChainID implements MultiSignatureAccountReferenceGetter
func (cp *ChainLinkJSONInfoGetter) GetSignedChainID() (string, error) {
	return cp.getSignedChainID()
}
