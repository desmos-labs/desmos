package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profile/internal/types/models"
	"github.com/desmos-labs/desmos/x/profile/internal/types/msgs"
)

const (
	DefaultParamspace                 = models.DefaultParamspace
	ProposalTypeNameSurnameParamsEdit = models.ProposalTypeNameSurnameParamsEdit
	ProposalTypeMonikerParamsEdit     = models.ProposalTypeMonikerParamsEdit
	ProposalTypeBioParamsEdit         = models.ProposalTypeBioParamsEdit
	ModuleName                        = models.ModuleName
	RouterKey                         = models.RouterKey
	StoreKey                          = models.StoreKey
	MinNameSurnameLength              = models.MinNameSurnameLength
	MaxNameSurnameLength              = models.MaxNameSurnameLength
	MinMonikerLength                  = models.MinMonikerLength
	MaxMonikerLength                  = models.MaxMonikerLength
	MaxBioLength                      = models.MaxBioLength
	ActionSaveProfile                 = models.ActionSaveProfile
	ActionDeleteProfile               = models.ActionDeleteProfile
	QuerierRoute                      = models.QuerierRoute
	QueryProfile                      = models.QueryProfile
	QueryProfiles                     = models.QueryProfiles
	QueryParams                       = models.QueryParams
)

var (
	// functions aliases
	ParamKeyTable                    = models.ParamKeyTable
	NewNameSurnameLenParams          = models.NewNameSurnameLenParams
	DefaultNameSurnameLenParams      = models.DefaultNameSurnameLenParams
	ValidateNameSurnameLenParams     = models.ValidateNameSurnameLenParams
	NewMonikerLenParams              = models.NewMonikerLenParams
	DefaultMonikerLenParams          = models.DefaultMonikerLenParams
	ValidateMonikerLenParams         = models.ValidateMonikerLenParams
	NewBioLenParams                  = models.NewBioLenParams
	DefaultBioLenParams              = models.DefaultBioLenParams
	ValidateBioLenParams             = models.ValidateBioLenParams
	NewParamsQueryResponse           = models.NewParamsQueryResponse
	NewNameSurnameParamsEditProposal = models.NewNameSurnameParamsEditProposal
	NewMonikerParamsEditProposal     = models.NewMonikerParamsEditProposal
	NewBioParamsEditProposal         = models.NewBioParamsEditProposal
	ProfileStoreKey                  = models.ProfileStoreKey
	MonikerStoreKey                  = models.MonikerStoreKey
	NewProfile                       = models.NewProfile
	NewPictures                      = models.NewPictures
	RegisterModelsCodec              = models.RegisterModelsCodec
	NewMsgSaveProfile                = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile              = msgs.NewMsgDeleteProfile
	RegisterMessagesCodec            = msgs.RegisterMessagesCodec

	// variable aliases
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
	MsgsCodec                   = msgs.MsgsCodec
)

type (
	NameSurnameLenParams          = models.NameSurnameLenParams
	MonikerLenParams              = models.MonikerLenParams
	BioLenParams                  = models.BioLenParams
	ParamsQueryResponse           = models.ParamsQueryResponse
	NameSurnameParamsEditProposal = models.NameSurnameParamsEditProposal
	MonikerParamsEditProposal     = models.MonikerParamsEditProposal
	BioParamsEditProposal         = models.BioParamsEditProposal
	Profile                       = models.Profile
	Profiles                      = models.Profiles
	Pictures                      = models.Pictures
	MsgSaveProfile                = msgs.MsgSaveProfile
	MsgDeleteProfile              = msgs.MsgDeleteProfile
)
