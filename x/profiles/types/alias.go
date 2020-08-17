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
	ActionDeleteRelationship                = models.ActionDeleteRelationship
	QuerierRoute                            = models.QuerierRoute
	QueryProfile                            = models.QueryProfile
	QueryProfiles                           = models.QueryProfiles
	QueryRelationships                      = models.QueryRelationships
	QueryParams                             = models.QueryParams
)

var (
	// functions aliases
	NewProfile                              = models.NewProfile
	NewProfiles                             = models.NewProfiles
	NewPictures                             = models.NewPictures
	RegisterModelsCodec                     = models.RegisterModelsCodec
	NewMonodirectionalRelationship          = models.NewMonodirectionalRelationship
	NewBiDirectionalRelationship            = models.NewBiDirectionalRelationship
	ProfileStoreKey                         = models.ProfileStoreKey
	DtagStoreKey                            = models.DtagStoreKey
	RelationshipsStoreKey                   = models.RelationshipsStoreKey
	UserRelationshipsStoreKey               = models.UserRelationshipsStoreKey
	NewMsgCreateMonoDirectionalRelationship = msgs.NewMsgCreateMonoDirectionalRelationship
	NewMsgRequestBidirectionalRelationship  = msgs.NewMsgRequestBidirectionalRelationship
	NewMsgAcceptBidirectionalRelationship   = msgs.NewMsgAcceptBidirectionalRelationship
	NewMsgDenyBidirectionalRelationship     = msgs.NewMsgDenyBidirectionalRelationship
	NewMsgDeleteRelationship                = msgs.NewMsgDeleteRelationship
	NewMsgSaveProfile                       = msgs.NewMsgSaveProfile
	NewMsgDeleteProfile                     = msgs.NewMsgDeleteProfile
	RegisterMessagesCodec                   = msgs.RegisterMessagesCodec

	// variable aliases
	ModelsCdc               = models.ModelsCdc
	ProfileStorePrefix      = models.ProfileStorePrefix
	DtagStorePrefix         = models.DtagStorePrefix
	RelationshipsPrefix     = models.RelationshipsPrefix
	UserRelationshipsPrefix = models.UserRelationshipsPrefix
	MsgsCodec               = msgs.MsgsCodec
)

type (
	Profile                              = models.Profile
	Profiles                             = models.Profiles
	Pictures                             = models.Pictures
	RelationshipID                       = models.RelationshipID
	Relationship                         = models.Relationship
	Relationships                        = models.Relationships
	MonodirectionalRelationship          = models.MonodirectionalRelationship
	MonoDirectionalRelationships         = models.MonoDirectionalRelationships
	BidirectionalRelationship            = models.BidirectionalRelationship
	RelationshipStatus                   = models.RelationshipStatus
	BidirectionalRelationships           = models.BidirectionalRelationships
	MsgCreateMonoDirectionalRelationship = msgs.MsgCreateMonoDirectionalRelationship
	MsgRequestBidirectionalRelationship  = msgs.MsgRequestBidirectionalRelationship
	MsgAcceptBidirectionalRelationship   = msgs.MsgAcceptBidirectionalRelationship
	MsgDenyBidirectionalRelationship     = msgs.MsgDenyBidirectionalRelationship
	MsgDeleteRelationship                = msgs.MsgDeleteRelationship
	MsgSaveProfile                       = msgs.MsgSaveProfile
	MsgDeleteProfile                     = msgs.MsgDeleteProfile
)
