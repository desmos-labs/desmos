package types

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

// ReportsStoreKey turn an id to a key used to store a reports inside the reports store
func ReportStoreKey(id string) []byte {
	return append(ReportsStorePrefix, []byte(id)...)
}
