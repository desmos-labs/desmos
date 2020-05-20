package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/report/internal/types/models"
	"github.com/desmos-labs/desmos/x/report/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/report/internal/types/msgs"
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
	ModelsCdc          = models.ModelsCdc
	ReportsStorePrefix = common.ReportsStorePrefix
	MsgsCodec          = msgs.MsgsCodec
)

type (
	ReportsQueryResponse = models.ReportsQueryResponse
	Report               = models.Report
	Reports              = models.Reports
	MsgReportPost        = msgs.MsgReportPost
)
