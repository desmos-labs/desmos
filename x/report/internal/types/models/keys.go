package models

const (
	ModuleName = "reports"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionReportPost = "report_post"

	// Queries
	QuerierRoute = ModuleName
	QueryReport  = "report"
	QueryReports = "reports"
)

var (
	ReportsStorePrefix = []byte("reports")
)
