package types

// NewGenesisState creates a new genesis state
func NewGenesisState(
	profiles []Profile, request []DTagTransferRequest,
	params Params,
) *GenesisState {
	return &GenesisState{
		Profiles:             profiles,
		Params:               params,
		DtagTransferRequests: request,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, profile := range data.Profiles {
		err := profile.Validate()
		if err != nil {
			return err
		}
	}

	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, transferReq := range data.DtagTransferRequests {
		err := transferReq.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
