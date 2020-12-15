package v0150

import v0130reports "github.com/desmos-labs/desmos/x/reports/legacy/v0.13.0"

// Migrate accepts exported genesis state from v0.13.0 and migrates it to v0.15.0
// genesis state.
func Migrate(oldGenState v0130reports.GenesisState) GenesisState {
	return GenesisState{
		Reports: ConvertReports(oldGenState.Reports),
	}
}

func ConvertReports(oldReportsMap map[string][]v0130reports.Report) []Report {
	var reports []Report

	for postID, oldReports := range oldReportsMap {
		for _, oldReport := range oldReports {
			report := Report{
				PostID:  postID,
				Type:    oldReport.Type,
				Message: oldReport.Message,
				User:    oldReport.User.String(),
			}
			reports = append(reports, report)
		}
	}

	return reports
}
