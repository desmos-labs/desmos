package types

import "fmt"

// NewGenesisState creates a new genesis state
func NewGenesisState(reports []Report) *GenesisState {
	return &GenesisState{
		Reports: reports,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState([]Report{})
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, report := range data.Reports {
		if containDuplicates(data.Reports, report) {
			return fmt.Errorf("duplicate report: %s", report)
		}

		err := report.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}

// containDuplicates tells whether the given reports slice contain duplicates of the provided report
func containDuplicates(reports []Report, report Report) bool {
	var count = 0
	for _, r := range reports {
		if r.Equal(report) {
			count++
		}
	}
	return count > 1
}
