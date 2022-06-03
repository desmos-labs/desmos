package types

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
)

// NewGenesisState returns a new genesis state instance
func NewGenesisState(subspaces []SubspaceDataEntry, reasons []Reason, reports []Report, params Params) *GenesisState {
	return &GenesisState{
		SubspacesData: subspaces,
		Reasons:       reasons,
		Reports:       reports,
		Params:        params,
	}
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
func (data GenesisState) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, report := range data.Reports {
		err := report.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

// DefaultGenesisState returns a DefaultGenesisState
func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, nil, nil, DefaultParams())
}

// ValidateGenesis validates the given genesis state and returns an error if something is invalid
func ValidateGenesis(data *GenesisState) error {
	for _, subspaceData := range data.SubspacesData {
		if containsDuplicatedSubspacesData(data.SubspacesData, subspaceData) {
			return fmt.Errorf("duplicated subspace entry for subspace id %d", subspaceData.SubspaceID)
		}

		err := subspaceData.Validate()
		if err != nil {
			return fmt.Errorf("invalid subspace data: %s", err)
		}
	}

	for _, reason := range data.Reasons {
		if containsDuplicatedReason(data.Reasons, reason) {
			return fmt.Errorf("duplicate reason: subspace id %d, reason id %d", reason.SubspaceID, reason.ID)
		}

		err := reason.Validate()
		if err != nil {
			return fmt.Errorf("invalid reason: %s", err)
		}
	}

	for _, report := range data.Reports {
		if containsDuplicatedReport(data.Reports, report) {
			return fmt.Errorf("duplicated report: subspace id %d, report id %d", report.SubspaceID, report.ID)
		}

		err := report.Validate()
		if err != nil {
			return fmt.Errorf("invalid report: %s", err)
		}
	}

	return data.Params.Validate()
}

// containsDuplicatedSubspacesData tells whether the given subspaces data slice
// contains a duplicated entry for the same subspace id as the one given
func containsDuplicatedSubspacesData(subspaces []SubspaceDataEntry, data SubspaceDataEntry) bool {
	var count = 0
	for _, r := range subspaces {
		if r.SubspaceID == data.SubspaceID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedReason tells whether the given reasons slice contains
// a duplicated reason based on the same subspace id and reason id of the given one
func containsDuplicatedReason(reasons []Reason, reason Reason) bool {
	var count = 0
	for _, r := range reasons {
		if r.SubspaceID == reason.SubspaceID && r.ID == reason.ID {
			count++
		}
	}
	return count > 1
}

// containsDuplicatedReport tells whether the given reports slice contains
// a duplicated report based on the same subspace id and report id of the given one
func containsDuplicatedReport(reports []Report, report Report) bool {
	var count = 0
	for _, r := range reports {
		if r.SubspaceID == report.SubspaceID && r.ID == report.ID {
			count++
		}
	}
	return count > 1
}

// --------------------------------------------------------------------------------------------------------------------

// NewSubspacesDataEntry returns a new SubspaceDataEntry instance
func NewSubspacesDataEntry(subspaceID uint64, reasonID uint32, reportID uint64) SubspaceDataEntry {
	return SubspaceDataEntry{
		SubspaceID: subspaceID,
		ReportID:   reportID,
		ReasonID:   reasonID,
	}
}

// Validate implements fmt.Validator
func (data SubspaceDataEntry) Validate() error {
	if data.SubspaceID == 0 {
		return fmt.Errorf("invalid subspace id: %d", data.SubspaceID)
	}

	if data.ReasonID == 0 {
		return fmt.Errorf("invalid reason id: %d", data.ReasonID)
	}

	if data.ReportID == 0 {
		return fmt.Errorf("invalid report id: %d", data.ReportID)
	}

	return nil
}
