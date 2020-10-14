package types

const (
	// Relationships events
	EventTypeRelationshipCreated  = "relationship_created"
	EventTypeRelationshipsDeleted = "relationships_deleted"

	// Relationships attributes
	AttributeRelationshipSender   = "relationship_sender"
	AttributeRelationshipReceiver = "relationship_receiver"
	AttributeRelationshipSubspace = "relationship_subspace"

	// UserBlocks events
	EventTypeBlockUser   = "block_user"
	EventTypeUnblockUser = "unblock_user"

	// UserBlocks attributes
	AttributeUserBlockBlocker = "blocker"
	AttributeUserBlockBlocked = "blocked"
	AttributeUserBlockReason  = "reason"
	AttributeSubspace         = "subspace"
)
