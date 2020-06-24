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
	RegisterMessagesCodec  = msgs.RegisterMessagesCodec
	NewMsgReportPost       = msgs.NewMsgReportPost
	GetQueryCmd            = cli.GetQueryCmd
	GetCmdQueryPostReports = cli.GetCmdQueryPostReports
	GetTxCmd               = cli.GetTxCmd
	GetCmdReportPost       = cli.GetCmdReportPost
	RegisterRoutes         = rest.RegisterRoutes
	RegisterInvariants     = keeper.RegisterInvariants
	AllInvariants          = keeper.AllInvariants
	ValidReportsIDs        = keeper.ValidReportsIDs
	NewHandler             = keeper.NewHandler
	NewKeeper              = keeper.NewKeeper
	NewQuerier             = keeper.NewQuerier
	DecodeStore            = simulation.DecodeStore
	SimulateMsgReportPost  = simulation.SimulateMsgReportPost
	RandomReportsData      = simulation.RandomReportsData
	RandomPostID           = simulation.RandomPostID
	RandomReportMessage    = simulation.RandomReportMessage
	RandomReportTypes      = simulation.RandomReportTypes
	WeightedOperations     = simulation.WeightedOperations
	RandomizedGenState     = simulation.RandomizedGenState
	NewGenesisState        = types.NewGenesisState
	DefaultGenesisState    = types.DefaultGenesisState
	ValidateGenesis        = types.ValidateGenesis
	RegisterCodec          = types.RegisterCodec
	RegisterModelsCodec    = models.RegisterModelsCodec
	ReportStoreKey         = models.ReportStoreKey
	NewReportResponse      = models.NewReportResponse
	NewReport              = models.NewReport

	// variable aliases
	ReportsStorePrefix     = common.ReportsStorePrefix
	ReportsTypeStorePrefix = common.ReportsTypeStorePrefix
	MsgsCodec              = msgs.MsgsCodec
	ModuleCdc              = types.ModuleCdc
	ModelsCdc              = models.ModelsCdc
)

type (
	ReportsData          = simulation.ReportsData
	GenesisState         = types.GenesisState
	ReportsQueryResponse = models.ReportsQueryResponse
	Report               = models.Report
	Reports              = models.Reports
	MsgReportPost        = msgs.MsgReportPost
	ReportPostReq        = rest.ReportPostReq
	Keeper               = keeper.Keeper
)
