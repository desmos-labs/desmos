package types

// DONTCOVER

const (
	ModuleName = "reports"
	RouterKey  = ModuleName
	StoreKey   = ModuleName

	ActionReportPost = "report_post"

	QuerierRoute = ModuleName
	QueryReports = "reports"

	// DefaultParamsSpace represents the default paramspace for the Params keeper
	DefaultParamsSpace = ModuleName
)

var (
	ReportsStorePrefix = []byte("reports")
)

// ReportStoreKey turns an id into a key used to store a report inside the reports store
func ReportStoreKey(id string) []byte {
	return append(ReportsStorePrefix, []byte(id)...)
}
