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
	OpWeightMsgCreateProfile = simulation.OpWeightMsgCreateProfile
	OpWeightMsgEditProfile   = simulation.OpWeightMsgEditProfile
	OpWeightMsgDeleteProfile = simulation.OpWeightMsgDeleteProfile
	DefaultGasValue          = simulation.DefaultGasValue
	ModuleName               = types.ModuleName
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	MinNameSurnameLength     = types.MinNameSurnameLength
	MaxNameSurnameLength     = types.MaxNameSurnameLength
	MaxMonikerLength         = types.MaxMonikerLength
	MaxBioLength             = types.MaxBioLength
	ActionCreateProfile      = types.ActionCreateProfile
	ActionEditProfile        = types.ActionEditProfile
	ActionDeleteProfile      = types.ActionDeleteProfile
	QuerierRoute             = types.QuerierRoute
	QueryProfile             = types.QueryProfile
	QueryProfiles            = types.QueryProfiles
	EventTypeProfileCreated  = types.EventTypeProfileCreated
	EventTypeProfileEdited   = types.EventTypeProfileEdited
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileMoniker  = types.AttributeProfileMoniker
	AttributeProfileCreator  = types.AttributeProfileCreator
)

var (
	// functions aliases
	NewHandler               = keeper.NewHandler
	GetEditedProfile         = keeper.GetEditedProfile
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	RegisterInvariants       = keeper.RegisterInvariants
	AllInvariants            = keeper.AllInvariants
	ValidProfileInvariant    = keeper.ValidProfileInvariant
	DecodeStore              = simulation.DecodeStore
	SimulateMsgCreateProfile = simulation.SimulateMsgCreateProfile
	SimulateMsgEditProfile   = simulation.SimulateMsgEditProfile
	SimulateMsgDeleteProfile = simulation.SimulateMsgDeleteProfile
	RandomProfileData        = simulation.RandomProfileData
	RandomProfile            = simulation.RandomProfile
	RandomMoniker            = simulation.RandomMoniker
	RandomName               = simulation.RandomName
	RandomSurname            = simulation.RandomSurname
	RandomBio                = simulation.RandomBio
	RandomProfilePic         = simulation.RandomProfilePic
	RandomProfileCover       = simulation.RandomProfileCover
	GetSimAccount            = simulation.GetSimAccount
	WeightedOperations       = simulation.WeightedOperations
	RandomizedGenState       = simulation.RandomizedGenState
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	RegisterCodec            = types.RegisterCodec
	NewMsgCreateProfile      = types.NewMsgCreateProfile
	NewMsgEditProfile        = types.NewMsgEditProfile
	NewMsgDeleteProfile      = types.NewMsgDeleteProfile
	ProfileStoreKey          = types.ProfileStoreKey
	MonikerStoreKey          = types.MonikerStoreKey
	NewProfile               = types.NewProfile
	NewPictures              = types.NewPictures
	ValidateURI              = types.ValidateURI

	// variable aliases
	ModuleCdc          = types.ModuleCdc
	TxHashRegEx        = types.TxHashRegEx
	URIRegEx           = types.URIRegEx
	ProfileStorePrefix = types.ProfileStorePrefix
	MonikerStorePrefix = types.MonikerStorePrefix
)

type (
	Keeper           = keeper.Keeper
	ProfileData      = simulation.ProfileData
	GenesisState     = types.GenesisState
	MsgCreateProfile = types.MsgCreateProfile
	MsgEditProfile   = types.MsgEditProfile
	MsgDeleteProfile = types.MsgDeleteProfile
	Profile          = types.Profile
	Profiles         = types.Profiles
	Pictures         = types.Pictures
)
