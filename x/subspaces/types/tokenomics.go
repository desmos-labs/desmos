package types

import (
	"fmt"
	"strings"
)

// NewTokenomics is a constructor of the Tokenomics type
func NewTokenomics(subspaceID, contractAddress, admin string, message []byte) Tokenomics {
	return Tokenomics{
		SubspaceID:      subspaceID,
		ContractAddress: contractAddress,
		Admin:           admin,
		Message:         message,
	}
}

// Validate performs some checks on tokenomics to ensure its validity
func (tp Tokenomics) Validate() error {
	if !IsValidSubspace(tp.SubspaceID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid SHA-256 hash", tp.SubspaceID)
	}

	if strings.TrimSpace(tp.ContractAddress) == "" {
		return fmt.Errorf("contract address cannot be empty or blank")
	}

	if tp.Admin == "" {
		return fmt.Errorf("invalid admin address")
	}

	if tp.ContractAddress == tp.Admin {
		return fmt.Errorf("contract address and admin address cannot be the same")
	}

	if tp.Message == nil {
		return fmt.Errorf("empty message bytes")
	}

	return nil
}
