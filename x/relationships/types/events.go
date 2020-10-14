package types

const (
	// Relationships events
	EventTypeRelationshipCreated  = "relationship_created"
	EventTypeRelationshipsDeleted = "relationships_deleted"

	// Relationships attributes
	AttributeRelationshipSender   = "sender"
	AttributeRelationshipReceiver = "receiver"
	AttributeRelationshipSubspace = "subspace"

	// UserBlocks events
	EventTypeBlockUser   = "block_user"
	EventTypeUnblockUser = "unblock_user"

	// UserBlocks attributes
	AttributeUserBlockBlocker = "userBlock_blocker"
	AttributeUserBlockBlocked = "userBlock_blocked"
	AttributeUserBlockReason  = "userBlock_reason"
	AttributeSubspace         = "subspace"
)
