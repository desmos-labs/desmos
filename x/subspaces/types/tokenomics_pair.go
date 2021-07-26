package types

import (
	"fmt"
	"strings"
)

// NewTokenomicsPair is a constructor of the TokenomicsPair type
func NewTokenomicsPair(subspaceID, contractAddress, admin string) TokenomicsPair {
	return TokenomicsPair{
		SubspaceID:      subspaceID,
		ContractAddress: contractAddress,
		Admin:           admin,
	}
}

// Validate performs some checks on tokenomics pair to ensure its validity
func (tp TokenomicsPair) Validate() error {
	if !IsValidSubspace(tp.SubspaceID) {
		return fmt.Errorf("invalid subspace id: %s it must be a valid SHA-256 hash", tp.SubspaceID)
	}

	if strings.TrimSpace(tp.ContractAddress) == "" {
		return fmt.Errorf("contract address cannot be empty or blank")
	}

	if tp.Admin == "" {
		return fmt.Errorf("invalid admin address")
	}

	return nil
}
