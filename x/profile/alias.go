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
	EventTypeProfileSaved    = types.EventTypeProfileSaved
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileMoniker  = types.AttributeProfileMoniker
	AttributeProfileCreator  = types.AttributeProfileCreator
	DefaultParamspace        = models.DefaultParamspace
	ModuleName               = models.ModuleName
	RouterKey                = models.RouterKey
	StoreKey                 = models.StoreKey
	MinNameSurnameLength     = models.MinNameSurnameLength
	MaxNameSurnameLength     = models.MaxNameSurnameLength
	MinMonikerLength         = models.MinMonikerLength
	MaxMonikerLength         = models.MaxMonikerLength
	MaxBioLength             = models.MaxBioLength
	ActionSaveProfile        = models.ActionSaveProfile
	ActionDeleteProfile      = models.ActionDeleteProfile
	QuerierRoute             = models.QuerierRoute
	QueryProfile             = models.QueryProfile
	QueryProfiles            = models.QueryProfiles
	QueryParams              = models.QueryParams
	OpWeightMsgSaveProfile   = simulation.OpWeightMsgSaveProfile
	OpWeightMsgDeleteProfile = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue          = simulation.DefaultGasValue
)

var (
	// functions aliases
	NewHandler                  = keeper.NewHandler
	NewKeeper                   = keeper.NewKeeper
	NewQuerier                  = keeper.NewQuerier
	RegisterInvariants          = keeper.RegisterInvariants
	AllInvariants               = keeper.AllInvariants
	ValidProfileInvariant       = keeper.ValidProfileInvariant
	RandomProfileData           = simulation.RandomProfileData
	RandomProfile               = simulation.RandomProfile
	RandomMoniker               = simulation.RandomMoniker
	RandomName                  = simulation.RandomName
	RandomSurname               = simulation.RandomSurname
	RandomBio                   = simulation.RandomBio
	RandomProfilePic            = simulation.RandomProfilePic
	RandomProfileCover          = simulation.RandomProfileCover
	GetSimAccount               = simulation.GetSimAccount
	RandomProfileParams         = simulation.RandomProfileParams
	WeightedOperations          = simulation.WeightedOperations
	RandomizedGenState          = simulation.RandomizedGenState
	DecodeStore                 = simulation.DecodeStore
	SimulateMsgSaveProfile      = simulation.SimulateMsgSaveProfile
	SimulateMsgDeleteProfile    = simulation.SimulateMsgDeleteProfile
	RegisterCodec               = types.RegisterCodec
	NewGenesisState             = types.NewGenesisState
	DefaultGenesisState         = types.DefaultGenesisState
	ValidateGenesis             = types.ValidateGenesis
	ParamKeyTable               = models.ParamKeyTable
	NewNameSurnameLenParams     = models.NewNameSurnameLenParams
	DefaultNameSurnameLenParams = models.DefaultNameSurnameLenParams
	NewMonikerLenParams         = models.NewMonikerLenParams
	DefaultMonikerLenParams     = models.DefaultMonikerLenParams
	NewBioLenParams             = models.NewBioLenParams
	DefaultBioLenParams         = models.DefaultBioLenParams
	NewParamsQueryResponse      = models.NewParamsQueryResponse
	ProfileStoreKey             = models.ProfileStoreKey
	MonikerStoreKey             = models.MonikerStoreKey
	NewProfile                  = models.NewProfile
	NewPictures                 = models.NewPictures
	RegisterModelsCodec         = models.RegisterModelsCodec
	NewMsgSaveProfile           = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile         = msgs.NewMsgDeleteProfile
	RegisterMessagesCodec       = msgs.RegisterMessagesCodec

	// variable aliases
	MsgsCodec                   = msgs.MsgsCodec
	ModuleCdc                   = types.ModuleCdc
	DefaultMinNameSurnameLength = models.DefaultMinNameSurnameLength
	DefaultMaxNameSurnameLength = models.DefaultMaxNameSurnameLength
	DefaultMinMonikerLength     = models.DefaultMinMonikerLength
	DefaultMaxMonikerLength     = models.DefaultMaxMonikerLength
	DefaultMaxBioLength         = models.DefaultMaxBioLength
	ParamStoreKeyNameSurnameLen = models.ParamStoreKeyNameSurnameLen
	ParamStoreKeyMonikerLen     = models.ParamStoreKeyMonikerLen
	ParamStoreKeyMaxBioLen      = models.ParamStoreKeyMaxBioLen
	TxHashRegEx                 = models.TxHashRegEx
	URIRegEx                    = models.URIRegEx
	ProfileStorePrefix          = models.ProfileStorePrefix
	MonikerStorePrefix          = models.MonikerStorePrefix
	ModelsCdc                   = models.ModelsCdc
)

type (
	Keeper               = keeper.Keeper
	ProfileData          = simulation.ProfileData
	ProfileParams        = simulation.ProfileParams
	GenesisState         = types.GenesisState
	NameSurnameLenParams = models.NameSurnameLenParams
	MonikerLenParams     = models.MonikerLenParams
	BioLenParams         = models.BioLenParams
	ParamsQueryResponse  = models.ParamsQueryResponse
	Profile              = models.Profile
	Profiles             = models.Profiles
	Pictures             = models.Pictures
	MsgSaveProfile       = msgs.MsgSaveProfile
	MsgDeleteProfile     = msgs.MsgDeleteProfile
)
