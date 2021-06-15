package types

import (
	"fmt"

	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(
	requests []DTagTransferRequest, relationships []Relationship, blocks []UserBlock, params Params, portID string, chainLinks []ChainLinkEntry,
) *GenesisState {
	return &GenesisState{
		Params:               params,
		DTagTransferRequests: requests,
		Relationships:        relationships,
		Blocks:               blocks,
		IBCPortID:            portID,
		ChainLinks:           chainLinks,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, DefaultParams(), IBCPortID, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, req := range data.DTagTransferRequests {
		err := req.Validate()
		if err != nil {
			return err
		}
	}

	for _, rel := range data.Relationships {
		if containDuplicates(data.Relationships, rel) {
			return fmt.Errorf("duplicated relationship: %s", rel)
		}

		err := rel.Validate()
		if err != nil {
			return err
		}
	}

	for _, ub := range data.Blocks {
		err := ub.Validate()
		if err != nil {
			return err
		}
	}

	for _, entry := range data.ChainLinks {
		err := entry.Link.Validate()
		if err != nil {
			return err
		}
	}

	if err := host.PortIdentifierValidator(data.IBCPortID); err != nil {
		return err
	}

	return nil
}

// containDuplicates tells whether the given relationships slice contain duplicates of the provided relationship
func containDuplicates(relationships []Relationship, relationship Relationship) bool {
	var count = 0
	for _, r := range relationships {
		if r.Equal(relationship) {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewChainLinkEntry allows to build a new ChainLinkEntry instance
func NewChainLinkEntry(user string, link ChainLink) ChainLinkEntry {
	return ChainLinkEntry{
		User: user,
		Link: link,
	}
}
