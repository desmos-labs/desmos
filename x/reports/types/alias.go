package types

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/reports/types/models"
	"github.com/desmos-labs/desmos/x/reports/types/models/common"
	"github.com/desmos-labs/desmos/x/reports/types/msgs"
)

const (
	ModuleName       = common.ModuleName
	RouterKey        = common.RouterKey
	StoreKey         = common.StoreKey
	ActionReportPost = common.ActionReportPost
	QuerierRoute     = common.QuerierRoute
	QueryReports     = common.QueryReports
)

var (
	// functions aliases
	ReportStoreKey        = models.ReportStoreKey
	NewReportResponse     = models.NewReportResponse
	NewReport             = models.NewReport
	RegisterModelsCodec   = models.RegisterModelsCodec
	NewMsgReportPost      = msgs.NewMsgReportPost
	RegisterMessagesCodec = msgs.RegisterMessagesCodec

	// variable aliases
	ModelsCdc              = models.ModelsCdc
	ReportsStorePrefix     = common.ReportsStorePrefix
	ReportsTypeStorePrefix = common.ReportsTypeStorePrefix
	MsgsCodec              = msgs.MsgsCodec
)

type (
	MsgReportPost        = msgs.MsgReportPost
	ReportsQueryResponse = models.ReportsQueryResponse
	Report               = models.Report
	Reports              = models.Reports
)
