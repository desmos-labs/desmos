package types

import (
	"fmt"
	"strings"
)

// ChainLink represents the data that have been used to prove an external chain account is owned by this user
type ChainLink struct {
	Name   string `json:"chain_name"` // Name of the chain on which the tx has been performed (eg. cosmoshub-3)
	TxHash string `json:"tx_hash"`    // Hex representation of the tx hash used to verify the ownership
}

// NewChainLink is a constructor for a new ChainLink
func NewChainLink(name, txHash string) ChainLink {
	return ChainLink{
		Name:   name,
		TxHash: txHash,
	}
}

// Validate implements validator
func (cl ChainLink) Validate() error {
	if len(strings.TrimSpace(cl.Name)) == 0 {
		return fmt.Errorf("chain name cannot be empty or blank")
	}

	if !TxHashRegEx.MatchString(cl.TxHash) {
		return fmt.Errorf("transaction hash of %s chain must be a valid sha-256 hash", cl.Name)
	}

	return nil
}

// Equals allows to check whether the contents of cl are the same of other
func (cl ChainLink) Equals(other ChainLink) bool {
	return cl.Name == other.Name &&
		cl.TxHash == other.TxHash
}
