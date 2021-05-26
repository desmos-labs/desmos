package wasm

import reportsTypes "github.com/desmos-labs/desmos/x/staging/reports/types"

type ReportsModuleQuery struct {
	Reports *ReportsQuery `json:"reports"`
}

type ReportsQuery struct {
	PostID string `json:"post_id"`
}

type Report struct {
	PostID  string `json:"post_id"`
	Kind    string `json:"kind"`
	Message string `json:"message"`
	User    string `json:"user"`
}

func convertReports(reports []reportsTypes.Report) []Report {
	convertedReports := make([]Report, len(reports))
	for index, report := range reports {
		convertedReports[index] = Report{
			PostID:  report.PostID,
			Kind:    report.Type,
			Message: report.Message,
			User:    report.User,
		}
	}
	return convertedReports
}

type ReportsResponse struct {
	Reports []Report `json:"reports"`
}
