package reports

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/reports/client/cli"
	"github.com/desmos-labs/desmos/x/reports/client/rest"
	"github.com/desmos-labs/desmos/x/reports/keeper"
	"github.com/desmos-labs/desmos/x/reports/simulation"
	"github.com/desmos-labs/desmos/x/reports/types"
	"github.com/desmos-labs/desmos/x/reports/types/models"
	"github.com/desmos-labs/desmos/x/reports/types/models/common"
	"github.com/desmos-labs/desmos/x/reports/types/msgs"
)

const (
	OpWeightMsgReportPost   = simulation.OpWeightMsgReportPost
	DefaultGasValue         = simulation.DefaultGasValue
	EventTypePostReported   = types.EventTypePostReported
	AttributeKeyPostID      = types.AttributeKeyPostID
	AttributeKeyReportOwner = types.AttributeKeyReportOwner
	ModuleName              = common.ModuleName
	RouterKey               = common.RouterKey
	StoreKey                = common.StoreKey
	ActionReportPost        = common.ActionReportPost
	QuerierRoute            = common.QuerierRoute
	QueryReports            = common.QueryReports
)

var (
	// functions aliases
	RegisterRoutes         = rest.RegisterRoutes
	NewHandler             = keeper.NewHandler
	NewKeeper              = keeper.NewKeeper
	NewQuerier             = keeper.NewQuerier
	RegisterInvariants     = keeper.RegisterInvariants
	AllInvariants          = keeper.AllInvariants
	ValidReportsIDs        = keeper.ValidReportsIDs
	WeightedOperations     = simulation.WeightedOperations
	RandomizedGenState     = simulation.RandomizedGenState
	DecodeStore            = simulation.DecodeStore
	SimulateMsgReportPost  = simulation.SimulateMsgReportPost
	RandomReportsData      = simulation.RandomReportsData
	RandomPostID           = simulation.RandomPostID
	RandomReportMessage    = simulation.RandomReportMessage
	RandomReportTypes      = simulation.RandomReportTypes
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	ValidateGenesis        = types.ValidateGenesis
	RegisterCodec          = types.RegisterCodec
	ReportStoreKey         = models.ReportStoreKey
	NewReportResponse      = models.NewReportResponse
	NewReport              = models.NewReport
	RegisterModelsCodec    = models.RegisterModelsCodec
	NewMsgReportPost       = msgs.NewMsgReportPost
	RegisterMessagesCodec  = msgs.RegisterMessagesCodec
	GetQueryCmd            = cli.GetQueryCmd
	GetCmdQueryPostReports = cli.GetCmdQueryPostReports
	GetTxCmd               = cli.GetTxCmd
	GetCmdReportPost       = cli.GetCmdReportPost

	// variable aliases
	ModuleCdc              = types.ModuleCdc
	ModelsCdc              = models.ModelsCdc
	ReportsStorePrefix     = common.ReportsStorePrefix
	ReportsTypeStorePrefix = common.ReportsTypeStorePrefix
	MsgsCodec              = msgs.MsgsCodec
)

type (
	GenesisState         = types.GenesisState
	ReportsQueryResponse = models.ReportsQueryResponse
	Report               = models.Report
	Reports              = models.Reports
	MsgReportPost        = msgs.MsgReportPost
	ReportPostReq        = rest.ReportPostReq
	Keeper               = keeper.Keeper
	ReportsData          = simulation.ReportsData
)
