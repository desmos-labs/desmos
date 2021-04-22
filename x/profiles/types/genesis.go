package types

// NewGenesisState creates a new genesis state
func NewGenesisState(
	request []DTagTransferRequest,
	params Params,
) *GenesisState {
	return &GenesisState{
		Params:              params,
		DTagTransferRequest: request,
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

	for _, req := range data.DTagTransferRequest {
		err := req.Validate()
		if err != nil {
			return err
		}
	}

	return nil
}
