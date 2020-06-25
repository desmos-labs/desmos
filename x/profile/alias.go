package profile

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profile/internal/keeper"
	"github.com/desmos-labs/desmos/x/profile/internal/simulation"
	"github.com/desmos-labs/desmos/x/profile/internal/types"
)

const (
	OpWeightMsgSaveProfile   = simulation.OpWeightMsgSaveProfile
	OpWeightMsgDeleteProfile = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue          = simulation.DefaultGasValue
	ParamsKey                = simulation.ParamsKey
	DefaultParamspace        = types.DefaultParamspace
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	MinMonikerLength         = types.MinMonikerLength
	MaxMonikerLength         = types.MaxMonikerLength
	MaxBioLength             = types.MaxBioLength
	ActionSaveProfile        = types.ActionSaveProfile
	ActionDeleteProfile      = types.ActionDeleteProfile
	QuerierRoute             = types.QuerierRoute
	QueryProfile             = types.QueryProfile
	QueryProfiles            = types.QueryProfiles
	QueryParams              = types.QueryParams
	EventTypeProfileSaved    = types.EventTypeProfileSaved
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileDtag     = types.AttributeProfileDtag
	AttributeProfileCreator  = types.AttributeProfileCreator
)

var (
	// functions aliases
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RegisterInvariants       = keeper.RegisterInvariants
	AllInvariants            = keeper.AllInvariants
	ValidProfileInvariant    = keeper.ValidProfileInvariant
	NewHandler               = keeper.NewHandler
	ValidateProfile          = keeper.ValidateProfile
	SimulateMsgSaveProfile   = simulation.SimulateMsgSaveProfile
	SimulateMsgDeleteProfile = simulation.SimulateMsgDeleteProfile
	NewRandomProfile         = simulation.NewRandomProfile
	RandomProfile            = simulation.RandomProfile
	RandomDTag               = simulation.RandomDTag
	RandomMoniker            = simulation.RandomMoniker
	RandomBio                = simulation.RandomBio
	RandomProfilePic         = simulation.RandomProfilePic
	RandomProfileCover       = simulation.RandomProfileCover
	GetSimAccount            = simulation.GetSimAccount
	RandomNameSurnameParams  = simulation.RandomNameSurnameParams
	RandomMonikerParams      = simulation.RandomMonikerParams
	RandomBioParams          = simulation.RandomBioParams
	WeightedOperations       = simulation.WeightedOperations
	RandomizedGenState       = simulation.RandomizedGenState
	ParamChanges             = simulation.ParamChanges
	DecodeStore              = simulation.DecodeStore
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	RegisterCodec            = types.RegisterCodec
	ParamKeyTable            = types.ParamKeyTable
	NewParams                = types.NewParams
	DefaultParams            = types.DefaultParams
	NewMonikerLenParams      = types.NewMonikerLenParams
	DefaultMonikerLenParams  = types.DefaultMonikerLenParams
	ValidateMonikerLenParams = types.ValidateMonikerLenParams
	NewDtagLenParams         = types.NewDtagLenParams
	DefaultDtagLenParams     = types.DefaultDtagLenParams
	ValidateDtagLenParams    = types.ValidateDtagLenParams
	ValidateBioLenParams     = types.ValidateBioLenParams
	NewMsgSaveProfile        = types.NewMsgSaveProfile
	NewMsgDeleteProfile      = types.NewMsgDeleteProfile
	ProfileStoreKey          = types.ProfileStoreKey
	DtagStoreKey             = types.DtagStoreKey
	NewProfile               = types.NewProfile
	NewProfiles              = types.NewProfiles
	NewPictures              = types.NewPictures

	// variable aliases
	ModuleCdc               = types.ModuleCdc
	DefaultMinMonikerLength = types.DefaultMinMonikerLength
	DefaultMaxMonikerLength = types.DefaultMaxMonikerLength
	DefaultMinDtagLength    = types.DefaultMinDtagLength
	DefaultMaxDtagLength    = types.DefaultMaxDtagLength
	DefaultMaxBioLength     = types.DefaultMaxBioLength
	MonikerLenParamsKey     = types.MonikerLenParamsKey
	DtagLenParamsKey        = types.DtagLenParamsKey
	MaxBioLenParamsKey      = types.MaxBioLenParamsKey
	DTagRegEx               = types.DTagRegEx
	URIRegEx                = types.URIRegEx
	ProfileStorePrefix      = types.ProfileStorePrefix
	DtagStorePrefix         = types.DtagStorePrefix
)

type (
	Keeper           = keeper.Keeper
	ProfileParams    = simulation.ProfileParams
	GenesisState     = types.GenesisState
	Params           = types.Params
	MonikerLengths   = types.MonikerLengths
	DtagLengths      = types.DtagLengths
	MsgSaveProfile   = types.MsgSaveProfile
	MsgDeleteProfile = types.MsgDeleteProfile
	Profile          = types.Profile
	Profiles         = types.Profiles
	Pictures         = types.Pictures
)
