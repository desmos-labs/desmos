package types

// DONTCOVER

const (
	EventTypeRelationshipCreated  = "relationship_created"
	EventTypeRelationshipsDeleted = "relationships_deleted"
	EventTypeBlockUser            = "block_user"
	EventTypeUnblockUser          = "unblock_user"

	AttributeValueCategory            = ModuleName
	AttributeRelationshipCreator      = "creator"
	AttributeRelationshipCounterparty = "counterparty"
	AttributeKeyUserBlockBlocker      = "blocker"
	AttributeKeyUserBlockBlocked      = "blocked"
	AttributeKeySubspace              = "subspace"
)
