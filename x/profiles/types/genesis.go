package types

import (
	"fmt"

	host "github.com/cosmos/ibc-go/v2/modules/core/24-host"
)

// NewGenesisState creates a new genesis state
func NewGenesisState(
	requests []DTagTransferRequest, relationships []Relationship, blocks []UserBlock,
	params Params, portID string,
	chainLinks []ChainLink, applicationLinks []ApplicationLink,
) *GenesisState {
	return &GenesisState{
		Params:               params,
		DTagTransferRequests: requests,
		Relationships:        relationships,
		Blocks:               blocks,
		IBCPortID:            portID,
		ChainLinks:           chainLinks,
		ApplicationLinks:     applicationLinks,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, DefaultParams(), IBCPortID, nil, nil)
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, req := range data.DTagTransferRequests {
		err = req.Validate()
		if err != nil {
			return err
		}
	}

	for i, rel := range data.Relationships {
		if containDuplicates(data.Relationships, rel) {
			return fmt.Errorf("duplicated relationship: %s", &data.Relationships[i])
		}

		err = rel.Validate()
		if err != nil {
			return err
		}
	}

	for _, ub := range data.Blocks {
		err = ub.Validate()
		if err != nil {
			return err
		}
	}

	err = host.PortIdentifierValidator(data.IBCPortID)
	if err != nil {
		return err
	}

	for _, l := range data.ChainLinks {
		err := l.Validate()
		if err != nil {
			return err
		}
	}

	for _, link := range data.ApplicationLinks {
		err = link.Validate()
		if err != nil {
			return err
		}
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
