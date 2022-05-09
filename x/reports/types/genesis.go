package types

import "fmt"

// NewGenesisState returns a new genesis state instance
func NewGenesisState(subspaces []SubspaceData, reports []Report, params Params) *GenesisState {
	return &GenesisState{
		SubspacesData: subspaces,
		Reports:       reports,
		Params:        params,
	}
}

// GetSubspaceReportID returns the next report id associated to the given subspace
func (data *GenesisState) GetSubspaceReportID(subspaceID uint64) uint64 {
	for _, subspace := range data.SubspacesData {
		if subspace.SubspaceID == subspaceID {
			return subspace.ReportID
		}
	}

	return 0
}

// DefaultGenesisState returns a DefaultGenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, subspaceData := range data.SubspacesData {
		if ContainsDuplicatedSubspacesData(data.SubspacesData, subspaceData) {
			return fmt.Errorf("duplicated subspace entry for subspace id %d", subspaceData.SubspaceID)
		}

		err := subspaceData.Validate()
		if err != nil {
			return fmt.Errorf("invalid subspace data: %s", err)
		}
	}

	for _, report := range data.Reports {
		if ContainsDuplicatedReport(data.Reports, report) {
			return fmt.Errorf("duplicated report for subspace id %d with id %d", report.SubspaceID, report.ID)
		}

		reportID := data.GetSubspaceReportID(report.SubspaceID)
		if reportID == 0 {
			return fmt.Errorf("next report id not found for subspace %d", report.SubspaceID)
		}

		if report.ID >= reportID {
			return fmt.Errorf("report id must be lower than next report id: %d", report.ID)
		}

		err := report.Validate()
		if err != nil {
			return fmt.Errorf("invalid report: %s", err)
		}
	}

	return data.Params.Validate()
}

// ContainsDuplicatedSubspacesData tells whether the given subspaces data slice
// contains a duplicated entry for the same subspace id as the one given
func ContainsDuplicatedSubspacesData(subspaces []SubspaceData, data SubspaceData) bool {
	var count = 0
	for _, r := range subspaces {
		if r.SubspaceID == data.SubspaceID {
			count++
		}
	}
	return count > 1
}

// ContainsDuplicatedReport tells whether the given reports slice contains
// a duplicated report based on the same subspace id and report id of the given one
func ContainsDuplicatedReport(reports []Report, report Report) bool {
	var count = 0
	for _, r := range reports {
		if r.SubspaceID == report.SubspaceID && r.ID == report.ID {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspacesData returns a new SubspacesData instance
func NewSubspacesData(subspaceID uint64, reportID uint64, reasonID uint32, reasons []Reason) SubspaceData {
	return SubspaceData{
		SubspaceID: subspaceID,
		ReportID:   reportID,
		ReasonID:   reasonID,
		Reasons:    reasons,
	}
}

// Validate implements fmt.Validator
func (data SubspaceData) Validate() error {
	if data.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", data.SubspaceID)
	}

	if data.ReportID == 0 {
		return fmt.Errorf("invalid report id: %d", data.ReportID)
	}

	if data.ReasonID == 0 {
		return fmt.Errorf("invalid reason id: %d", data.ReasonID)
	}

	// Check reason ids
	for _, reason := range data.Reasons {
		if reason.ID >= data.ReasonID {
			return fmt.Errorf("reason id must be lower than next reason id: %d", reason.ID)
		}
	}

	return NewReasons(data.Reasons...).Validate()
}
