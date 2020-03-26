package types

// GenesisState contains the data of the genesis state for the profile module
type GenesisState struct {
	Accounts []Account `json:"accounts"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(accounts []Account) GenesisState {
	return GenesisState{Accounts: accounts}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, account := range data.Accounts {
		if err := account.Validate(); err != nil {
			return err
		}
	}
	return nil
}
