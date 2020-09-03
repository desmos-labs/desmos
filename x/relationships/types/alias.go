package types

// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/desmos-labs/desmos/x/relationships/types/models"
	"github.com/desmos-labs/desmos/x/relationships/types/msgs"
)

const (
	ModuleName               = models.ModuleName
	RouterKey                = models.RouterKey
	StoreKey                 = models.StoreKey
	ActionCreateRelationship = models.ActionCreateRelationship
	ActionDeleteRelationship = models.ActionDeleteRelationship
	QuerierRoute             = models.QuerierRoute
	QueryUserRelationships   = models.QueryUserRelationships
	QueryRelationships       = models.QueryRelationships
)

var (
	// functions aliases
	NewRelationship          = models.NewRelationship
	RelationshipsStoreKey    = models.RelationshipsStoreKey
	RegisterModelsCodec      = models.RegisterModelsCodec
	NewMsgCreateRelationship = msgs.NewMsgCreateRelationship
	NewMsgDeleteRelationship = msgs.NewMsgDeleteRelationship
	RegisterMessagesCodec    = msgs.RegisterMessagesCodec

	// variable aliases
	RelationshipsStorePrefix = models.RelationshipsStorePrefix
	ModelsCdc                = models.ModelsCdc
	MsgsCodec                = msgs.MsgsCodec
)

type (
	MsgCreateRelationship = msgs.MsgCreateRelationship
	MsgDeleteRelationship = msgs.MsgDeleteRelationship
	Relationship          = models.Relationship
	Relationships         = models.Relationships
)
