package common

const (
	ModuleName = "reports"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionReportPost = "report_post"

	// Queries
	QuerierRoute = ModuleName
	QueryReports = "reports"
)

var (
	ReportsStorePrefix     = []byte("reports")
	ReportsTypeStorePrefix = []byte("report_type")
)
