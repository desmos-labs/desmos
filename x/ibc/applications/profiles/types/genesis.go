package types

import host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"

// NewGenesisState creates a new ibc-profiles GenesisState instance.
func NewGenesisState(portID string) *GenesisState {
	return &GenesisState{
		PortId: portID,
	}
}

// DefaultGenesisState returns a GenesisState with "profiles" as the default PortID.
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		PortId: PortID,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return host.PortIdentifierValidator(gs.PortId)
}
