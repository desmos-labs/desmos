package types

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Posts Posts            `json:"posts"`
	Likes map[string]Likes `json:"likes"`
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, record := range data.Posts {
		if err := record.Validate(); err != nil {
			return err
		}
	}

	for _, likes := range data.Likes {
		for _, record := range likes {
			if err := record.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
