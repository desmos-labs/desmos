package v0150

import (
	v0150reports "github.com/desmos-labs/desmos/x/staging/reports/types"
)

type GenesisState = v0150reports.GenesisState

func FindReportsForPostWithID(state GenesisState, id string) []Report {
	var reports []Report
	for _, report := range state.Reports {
		if report.PostId == id {
			reports = append(reports, report)
		}
	}
	return reports
}

type Report = v0150reports.Report
