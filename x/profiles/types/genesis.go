package types

// NewGenesisState creates a new genesis state
func NewGenesisState(
	request []DTagTransferRequest,
	params Params,
) *GenesisState {
	return &GenesisState{
		Params:               params,
		DtagTransferRequests: request,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, req := range data.DtagTransferRequests {
		err := req.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}

// containDuplicates tells whether or not the profiles slice contain duplicates of the given profile
func containDuplicates(profiles []Profile, profile Profile) bool {
	var count = 0
	for _, p := range profiles {
		if p.GetAddress().Equals(profile.GetAddress()) || p.Dtag == profile.Dtag {
			count++
		}
	}
	return count > 1
}

// profileExists tells whether the given profiles slice contain a profile associated to the given address
func profileExists(profiles []Profile, address string) bool {
	for _, profile := range profiles {
		if profile.GetAddress().String() == address {
			return true
		}
	}
	return false
}
