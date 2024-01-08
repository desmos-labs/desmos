package types

// DONTCOVER

const (
	EventTypeRelationshipCreated = "create_relationship"
	EventTypeRelationshipDeleted = "delete_relationship"
	EventTypeBlockUser           = "block_user"
	EventTypeUnblockUser         = "unblock_user"

	AttributeRelationshipCreator      = "creator"
	AttributeRelationshipCounterparty = "counterparty"
	AttributeKeyUserBlockBlocker      = "blocker"
	AttributeKeyUserBlockBlocked      = "blocked"
	AttributeKeySubspace              = "subspace"
)
