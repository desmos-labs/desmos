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
	ReportsStorePrefix = []byte("reports")

	// Report's type
	// TODO should this lists be saved on chain-state?
	ReportsTypes = []string{
		"nudity",
		"violence",
		"intimidation",
		"suicide or self-harm",
		"fake news",
		"spam",
		"unauthorized sale",
		"hatred incitement",
		"promotion of drug use",
		"non-consensual intimate images",
		"pornography",
		"children abuse",
		"animals abuse",
		"bullying",
		"scam",
	}
)
