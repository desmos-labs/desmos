package reports

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/reports/client/cli"
	"github.com/desmos-labs/desmos/x/reports/client/rest"
	"github.com/desmos-labs/desmos/x/reports/internal/keeper"
	"github.com/desmos-labs/desmos/x/reports/internal/simulation"
	"github.com/desmos-labs/desmos/x/reports/internal/types"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models"
	"github.com/desmos-labs/desmos/x/reports/internal/types/models/common"
	"github.com/desmos-labs/desmos/x/reports/internal/types/msgs"
)

const (
	ModuleName              = common.ModuleName
	RouterKey               = common.RouterKey
	StoreKey                = common.StoreKey
	ActionReportPost        = common.ActionReportPost
	QuerierRoute            = common.QuerierRoute
	QueryReports            = common.QueryReports
	OpWeightMsgReportPost   = simulation.OpWeightMsgReportPost
	DefaultGasValue         = simulation.DefaultGasValue
	EventTypePostReported   = types.EventTypePostReported
	AttributeKeyPostID      = types.AttributeKeyPostID
	AttributeKeyReportOwner = types.AttributeKeyReportOwner
)

var (
	// functions aliases
	RegisterRoutes         = rest.RegisterRoutes
	NewQuerier             = keeper.NewQuerier
	RegisterInvariants     = keeper.RegisterInvariants
	AllInvariants          = keeper.AllInvariants
	ValidReportsIDs        = keeper.ValidReportsIDs
	NewHandler             = keeper.NewHandler
	NewKeeper              = keeper.NewKeeper
	SimulateMsgReportPost  = simulation.SimulateMsgReportPost
	RandomReportsData      = simulation.RandomReportsData
	RandomPostID           = simulation.RandomPostID
	RandomReportMessage    = simulation.RandomReportMessage
	RandomReportTypes      = simulation.RandomReportTypes
	WeightedOperations     = simulation.WeightedOperations
	RandomizedGenState     = simulation.RandomizedGenState
	DecodeStore            = simulation.DecodeStore
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	ValidateGenesis        = types.ValidateGenesis
	RegisterCodec          = types.RegisterCodec
	NewReport              = models.NewReport
	RegisterModelsCodec    = models.RegisterModelsCodec
	ReportStoreKey         = models.ReportStoreKey
	NewReportResponse      = models.NewReportResponse
	NewMsgReportPost       = msgs.NewMsgReportPost
	RegisterMessagesCodec  = msgs.RegisterMessagesCodec
	GetQueryCmd            = cli.GetQueryCmd
	GetCmdQueryPostReports = cli.GetCmdQueryPostReports
	GetTxCmd               = cli.GetTxCmd
	GetCmdReportPost       = cli.GetCmdReportPost

	// variable aliases
	ReportsStorePrefix     = common.ReportsStorePrefix
	ReportsTypeStorePrefix = common.ReportsTypeStorePrefix
	MsgsCodec              = msgs.MsgsCodec
	ModuleCdc              = types.ModuleCdc
	ModelsCdc              = models.ModelsCdc
)

type (
	Report               = models.Report
	Reports              = models.Reports
	ReportsQueryResponse = models.ReportsQueryResponse
	MsgReportPost        = msgs.MsgReportPost
	ReportPostReq        = rest.ReportPostReq
	Keeper               = keeper.Keeper
	ReportsData          = simulation.ReportsData
	GenesisState         = types.GenesisState
)
