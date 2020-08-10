package types

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/profiles/types/models"
	"github.com/desmos-labs/desmos/x/profiles/types/msgs"
)

const (
	Sent                                    = models.Sent
	Accepted                                = models.Accepted
	Denied                                  = models.Denied
	ModuleName                              = models.ModuleName
	RouterKey                               = models.RouterKey
	StoreKey                                = models.StoreKey
	ActionSaveProfile                       = models.ActionSaveProfile
	ActionDeleteProfile                     = models.ActionDeleteProfile
	ActionCreateMonoDirectionalRelationship = models.ActionCreateMonoDirectionalRelationship
	ActionRequestBiDirectionalRelationship  = models.ActionRequestBiDirectionalRelationship
	ActionAcceptBiDirectionalRelationship   = models.ActionAcceptBiDirectionalRelationship
	ActionDenyBiDirectionalRelationship     = models.ActionDenyBiDirectionalRelationship
	ActionDeleteRelationships               = models.ActionDeleteRelationships
	QuerierRoute                            = models.QuerierRoute
	QueryProfile                            = models.QueryProfile
	QueryProfiles                           = models.QueryProfiles
	QueryParams                             = models.QueryParams
	QueryMonoDirectionalRelationships       = models.QueryMonoDirectionalRelationships
	QueryBiDirectionalRelationships         = models.QueryBiDirectionalRelationships
)

var (
	// functions aliases
	NewMonodirectionalRelationship          = models.NewMonodirectionalRelationship
	NewBiDirectionalRelationship            = models.NewBiDirectionalRelationship
	ProfileStoreKey                         = models.ProfileStoreKey
	DtagStoreKey                            = models.DtagStoreKey
	RelationshipsStoreKey                   = models.RelationshipsStoreKey
	NewProfile                              = models.NewProfile
	NewProfiles                             = models.NewProfiles
	NewPictures                             = models.NewPictures
	RegisterModelsCodec                     = models.RegisterModelsCodec
	RegisterMessagesCodec                   = msgs.RegisterMessagesCodec
	NewMsgCreateMonoDirectionalRelationship = msgs.NewMsgCreateMonoDirectionalRelationship
	NewMsgRequestBidirectionalRelationship  = msgs.NewMsgRequestBidirectionalRelationship
	NewMsgAcceptBidirectionalRelationship   = msgs.NewMsgAcceptBidirectionalRelationship
	NewMsgDenyBidirectionalRelationship     = msgs.NewMsgDenyBidirectionalRelationship
	NewMsgDeleteRelationships               = msgs.NewMsgDeleteRelationships
	NewMsgSaveProfile                       = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile                     = msgs.NewMsgDeleteProfile

	// variable aliases
	ProfileStorePrefix  = models.ProfileStorePrefix
	DtagStorePrefix     = models.DtagStorePrefix
	RelationshipsPrefix = models.RelationshipsPrefix
	ModelsCdc           = models.ModelsCdc
	MsgsCodec           = msgs.MsgsCodec
)

type (
	Relationship                         = models.Relationship
	Relationships                        = models.Relationships
	MonodirectionalRelationship          = models.MonodirectionalRelationship
	MonoDirectionalRelationships         = models.MonoDirectionalRelationships
	BidirectionalRelationship            = models.BidirectionalRelationship
	RelationshipStatus                   = models.RelationshipStatus
	BidirectionalRelationships           = models.BidirectionalRelationships
	Profile                              = models.Profile
	Profiles                             = models.Profiles
	Pictures                             = models.Pictures
	MsgCreateMonoDirectionalRelationship = msgs.MsgCreateMonoDirectionalRelationship
	MsgRequestBidirectionalRelationship  = msgs.MsgRequestBidirectionalRelationship
	MsgAcceptBidirectionalRelationship   = msgs.MsgAcceptBidirectionalRelationship
	MsgDenyBidirectionalRelationship     = msgs.MsgDenyBidirectionalRelationship
	MsgDeleteRelationships               = msgs.MsgDeleteRelationships
	MsgSaveProfile                       = msgs.MsgSaveProfile
	MsgDeleteProfile                     = msgs.MsgDeleteProfile
)
