package types

import "fmt"

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
		if containDuplicates(data.Profiles, profile) {
			return fmt.Errorf("duplicated profile: %s", profile)
		}

		err := profile.Validate()
		if err != nil {
			return err
		}
	}

	err := data.Params.Validate()
	if err != nil {
		return err
	}

	for _, req := range data.DtagTransferRequests {
		if !profileExists(data.Profiles, req.Sender) {
			return fmt.Errorf("invalid DTag transfer request; sender does not exist: %s", req.Sender)
		}

		if !profileExists(data.Profiles, req.Receiver) {
			return fmt.Errorf("invalid DTag transfer request; receiver does not exist: %s", req.Receiver)
		}

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
		if p.Equal(profile) {
			count += 1
		}
	}
	return count > 1
}

// profileExists tells whether the given profiles slice contain a profile associated to the given address
func profileExists(profiles []Profile, address string) bool {
	for _, profile := range profiles {
		if profile.Creator == address {
			return true
		}
	}
	return false
}
