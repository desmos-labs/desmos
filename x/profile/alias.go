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
	EventTypeProfileCreated  = types.EventTypeProfileCreated
	EventTypeProfileEdited   = types.EventTypeProfileEdited
	EventTypeProfileDeleted  = types.EventTypeProfileDeleted
	AttributeProfileMoniker  = types.AttributeProfileMoniker
	AttributeProfileCreator  = types.AttributeProfileCreator
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
)

var (
	// functions aliases

	NewHandler               = keeper.NewHandler
	GetEditedProfile         = keeper.GetEditedProfile
	NewKeeper                = keeper.NewKeeper
	NewQuerier               = keeper.NewQuerier
	DecodeStore              = simulation.DecodeStore
	RandomizedGenState       = simulation.RandomizedGenState
	WeightedOperations       = simulation.WeightedOperations
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
	NewGenesisState          = types.NewGenesisState
	DefaultGenesisState      = types.DefaultGenesisState
	ValidateGenesis          = types.ValidateGenesis
	ProfileStoreKey          = types.ProfileStoreKey
	MonikerStoreKey          = types.MonikerStoreKey
	NewMsgCreateProfile      = types.NewMsgCreateProfile
	NewMsgEditProfile        = types.NewMsgEditProfile
	NewMsgDeleteProfile      = types.NewMsgDeleteProfile
	NewPictures              = types.NewPictures
	ValidateURI              = types.ValidateURI
	NewProfile               = types.NewProfile
	RegisterCodec            = types.RegisterCodec

	// variable aliases

	TxHashRegEx        = types.TxHashRegEx
	URIRegEx           = types.URIRegEx
	ProfileStorePrefix = types.ProfileStorePrefix
	MonikerStorePrefix = types.MonikerStorePrefix
	ModuleCdc          = types.ModuleCdc
)

type (
	GenesisState     = types.GenesisState
	MsgCreateProfile = types.MsgCreateProfile
	MsgEditProfile   = types.MsgEditProfile
	MsgDeleteProfile = types.MsgDeleteProfile
	Pictures         = types.Pictures
	Profile          = types.Profile
	Profiles         = types.Profiles
	Keeper           = keeper.Keeper
)
