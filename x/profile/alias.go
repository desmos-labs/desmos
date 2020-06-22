package profile

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/desmos-labs/desmos/x/profile/internal/types/msgs"
)

const (
	OpWeightMsgSaveProfile                      = simulation.OpWeightMsgSaveProfile
	OpWeightMsgDeleteProfile                    = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue                             = simulation.DefaultGasValue
	ParamsKey                                   = simulation.ParamsKey
	OpWeightSubmitNameSurnameParamsEditProposal = simulation.OpWeightSubmitNameSurnameParamsEditProposal
	OpWeightSubmitMonikerParamsEditProposal     = simulation.OpWeightSubmitMonikerParamsEditProposal
	OpWeightSubmitBiographyParamsEditProposal   = simulation.OpWeightSubmitBiographyParamsEditProposal
	EventTypeProfileSaved                       = types.EventTypeProfileSaved
	EventTypeProfileDeleted                     = types.EventTypeProfileDeleted
	AttributeProfileMoniker                     = types.AttributeProfileMoniker
	AttributeProfileCreator                     = types.AttributeProfileCreator
	ModuleName                                  = models.ModuleName
	RouterKey                                   = models.RouterKey
	StoreKey                                    = models.StoreKey
	MinNameSurnameLength                        = models.MinNameSurnameLength
	MaxNameSurnameLength                        = models.MaxNameSurnameLength
	MinMonikerLength                            = models.MinMonikerLength
	MaxMonikerLength                            = models.MaxMonikerLength
	MaxBioLength                                = models.MaxBioLength
	ActionSaveProfile                           = models.ActionSaveProfile
	ActionDeleteProfile                         = models.ActionDeleteProfile
	QuerierRoute                                = models.QuerierRoute
	QueryProfile                                = models.QueryProfile
	QueryProfiles                               = models.QueryProfiles
	QueryParams                                 = models.QueryParams
	DefaultParamspace                           = models.DefaultParamspace
	ProposalTypeNameSurnameParamsEdit           = models.ProposalTypeNameSurnameParamsEdit
	ProposalTypeMonikerParamsEdit               = models.ProposalTypeMonikerParamsEdit
	ProposalTypeBioParamsEdit                   = models.ProposalTypeBioParamsEdit
)

var (
	// functions aliases
	NewEditParamsProposalHandler          = keeper.NewEditParamsProposalHandler
	NewKeeper                             = keeper.NewKeeper
	NewQuerier                            = keeper.NewQuerier
	RegisterInvariants                    = keeper.RegisterInvariants
	AllInvariants                         = keeper.AllInvariants
	ValidProfileInvariant                 = keeper.ValidProfileInvariant
	NewHandler                            = keeper.NewHandler
	ValidateProfile                       = keeper.ValidateProfile
	SimulateMsgSaveProfile                = simulation.SimulateMsgSaveProfile
	SimulateMsgDeleteProfile              = simulation.SimulateMsgDeleteProfile
	RandomProfileData                     = simulation.RandomProfileData
	RandomProfile                         = simulation.RandomProfile
	RandomMoniker                         = simulation.RandomMoniker
	RandomName                            = simulation.RandomName
	RandomSurname                         = simulation.RandomSurname
	RandomBio                             = simulation.RandomBio
	RandomProfilePic                      = simulation.RandomProfilePic
	RandomProfileCover                    = simulation.RandomProfileCover
	GetSimAccount                         = simulation.GetSimAccount
	RandomNameSurnameParams               = simulation.RandomNameSurnameParams
	RandomMonikerParams                   = simulation.RandomMonikerParams
	RandomBioParams                       = simulation.RandomBioParams
	WeightedOperations                    = simulation.WeightedOperations
	RandomizedGenState                    = simulation.RandomizedGenState
	ProposalContents                      = simulation.ProposalContents
	SimulateNameSurnameEditParamsProposal = simulation.SimulateNameSurnameEditParamsProposal
	SimulateMonikerEditParamsProposal     = simulation.SimulateMonikerEditParamsProposal
	SimulateBiographyEditParamsProposal   = simulation.SimulateBiographyEditParamsProposal
	ParamChanges                          = simulation.ParamChanges
	DecodeStore                           = simulation.DecodeStore
	RegisterCodec                         = types.RegisterCodec
	NewGenesisState                       = types.NewGenesisState
	DefaultGenesisState                   = types.DefaultGenesisState
	ValidateGenesis                       = types.ValidateGenesis
	ProfileStoreKey                       = models.ProfileStoreKey
	MonikerStoreKey                       = models.MonikerStoreKey
	NewProfile                            = models.NewProfile
	NewPictures                           = models.NewPictures
	RegisterModelsCodec                   = models.RegisterModelsCodec
	ParamKeyTable                         = models.ParamKeyTable
	NewParams                             = models.NewParams
	DefaultParams                         = models.DefaultParams
	NewNameSurnameLenParams               = models.NewNameSurnameLenParams
	DefaultNameSurnameLenParams           = models.DefaultNameSurnameLenParams
	ValidateNameSurnameLenParams          = models.ValidateNameSurnameLenParams
	NewMonikerLenParams                   = models.NewMonikerLenParams
	DefaultMonikerLenParams               = models.DefaultMonikerLenParams
	ValidateMonikerLenParams              = models.ValidateMonikerLenParams
	NewBioLenParams                       = models.NewBioLenParams
	DefaultBioLenParams                   = models.DefaultBioLenParams
	ValidateBioLenParams                  = models.ValidateBioLenParams
	NewNameSurnameParamsEditProposal      = models.NewNameSurnameParamsEditProposal
	NewMonikerParamsEditProposal          = models.NewMonikerParamsEditProposal
	NewBioParamsEditProposal              = models.NewBioParamsEditProposal
	NewMsgSaveProfile                     = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile                   = msgs.NewMsgDeleteProfile
	RegisterMessagesCodec                 = msgs.RegisterMessagesCodec

	// variable aliases
	MsgsCodec                   = msgs.MsgsCodec
	ModuleCdc                   = types.ModuleCdc
	TxHashRegEx                 = models.TxHashRegEx
	URIRegEx                    = models.URIRegEx
	ProfileStorePrefix          = models.ProfileStorePrefix
	MonikerStorePrefix          = models.MonikerStorePrefix
	ModelsCdc                   = models.ModelsCdc
	DefaultMinNameSurnameLength = models.DefaultMinNameSurnameLength
	DefaultMaxNameSurnameLength = models.DefaultMaxNameSurnameLength
	DefaultMinMonikerLength     = models.DefaultMinMonikerLength
	DefaultMaxMonikerLength     = models.DefaultMaxMonikerLength
	DefaultMaxBioLength         = models.DefaultMaxBioLength
	NameSurnameLenParamsKey     = models.NameSurnameLenParamsKey
	MonikerLenParamsKey         = models.MonikerLenParamsKey
	BioLenParamsKey             = models.BioLenParamsKey
)

type (
	Keeper                        = keeper.Keeper
	ProfileData                   = simulation.ProfileData
	ProfileParams                 = simulation.ProfileParams
	GenesisState                  = types.GenesisState
	Profile                       = models.Profile
	Profiles                      = models.Profiles
	Pictures                      = models.Pictures
	Params                        = models.Params
	NameSurnameLengths            = models.NameSurnameLengths
	MonikerLengths                = models.MonikerLengths
	BiographyLengths              = models.BiographyLengths
	EditNameSurnameParamsProposal = models.EditNameSurnameParamsProposal
	EditMonikerParamsProposal     = models.EditMonikerParamsProposal
	EditBioParamsProposal         = models.EditBioParamsProposal
	MsgSaveProfile                = msgs.MsgSaveProfile
	MsgDeleteProfile              = msgs.MsgDeleteProfile
)
