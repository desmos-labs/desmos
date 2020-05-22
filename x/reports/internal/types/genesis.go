package types

import "fmt"

// GenesisState contains the data of the genesis state for the posts module
type GenesisState struct {
	Reports      map[string]Reports `json:"reports" yaml:"reports"`
	ReportsTypes ReportTypes        `json:"reports_types" yaml:"reports"`
}

// NewGenesisState creates a new genesis state
func NewGenesisState(reports map[string]Reports, reportsTypes ReportTypes) GenesisState {
	return GenesisState{
		Reports:      reports,
		ReportsTypes: reportsTypes,
	}
}

// DefaultGenesisState returns a default GenesisState
func DefaultGenesisState() GenesisState {
	return GenesisState{}
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data GenesisState) error {
	for _, rT := range data.ReportsTypes {
		if rT.Empty() {
			return fmt.Errorf("report type cannot be empty")
		}
	}

	for _, reports := range data.Reports {
		if err := reports.Validate(); err != nil {
			return err
		}
	}
	return nil
}
