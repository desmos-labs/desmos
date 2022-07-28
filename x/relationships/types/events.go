package types

// DONTCOVER

const (
	EventTypeRelationshipCreated = "create_relationship"
	EventTypeRelationshipDeleted = "delete_relationship"
	EventTypeBlockUser           = "block_user"
	EventTypeUnblockUser         = "unblock_user"

	AttributeValueCategory            = ModuleName
	AttributeRelationshipCreator      = "creator"
	AttributeRelationshipCounterparty = "counterparty"
	AttributeKeyUserBlockBlocker      = "blocker"
	AttributeKeyUserBlockBlocked      = "blocked"
	AttributeKeySubspace              = "subspace"
)
