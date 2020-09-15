package types

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Profiles             []Profile             `json:"profiles" yaml:"profiles"`
	Params               Params                `json:"params" yaml:"params"`
	DTagTransferRequests []DTagTransferRequest `json:"dtag_transfer_requests" yaml:"dtag_transfer_requests"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(profiles []Profile, params Params, request []DTagTransferRequest) GenesisState {
	return GenesisState{
		Profiles:             profiles,
		Params:               params,
		DTagTransferRequests: request,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{
		Profiles:             Profiles{},
		Params:               DefaultParams(),
		DTagTransferRequests: []DTagTransferRequest{},
	}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, profile := range data.Profiles {
		if err := profile.Validate(); err != nil {
			return err
		}
	}

	if err := data.Params.Validate(); err != nil {
		return err
	}

	for _, transferReq := range data.DTagTransferRequests {
		if err := transferReq.Validate(); err != nil {
			return err
		}
	}

	return nil
}
