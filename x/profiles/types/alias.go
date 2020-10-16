package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/desmos-labs/desmos/x/profiles/types/msgs"
)

const (
	ModuleName                = models.ModuleName
	RouterKey                 = models.RouterKey
	StoreKey                  = models.StoreKey
	ActionSaveProfile         = models.ActionSaveProfile
	ActionDeleteProfile       = models.ActionDeleteProfile
	ActionRequestDtag         = models.ActionRequestDtag
	ActionAcceptDtagTransfer  = models.ActionAcceptDtagTransfer
	RejectDTagTransferRequest = models.RejectDTagTransferRequest
	CancelDTagTransferRequest = models.CancelDTagTransferRequest
	QuerierRoute              = models.QuerierRoute
	QueryProfile              = models.QueryProfile
	QueryProfiles             = models.QueryProfiles
	QueryDTagRequests         = models.QueryDTagRequests
	QueryParams               = models.QueryParams
)

var (
	// functions aliases
	ProfileStoreKey             = models.ProfileStoreKey
	DtagStoreKey                = models.DtagStoreKey
	DtagTransferRequestStoreKey = models.DtagTransferRequestStoreKey
	NewProfile                  = models.NewProfile
	NewProfiles                 = models.NewProfiles
	NewPictures                 = models.NewPictures
	NewDTagTransferRequest      = models.NewDTagTransferRequest
	RegisterModelsCodec         = models.RegisterModelsCodec
	NewMsgSaveProfile           = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile         = msgs.NewMsgDeleteProfile
	NewMsgRequestDTagTransfer   = msgs.NewMsgRequestDTagTransfer
	NewMsgAcceptDTagTransfer    = msgs.NewMsgAcceptDTagTransfer
	NewMsgRefuseDTagRequest     = msgs.NewMsgRefuseDTagRequest
	NewMsgCancelDTagRequest     = msgs.NewMsgCancelDTagRequest
	RegisterMessagesCodec       = msgs.RegisterMessagesCodec

	// variable aliases
	ProfileStorePrefix         = models.ProfileStorePrefix
	DtagStorePrefix            = models.DtagStorePrefix
	DTagTransferRequestsPrefix = models.DTagTransferRequestsPrefix
	ModelsCdc                  = models.ModelsCdc
	MsgsCodec                  = msgs.MsgsCodec
)

type (
	Profile                = models.Profile
	Profiles               = models.Profiles
	Pictures               = models.Pictures
	DTagTransferRequest    = models.DTagTransferRequest
	MsgSaveProfile         = msgs.MsgSaveProfile
	MsgDeleteProfile       = msgs.MsgDeleteProfile
	MsgRequestDTagTransfer = msgs.MsgRequestDTagTransfer
	MsgAcceptDTagTransfer  = msgs.MsgAcceptDTagTransfer
	MsgRefuseDTagRequest   = msgs.MsgRefuseDTagRequest
	MsgCancelDTagRequest   = msgs.MsgCancelDTagRequest
)
