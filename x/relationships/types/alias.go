package types

// nolint
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
	ActionBlockUser          = models.ActionBlockUser
	ActionUnblockUser        = models.ActionUnblockUser
	QuerierRoute             = models.QuerierRoute
	QueryUserRelationships   = models.QueryUserRelationships
	QueryRelationships       = models.QueryRelationships
	QueryUserBlocks          = models.QueryUserBlocks
)

var (
	// functions aliases
	NewRelationshipResponse  = models.NewRelationshipResponse
	NewUserBlock             = models.NewUserBlock
	RelationshipsStoreKey    = models.RelationshipsStoreKey
	UsersBlocksStoreKey      = models.UsersBlocksStoreKey
	RegisterModelsCodec      = models.RegisterModelsCodec
	NewMsgCreateRelationship = msgs.NewMsgCreateRelationship
	NewMsgDeleteRelationship = msgs.NewMsgDeleteRelationship
	NewMsgBlockUser          = msgs.NewMsgBlockUser
	NewMsgUnblockUser        = msgs.NewMsgUnblockUser
	RegisterMessagesCodec    = msgs.RegisterMessagesCodec

	// variable aliases
	RelationshipsStorePrefix = models.RelationshipsStorePrefix
	UsersBlocksStorePrefix   = models.UsersBlocksStorePrefix
	ModelsCdc                = models.ModelsCdc
	MsgsCodec                = msgs.MsgsCodec
)

type (
	RelationshipsResponse = models.RelationshipsResponse
	UserBlock             = models.UserBlock
	MsgCreateRelationship = msgs.MsgCreateRelationship
	MsgDeleteRelationship = msgs.MsgDeleteRelationship
	MsgBlockUser          = msgs.MsgBlockUser
	MsgUnblockUser        = msgs.MsgUnblockUser
)
